import fs from 'fs';
import cheerio from 'cheerio';
import { Fetcher } from '../imageFetcher/wikiImageFetcher';
import { Parser } from '../jsonParser/aircraftDataParser';
import { Database } from '../databaseConnection/dbConnector';
import { AircraftType, AircraftMapping } from '../types/aircraft';

export interface Scraper {
  scrapeImages(
    aircraftTypes: AircraftType[],
    fetcher: Fetcher,
    parser: Parser,
    s3Url: string
  ): Promise<{ [key: string]: { imageUrls: string[]; failed: boolean } }>;
}

// Function to create a scraper object with necessary dependencies and mapping data.
export const createScraper = (db: Database): Scraper => {
  // Function to load the aircraft type mapping data from a JSON file.
  const loadAircraftTypeMap = (
    filePath: string
  ): Map<string, AircraftMapping> => {
    try {
      const data = fs.readFileSync(filePath);
      const jsonData = JSON.parse(data.toString());
      return new Map(Object.entries(jsonData));
    } catch (error) {
      console.error(
        `Error reading or parsing aircraft type map JSON file: ${error}`
      );
      return new Map();
    }
  };

  // This map will be used to find URLs for aircraft images based on their type.
  const aircraftTypeMap = loadAircraftTypeMap(
    'src/jsonParser/aircraftMapping.json'
  );

  // Function to fetch the original high-resolution image URL from a Wikipedia file page URL
  const fetchOriginalImageUrl = async (
    filePageUrl: string,
    fetcher: Fetcher
  ): Promise<string> => {
    try {
      const htmlContent = await fetcher.fetchHtmlFromUrl(filePageUrl);
      const $ = cheerio.load(htmlContent);
      const originalFileLink = $('.fullImageLink a').attr('href');
      return originalFileLink ? `https:${originalFileLink}` : '';
    } catch (error) {
      console.error(
        `Error fetching original image URL from ${filePageUrl}: ${error}`
      );
      return '';
    }
  };

  // Function to fetch images from Wikipedia based on an aircraft's make and model
  const handleIndirectMapping = async (
    mappingUrl: string,
    fetcher: Fetcher,
    parser: Parser
  ): Promise<string> => {
    const indirectHtmlContent = await fetcher.fetchHtmlFromUrl(
      `https://en.wikipedia.org/wiki/${mappingUrl}`
    );
    const indirectImageUrls = await parser.extractImageUrls(
      indirectHtmlContent
    );
    const firstRelevantThumbnailUrl = indirectImageUrls.find(
      (url) => url.includes('File:') && !url.includes('Commons-logo')
    );

    if (firstRelevantThumbnailUrl) {
      const filePageUrl = firstRelevantThumbnailUrl.startsWith('/wiki/File:')
        ? `https://en.wikipedia.org${firstRelevantThumbnailUrl}`
        : firstRelevantThumbnailUrl;
      return await fetchOriginalImageUrl(filePageUrl, fetcher);
    }

    return '';
  };

  // Function to fetch images from Wikipedia based on an aircraft's make and model
  const fetchImagesFromWikipedia = async (
    aircraft: AircraftType,
    fetcher: Fetcher,
    parser: Parser
  ): Promise<string[]> => {
    const searchTerms = `${aircraft.make} ${aircraft.model}`;
    const searchUrl = `https://en.wikipedia.org/w/index.php?search=${encodeURIComponent(
      searchTerms
    )}`;

    try {
      const htmlContent = await fetcher.fetchHtmlFromUrl(searchUrl);
      const $ = cheerio.load(htmlContent);
      const searchResults = $('.mw-search-result-heading a');

      if (searchResults.length === 0) {
        // If search results list doesn't exist, treat it as a direct mapping
        console.log(
          `No search results list found for ${searchTerms}. Assuming it directed us right to the page.`
        );
        // Directly fetch image URLs from the content page
        const imageUrls = await parser.extractImageUrls(htmlContent);
        const filteredImageUrls = imageUrls.filter((url) =>
          url.includes('File:')
        );
        return filteredImageUrls.map((url) => {
          if (url.startsWith('https://')) {
            return url; // Already a full URL
          } else {
            return `https://en.wikipedia.org${url}`; // Append Wikipedia domain
          }
        });
      }

      // If search results list exists, proceed with parsing and extracting image URLs
      const topResultUrl = `https://en.wikipedia.org${searchResults
        .eq(0)
        .attr('href')}`;
      const topResultHtml = await fetcher.fetchHtmlFromUrl(topResultUrl);
      const imageUrls = await parser.extractImageUrls(topResultHtml);

      const filteredImageUrls = imageUrls.filter((url) =>
        url.includes('File:')
      );
      return filteredImageUrls.map((url) => {
        if (url.startsWith('https://')) {
          return url; // Already a full URL
        } else {
          return `https://en.wikipedia.org${url}`; // Append Wikipedia domain
        }
      });
    } catch (error) {
      console.error(
        `Error fetching images from Wikipedia for ${searchTerms}: ${error}`
      );
      return [];
    }
  };

  // Function to fetch images for a given aircraft. It checks the aircraftTypeMap for 'exact' or 'indirect' URLs,
  // and falls back to fetching from Wikipedia if no mapping is found.
  const fetchImages = async (
    aircraft: AircraftType,
    fetcher: Fetcher,
    parser: Parser
  ): Promise<string[]> => {
    const searchTermsForMap = `${aircraft.make} ${aircraft.model}`;
    console.log(`Attempting to fetch images for ${searchTermsForMap}`);

    try {
      let highResImageUrl = '';

      if (aircraftTypeMap.has(searchTermsForMap)) {
        const mapping = aircraftTypeMap.get(searchTermsForMap);
        if (mapping) {
          switch (mapping.type) {
            case 'exact':
              highResImageUrl = mapping.url;
              break;
            case 'indirect':
              highResImageUrl = await handleIndirectMapping(
                mapping.url,
                fetcher,
                parser
              );
              break;
            default:
              break;
          }
        }
      }

      if (!highResImageUrl) {
        // Fallback to Wikipedia search if no URLs found
        console.log(
          `No direct mapping found for ${searchTermsForMap}. Searching Wikipedia...`
        );
        const wikipediaImages = await fetchImagesFromWikipedia(
          aircraft,
          fetcher,
          parser
        );
        if (wikipediaImages.length > 0) {
          let filePageUrl = wikipediaImages[0];
          if (!filePageUrl.startsWith('https://')) {
            filePageUrl = `https://en.wikipedia.org${filePageUrl}`;
          }
          console.log(`Using image from Wikipedia for ${searchTermsForMap}`);
          highResImageUrl = await fetchOriginalImageUrl(filePageUrl, fetcher);
        }
      }

      console.log(
        `Determined URL for ${searchTermsForMap}: ${highResImageUrl}`
      );
      return highResImageUrl ? [highResImageUrl] : [];
    } catch (error) {
      console.error(`Error fetching images for ${searchTermsForMap}: ${error}`);
      return [];
    }
  };

  // Function to scrape images for multiple aircraft types. This function iterates through each aircraft type,
  // fetches images using the fetchImages function, and then processes the results.
  const scrapeImages = async (
    aircraftTypes: AircraftType[],
    fetcher: Fetcher,
    parser: Parser
  ): Promise<{ [key: string]: { imageUrls: string[]; failed: boolean } }> => {
    const aircraftDetails: {
      [key: string]: { imageUrls: string[]; failed: boolean };
    } = {};

    for (const aircraft of aircraftTypes) {
      const images = await fetchImages(aircraft, fetcher, parser);
      const typeKey = `${aircraft.make} ${aircraft.model} ${aircraft.registrationNumber}`;

      // Debug log
      console.log(typeKey, images.length);

      const failed = images.length === 0;
      aircraftDetails[typeKey] = {
        imageUrls: images,
        failed: failed,
      };

      // Save image URLs to the database
      if (!failed) {
        try {
          const dbAircraftId = await db.getAircraftIdByType(
            aircraft.make,
            aircraft.model,
            aircraft.registrationNumber
          );
          if (dbAircraftId !== null) {
            for (const imageUrl of images) {
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
        } catch (error) {
          console.error(
            `Error saving images to the database for aircraft ${typeKey}: ${error}`
          );
        }
      }
    }

    return aircraftDetails;
  };

  return { scrapeImages };
};
