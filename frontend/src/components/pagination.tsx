import React from 'react';
import { getPageNumbers } from '@/lib/utils';

interface PaginationProps {
  currentPage: number;
  totalPages: number;
  onPageChange: (page: number) => void;
  darkMode?: boolean;
}

const Pagination: React.FC<PaginationProps> = ({
  currentPage,
  totalPages,
  onPageChange,
  darkMode,
}) => {
  const pageNumbers = getPageNumbers(currentPage, totalPages);

  return (
    <div className="flex items-center justify-center space-x-1 my-8">
      {pageNumbers.map((page, index) =>
        page > 0 ? (
          <button
            key={index}
            onClick={() => onPageChange(page)}
            className={`px-3 py-2 text-sm ${
              currentPage === page
                ? darkMode
                  ? 'bg-gray-800 text-white'
                  : 'bg-foreground text-white'
                : 'bg-white text-gray-700'
            } rounded hover:bg-foreground hover:text-white focus:outline-none focus:ring-2 focus:ring-foreground focus:ring-opacity-50`}
          >
            {page}
          </button>
        ) : (
          <span key={index} className="px-3 py-2 text-sm text-gray-500">
            ...
          </span>
        )
      )}
    </div>
  );
};

export default Pagination;
