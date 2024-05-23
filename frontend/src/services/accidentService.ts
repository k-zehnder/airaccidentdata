import { AxiosInstance } from 'axios';
import { Accident, Aircraft, Injury } from '@/types/aviationTypes';

// Get the base URL based on the environment
const getBaseUrl = (env: string): string =>
  env === 'development'
    ? 'http://localhost:8080'
    : 'https://airaccidentdata.com';

// Factory function to create an accident service
export const createAccidentService = (
  httpClient: AxiosInstance,
  env: string
) => {
  // Determine the base URL based on the environment
  const baseUrl = getBaseUrl(env);

  // Fetches accident data from the API
  const fetchAccidents = async (
    page: number
  ): Promise<{ accidents: Accident[]; total: number }> => {
    const apiUrl = `${baseUrl}/api/v1/accidents?page=${Math.max(
      1,
      Math.floor(page)
    )}`;
    const response = await httpClient.get<{
      accidents: Accident[];
      total: number;
    }>(apiUrl);
    return response.data;
  };

  // Fetches detailed information about a specific accident
  const fetchAccidentDetails = async (
    accidentId: number
  ): Promise<Accident> => {
    const apiUrl = `${baseUrl}/api/v1/accidents/${accidentId}`;
    const response = await httpClient.get<Accident>(apiUrl);
    return response.data;
  };

  // Fetches details about a specific aircraft
  const fetchAircraftDetails = async (
    aircraftId: number
  ): Promise<Aircraft> => {
    const apiUrl = `${baseUrl}/api/v1/aircrafts/${aircraftId}`;
    const response = await httpClient.get<Aircraft>(apiUrl);
    return response.data;
  };

  // Fetches image URL from the database
  const fetchImageUrlFromDB = async (
    aircraftId: number
  ): Promise<string | null> => {
    const apiUrl = `${baseUrl}/api/v1/aircrafts/${aircraftId}/images`;
    try {
      const response = await httpClient.get<{
        images: {
          id: number;
          aircraft_id: number;
          image_url: string;
          s3_url: string;
        }[];
      }>(apiUrl);
      if (response.data.images && response.data.images.length > 0) {
        return response.data.images[0].image_url;
      } else {
        return null;
      }
    } catch (error) {
      console.error('Error fetching aircraft image URL from DB:', error);
      return null;
    }
  };

  // Fetches images from S3
  const fetchImagesFromS3 = async (aircraftId: number): Promise<string[]> => {
    const apiUrl = `${baseUrl}/api/v1/aircrafts/${aircraftId}/images`;
    try {
      const response = await httpClient.get<{
        images: { s3_url: string }[] | null;
      }>(apiUrl);
      if (response.data.images && Array.isArray(response.data.images)) {
        return response.data.images.map((img) => img.s3_url);
      } else {
        return [];
      }
    } catch (error) {
      console.error('Error fetching aircraft images:', error);
      return [];
    }
  };

  // Fetches images of a specific aircraft. In development mode, it skips S3 fetching
  const fetchAircraftImages = async (aircraftId: number): Promise<string[]> => {
    if (env === 'development') {
      const imageUrl = await fetchImageUrlFromDB(aircraftId);
      if (imageUrl) {
        return [imageUrl];
      } else {
        return [
          'https://upload.wikimedia.org/wikipedia/commons/e/e2/BK-117_Polizei-NRW_D-HNWL.jpg',
        ];
      }
    }

    const s3Images = await fetchImagesFromS3(aircraftId);
    if (s3Images.length > 0) {
      return s3Images;
    } else {
      return [
        'https://upload.wikimedia.org/wikipedia/commons/e/e2/BK-117_Polizei-NRW_D-HNWL.jpg',
      ];
    }
  };

  // Fetches injury data related to a specific accident
  const fetchAccidentInjuries = async (
    accidentId: number
  ): Promise<Injury[]> => {
    const apiUrl = `${baseUrl}/api/v1/accidents/${accidentId}/injuries`;
    const response = await httpClient.get<{ injuries: Injury[] }>(apiUrl);
    return response.data.injuries;
  };

  // Fetches the location data for a specific accident
  const fetchLocation = async (accidentId: number): Promise<Location> => {
    const apiUrl = `${baseUrl}/api/v1/accidents/${accidentId}/location`;
    const response = await httpClient.get<Location>(apiUrl);
    return response.data;
  };

  return {
    fetchAccidents,
    fetchAccidentDetails,
    fetchAircraftDetails,
    fetchAircraftImages,
    fetchAccidentInjuries,
    fetchLocation,
  };
};
