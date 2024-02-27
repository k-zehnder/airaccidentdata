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
  const baseURL = process.env.NODE_ENV === 'development' ? 'http://localhost:8080/api/v1': 'https://airaccidentdata.com/api/v1'
  
  return axios.create({
    baseURL: baseURL,
    headers: {
      'Content-Type': 'application/json',
    },
  });
};
