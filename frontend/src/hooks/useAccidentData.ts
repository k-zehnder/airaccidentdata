import { useEffect, useState } from 'react';
import axios from 'axios';
import { Aircraft, Accident, Injury } from '@/types/accident';

export const useAccidentData = (currentPage: number) => {
  const [accidents, setAccidents] = useState<Accident[]>([]);
  const [totalPages, setTotalPages] = useState(0);
  const [isFetching, setFetching] = useState(false);
  const accidentsPerPage = 10;

  useEffect(() => {
    const fetchAccidents = async () => {
      setFetching(true);
      try {
        const apiUrl = `${
          process.env.NEXT_PUBLIC_ENV === 'development'
            ? 'http://localhost:8080'
            : 'https://airaccidentdata.com'
        }/api/v1/accidents?page=${Math.max(1, Math.floor(currentPage))}`;
        const response = await axios.get<{
          accidents: Accident[];
          total: number;
        }>(apiUrl);

        // Fetch additional details for each accident, including injuries
        const accidentsWithDetails = await Promise.all(
          response.data.accidents.map(async (accident) => {
            try {
              const aircraftApiUrl = `${
                process.env.NEXT_PUBLIC_ENV === 'development'
                  ? 'http://localhost:8080'
                  : 'https://airaccidentdata.com'
              }/api/v1/aircrafts/${accident.aircraft_id}`;
              const aircraftResponse = await axios.get<Aircraft>(
                aircraftApiUrl
              );

              const imageUrl = `${
                process.env.NEXT_PUBLIC_ENV === 'development'
                  ? 'http://localhost:8080'
                  : 'https://airaccidentdata.com'
              }/api/v1/aircrafts/${accident.aircraft_id}/images`;
              const imageResponse = await axios.get<{
                images: { s3_url: string }[];
              }>(imageUrl);
              const images = imageResponse.data.images;
              const aircraftImageUrl =
                images.length > 0 ? images[0].s3_url : '';

              // Fetch injury information
              const injuriesUrl = `${
                process.env.NEXT_PUBLIC_ENV === 'development'
                  ? 'http://localhost:8080'
                  : 'https://airaccidentdata.com'
              }/api/v1/injuries/${accident.id}`;
              const injuriesResponse = await axios.get<{
                injuries: Injury[];
              }>(injuriesUrl);
              const injuries = injuriesResponse.data.injuries;

              return {
                ...accident,
                aircraftDetails: aircraftResponse.data,
                imageUrl: aircraftImageUrl,
                injuries,
              };
            } catch (error) {
              console.error(
                `Error fetching details for accident ID ${accident.id}:`,
                error
              );
              return null;
            }
          })
        );

        // Filtering out null values from accidentsWithDetails array, assuming that any null values represent failed fetches or missing details
        const filteredAccidents = accidentsWithDetails.filter(
          Boolean
        ) as Accident[];

        setAccidents(filteredAccidents);
        setTotalPages(Math.ceil(response.data.total / accidentsPerPage));
      } catch (error) {
        console.error('Error fetching accidents:', error);
      }
      setFetching(false);
    };

    fetchAccidents();
  }, [currentPage]);

  return { accidents, totalPages, isFetching };
};

export const useFetchAccidentDetails = (accidentId: string) => {
  const [accidentDetails, setAccidentDetails] = useState<Accident | null>(null);
  const [isLoading, setLoading] = useState(false);

  useEffect(() => {
    const fetchAccidentDetails = async () => {
      setLoading(true);
      try {
        // Determine the base URL based on environment
        const baseUrl =
          process.env.NEXT_PUBLIC_ENV === 'development'
            ? 'http://localhost:8080'
            : 'https://airaccidentdata.com';

        const [accidentResponse, aircraftResponse, imageResponse] =
          await Promise.all([
            axios.get<Accident>(`${baseUrl}/api/v1/accidents/${accidentId}`),
            axios.get<Aircraft>(`${baseUrl}/api/v1/aircrafts/${accidentId}`),
            axios.get<{ images: { s3_url: string }[] }>(
              `${baseUrl}/api/v1/aircrafts/${accidentId}/images`
            ),
          ]);

        const accidentData = accidentResponse.data;
        const aircraftData = aircraftResponse.data;
        const imageData = imageResponse.data;

        // Fetching the stored S3 URL for the aircraft image
        const aircraftImageUrl =
          imageData.images.length > 0 ? imageData.images[0].s3_url : '';

        // Combine accident, aircraft details, and image URL
        const combinedData = {
          ...accidentData,
          aircraftDetails: aircraftData,
          imageUrl: aircraftImageUrl,
        };

        setAccidentDetails(combinedData);
      } catch (error) {
        console.error('Error fetching accident details:', error);
        setAccidentDetails(null);
      }
      setLoading(false);
    };

    fetchAccidentDetails();
  }, [accidentId]);

  return { accidentDetails, isLoading };
};
