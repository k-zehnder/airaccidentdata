'use client';

import React from 'react';
import { ThemeToggle } from './Theme-Toggle';
import Link from 'next/link';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPlane } from '@fortawesome/free-solid-svg-icons';

const Header = () => {
  return (
    <header className="bg-background sticky top-0 z-40 w-full border-b">
      <div className="flex justify-between items-center py-4 px-6">
        {/* Left Section */}
        <div className="flex items-center space-x-2">
          <Link legacyBehavior href="/">
            <a className="flex items-center space-x-2">
              <FontAwesomeIcon icon={faPlane} />
              <span className="inline-block font-bold">
                airaccidentdata.com
              </span>
            </a>
          </Link>
        </div>

        {/* Right Section */}
        <div className="flex items-center space-x-4">
          {/* API Link */}
          <Link
            legacyBehavior
            href="https://airaccidentdata.com/swagger/index.html"
          >
            <a
              target="_blank"
              rel="noopener noreferrer"
              className="flex items-center"
            >
              <span className="text-sm font-medium text-muted-foreground">
                API
              </span>
            </a>
          </Link>

          {/* Theme Toggle */}
          <ThemeToggle />
        </div>
      </div>
    </header>
  );
};

export default Header;
