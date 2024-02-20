"use client";

import React, { useEffect, useState } from "react";
import Link from "next/link";
import { Header } from "@/components/header";
import axios from 'axios';

interface Accident {
  id: number;
  registrationNumber: string;
  remarkText: string;
  aircraftMakeName: string;
  aircraftModelName: string;
  link: string;
}

export default function Home() {
  const [accidents, setAccidents] = useState([]);
  const [isFetching, setFetching] = useState(false);

  useEffect(() => {
    const fetchAccidents = async () => {
      try {
        const apiUrl =
          process.env.NEXT_PUBLIC_ENV === 'development'
            ? 'http://localhost:8080/api/v1/accidents'
            : 'https://airaccidentdata.com/api/v1/accidents';
        const response = await axios.get(apiUrl);
        setAccidents(response.data);
        setFetching(false);
      } catch (error) {
        console.error('Error fetching accidents:', error);
        setFetching(false);
      }
    };

    fetchAccidents();
  }, []);

  if (isFetching) return <div>Loading...</div>;

  return (
    <>
      {/* Header */}
      <Header />

      <div className="container mx-auto px-4 py-6">
        <div className="grid grid-cols-1 lg:grid-cols-4 gap-6">
          {/* Sidebar - make sure it's only visible on large screens */}
          <div className="hidden lg:block lg:col-span-1 p-4 rounded-lg border">
            <h3 className="text-xl font-bold mb-4">Sidebar</h3>
            <ul className="space-y-2">
              <li><a href="#" className="text-blue-600 hover:underline">Recent Accidents</a></li>
              <li><a href="#" className="text-blue-600 hover:underline">Top Categories</a></li>
              <li><a href="#" className="text-blue-600 hover:underline">Search Reports</a></li>
              {/* Add more sidebar items here */}
            </ul>
          </div>

          {/* Main Content - takes up the full width on small screens, and 3 columns on large screens */}
          <div className="lg:col-span-3 p-4">
            <h1 className="text-4xl font-bold tracking-tight mb-6">
              Explore Aviation Accidents and Insights
            </h1>

            {/* Accident Posts */}
            {accidents.map((accident: Accident) => (
              <article key={accident.id} className="border-b-2 py-4">
                <Link legacyBehavior href={`/accidents/${accident.id}`}>
                  <a>
                    <h2 className="text-2xl font-semibold">{accident.registrationNumber}: {accident.aircraftMakeName} {accident.aircraftModelName}</h2>
                    <p className="text-gray-600">{accident.remarkText}</p>
                  </a>
                </Link>
              </article>
            ))}
          </div>
        </div>
      </div>
    </>
  );
}
