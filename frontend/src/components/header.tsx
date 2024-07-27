'use client';

import React from 'react';
import Link from 'next/link';
import { ThemeToggle } from './theme-toggle';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPlane } from '@fortawesome/free-solid-svg-icons';

interface HeaderProps {
  onLogoClick: () => void;
}

export const Header: React.FC<HeaderProps> = ({ onLogoClick }) => {
  return (
    <header className="bg-background sticky top-0 z-40 w-full border-b">
      <div className="flex justify-between items-center py-4 px-6">
        <div className="flex items-center space-x-2">
          <div
            className="flex items-center space-x-2 cursor-pointer"
            onClick={onLogoClick}
          >
            <FontAwesomeIcon icon={faPlane} />
            <span className="inline-block font-bold">airaccidentdata.com</span>
          </div>
        </div>
        <div className="flex items-center space-x-4">
          <Link
            href="https://airaccidentdata.com/swagger/index.html"
            passHref
            target="_blank"
            rel="noopener noreferrer"
          >
            <span className="text-sm font-medium text-muted-foreground cursor-pointer">
              API
            </span>
          </Link>
          <ThemeToggle />
        </div>
      </div>
    </header>
  );
};
