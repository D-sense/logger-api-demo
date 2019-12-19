// Provides interface for extracting texts from an image
package ocr

import (
	"context"
	"fmt"
	"github.com/otiai10/gosseract"
)

func (ImageRepo) TextExtractor(ctx context.Context, imgPath string ) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	err := client.SetImage(imgPath)
	if err != nil {
		return "", err
	}
	client.Languages = []string{"eng"}

	text, err := client.Text()
    if err != nil {
		return "", err
	}
    fmt.Println(text)
	return text, nil
}
