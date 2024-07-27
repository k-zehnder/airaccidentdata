import { useState, useEffect } from 'react';
import axios from 'axios';
import { Accident } from '@/types/aviationTypes';

export const useAccidentData = (currentPage: number, searchQuery: string) => {
  const [accidents, setAccidents] = useState<Accident[]>([]);
  const [totalPages, setTotalPages] = useState<number>(0);
  const [isLoading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    const fetchAccidents = async () => {
      setLoading(true);
      try {
        const response = await axios.post('/api/search', {
          query: searchQuery,
          size: 10,
          page: currentPage,
        });
        setAccidents(response.data.results);
        setTotalPages(Math.ceil(response.data.total / 10));
      } catch (error) {
        console.error('Error fetching accidents:', error);
      }
      setLoading(false);
    };

    fetchAccidents();
  }, [currentPage, searchQuery]);

  return { accidents, totalPages, isLoading };
};
