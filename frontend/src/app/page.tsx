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
          {/* API documentation link */}
        </div>
        <div className="max-w-4xl mx-auto mb-8">
          {/* Title and description */}
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
                  {/* Date and aircraft details */}
                  <p className="text-gray-500">
                    {formatDate(accident.entry_date)}
                  </p>
                  <h3 className="text-xl font-semibold mb-1">
                    {accident.aircraftDetails?.registration_number}:{' '}
                    {accident.aircraftDetails?.aircraft_make_name}{' '}
                    {accident.aircraftDetails?.aircraft_model_name}
                  </h3>
                  {/* Original badges */}
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
                  {/* Badge for birds */}
                  {accident.remark_text.toLowerCase().includes('bird') && (
                    <Badge className="bg-blue-500 mb-1">Birds Present</Badge>
                  )}
                  {/* Badge for stall */}
                  {(accident.remark_text.toLowerCase().includes('stall') ||
                    accident.remark_text.toLowerCase().includes('stalled')) && (
                    <Badge className="bg-orange-500 mb-1">Stall</Badge>
                  )}
                  <p className="text-gray-500">{accident.remark_text}</p>
                </div>
                <div className="w-1/4 flex justify-end">
                  {/* Image */}
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
        {/* Pagination component */}
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
