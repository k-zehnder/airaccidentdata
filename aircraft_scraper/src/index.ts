/**
 * Orchestrates the image scraping and uploading pipeline for aircraft data.
 */
import config from './config/config';
import { createAxiosFetcher } from './imageFetcher/wikiImageFetcher';
import { createCheerioParser } from './htmlParser/aircraftDataParser';
import { createScraper } from './wikiScraper/aircraftImageScraper';
import { createS3BucketUploader } from './awsIntegration/awsClient';
import { createDatabaseConnection } from './database/dbConnector';
import { processImages } from './imageProcessor/imageProcessor';
import { uploadImagesAndUpdateDb } from './imageProcessor/imageUploader';

const main = async (): Promise<void> => {
  try {
    // Initialize database and functional component instances
    const db = await createDatabaseConnection(config);
    const imageFetcher = createAxiosFetcher();
    const htmlParser = createCheerioParser();
    const aviationScraper = createScraper(db);
    const awsClient = createS3BucketUploader(config);

    // Retrieve images from Wikipedia and store them in the database
    await processImages(db, imageFetcher, htmlParser, aviationScraper, config);

    // Upload images to S3 and update the database
    await uploadImagesAndUpdateDb(db, awsClient, imageFetcher);

    // Close database to free resources
    await db.close();
  } catch (error) {
    console.error('Error in main process:', error);
  }
};

main();
