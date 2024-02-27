import axios, { AxiosInstance } from 'axios';

// Define the structure of the accident details
interface AccidentDetails {
  date: string;
  aircraftModel: string;
  location: string;
  summary: string;
  recommendations: string[];
}

// Define the structure of the Axios client
interface APIClient extends AxiosInstance {}

export const useClient = (): APIClient => {
  return axios.create({
    baseURL: 'http://localhost:8080/api/v1',
    headers: {
      'Content-Type': 'application/json',
    },
  });
};
