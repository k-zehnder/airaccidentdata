// Package main provides functionality to fetch and download a CSV file from a given URL.
package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"

	"github.com/chromedp/chromedp"
)

// main is the entry point of the application. It loads environment variables from a .env file,
// sets up a connection to the MySQL database, and processes a CSV file to insert data into the database.
func main() {
	pageURL := "https://www.asias.faa.gov/apex/f?p=100:93:::NO:::"
	outputPath := "downloaded_file.csv"

	downloadLink, err := fetchCsvDownloadLink(pageURL)
	if err != nil {
		log.Fatalf("Error fetching CSV download link: %v", err)
	}

	if err := downloadCsvFile(downloadLink, outputPath); err != nil {
		log.Fatalf("Error downloading CSV file: %v", err)
	}

	fmt.Println("CSV file downloaded successfully")
}

// fetchCsvDownloadLink fetches the CSV download link from a given URL using chromedp.
func fetchCsvDownloadLink(url string) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var downloadLink string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`#div_wid a.bodytextlink`),
		chromedp.Evaluate(`document.querySelector('#div_wid a.bodytextlink').href`, &downloadLink),
	); err != nil {
		return "", err
	}

	if downloadLink == "" {
		return "", fmt.Errorf("CSV download link not found")
	}

	return downloadLink, nil
}

// downloadCsvFile downloads a CSV file using the `curl` shell command.
func downloadCsvFile(downloadLink, outputPath string) error {
	cmd := exec.Command("curl", "-o", outputPath, downloadLink)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error downloading CSV file: %v, output: %s", err, output)
	}
	return nil
}
