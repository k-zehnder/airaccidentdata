/**
 * Orchestrates the image scraping and uploading pipeline for aircraft data.
*/
import config from './config/config';
import { createAxiosFetcher } from './fetcher/fetcher';
import { createCheerioParser } from './parser/parser';
import { createScraper } from './scraper/scraper';
import { createS3BucketUploader } from './aws/aws';
import { createDatabaseConnection } from './database/connection';
import { processImages } from './processor/processImages';
import { uploadImagesAndUpdateDb } from './processor/uploadImages';

const main = async (): Promise<void> => {
  try {
    // Initialize database and functional component instances
    const db = await createDatabaseConnection(config);
    const fetcher = createAxiosFetcher();
    const parser = createCheerioParser();
    const scraper = createScraper(db);
    const awsUploader = createS3BucketUploader(config);

    // Retrieve images from Wikipedia and store them in the database
    await processImages(db, fetcher, parser, scraper, config);

    // Upload images to S3 and update the database
    await uploadImagesAndUpdateDb(db, awsUploader, fetcher);

    // Close database to free resources
    await db.close();
  } catch (error) {
    console.error('Error in main process:', error);
  }
};

main();
