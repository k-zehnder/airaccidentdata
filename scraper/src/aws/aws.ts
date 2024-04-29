import { S3Client, PutObjectCommand } from "@aws-sdk/client-s3";
import { Fetcher } from "../fetcher/fetcher";
import crypto from 'crypto';

// Function to generate a hash from the image URL
function generateHash(input: string): string {
    return crypto.createHash('md5').update(input).digest('hex');
}

// Function to determine the file name based on the image URL
function determineFileName(imageUrl: string): string {
    const urlHash = generateHash(imageUrl);
    const timestamp = Date.now(); // Use current timestamp for uniqueness
    return `${urlHash}-${timestamp}.jpg`; // Example: 'f84f305a6f1e484017eaf3c2c60adc12-1597765337031.jpg'
}

export interface AWSUploader {
    uploadImages(urls: string[], fetcher: Fetcher): Promise<string[]>;
}

export const createS3BucketUploader = (config: typeof import("../config/config").default): AWSUploader => {
    const s3Client = new S3Client({
        region: config.aws.region,
        credentials: {
            accessKeyId: config.aws.accessKeyId,
            secretAccessKey: config.aws.secretAccessKey
        }
    });
    
    const uploadImageToS3 = async (imageUrl: string, fetcher: Fetcher): Promise<string> => {
        console.log('Uploading image to S3:', imageUrl); // Log the image URL being uploaded

        // Use the fetcher to fetch the image data
        const imageData = await fetcher.fetchImageFromUrl(imageUrl);

        const fileName = determineFileName(imageUrl); 
        const contentType = 'image/jpeg'; 

        const uploadParams = {
            Bucket: config.aws.s3Bucket,
            Key: fileName,
            Body: imageData,
            ContentType: contentType
        };

        try {
            console.log('Sending PutObjectCommand:', uploadParams); // Log the upload params before sending

            await s3Client.send(new PutObjectCommand(uploadParams));

            const s3Url = `https://${config.aws.s3Bucket}.s3.${config.aws.region}.amazonaws.com/${fileName}`;
            console.log('Image uploaded successfully. S3 URL:', s3Url); // Log the S3 URL after successful upload

            return s3Url;
        } catch (error) {
            console.error('Error uploading image to S3:', error);
            throw error;
        }
    };

    const uploadImages = async (urls: string[], fetcher: Fetcher): Promise<string[]> => {
        console.log('Uploading multiple images to S3...');
        console.log('Image URLs:', urls); // Log the image URLs being uploaded
        
        const s3Urls: string[] = [];
        
        for (const url of urls) {
            try {
                const s3Url = await uploadImageToS3(url, fetcher);
                s3Urls.push(s3Url);
            } catch (error) {
                console.error('Error uploading image to S3:', error);
                s3Urls.push(''); // Push an empty string for failed uploads
            }
        }
        
        return s3Urls;
    };

    return { uploadImages };
};
