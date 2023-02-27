package pdfextractor

import (
	"bytes"

	"github.com/dslipak/pdf"
)

// GetPdfText prints out contents of PDF file to stdout.
func GetPdfText(inputPath string) (*string, error) {
	r, err := pdf.Open(inputPath)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return nil, err
	}
	buf.ReadFrom(b)
	textString := buf.String()
	return &textString, nil
}
