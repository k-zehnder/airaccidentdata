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
    const db = await createDatabaseConnection(config);
    const fetcher = createAxiosFetcher();
    const parser = createCheerioParser();
    const scraper = createScraper(db);
    const awsUploader = createS3BucketUploader(config);

    await processImages(db, fetcher, parser, scraper, config);
    await uploadImagesAndUpdateDb(db, awsUploader, fetcher);

    await db.close();
  } catch (error) {
    console.error('Error:', error);
  }
};

main();
