// Handles AWS S3 uploads and database updates for aircraft image data.
import { S3Client, PutObjectCommand } from '@aws-sdk/client-s3';
import { Database } from '../../database/dbConnector';
import { determineFileName } from '../../helpers/fileUtils';

export interface AWSClient {
  uploadImagesAndHandleDb(
    images: { imageUrl: string; imageData: Buffer; aircraftId: number }[]
  ): Promise<void>;
}

// Creates an AWS client for handling uploads to S3 and database interactions
export const createAWSClient = (
  config: typeof import('../../config/config').default,
  db: Database
): AWSClient => {
  const s3Client = new S3Client({
    region: config.aws.region,
    credentials: {
      accessKeyId: config.aws.accessKeyId,
      secretAccessKey: config.aws.secretAccessKey,
    },
  });

  // Uploads images to S3 and updates the corresponding database records
  const uploadImagesAndHandleDb = async (
    images: { imageUrl: string; imageData: Buffer; aircraftId: number }[]
  ) => {
    for (const { imageUrl, imageData, aircraftId } of images) {
      console.log('Uploading image to S3:', imageUrl);
      const fileName = determineFileName(imageUrl);
      const contentType = 'image/jpeg';
      const uploadParams = {
        Bucket: config.aws.s3Bucket,
        Key: fileName,
        Body: imageData,
        ContentType: contentType,
      };

      try {
        await s3Client.send(new PutObjectCommand(uploadParams));
        const s3Url = `https://s.airaccidentdata.com/${fileName}`;
        await db.updateAircraftImageWithS3Url(aircraftId, imageUrl, s3Url);
      } catch (error) {
        console.error('Error uploading image to S3:', error);
      }
    }
  };

  return { uploadImagesAndHandleDb };
};
