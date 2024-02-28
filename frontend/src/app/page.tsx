"use client";

import React, { useState } from 'react';
import Link from 'next/link';
import { Header } from '@/components/header';
import Pagination from '@/components/pagination';
import { Badge } from '@/components/ui/badge';
import { useAccidentData } from '../hooks/useAccidentData';
import { formatDate } from '../lib/utils';

const Home = () => {
  const [currentPage, setCurrentPage] = useState(1);
  const { accidents, totalPages, isFetching } = useAccidentData(currentPage);

  if (isFetching) return <div>Loading...</div>;

  return (
    <>
      <Header />
      <main className="container mx-auto px-4 py-6">
        <div className="max-w-4xl mx-auto mb-8 text-center">
          <a
            href="https://airaccidentdata.com/swagger/index.html"
            target="_blank"
            rel="noopener noreferrer"
            className="inline-flex items-center rounded-lg bg-muted px-3 py-1 text-sm font-medium justify-center"
          >
            🔗
            <div className="ml-2 w-1 h-4 bg-border" data-orientation="vertical" role="none"></div> 
            <span className="sm:hidden">Check out the API Documentation</span>
            <span className="hidden sm:inline">Check out the API Documentation</span>
            <svg
              width="15"
              height="15"
              viewBox="0 0 15 15"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
              className="ml-1 h-4 w-4"
            >
              <path
                d="M8.14645 3.14645C8.34171 2.95118 8.65829 2.95118 8.85355 3.14645L12.8536 7.14645C13.0488 7.34171 13.0488 7.65829 12.8536 7.85355L8.85355 11.8536C8.65829 12.0488 8.34171 12.0488 8.14645 11.8536C7.95118 11.6583 7.95118 11.3417 8.14645 11.1464L11.2929 8H2.5C2.22386 8 2 7.77614 2 7.5C2 7.22386 2.22386 7 2.5 7H11.2929L8.14645 3.85355C7.95118 3.65829 7.95118 3.34171 8.14645 3.14645Z"
                fill="currentColor"
                fillRule="evenodd"
                clipRule="evenodd"
              ></path>
            </svg>
          </a>
        </div>
        <div className="max-w-4xl mx-auto mb-8">
          <h1 className="text-4xl lg:text-5xl font-bold mb-4 text-center">
            Explore Aviation Accidents and Insights
          </h1>
          <p className="text-lg text-muted-foreground mb-8 text-center">
            Your gateway to understanding air travel incidents and promoting a safer flying
            future.
          </p>
        </div>
        <div className="max-w-4xl mx-auto">
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
                  <h3 className="text-xl font-semibold mb-1">
                    {accident.registrationNumber}: {accident.aircraftMakeName}{' '}
                    {accident.aircraftModelName}
                  </h3>
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
