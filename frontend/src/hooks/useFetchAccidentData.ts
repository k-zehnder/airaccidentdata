import { useEffect, useState } from 'react';
import { Accident } from '@/types/aviationTypes';

interface AccidentService {
  fetchAccidents: (
    page: number
  ) => Promise<{ accidents: Accident[]; total: number }>;
  fetchAircraftDetails: (aircraftId: number) => Promise<any>;
  fetchAircraftImages: (aircraftId: number) => Promise<any>;
  fetchAccidentInjuries: (accidentId: number) => Promise<any>;
}

export const useFetchAccidentData = (
  currentPage: number,
  accidentService: AccidentService
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
        const detailedAccidents = await getDetailedAccidents(accidents);
        const filteredAccidents = filterValidAccidents(detailedAccidents);
        setAccidents(filteredAccidents);
        setTotalPages(Math.ceil(total / accidentsPerPage));
      } catch (error) {
        console.error('Error fetching accidents:', error);
      }
      setFetching(false);
    };

    fetchAccidentsData();
  }, [currentPage, accidentService]);

  const getDetailedAccidents = async (accidents: Accident[]) => {
    return await Promise.all(
      accidents.map(async (accident) => {
        try {
          const [aircraftDetails, images, injuries] = await Promise.all([
            accidentService.fetchAircraftDetails(accident.aircraft_id),
            accidentService.fetchAircraftImages(accident.aircraft_id),
            accidentService.fetchAccidentInjuries(accident.id),
          ]);
          return {
            ...accident,
            aircraftDetails,
            imageUrl:
              images[0] ||
              'https://upload.wikimedia.org/wikipedia/commons/e/e2/BK-117_Polizei-NRW_D-HNWL.jpg',
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
  };

  const filterValidAccidents = (accidents: (Accident | null)[]): Accident[] => {
    return accidents.filter(
      (accident): accident is Accident =>
        accident !== null && accident.aircraftDetails !== null
    );
  };

  return { accidents, totalPages, isFetching };
};
