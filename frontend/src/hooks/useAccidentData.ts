import { useEffect, useState } from 'react';
import axios from 'axios';
import { Aircraft, Accident } from '@/types/accident';

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

        // Fetch aircraft details for each accident
        const accidentsWithAircraftDetails = await Promise.all(
          response.data.accidents.map(async (accident: any) => {
            try {
              const aircraftApiUrl = `${
                process.env.NEXT_PUBLIC_ENV === 'development'
                  ? 'http://localhost:8080'
                  : 'https://airaccidentdata.com'
              }/api/v1/aircrafts/${accident.aircraft_id}`;
              const aircraftResponse =
                await axios.get<Aircraft>(aircraftApiUrl);

              // Fetching the stored S3 URL for the aircraft image
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

              return {
                ...accident,
                aircraftDetails: aircraftResponse.data,
                imageUrl: aircraftImageUrl,
              };
            } catch (error) {
              console.error(
                `Error fetching accident details for ID ${accident.id}:`,
                error,
              );
              return null;
            }
          }),
        );

        // Filter out null values (if any) and cast to Accident[]
        const filteredAccidents = accidentsWithAircraftDetails.filter(
          (accident): accident is Accident => accident !== null,
        ) as Accident[];

        setAccidents(filteredAccidents);

        // Calculate total pages based on the total number of accidents and accidents per page
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

        const apiUrl = `${baseUrl}/api/v1/accidents/${accidentId}`;
        const response = await axios.get<Accident>(apiUrl);
        const accidentData = response.data;

        // Fetch aircraft details using aircraft ID
        const aircraftApiUrl = `${baseUrl}/api/v1/aircrafts/${accidentData.aircraft_id}`;
        const aircraftResponse = await axios.get<Aircraft>(aircraftApiUrl);
        const aircraftData = aircraftResponse.data;

        // Fetching the stored S3 URL for the aircraft image
        const imageUrl = `${baseUrl}/api/v1/aircrafts/${accidentData.aircraft_id}/images`;
        const imageResponse = await axios.get<{ images: { s3_url: string }[] }>(
          imageUrl,
        );
        const images = imageResponse.data.images;
        const aircraftImageUrl = images.length > 0 ? images[0].s3_url : '';

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
