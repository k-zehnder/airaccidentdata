import axios from 'axios';
import { useEffect, useState } from 'react';
import { Accident } from '@/types/aviationTypes';
import { createAccidentService } from '../services/accidentService';
import config from '../config/config';

// Default instance of the accident service
const defaultAccidentService = createAccidentService(axios, config.nodeEnv);

export const useFetchAccidentDetails = (
  accidentId: number,
  accidentService = defaultAccidentService
) => {
  const [accidentDetails, setAccidentDetails] = useState<Accident | null>(null);
  const [isLoading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    const fetchDetails = async () => {
      setLoading(true);
      try {
        const accident: Accident = await accidentService.fetchAccidentDetails(
          accidentId
        );
        const [aircraftDetails, images, location] = await Promise.all([
          accidentService.fetchAircraftDetails(accident.aircraft_id),
          accidentService.fetchAircraftImages(accident.aircraft_id),
          accidentService.fetchLocation(accidentId),
        ]);

        const imageUrl = images[0] || '';

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
  }, [accidentId, accidentService]);

  return { accidentDetails, isLoading };
};
