package services

import(
	"context"
	"go-apps/ocr-web-service/ocr"
)

type OCRService interface{
	UploadImage(ctx context.Context, imgInfo *ocr.ImageRepo) (string, error)
    TextExtractor(ctx context.Context, imgPath string) (string, error)
}

type ImageService struct {
	Ocr OCRService
}