// Provides functions to fetch HTML, images, and specific aircraft images from URLs using Axios and Cheerio.
import axios from 'axios';
import cheerio from 'cheerio';
import { loadAircraftTypeMap } from './mappingLoader';
import { Parser } from './aircraftHtmlParser';
import { AircraftType } from '../types/aircraft';
import config from '../config/config';

export interface Fetcher {
  fetchHtmlFromUrl(url: string): Promise<string>;
  fetchImageFromUrl(url: string): Promise<Buffer>;
  fetchOriginalImageUrl(filePageUrl: string): Promise<string>;
  fetchImages(aircraft: AircraftType): Promise<string[]>;
}

// Creates an Axios fetcher
export const createAxiosFetcher = (parser: Parser): Fetcher => {
  // Fetches HTML content from a URL
  const fetchHtmlFromUrl = async (url: string): Promise<string> => {
    const response = await axios.get(url);
    return response.data;
  };

  // Fetches image data from a URL
  const fetchImageFromUrl = async (url: string): Promise<Buffer> => {
    const response = await axios.get(url, { responseType: 'arraybuffer' });
    return response.data;
  };

  // Fetches the original high-resolution image URL from a Wikipedia file page URL
  const fetchOriginalImageUrl = async (
    filePageUrl: string
  ): Promise<string> => {
    try {
      const htmlContent = await fetchHtmlFromUrl(filePageUrl);
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
  const fetchImagesFromWikipedia = async (
    aircraft: AircraftType
  ): Promise<string[]> => {
    const searchTerms = `${aircraft.make} ${aircraft.model}`;
    const searchUrl = `https://en.wikipedia.org/w/index.php?search=${encodeURIComponent(
      searchTerms
    )}`;

    try {
      const htmlContent = await fetchHtmlFromUrl(searchUrl);
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
      const topResultHtml = await fetchHtmlFromUrl(topResultUrl);
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

  // Function to fetch images for a given aircraft. It checks the aircraftTypeMap for URLs,
  // and falls back to fetching from Wikipedia if no mapping is found.
  const fetchImages = async (aircraft: AircraftType): Promise<string[]> => {
    // Load the JSON data which maps aircraft types to Wikipedia URLs
    const aircraftTypeMap = loadAircraftTypeMap(config.typeMapPath);

    const searchTermsForMap = `${aircraft.make} ${aircraft.model}`;
    console.log(`Attempting to fetch images for ${searchTermsForMap}`);

    try {
      let highResImageUrl = '';

      if (aircraftTypeMap.has(searchTermsForMap)) {
        const mapping = aircraftTypeMap.get(searchTermsForMap);
        if (mapping) {
          highResImageUrl = mapping.url;
        }
      }

      if (!highResImageUrl) {
        // Fallback to Wikipedia search if no URLs found
        console.log(
          `No direct mapping found for ${searchTermsForMap}. Searching Wikipedia...`
        );
        const wikipediaImages = await fetchImagesFromWikipedia(aircraft);
        if (wikipediaImages.length > 0) {
          let filePageUrl = wikipediaImages[0];
          if (!filePageUrl.startsWith('https://')) {
            filePageUrl = `https://en.wikipedia.org${filePageUrl}`;
          }
          console.log(`Using image from Wikipedia for ${searchTermsForMap}`);
          highResImageUrl = await fetchOriginalImageUrl(filePageUrl);
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

  return {
    fetchHtmlFromUrl,
    fetchImageFromUrl,
    fetchOriginalImageUrl,
    fetchImages,
  };
};
