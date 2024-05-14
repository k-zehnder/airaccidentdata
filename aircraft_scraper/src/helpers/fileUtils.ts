// Provides utility functions for generating hashes and determining file names for image URLs.
import crypto from 'crypto';

// Generates a hash from the input string.
export function generateHash(input: string): string {
  return crypto.createHash('md5').update(input).digest('hex');
}

// Determines the file name based on the image URL.
export function determineFileName(imageUrl: string): string {
  const urlHash = generateHash(imageUrl);
  const timestamp = Date.now(); // Use current timestamp for uniqueness
  return `${urlHash}-${timestamp}.jpg`; // Example: 'f84f305a6f1e484017eaf3c2c60adc12-1597765337031.jpg'
}
