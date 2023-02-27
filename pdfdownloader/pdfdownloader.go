package pdfdownloader

import (
	"io"
	"net/http"
	"os"
)

// DownloadPdf downloads a PDF file from a URL.
func DownloadPdf(url string, fileName string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
