"use client";

import { useEffect } from 'react';
import Link from 'next/link';
import { Header } from '@/components/header';
import { buttonVariants } from '@/components/ui/button';
import { useFetchAccidentDetails } from '@/hooks/useAccidentData';
import MapComponent from '@/components/MapComponent';

const AccidentDetail = ({
  params,
}: {
  params: { accidentId: string };
}) => {
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
        <div>Loading accident details...</div>
      </>
    );
  }

  // Hardcoded safety recommendations
  const safetyRecommendations: string[] = [
    "Ensure proper maintenance checks are performed regularly.",
    "Implement additional training programs for flight crew members."
  ];

  return (
    <>
      <Header />
      <section className="container flex flex-col space-y-10 mt-10">
        <div className="text-4xl font-bold">Air Accident Details</div>
        {/* <div className="mt-10">Some information about {registrationNumber}</div> */}

        {/* Main Content Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {/* Left Column */}
          <div className="md:col-span-1 bg-white dark:bg-background/50 shadow-md rounded border p-6 mb-6">
            <h2 className="text-2xl font-semibold mb-4">Incident Overview</h2>
            <p>
              <strong>Date: </strong>{accidentDetails?.entry_date}
            </p>
            <p>
              <strong>Aircraft Make:</strong> {accidentDetails?.aircraftDetails?.aircraft_make_name} {accidentDetails?.aircraftDetails?.aircraft_model_name}
            </p>
            <p>
              <strong>Event Type:</strong> {accidentDetails?.event_type_description}
            </p>
            <p>
              <strong>Remark:</strong> {accidentDetails?.remark_text}
            </p>
            {/* <p>{accidentDetails.summary}</p> */}
          </div>

          {/* Right Column */}
          <div className="md:col-span-1 bg-white dark:bg-background/50 shadow-md rounded border p-6 mb-6">
            <h2 className="text-2xl font-semibold mb-4">
              Location Information
            </h2>
            <p>
              {accidentDetails?.location_city_name}, {accidentDetails?.location_state_name}, {accidentDetails?.location_country_name}
            </p>
            {/* Insert map or location details here */}
            <MapComponent />
          </div>
        </div>

        {/* Reports and Analysis Section */}
        <section className="bg-white dark:bg-background/50 shadow-md rounded border p-6 mb-6">
          <h2 className="text-2xl font-semibold mb-4">Reports & Analysis</h2>
          {/* Insert links to reports or summaries here */}
        </section>

        {/* Safety Measures / Recommendations Section */}
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
