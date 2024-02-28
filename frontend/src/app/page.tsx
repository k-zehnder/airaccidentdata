"use client";
// Home component
import React, { useState } from 'react';
import Link from 'next/link';
import { Header } from '@/components/header';
import Pagination from '@/components/pagination';
import { Badge } from '@/components/ui/badge';
import { useAccidentData } from '../hooks/useAccidentData';

const Home = () => {
  const [currentPage, setCurrentPage] = useState(1);
  const { accidents, totalPages, isFetching } = useAccidentData(currentPage);

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
                    accident.cabinCrewInjurySerious !== 0 ||
                    accident.passengerInjurySerious !== 0 ||
                    accident.groundInjurySerious !== 0) && (
                    <Badge key={accident.id} className="bg-yellow-500 mb-1">
                      Serious Injuries
                    </Badge>
                  )}
                  {(accident.flightCrewInjuryMinor !== 0 ||
                    accident.cabinCrewInjuryMinor !== 0 ||
                    accident.passengerInjuryMinor !== 0 ||
                    accident.groundInjuryMinor !== 0) && (
                    <Badge key={accident.id} className="bg-green-500 mb-1">
                      Minor Injuries
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
          darkMode={true} 
        />
      </main>
    </>
  );
};

export default Home;
