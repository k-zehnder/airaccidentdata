'use client';

import React, { useEffect } from 'react';
import Link from 'next/link';
import { Header } from '@/components/header';
import { buttonVariants } from '@/components/ui/button';
import { useFetchAccidentDetails } from '@/hooks/useAccidentData';
import MapComponent from '@/components/MapComponent';
import Loader from '@/components/Loader';

interface AccidentDetailProps {
  params: {
    accidentId: string;
  };
}

const AccidentDetail: React.FC<AccidentDetailProps> = ({ params }) => {
  const { accidentId } = params;
  const { accidentDetails, isLoading } = useFetchAccidentDetails(accidentId);

  useEffect(() => {
    if (accidentDetails) {
      console.log('Accident details:', accidentDetails);
    }
  }, [accidentDetails]);

  if (isLoading) {
    return (
      <>
        <Header />
        <Loader />
      </>
    );
  }

  // Hardcoded safety recommendations
  const safetyRecommendations: string[] = [
    'Ensure proper maintenance checks are performed regularly.',
    'Implement additional training programs for flight crew members.',
  ];

  return (
    <>
      <Header />
      <section className="container flex flex-col space-y-5 mt-10">
        <div className="text-4xl font-bold">Air Accident Details</div>
        <div className="text-xl mt-0 text-muted-foreground dark:text-slate-200">
          {accidentDetails?.remark_text}
        </div>
        <hr className="my-4"></hr>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div className="md:col-span-1">
            <img
              src={accidentDetails?.imageUrl}
              alt="Aircraft Image"
              className="max-w-full h-auto rounded-lg shadow-md mb-6"
            />
          </div>

          <div className="md:col-span-1 bg-white dark:bg-background/50 shadow-md rounded border p-6 mb-6">
            <h2 className="text-2xl font-semibold mb-4">Incident Overview</h2>
            <p>
              <strong>Date: </strong>
              {accidentDetails?.entry_date}
            </p>
            <p>
              <strong>Aircraft Make:</strong>{' '}
              {accidentDetails?.aircraftDetails?.aircraft_make_name}{' '}
              {accidentDetails?.aircraftDetails?.aircraft_model_name}
            </p>
            <p>
              <strong>Event Type:</strong>{' '}
              {accidentDetails?.event_type_description}
            </p>
            <p>
              <strong>Remark:</strong> {accidentDetails?.remark_text}
            </p>
          </div>
        </div>

        <section className="bg-white dark:bg-background/50 shadow-md rounded border p-6 mb-6">
          <h2 className="text-2xl font-semibold mb-4">Location Information</h2>
          <p>
            {accidentDetails?.location_city_name},{' '}
            {accidentDetails?.location_state_name},{' '}
            {accidentDetails?.location_country_name}
          </p>
          <MapComponent
            latitude={accidentDetails?.latitude}
            longitude={accidentDetails?.longitude}
          />
        </section>

        <section className="bg-white dark:bg-background/50 shadow-md rounded border p-6 mb-6">
          <h2 className="text-2xl font-semibold mb-4">Reports & Analysis</h2>
        </section>

        <section className="bg-white dark:bg-background/50 shadow-md rounded border p-6 mb-6">
          <h2 className="text-2xl font-semibold mb-4">
            Safety Recommendations
          </h2>
          <ul>
            {safetyRecommendations.map((recommendation, index) => (
              <li key={index}>{recommendation}</li>
            ))}
          </ul>
        </section>

        <Link legacyBehavior href="/" passHref>
          <a className={buttonVariants({ variant: 'outline' })}>Home</a>
        </Link>
        <div className="mb-2"></div>
      </section>
    </>
  );
};

export default AccidentDetail;
