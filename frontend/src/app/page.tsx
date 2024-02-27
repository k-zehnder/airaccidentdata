"use client";

import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Link from 'next/link';
import { Header } from '@/components/header';
import Pagination from '@/components/pagination';
import { Badge } from '@/components/ui/badge';

interface Accident {
  id: number;
  registrationNumber: string;
  remarkText: string;
  aircraftMakeName: string;
  aircraftModelName: string;
  entryDate: string;
  fatalFlag: string;
  flightCrewInjuryNone: number;
  flightCrewInjuryMinor: number;
  flightCrewInjurySerious: number;
  flightCrewInjuryFatal: number;
  flightCrewInjuryUnknown: number;
  cabinCrewInjuryNone: number;
  cabinCrewInjuryMinor: number;
  cabinCrewInjurySerious: number;
  cabinCrewInjuryFatal: number;
  cabinCrewInjuryUnknown: number;
  passengerInjuryNone: number;
  passengerInjuryMinor: number;
  passengerInjurySerious: number;
  passengerInjuryFatal: number;
  passengerInjuryUnknown: number;
  groundInjuryNone: number;
  groundInjuryMinor: number;
  groundInjurySerious: number;
  groundInjuryFatal: number;
  groundInjuryUnknown: number;
}

const Home = () => {
  const [accidents, setAccidents] = useState<Accident[]>([]);
  const [totalPages, setTotalPages] = useState(0);
  const [currentPage, setCurrentPage] = useState(1);
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

  const formatDate = (dateString: string) => {
    const options: Intl.DateTimeFormatOptions = {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
    };
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', options);
  };

  if (isFetching) return <div>Loading...</div>;

  return (
    <>
      <Header />
      <main className="container mx-auto px-4 py-6">
        <h1 className="text-4xl font-bold tracking-tight mb-6">
          Explore Aviation Accidents and Insights
        </h1>
        <div>
          {accidents.map((accident) => (
            <div key={accident.id} className="border-b-2 py-4">
              <span className="text-gray-500 text-sm block lg:text-base mb-1">
                {formatDate(accident.entryDate)}
              </span>
              <Link
                legacyBehavior
                href={`/accidents/${accident.registrationNumber}`}
              >
                <a>
                  <h2 className="text-2xl font-semibold">
                    {accident.registrationNumber}: {accident.aircraftMakeName}{' '}
                    {accident.aircraftModelName}
                  </h2>
                  {accident.fatalFlag === 'Yes' && (
                    <Badge key={accident.id} className="bg-red-500 mb-1">
                      Fatalities
                    </Badge>
                  )}
                  {(accident.flightCrewInjurySerious !== 0 ||
                    accident.flightCrewInjuryFatal !== 0 ||
                    accident.flightCrewInjuryUnknown !== 0 ||
                    accident.cabinCrewInjurySerious !== 0 ||
                    accident.cabinCrewInjuryFatal !== 0 ||
                    accident.passengerInjurySerious !== 0 ||
                    accident.passengerInjuryFatal !== 0 ||
                    accident.groundInjurySerious !== 0 ||
                    accident.groundInjuryFatal !== 0 ||
                    accident.groundInjuryUnknown !== 0) && (
                    <Badge key={accident.id} className="bg-yellow-500 mb-1">
                      Injuries
                    </Badge>
                  )}
                  <p className="text-gray-500">{accident.remarkText}</p>
                </a>
              </Link>
            </div>
          ))}
        </div>
        <Pagination
          currentPage={currentPage}
          totalPages={totalPages}
          onPageChange={setCurrentPage}
        />
      </main>
    </>
  );
};

export default Home;
