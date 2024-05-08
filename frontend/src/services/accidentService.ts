import axios from 'axios';
import { Accident, Aircraft, Injury } from '@/types/aviationTypes';

const getBaseUrl = (): string =>
  process.env.NEXT_PUBLIC_ENV === 'development'
    ? 'http://localhost:8080'
    : 'https://airaccidentdata.com';

export const fetchAccidents = async (
  page: number
): Promise<{ accidents: Accident[]; total: number }> => {
  const apiUrl = `${getBaseUrl()}/api/v1/accidents?page=${Math.max(
    1,
    Math.floor(page)
  )}`;
  const response = await axios.get<{ accidents: Accident[]; total: number }>(
    apiUrl
  );
  return response.data;
};

export const fetchAccidentDetails = async (
  accidentId: number
): Promise<Accident> => {
  const apiUrl = `${getBaseUrl()}/api/v1/accidents/${accidentId}`;
  const response = await axios.get<Accident>(apiUrl);
  return response.data;
};

export const fetchAircraftDetails = async (
  aircraftId: number
): Promise<Aircraft> => {
  const apiUrl = `${getBaseUrl()}/api/v1/aircrafts/${aircraftId}`;
  const response = await axios.get<Aircraft>(apiUrl);
  return response.data;
};

export const fetchAircraftImages = async (
  aircraftId: number
): Promise<string[]> => {
  const apiUrl = `${getBaseUrl()}/api/v1/aircrafts/${aircraftId}/images`;
  try {
    const response = await axios.get<{ images: { s3_url: string }[] | null }>(
      apiUrl
    );
    // Check if the response has the images array and if it's not null
    if (response.data.images && Array.isArray(response.data.images)) {
      return response.data.images.map((img) => img.s3_url);
    } else {
      // Handle case where images are null or not an array by returning an array with the default image
      return [
        'https://upload.wikimedia.org/wikipedia/commons/e/e2/BK-117_Polizei-NRW_D-HNWL.jpg',
      ];
    }
  } catch (error) {
    console.error('Error fetching aircraft images:', error);
    return [
      'https://upload.wikimedia.org/wikipedia/commons/e/e2/BK-117_Polizei-NRW_D-HNWL.jpg',
    ]; // Return an array with the default image if there's an error
  }
};

export const fetchAccidentInjuries = async (
  accidentId: number
): Promise<Injury[]> => {
  const apiUrl = `${getBaseUrl()}/api/v1/accidents/${accidentId}/injuries`;
  const response = await axios.get<{ injuries: Injury[] }>(apiUrl);
  return response.data.injuries;
};

export const fetchLocation = async (accidentId: number): Promise<Location> => {
  const apiUrl = `${getBaseUrl()}/api/v1/accidents/${accidentId}/location`;
  const response = await axios.get<Location>(apiUrl);
  return response.data;
};
