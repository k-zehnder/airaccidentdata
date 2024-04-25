'use client';

import React, { useState } from 'react';
import Link from 'next/link';
import { Header } from '@/components/header';
import Pagination from '@/components/pagination';
import { Badge } from '@/components/ui/badge';
import { useAccidentData } from '../hooks/useAccidentData';
import Loader from '@/components/Loader';
import { formatDate } from '../lib/utils';

const Home = () => {
  const [currentPage, setCurrentPage] = useState(1);
  const { accidents, totalPages, isFetching } = useAccidentData(currentPage);

  if (isFetching) {
    return (
      <>
        <Header />
        <Loader />
      </>
    );
  }

  return (
    <>
      <Header />
      <main className="container mx-auto px-4 py-6">
        <div className="max-w-4xl mx-auto mb-6 text-center">
          <a
            href="https://airaccidentdata.com/swagger/index.html"
            target="_blank"
            rel="noopener noreferrer"
            className="inline-flex items-center rounded-lg bg-muted px-3 py-1 text-sm font-medium justify-center"
          >
            <span className="inline-flex items-center">🔗</span>
            <span className="ml-2">Check out the API Documentation</span>
            <svg
              width="15"
              height="15"
              viewBox="0 0 15 15"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
              className="ml-2 h-4 w-4"
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
            Your gateway to understanding air travel incidents and promoting a
            safer flying future.
          </p>
        </div>
        <div className="max-w-4xl mx-auto">
          {accidents.map((accident) => (
            <Link
              key={accident.id}
              legacyBehavior
              href={`/accidents/${accident.id}`}
            >
              <a className="block border-b-2 py-4 flex items-center  hover:bg-gray-100 dark:hover:bg-gray-900 transition-colors">
                <div className="flex-1">
                  <span className="text-gray-500 text-sm block lg:text-base mb-1">
                    {formatDate(accident.entry_date)}
                  </span>
                  <h3 className="text-xl font-semibold mb-1">
                    {accident.aircraftDetails?.registration_number}:{' '}
                    {accident.aircraftDetails?.aircraft_make_name}{' '}
                    {accident.aircraftDetails?.aircraft_model_name}
                  </h3>
                  {accident.fatal_flag === 'Yes' && (
                    <Badge className="bg-red-500 mb-1">Fatalities</Badge>
                  )}
                  {(accident.flight_crew_injury_serious !== 0 ||
                    accident.cabin_crew_injury_serious !== 0 ||
                    accident.passenger_injury_serious !== 0 ||
                    accident.ground_injury_serious !== 0) && (
                    <Badge className="bg-yellow-500 mb-1">
                      Serious Injuries
                    </Badge>
                  )}
                  {(accident.flight_crew_injury_minor !== 0 ||
                    accident.cabin_crew_injury_minor !== 0 ||
                    accident.passenger_injury_minor !== 0 ||
                    accident.ground_injury_minor !== 0) && (
                    <Badge className="bg-green-500 mb-1">Minor Injuries</Badge>
                  )}
                  {accident.remark_text.toLowerCase().includes('bird') && (
                    <Badge className="bg-blue-500 mb-1">Birds Present</Badge>
                  )}
                  {(accident.remark_text.toLowerCase().includes('stall') ||
                    accident.remark_text.toLowerCase().includes('stalled')) && (
                    <Badge className="bg-orange-500 mb-1">Stall</Badge>
                  )}
                  <p className="text-gray-500">{accident.remark_text}</p>
                </div>
                <div className="w-1/4 flex justify-end">
                  <img
                    src={
                      accident.imageUrl ||
                      'https://upload.wikimedia.org/wikipedia/commons/e/e2/BK-117_Polizei-NRW_D-HNWL.jpg'
                    }
                    alt={`Thumbnail for ${accident.aircraft_id}`}
                    className="w-16 h-16 object-cover mb-2"
                  />
                </div>
              </a>
            </Link>
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
