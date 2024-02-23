'use client';

import Link from 'next/link';
import { Header } from '@/components/header';
import { buttonVariants } from '@/components/ui/button';

export default function AccidentDetail({
  params,
}: {
  params: { registrationNumber: string };
}) {
  const { registrationNumber } = params;

  return (
    <>
      <Header />
      <section className="container flex flex-col space-y-10 mt-10">
        <div className="text-4xl font-bold">Air Accident Details</div>
        <div className="mt-10">Some information about {registrationNumber}</div>

        {/* <!-- Main Content Grid --> */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {/* <!-- Left Column --> */}
          <div className="md:col-span-1 bg-white dark:bg-background/50 shadow-md rounded border p-6 mb-6">
            <h2 className="text-2xl font-semibold mb-4">Incident Overview</h2>
            {/* <!-- Details like date, aircraft model, etc. --> */}
          </div>

          {/* <!-- Right Column --> */}
          <div className="md:col-span-1 bg-white dark:bg-background/50 shadow-md rounded border p-6 mb-6">
            <h2 className="text-2xl font-semibold mb-4">
              Location Information
            </h2>
            {/* <!-- Map or details about the location --> */}
          </div>
        </div>

        {/* <!-- Reports and Analysis Section --> */}
        <section className="bg-white dark:bg-background/50 shadow-md rounded border p-6 mb-6">
          <h2 className="text-2xl font-semibold mb-4">Reports & Analysis</h2>
          {/* <!-- Links to reports or summaries --> */}
        </section>

        {/* <!-- Safety Measures / Recommendations Section --> */}
        <section className="bg-white dark:bg-background/50 shadow-md rounded border p-6 mb-6">
          <h2 className="text-2xl font-semibold mb-4">
            Safety Recommendations
          </h2>
          {/* <!-- Bullet points or paragraphs --> */}
        </section>
        <Link href="/" className={buttonVariants({ variant: 'outline' })}>
          Home
        </Link>
      </section>
    </>
  );
}
