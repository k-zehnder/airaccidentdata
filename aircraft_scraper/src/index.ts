/**
 * Orchestrates the image scraping and uploading pipeline for aircraft data.
 */
import config from './config/config';
import { createDatabaseConnection } from './database/dbConnector';
import { createAxiosFetcher } from './helpers/wikiImageFetcher';
import { createCheerioParser } from './helpers/aircraftHtmlParser';
import { createAircraftImageScraper } from './services/scraper/aircraftImageScraper';
import { createAWSClient } from './services/aws/awsClient';

// Main function to coordinate scraping and uploading images.
const main = async () => {
  // Initialize database connection
  const db = await createDatabaseConnection(config);

  // Initialize the necessary helpers and clients
  const parser = createCheerioParser();
  const fetcher = createAxiosFetcher(parser);
  const aircraftScraper = createAircraftImageScraper(db, fetcher);

  try {
    // Scrape images using the aircraft scraper
    const aircraftImages = await aircraftScraper.scrapeImages();

    if (config.nodeEnv !== 'development') {
      // Upload images and handle database updates in non-development environments
      const awsClient = createAWSClient(config, db);
      await awsClient.uploadImagesAndHandleDb(aircraftImages);
    } else {
      console.log('Development mode: Skipping S3 upload.');
    }

    console.log('All images processed.');
  } catch (error) {
    console.error('Error in main process:', error);
  } finally {
    // Ensure the database connection is closed
    await db.close();
  }
};

// Execute the main function
main();
