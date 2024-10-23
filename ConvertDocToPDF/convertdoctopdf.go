package convertdoctopdf

import (
	"fmt"
	"os"

	convertapi "github.com/ConvertAPI/convertapi-go/pkg"
	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"github.com/ConvertAPI/convertapi-go/pkg/param"
)

func ConvertDocToPDF(inputFilePath string, outputFilePath string) error {
	config.Default = config.NewDefault("secret_pZQl7qj5eYjjBY5x")
	_, err := convertapi.ConvDef("doc", "pdf",
		param.NewPath("File", inputFilePath, nil),
	).ToPath(outputFilePath)
	if err != nil {
		return fmt.Errorf("conversion failed: %v", err)
	}
	fileInfo, _ := os.Stat(outputFilePath)

	if fileInfo.Size() > 150*1024 {
		return fmt.Errorf("size is more than 150kb and the size is: %d bytes", fileInfo.Size())
	} else {
		fmt.Println("Size is less than 150kb and the size is ", fileInfo.Size())
	}
	return nil
}
