'use client';

import React, { createContext, useContext, useState, ReactNode } from 'react';
import { Accident } from '@/types/aviationTypes';
import { createAccidentService } from '../services/accidentService';
import axios from 'axios';
import config from '@/config/config';

const AccidentService = createAccidentService(axios, config.nodeEnv);

interface AccidentContextType {
  accidents: Accident[];
  totalPages: number;
  isLoading: boolean;
  fetchAccidents: (page: number, searchQuery: string) => void;
  fetchAccidentDetails: (accidentId: number) => Promise<Accident | null>;
}

const AccidentContext = createContext<AccidentContextType | undefined>(
  undefined
);

export const AccidentProvider = ({ children }: { children: ReactNode }) => {
  const [accidents, setAccidents] = useState<Accident[]>([]);
  const [totalPages, setTotalPages] = useState<number>(0);
  const [isLoading, setLoading] = useState<boolean>(false);

  const fetchAccidents = async (currentPage: number, searchQuery: string) => {
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

  const fetchAccidentDetails = async (
    accidentId: number
  ): Promise<Accident | null> => {
    try {
      // Fetch the basic accident details
      const accident: Accident = await AccidentService.fetchAccidentDetails(
        accidentId
      );

      // Fetch related details in parallel
      const [aircraftDetails, images, location] = await Promise.all([
        AccidentService.fetchAircraftDetails(accident.aircraft_id),
        AccidentService.fetchAircraftImages(accident.aircraft_id),
        AccidentService.fetchLocation(accidentId),
      ]);

      // Return the accident with all related details merged
      return {
        ...accident,
        aircraftDetails,
        imageUrl: images[0],
        location,
      };
    } catch (error) {
      console.error('Error fetching accident details:', error);
      return null;
    }
  };

  return (
    <AccidentContext.Provider
      value={{
        accidents,
        totalPages,
        isLoading,
        fetchAccidents,
        fetchAccidentDetails,
      }}
    >
      {children}
    </AccidentContext.Provider>
  );
};

export const useAccidentContext = (): AccidentContextType => {
  const context = useContext(AccidentContext);
  if (!context) {
    throw new Error(
      'useAccidentContext must be used within an AccidentProvider'
    );
  }
  return context;
};
