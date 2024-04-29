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
    const aircraftTypes = await db.getAircraftTypes();

    // Scrape images for each aircraft type
    const aircraftDetails = await scraper.scrapeImages(
      aircraftTypes,
      fetcher,
      parser,
      config.aws.s3Bucket,
    );

    // Log and save details for each aircraft type
    for (const [type, details] of Object.entries(aircraftDetails)) {
      console.log(`Aircraft Type: ${type}`);
      console.log(`Image URLs: ${details.imageUrls}`);
      console.log(`Scraping Failed: ${details.failed}`);

      if (!details.failed) {
        const [make, model, registrationNumber] = type.split(' ');
        const aircraftId = await db.getAircraftIdByType(
          make,
          model,
          registrationNumber,
        );

        if (aircraftId) {
          try {
            // Save the Wikipedia URLs to the database
            for (let i = 0; i < details.imageUrls.length; i++) {
              const wikiUrl = details.imageUrls[i];
              await db.insertAircraftImage(aircraftId, wikiUrl);
              console.log(
                `Saved image URL for aircraft: ${type} - Wiki URL: ${wikiUrl}`,
              );
            }
          } catch (saveError) {
            console.error(
              'Error saving image URLs to the database:',
              saveError,
            );
          }
        } else {
          console.log(`Aircraft ID not found for type: ${type}`);
        }
      }
    }
  } catch (error) {
    console.error('Error processing images:', error);
    throw error; // Rethrow the error for handling in the main function
  }
};
