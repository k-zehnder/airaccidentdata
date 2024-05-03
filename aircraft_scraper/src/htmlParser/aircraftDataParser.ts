import cheerio from 'cheerio';

export interface Parser {
  extractImageUrls(html: string): Promise<string[]>;
  modifyImageUrl(url: string): string;
  parseAircraftType(type: string): any;
}

const createCheerioParser = (): Parser => {
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

  const parseAircraftType = (type: string) => {
    // Extract registration number (last element)
    const registrationNumber = type.substring(type.lastIndexOf(' ') + 1);

    // Extract make and model from the remaining string
    const makeAndModel = type.substring(0, type.lastIndexOf(' ')).trim();

    // Split make and model by the last occurrence of a space
    const lastSpaceIndex = makeAndModel.lastIndexOf(' ');
    const make = makeAndModel.substring(0, lastSpaceIndex);
    const model = makeAndModel.substring(lastSpaceIndex + 1);

    return { make, model, registrationNumber };
  };

  return { extractImageUrls, modifyImageUrl, parseAircraftType };
};

export { createCheerioParser };
