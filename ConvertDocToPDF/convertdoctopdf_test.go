package convertdoctopdf_test

import (
	convertdoctopdf "conversion/ConvertDocToPDF"
	"testing"
)

func TestConvertDocToPDF(t *testing.T) {
	inputFile := "C:\\Users\\Soft suave\\Downloads\\Naukri_Prashantsahu[2y_7m].doc"
	outputFile := "C:\\Users\\Soft suave\\Downloads\\output18.pdf"

	err := convertdoctopdf.ConvertDocToPDF(inputFile, outputFile)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}
