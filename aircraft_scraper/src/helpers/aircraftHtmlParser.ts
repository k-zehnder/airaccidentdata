// Provides a parser using Cheerio to extract and modify image URLs from HTML and parse aircraft types.
import cheerio from 'cheerio';

export interface Parser {
  extractImageUrls(html: string): Promise<string[]>;
  modifyImageUrl(url: string): string;
  parseAircraftType(type: string): any;
}

// Creates a Cheerio parser
const createCheerioParser = (): Parser => {
  // Modifies the image URL to remove resolution and format it correctly
  const modifyImageUrl = (url: string): string => {
    if (!url.startsWith('//')) {
      return url;
    }
    const parts = url.split('/');
    let filename = parts[parts.length - 1];

    // Remove the resolution part from the filename
    const resolutionRegex = /\d+px-/;
    filename = filename.replace(resolutionRegex, '');

    return `https://en.wikipedia.org/wiki/File:${filename}`;
  };

  // Extracts image URLs from the provided HTML
  const extractImageUrls = async (html: string): Promise<string[]> => {
    const $ = cheerio.load(html);
    const imageUrls: string[] = [];

    $('img').each((_, element) => {
      const url = $(element).attr('src');
      if (url) {
        const modifiedUrl = modifyImageUrl(url);
        console.log(`Modified URL: ${modifiedUrl}`);
        imageUrls.push(modifiedUrl);
      }
    });

    return imageUrls;
  };

  // Parses the aircraft type string into make, model, and registration number
  const parseAircraftType = (type: string) => {
    const registrationNumber = type.substring(type.lastIndexOf(' ') + 1);
    const makeAndModel = type.substring(0, type.lastIndexOf(' ')).trim();
    const lastSpaceIndex = makeAndModel.lastIndexOf(' ');
    const make = makeAndModel.substring(0, lastSpaceIndex);
    const model = makeAndModel.substring(lastSpaceIndex + 1);

    return { make, model, registrationNumber };
  };

  return { extractImageUrls, modifyImageUrl, parseAircraftType };
};

export { createCheerioParser };
