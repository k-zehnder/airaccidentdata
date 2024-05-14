// Coordinates the scraping and handling of aircraft images.
import { Fetcher } from '../../helpers/wikiImageFetcher';
import { Database } from '../../database/dbConnector';

export interface AircraftImageScraper {
  scrapeImages(): Promise<
    { imageUrl: string; imageData: Buffer; aircraftId: number }[]
  >;
}

// Creates an aircraft image scraper with database and fetcher dependencies.
export const createAircraftImageScraper = (
  db: Database,
  fetcher: Fetcher
): AircraftImageScraper => {
  const scrapeImages = async (): Promise<
    { imageUrl: string; imageData: Buffer; aircraftId: number }[]
  > => {
    const aircraftTypes = await db.getAircraftTypes();
    const imagesToUpload: {
      imageUrl: string;
      imageData: Buffer;
      aircraftId: number;
    }[] = [];

    for (const aircraft of aircraftTypes) {
      try {
        const images = await fetcher.fetchImages(aircraft);
        const typeKey = `${aircraft.make} ${aircraft.model} ${aircraft.registrationNumber}`;
        console.log(typeKey, images.length);

        if (images.length > 0) {
          const dbAircraftId = await db.getAircraftIdByType(
            aircraft.make,
            aircraft.model,
            aircraft.registrationNumber
          );

          if (dbAircraftId !== null) {
            const imageDetails = await Promise.all(
              images.map(async (imageUrl) => {
                try {
                  const imageData = await fetcher.fetchImageFromUrl(imageUrl);
                  return { imageUrl, imageData, aircraftId: dbAircraftId };
                } catch (error) {
                  console.error(`Error fetching image ${imageUrl}:`, error);
                  return null; // Return null for failed fetches
                }
              })
            );

            // Filter out any null entries from failed fetches
            const validImageDetails = imageDetails.filter(Boolean) as {
              imageUrl: string;
              imageData: Buffer;
              aircraftId: number;
            }[];

            imagesToUpload.push(...validImageDetails);

            // Save image URLs to the database
            for (const { imageUrl } of validImageDetails) {
              await db.insertAircraftImage(dbAircraftId, imageUrl);
              console.log(
                `Successfully saved image for ${typeKey}: Original URL: ${imageUrl}`
              );
            }
          } else {
            console.log(
              `Skipping saving images for aircraft ${typeKey} as it was not found in the database.`
            );
          }
        }
      } catch (error) {
        console.error(
          `Error processing aircraft ${aircraft.make} ${aircraft.model} ${aircraft.registrationNumber}:`,
          error
        );
      }
    }

    return imagesToUpload;
  };

  return { scrapeImages };
};
