import { useEffect, useState } from 'react';
import axios from 'axios';
import { Aircraft, Accident, Injury, Location } from '@/types/aviationTypes';

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
              const images = imageResponse.data?.images;
              const aircraftImageUrl =
                images && images.length > 0 ? images[0].s3_url : '';

              const injuriesUrl = `${
                process.env.NEXT_PUBLIC_ENV === 'development'
                  ? 'http://localhost:8080'
                  : 'https://airaccidentdata.com'
              }/api/v1/accidents/${accident.id}/injuries`;
              const injuriesResponse = await axios.get<{ injuries: Injury[] }>(
                injuriesUrl
              );
              const injuries = injuriesResponse.data?.injuries;

              return {
                ...accident,
                aircraftDetails: aircraftResponse.data,
                imageUrl: aircraftImageUrl,
                injuries: injuries || [],
              };
            } catch (error) {
              console.error(
                `Error fetching details for accident ID ${accident.id}:`,
                error
              );
              return null; // Returning null if any error occurs during the fetch for an accident
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
      const baseUrl =
        process.env.NEXT_PUBLIC_ENV === 'development'
          ? 'http://localhost:8080'
          : 'https://airaccidentdata.com';

      try {
        // Fetch accident details
        const accidentResponse = await axios.get<Accident>(
          `${baseUrl}/api/v1/accidents/${accidentId}`
        );
        const aircraftId = accidentResponse.data.aircraft_id;

        // Fetch additional data in parallel: aircraft details, images, and location
        const [aircraftResponse, imageResponse, locationResponse] =
          await Promise.all([
            axios.get<Aircraft>(`${baseUrl}/api/v1/aircrafts/${aircraftId}`),
            axios.get<{ images: { s3_url: string }[] }>(
              `${baseUrl}/api/v1/aircrafts/${aircraftId}/images`
            ),
            axios.get<Location>(
              `${baseUrl}/api/v1/accidents/${accidentId}/location`
            ),
          ]);

        // Extract image URL if images are available
        const imageUrl =
          imageResponse.data.images && imageResponse.data.images.length > 0
            ? imageResponse.data.images[0].s3_url
            : '';

        // Combine all fetched data into a single object
        const combinedData = {
          ...accidentResponse.data,
          aircraftDetails: aircraftResponse.data,
          imageUrl: imageUrl,
          location: locationResponse.data,
        };

        setAccidentDetails(combinedData);
      } catch (error) {
        console.error('Error fetching accident details:', error);
        setAccidentDetails(null);
      } finally {
        setLoading(false);
      }
    };

    fetchAccidentDetails();
  }, [accidentId]);

  return { accidentDetails, isLoading };
};
