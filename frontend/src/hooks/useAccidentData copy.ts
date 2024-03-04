import { useEffect, useState } from 'react';
import axios from 'axios';
import { Aircraft, Accident } from '@/types/accident';

export const useAccidentData = (currentPage: number) => {
  const [accidents, setAccidents] = useState<Accident[]>([]);
  const [totalPages, setTotalPages] = useState(0);
  const [isFetching, setFetching] = useState(false);

  useEffect(() => {
    const fetchAccidents = async () => {
      setFetching(true);
      try {
        const apiUrl =
          process.env.NEXT_PUBLIC_ENV === 'development'
            ? `http://localhost:8080/api/v1/accidents?page=${currentPage}`
            : `https://airaccidentdata.com/api/v1/accidents?page=${currentPage}`;
        const response = await axios.get<{
          accidents: Accident[];
          total: number;
        }>(apiUrl);

        // Fetch aircraft details for each accident
        const accidentsWithAircraftDetails = await Promise.all(
          response.data.accidents.map(async (accident) => {
            const aircraftApiUrl = `https://airaccidentdata.com/api/v1/aircrafts/${accident.aircraft_id}`;
            const aircraftResponse = await axios.get<Aircraft>(aircraftApiUrl);
            return {
              ...accident,
              aircraftDetails: aircraftResponse.data,
            };
          })
        );

        setAccidents(accidentsWithAircraftDetails);
        setTotalPages(Math.ceil(response.data.total / 10)); // 10 accidents per page
      } catch (error) {
        console.error('Error fetching accidents:', error);
      }
      setFetching(false);
    };

    fetchAccidents();
  }, [currentPage]);

  return { accidents, totalPages, isFetching };
};
