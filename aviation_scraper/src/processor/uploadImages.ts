import { Database } from '../database/dbConnector';
import { AWSUploader } from '../awsIntegration/awsClient';
import { Fetcher } from '../imageFetcher/wikiImageFetcher';

export const uploadImagesAndUpdateDb = async (
  db: Database,
  awsUploader: AWSUploader,
  fetcher: Fetcher
): Promise<void> => {
  try {
    // Fetch images from the saved URLs
    const allAircraftImages = await db.getAllAircraftImages();

    for (const image of allAircraftImages) {
      const { aircraftId, imageUrl } = image;
      try {
        // Upload the fetched image to S3
        const s3Urls = await awsUploader.uploadImages([imageUrl], fetcher);

        if (s3Urls.length > 0) {
          const s3Url = s3Urls[0];

          // Extract the hash from the S3 URL
          const hash = s3Url.split('/').pop();

          // Construct the modified S3 URL with the fixed domain prefix
          const modifiedS3Url = `https://s.airaccidentdata.com/${hash}`;

          // Update the database with the modified S3 URL
          await db.updateAircraftImageWithS3Url(
            aircraftId,
            imageUrl,
            modifiedS3Url
          );

          console.log(
            `Updated image for aircraft with ID ${aircraftId} - S3 URL: ${modifiedS3Url}`
          );
        } else {
          console.error('Failed to upload image to S3.');
        }
      } catch (updateError) {
        console.error('Error updating image with S3 URL:', updateError);
      }
    }
  } catch (error) {
    console.error('Error uploading images and updating database:', error);
    throw error; // Rethrow the error for handling in the main function
  }
};
