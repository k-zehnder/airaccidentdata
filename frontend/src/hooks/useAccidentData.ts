import { useEffect, useState } from 'react';
import axios from 'axios';
import { Accident, AccidentDetails } from '@/types/accident';

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
        setAccidents(response.data.accidents);
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

export const useFetchAccidentDetails = (registrationNumber: string) => {
  const [accidentDetails, setAccidentDetails] = useState<AccidentDetails | null>(
    null
  );

  useEffect(() => {
    const fetchAccidentDetails = async () => {
      try {
        if (registrationNumber) {
          const apiUrl =
          process.env.NEXT_PUBLIC_ENV === 'development'
            ? `http://localhost:8080/api/v1/accidents/${registrationNumber}`
            : `https://airaccidentdata.com/api/v1/accidents/${registrationNumber}`;
          const response = await axios.get<AccidentDetails>(apiUrl);
          console.log(response.data);
          setAccidentDetails(response.data);
        }
      } catch (error) {
        console.error('Error fetching accident details:', error);
      }
    };

    fetchAccidentDetails();
  }, [registrationNumber]);

  return accidentDetails;
};
