import axios from 'axios';
import { useEffect, useState } from 'react';
import { Accident } from '@/types/aviationTypes';
import { createAccidentService } from '../services/accidentService';
import config from '../config/config';

// Default instance of the accident service
const defaultAccidentService = createAccidentService(axios, config.nodeEnv);

export const useFetchAccidentData = (
  currentPage: number,
  accidentService = defaultAccidentService
) => {
  const [accidents, setAccidents] = useState<Accident[]>([]);
  const [totalPages, setTotalPages] = useState<number>(0);
  const [isFetching, setFetching] = useState<boolean>(false);
  const accidentsPerPage = 10;

  useEffect(() => {
    const fetchAccidentsData = async () => {
      setFetching(true);
      try {
        const { accidents, total } = await accidentService.fetchAccidents(
          currentPage
        );
        const accidentsWithDetails = await Promise.all(
          accidents.map(async (accident) => {
            try {
              const aircraftDetails =
                await accidentService.fetchAircraftDetails(
                  accident.aircraft_id
                );
              const images = await accidentService.fetchAircraftImages(
                accident.aircraft_id
              );
              const injuries = await accidentService.fetchAccidentInjuries(
                accident.id
              );
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
  }, [currentPage, accidentService]);

  return { accidents, totalPages, isFetching };
};
