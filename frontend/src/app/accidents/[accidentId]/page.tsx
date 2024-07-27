'use client';

import React from 'react';
import Link from 'next/link';
import { buttonVariants } from '@/components/ui/button';
import { useFetchAccidentDetails } from '@/hooks/useFetchAccidentDetails';
import MapComponent from '@/components/MapComponent';
import Loader from '@/components/Loader';
import { Header } from '@/components/header';
import { useRouter } from 'next/navigation';

interface AccidentDetailProps {
  params: {
    accidentId: string;
  };
}

const AccidentDetail: React.FC<AccidentDetailProps> = ({ params }) => {
  const router = useRouter();
  const { accidentId } = params;
  const accidentIdNumber = parseInt(accidentId, 10);
  const { accidentDetails, isLoading } =
    useFetchAccidentDetails(accidentIdNumber);

  const handleLogoClick = () => {
    router.push('/');
  };

  if (isLoading) {
    return <Loader />;
  }

  const safetyRecommendations: string[] = [
    'Ensure proper maintenance checks are performed regularly.',
    'Implement additional training programs for flight crew members.',
  ];

  return (
    <>
      <Header onLogoClick={handleLogoClick} />
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
            {accidentDetails?.location?.city_name},{' '}
            {accidentDetails?.location?.state_name},{' '}
            {accidentDetails?.location?.country_name}
          </p>
          <MapComponent
            latitude={accidentDetails?.location?.latitude}
            longitude={accidentDetails?.location?.longitude}
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
        <Link href="/" passHref>
          <span className={`${buttonVariants({ variant: 'outline' })} w-full`}>
            Home
          </span>
        </Link>
        <div className="mb-2"></div>
      </section>
    </>
  );
};

export default AccidentDetail;
