import { AxiosInstance } from 'axios';
import { Accident, Aircraft, Injury } from '@/types/aviationTypes';

// Factory function to create an accident service
export const createAccidentService = (
  httpClient: AxiosInstance,
  env: string
) => {
  // Get the base URL based on the environment
  const getBaseUrl = (env: string): string =>
    env === 'development'
      ? 'http://localhost:8080'
      : 'https://airaccidentdata.com';

  const baseUrl = getBaseUrl(env);

  // Fetch accidents with pagination
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

  // Fetch details of a specific aircraft
  const fetchAircraftDetails = async (
    aircraftId: number
  ): Promise<Aircraft> => {
    const apiUrl = `${baseUrl}/api/v1/aircrafts/${aircraftId}`;
    const response = await httpClient.get<Aircraft>(apiUrl);
    return response.data;
  };

  // Fetch images of a specific aircraft
  const fetchAircraftImages = async (aircraftId: number): Promise<string[]> => {
    const apiUrl = `${baseUrl}/api/v1/aircrafts/${aircraftId}/images`;
    const response = await httpClient.get<{
      images: { s3_url?: string; image_url?: string }[] | null;
    }>(apiUrl);

    // Check for s3_url or use a local_url as a fallback
    return (
      response.data.images?.map(
        (img) =>
          img.s3_url ||
          img.image_url ||
          'https://upload.wikimedia.org/wikipedia/commons/e/e2/BK-117_Polizei-NRW_D-HNWL.jpg'
      ) || []
    );
  };

  // Fetch injury data for a specific accident
  const fetchAccidentInjuries = async (
    accidentId: number
  ): Promise<Injury[]> => {
    const apiUrl = `${baseUrl}/api/v1/accidents/${accidentId}/injuries`;
    const response = await httpClient.get<{ injuries: Injury[] }>(apiUrl);
    return response.data.injuries;
  };

  // Fetch details of a specific accident
  const fetchAccidentDetails = async (
    accidentId: number
  ): Promise<Accident> => {
    const apiUrl = `${baseUrl}/api/v1/accidents/${accidentId}`;
    const response = await httpClient.get<Accident>(apiUrl);
    return response.data;
  };

  // Fetch location data for a specific accident
  const fetchLocation = async (accidentId: number): Promise<any> => {
    const apiUrl = `${baseUrl}/api/v1/accidents/${accidentId}/location`;
    const response = await httpClient.get<any>(apiUrl);
    return response.data;
  };

  return {
    fetchAccidents,
    fetchAircraftDetails,
    fetchAircraftImages,
    fetchAccidentInjuries,
    fetchAccidentDetails,
    fetchLocation,
  };
};
