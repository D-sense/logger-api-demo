package services

import(
	"context"
	"go-apps/ocr-web-service/uploader"
)

type OCRService interface{
	UploadImage(ctx context.Context, imgInfo *uploader.ImageRepo) (string, error)
}

type ImageService struct {
	Ocr OCRService
}