package convertdoctopdf_test

import (
	convertdoctopdf "conversion/ConvertDocToPDF"
	"testing"
)

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}
func customError(msg string) error {
	return &MyError{Message: msg}
}
func TestConvertDocToPDF(t *testing.T) {
	testCases := []struct {
		name   string
		input  []string
		output error
	}{
		{
			name: "Success",
			input: []string{
				"C:\\Users\\Soft suave\\Downloads\\Naukri_Prashantsahu[2y_7m].doc",
				"C:\\Users\\Soft suave\\Downloads\\output13.pdf",
			},
			output: nil,
		},
		{
			name: "Failure - File too large",
			input: []string{
				"C:\\Users\\Soft suave\\Downloads\\file-sample_500kB.doc",
				"C:\\Users\\Soft suave\\Downloads\\output9.pdf",
			},
			output: customError("size is more than 150kb"),
		},
		{
			name: "Failure - Conversion error",
			input: []string{
				"C:\\Users\\Soft suave\\Downloads\\Naukri_Prashantsahu[2y_7m].doc",
				"C:\\Users\\Soft suave\\Downloads\\output9.pdf",
			},
			output: customError("conversion failed Correct Your Path"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := convertdoctopdf.ConvertDocToPDF(tc.input[0], tc.input[1])

			if got != nil {
				if got.Error() != tc.output.Error() {
					t.Errorf("This is Error")
				}
			}

		})
	}
}
