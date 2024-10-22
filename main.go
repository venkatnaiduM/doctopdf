package main

import (
	"fmt"
	"log"
	"os"

	// "github.com/ConvertAPI/convertapi-go"
	// convertapi "github.com/ConvertAPI/convertapi-go"
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
		return fmt.Errorf("conversion failed: %w", err)
	}
	fileInfo, _ := os.Stat(outputFilePath)

	if fileInfo.Size() > 60*1024 {
		return fmt.Errorf("generated PDF exceeds size limit: %d bytes", fileInfo.Size())
	} else {
		fmt.Println("Size is less than 150kb", fileInfo.Size())
	}
	return nil
}

func main() {
	inputFilePath := "C:\\Users\\Soft suave\\Downloads\\Naukri_Prashantsahu[2y_7m].doc"
	outputFilePath := "C:\\Users\\Soft suave\\Downloads\\output6.pdf"
	err := ConvertDocToPDF(inputFilePath, outputFilePath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
