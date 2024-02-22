'use client'

import Link from "next/link";
import { Header } from "@/components/header"; 
import { buttonVariants } from "@/components/ui/button";

export default function AccidentDetail({ params }: { params: { registrationNumber: string } }) {
  const { registrationNumber } = params;

  return (
    <>
      <Header />
      <section className="container flex flex-col space-y-10 mt-10">
        <div className="text-4xl">Information about this accident:</div> 
        <div className="mt-10">{registrationNumber}</div>
        <Link
          href="/"
          className={buttonVariants({ variant: "outline" })}
        >
          Home
        </Link>
      </section>
    </>
  );
}