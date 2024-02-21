'use client'

import Link from "next/link";
import { Header } from "@/components/header"; 

export default function AccidentDetail({ params }: { params: { registrationNumber: string } }) {
  const { registrationNumber } = params;

  return (
    <>
      <Header />
      <section className="container flex flex-col space-y-10 mt-10">
        <div className="text-4xl">Information about this accident:</div> 
        <div className="mt-10">{registrationNumber}</div>
          <Link legacyBehavior href="/">
              <a className="underline underline-offset-4 text-gray-800 dark:text-white hover:text-blue-600 dark:hover:text-blue-400">Back to Home</a>
          </Link>
      </section>
    </>
  );
}