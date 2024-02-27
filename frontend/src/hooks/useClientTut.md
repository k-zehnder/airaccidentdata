To modularize the fetch logic into a custom hook called `useClient`, you can create a separate file for the hook and then utilize it within your component. Here's how you can do it:

First, create a new file called `useClient.ts` in your components directory:

```typescript
// components/useClient.ts
import { useState, useEffect } from 'react';
import axios from 'axios';

interface Accident {
  id: number;
  registrationNumber: string;
  remarkText: string;
  aircraftMakeName: string;
  aircraftModelName: string;
  entryDate: string;
  fatalFlag: string;
  link: string;
}

export const useClient = (currentPage: number) => {
  const [accidents, setAccidents] = useState<Accident[]>([]);
  const [totalPages, setTotalPages] = useState(0);
  const [isFetching, setFetching] = useState(false);

  useEffect(() => {
    const fetchAccidents = async () => {
      setFetching(true);
      try {
        const apiUrl =
          process.env.NEXT_PUBLIC_ENV === 'development'
            ? `http://localhost:8080/api/v1/accidents?page=${currentPage}`
            : `https://airaccidentdata.com/api/v1/accidents?page=${currentPage}`;
        const response = await axios.get<{
          accidents: Accident[];
          total: number;
        }>(apiUrl);
        setAccidents(response.data.accidents);
        setTotalPages(Math.ceil(response.data.total / 10)); // 10 accidents per page
      } catch (error) {
        console.error('Error fetching accidents:', error);
      }
      setFetching(false);
    };

    fetchAccidents();
  }, [currentPage]);

  return { accidents, totalPages, isFetching };
};
```

Then, you can modify your `Home` component to use this hook:

```typescript
import React, { useState } from 'react';
import Link from 'next/link';
import { Header } from '@/components/header';
import Pagination from '@/components/pagination';
import { Badge } from '@/components/ui/badge';
import { useClient } from '@/components/useClient';

const formatDate = (dateString: string) => {
  const options: Intl.DateTimeFormatOptions = {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  };
  const date = new Date(dateString);
  return date.toLocaleDateString('en-US', options);
};

const Home = () => {
  const [currentPage, setCurrentPage] = useState(1);
  const { accidents, totalPages, isFetching } = useClient(currentPage);

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
                  {accident.fatalFlag === 'Yes' ? (
                    <Badge key={accident.id} className="bg-red-500 mb-2">
                      Fatalities
                    </Badge>
                  ) : null}
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
        />
      </main>
    </>
  );
};

export default Home;
```

This separates the fetching logic into a reusable hook, `useClient`, which you can then use in any component that needs to fetch data from the API.