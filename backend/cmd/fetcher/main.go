package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os/exec"

	"github.com/chromedp/chromedp"
)

// Function to fetch the CSV download link from a given URL using Puppeteer
func fetchCsvDownloadLink(url string) (string, error) {
	// Create a new Chrome instance
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Navigate to the URL
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`#div_wid a.bodytextlink`),
	); err != nil {
		return "", err
	}

	// Get the href attribute of the download link
	var downloadLink string
	if err := chromedp.Run(ctx, chromedp.Evaluate(`document.querySelector('#div_wid a.bodytextlink').href`, &downloadLink)); err != nil {
		return "", err
	}

	if downloadLink == "" {
		return "", errors.New("CSV download link not found")
	}

	return downloadLink, nil
}

// Function to download a CSV file using the `curl` shell command.
func downloadCsvFile(downloadLink, outputPath string) error {
	// Run curl command to download the file
	cmd := exec.Command("curl", "-o", outputPath, downloadLink)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error downloading CSV file: %v, output: %s", err, output)
	}
	return nil
}

func main() {
	// URL to fetch the CSV download link from
	pageURL := "https://www.asias.faa.gov/apex/f?p=100:93:::NO:::"

	// Output path for the downloaded CSV file
	outputPath := "downloaded_file.csv"

	// Fetch the CSV download link
	downloadLink, err := fetchCsvDownloadLink(pageURL)
	if err != nil {
		log.Fatalf("Error fetching CSV download link: %v", err)
	}

	// Download the CSV file
	if err := downloadCsvFile(downloadLink, outputPath); err != nil {
		log.Fatalf("Error downloading CSV file: %v", err)
	}

	fmt.Println("CSV file downloaded successfully")
}
