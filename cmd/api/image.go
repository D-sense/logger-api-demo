package api

import (
	"context"
	"encoding/json"
	logger "go-apps/ocr-web-service/log"
	"go-apps/ocr-web-service/model"
	"go-apps/ocr-web-service/services"
	"go-apps/ocr-web-service/uploader"
	"go-apps/ocr-web-service/utils"
	"net/http"
)

type ImageController struct {
	services.ImageService
}

var imgModel model.Image
var successResp utils.Success
var errorResp utils.Error

//logger
var standardLogger = logger.NewLogger()

func (image *ImageController) ImageAsync(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&imgModel)

		// validate if image input
		if validationResult := utils.ValidateImageInput(imgModel.ImageData); len(validationResult) != 0 {
			standardLogger.MissingArg("imgModel.ImageData")
			errorResp.RespondWithErrorJSON(w, utils.DataValidationErr, validationResult)
			return
		}

		//upload image
		imgInfo := uploader.NewImage(utils.ImagePath, imgModel.ImageData)
		imgPath, err := image.Ocr.UploadImage(ctx, imgInfo)
		if err != nil {
			standardLogger.InvalidArgValue("imgModel.ImageData", imgModel.ImageData)
			errorResp.RespondWithErrorJSON(w, utils.BadRequest, err.Error())
			return
		}

		//respond to request
		successResp.RespondWithSuccessTextJSON(w, imgPath)
	}
}
