import { Database } from '../database/connection';
import { Fetcher } from '../fetcher/fetcher';
import { Parser } from '../parser/parser';
import { Scraper } from '../scraper/scraper';

export const processImages = async (
  db: Database,
  fetcher: Fetcher,
  parser: Parser,
  scraper: Scraper,
  config: any
): Promise<void> => {
  try {
    // Fetch all aircraft types including their IDs at the start
    const aircraftTypes = await db.getAircraftTypes();

    // Scrape images for each aircraft type
    const aircraftDetails = await scraper.scrapeImages(
      aircraftTypes,
      fetcher,
      parser,
      config.aws.s3Bucket
    );

    // Log and save details for each aircraft
    for (const aircraft of aircraftTypes) {
      const details =
        aircraftDetails[
          `${aircraft.make} ${aircraft.model} ${aircraft.registrationNumber}`
        ];

      if (details && !details.failed) {
        console.log(
          `Aircraft ID: ${aircraft.id}, Image URLs: ${details.imageUrls.join(
            ', '
          )}`
        );

        // Save the image URLs to the database using the ID
        for (const imageUrl of details.imageUrls) {
          await db.insertAircraftImage(aircraft.id, imageUrl);
          console.log(
            `Saved image URL for aircraft ID: ${aircraft.id} - Image URL: ${imageUrl}`
          );
        }
      } else {
        console.log(
          `Scraping failed or no details found for Aircraft ID: ${aircraft.id}`
        );
      }
    }
  } catch (error) {
    console.error('Error processing images:', error);
    throw error; // Rethrow the error for handling in the main function
  }
};
