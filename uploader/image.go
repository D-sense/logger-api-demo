// Provides interface for uploading, and Fetching images
package uploader

import (
	"context"
	"encoding/base64"
	"go-apps/ocr-web-service/utils"
	"os"
	"strings"
)

type ImageRepo struct {
	pathName string
    image    string
}

func NewImage(pathName, image string) *ImageRepo {
    return &ImageRepo{
		pathName, image,
	}
}

func (ImageRepo) UploadImage(ctx context.Context, imgInfo *ImageRepo) (string, error) {
	extension := ".png"
	base64image := imgInfo.image[strings.IndexByte(imgInfo.image, ',')+1:]
	dec, err := base64.StdEncoding.DecodeString(base64image)
	if err != nil {
		return "", err
	}

	fileName := (utils.GenerateRandomWord(50)) + extension
	f, err := os.Create(imgInfo.pathName + "/" + fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		return "", err
	}
	if err := f.Sync(); err != nil {
		return "", err
	}

	f.Seek(0, 0)

	return f.Name(), nil
}


