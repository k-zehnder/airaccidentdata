import { useEffect, useState } from 'react';
import { Accident } from '@/types/aviationTypes';
import {
  fetchAccidents,
  fetchAccidentDetails,
  fetchAircraftDetails,
  fetchAircraftImages,
  fetchAccidentInjuries,
  fetchLocation,
} from '../services/accidentService';

export const useAccidentData = (currentPage: number) => {
  const [accidents, setAccidents] = useState<Accident[]>([]);
  const [totalPages, setTotalPages] = useState<number>(0);
  const [isFetching, setFetching] = useState<boolean>(false);
  const accidentsPerPage = 10;

  useEffect(() => {
    const fetchAccidentsData = async () => {
      setFetching(true);
      try {
        const { accidents, total } = await fetchAccidents(currentPage);
        const accidentsWithDetails = await Promise.all(
          accidents.map(async (accident) => {
            try {
              const aircraftDetails = await fetchAircraftDetails(
                accident.aircraft_id
              );
              const images = await fetchAircraftImages(accident.aircraft_id);
              const injuries = await fetchAccidentInjuries(accident.id);
              return {
                ...accident,
                aircraftDetails,
                imageUrl: images[0] || '',
                injuries,
              };
            } catch (error) {
              console.error(
                `Error fetching details for accident ID ${accident.id}:`,
                error
              );
              return null; // Important to return null in case of error
            }
          })
        );

        // Clean up the array by removing any null or falsy values
        const filteredAccidents = accidentsWithDetails.filter(
          (accident) => accident !== null && accident.aircraftDetails !== null
        ) as Accident[];
        setAccidents(filteredAccidents);
        setTotalPages(Math.ceil(total / accidentsPerPage));
      } catch (error) {
        console.error('Error fetching accidents:', error);
      }
      setFetching(false);
    };

    fetchAccidentsData();
  }, [currentPage]);

  return { accidents, totalPages, isFetching };
};

export const useFetchAccidentDetails = (accidentId: number) => {
  const [accidentDetails, setAccidentDetails] = useState<Accident | null>(null);
  const [isLoading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    const fetchDetails = async () => {
      setLoading(true);
      try {
        const accident: Accident = await fetchAccidentDetails(accidentId);
        const [aircraftDetails, images, location] = await Promise.all([
          fetchAircraftDetails(accident.aircraft_id),
          fetchAircraftImages(accident.aircraft_id),
          fetchLocation(accidentId),
        ]);

        const imageUrl = images[0] || '';

        // TODO: Correct this type
        setAccidentDetails((prev: any) => ({
          ...prev,
          ...accident,
          aircraftDetails,
          imageUrl,
          location,
        }));
      } catch (error) {
        console.error('Error fetching accident details:', error);
        setAccidentDetails(null);
      } finally {
        setLoading(false);
      }
    };

    fetchDetails();
  }, [accidentId]);

  return { accidentDetails, isLoading };
};
