'use client';

import React, { useEffect, useState } from 'react';
import Link from 'next/link';
import Pagination from '@/components/pagination';
import AccidentBadges from '@/components/AccidentBadges';
import SearchBar from '@/components/SearchBar';
import Loader from '@/components/Loader';
import { formatDate } from '@/lib/utils';
import { Header } from '@/components/header';
import { useRouter, useSearchParams } from 'next/navigation';
import { useAccidentData } from '@/hooks/useAccidentData';
import SuspenseWrapper from '@/components/SuspenseWrapper';

const Home: React.FC = () => {
  const router = useRouter();
  const searchParams = useSearchParams() as URLSearchParams;
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [searchQuery, setSearchQuery] = useState<string>('');

  useEffect(() => {
    const page = searchParams.get('page');
    const query = searchParams.get('query');
    if (page) setCurrentPage(Number(page));
    if (query) setSearchQuery(query);
  }, [searchParams]);

  const { accidents, totalPages, isLoading } = useAccidentData(
    currentPage,
    searchQuery
  );

  const handleLogoClick = () => {
    setSearchQuery('');
    setCurrentPage(1);
    router.push('/');
  };

  const handleSearch = (query: string) => {
    setSearchQuery(query);
    setCurrentPage(1);
    router.push(`/?page=1&query=${query}`);
  };

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
    router.push(`/?page=${page}&query=${searchQuery}`);
  };

  if (isLoading) {
    return <Loader />;
  }

  return (
    <>
      <Header onLogoClick={handleLogoClick} />

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

        <SearchBar onSearch={handleSearch} searchQuery={searchQuery} />

        <div className="max-w-4xl mx-auto">
          {accidents.map((accident) => (
            <Link
              key={accident.id}
              href={`/accidents/${accident.id}?page=${currentPage}`}
              className="block border-b-2 py-4 flex items-center hover:bg-gray-100 dark:hover:bg-gray-900 transition-colors"
            >
              <div className="flex-1">
                <span className="text-gray-500 text-sm block lg:text-base mb-1">
                  {formatDate(accident.entry_date)}
                </span>
                <h3 className="text-xl font-semibold mb-1">
                  {accident.aircraftDetails?.registration_number}:{' '}
                  {accident.aircraftDetails?.aircraft_make_name}{' '}
                  {accident.aircraftDetails?.aircraft_model_name}
                </h3>
                <AccidentBadges accident={accident} />
                <p className="text-gray-500">{accident.remark_text}</p>
              </div>
              <div className="w-1/4 flex justify-end">
                <img
                  src={
                    accident.imageUrl ||
                    'https://upload.wikimedia.org/wikipedia/commons/e/e2/BK-117_Polizei-NRW_D-HNWL.jpg'
                  }
                  alt={`Thumbnail for ${accident.aircraft_id}`}
                  className="w-16 h-16 object-contain mb-2"
                />
              </div>
            </Link>
          ))}
        </div>

        <Pagination
          currentPage={currentPage}
          totalPages={totalPages}
          onPageChange={handlePageChange}
          darkMode={true}
        />
      </main>
    </>
  );
};

const Page: React.FC = () => {
  return (
    <SuspenseWrapper>
      <Home />
    </SuspenseWrapper>
  );
};

export default Page;
