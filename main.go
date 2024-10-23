package main

import (
	convertdoctopdf "conversion/ConvertDocToPDF"
	"log"
)

var inputFilePath string = "C:\\Users\\Soft suave\\Downloads\\Naukri_Prashantsahu[2y_7m].doc"
var outputFilePath string = "C:\\Users\\Soft suave\\Downloads\\output9.pdf"

func main() {

	err := convertdoctopdf.ConvertDocToPDF(inputFilePath, outputFilePath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
