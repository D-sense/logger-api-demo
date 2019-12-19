package api

import (
	"encoding/json"
	"fmt"
	logger "go-apps/ocr-web-service/log"
	"go-apps/ocr-web-service/ocr"
	"go-apps/ocr-web-service/model"
	"go-apps/ocr-web-service/services"
	"go-apps/ocr-web-service/utils"
	"net/http"
	"context"
)

type ImageController struct {
	services.ImageService
}

var imgModel model.Image
var successResp utils.Success
var errorResp utils.Error
var standardLogger = logger.NewLogger()

func (image *ImageController) ImageAsync(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&imgModel)

		// validate if image input
		if validationResult := utils.ValidateImageInput(imgModel.ImageData); len(validationResult) != 0 {
			standardLogger.InvalidArgValue(imgModel.ImageData, imgModel.ImageData)
			errorResp.RespondWithErrorJSON(w, utils.DataValidationErr, validationResult)
			return
		}

		//upload image
		imgInfo := ocr.NewImage(utils.ImagePath, imgModel.ImageData)
		imgPath, err := image.Ocr.UploadImage(ctx, imgInfo)
		if err != nil {
			errorResp.RespondWithErrorJSON(w, utils.BadRequest, err.Error())
			return
		}

		//extract text from the image
		text, err := image.Ocr.TextExtractor(ctx, imgPath)
		if err != nil {
			errorResp.RespondWithErrorJSON(w, utils.BadRequest, err.Error())
			return
		}

		//respond to request
		successResp.RespondWithSuccessTextJSON(w, text)
	}
}

func (image *ImageController) CreateImage(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&imgModel)

		// validate if image input
		if validationResult := utils.ValidateImageInput(imgModel.ImageData); len(validationResult) != 0 {
			//standardLogger.InvalidArgValue(imgModel.ImageData, imgModel.ImageData)
			errorResp.RespondWithErrorJSON(w, utils.DataValidationErr, validationResult)
			return
		}

		//upload image
		imgInfo := ocr.NewImage(utils.ImagePath, imgModel.ImageData)
		imageUpload, err := image.Ocr.UploadImage(ctx, imgInfo)
		if err != nil {
			errorResp.RespondWithErrorJSON(w, utils.BadRequest, err.Error())
			return
		}

		// base64 encoding image name
		encoded := utils.EncodeStringToBase64(imageUpload)
		fmt.Println("encoded Id: ", encoded)

		//respond to request with ID
		successResp.RespondWithSuccessTaskIdJSON(w, encoded)
	}
}

func (image *ImageController) GetTextByID(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&imgModel)

		// validate if image input
		if validationResult := utils.ValidateImageInput(imgModel.TaskId); len(validationResult) != 0 {
			standardLogger.InvalidArgValue(imgModel.ImageData, imgModel.ImageData)
			errorResp.RespondWithErrorJSON(w, utils.DataValidationErr, validationResult)
			return
		}

		// check taskId provided
		decodedId, err := ocr.CheckImageID(imgModel.TaskId)
		if err != nil {
			errorResp.RespondWithErrorJSON(w, utils.BadRequest, utils.IdNotFound)
			return
		}

		//extract text from the image
		text, err := image.Ocr.TextExtractor(ctx, decodedId)
		if err != nil {
			errorResp.RespondWithErrorJSON(w, utils.BadRequest, err.Error())
			return
		}

		//respond to request
		successResp.RespondWithSuccessTaskIdJSON(w, text)
	}
}

