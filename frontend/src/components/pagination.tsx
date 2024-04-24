import React from 'react';

interface PaginationProps {
  currentPage: number;
  totalPages: number;
  onPageChange: (page: number) => void;
  darkMode?: boolean;
}

// Pagination component for navigating through pages
const Pagination: React.FC<PaginationProps> = ({
  currentPage,
  totalPages,
  onPageChange,
  darkMode,
}) => {
  // Calculates the range of page numbers to display
  const getPageNumbers = () => {
    let pages = [];
    let startPage = Math.max(currentPage - 2, 1);
    let endPage = Math.min(startPage + 4, totalPages);

    // Adjust the starting page number if the range is less than 5
    if (endPage - startPage < 4) {
      startPage = Math.max(endPage - 4, 1);
    }

    // Populate the array with page numbers within the range
    for (let i = startPage; i <= endPage; i++) {
      pages.push(Number(Math.round(i))); // Ensure page numbers are integers
    }

    return pages;
  };

  return (
    <div className="flex items-center justify-center space-x-1 my-8">
      {/* Display First and Previous buttons if not on the first page */}
      {currentPage > 1 && (
        <>
          <button
            onClick={() => onPageChange(1)}
            className={`px-3 py-2 ${darkMode ? 'bg-gray-800 text-white' : 'bg-foreground text-white'} rounded hover:bg-foreground focus:outline-none focus:ring-2 focus:ring-opacity-50`}
          >
            First
          </button>
          <button
            onClick={() => onPageChange(currentPage - 1)}
            className={`px-3 py-2 ${darkMode ? 'bg-gray-800 text-white' : 'bg-foreground text-white'} rounded hover:bg-foreground focus:outline-none focus:ring-2 focus:ring-opacity-50`}
          >
            Previous
          </button>
        </>
      )}

      {/* Display page numbers */}
      {getPageNumbers().map((page) => (
        <button
          key={page}
          onClick={() => onPageChange(page)}
          className={`px-3 py-2 text-sm ${currentPage === page ? (darkMode ? 'bg-gray-800 text-white' : 'bg-foreground text-white') : 'bg-white text-gray-700'} rounded hover:bg-foreground hover:text-white focus:outline-none focus:ring-2 focus:ring-foreground focus:ring-opacity-50`}
        >
          {page}
        </button>
      ))}

      {/* Display Next and Last buttons if not on the last page */}
      {currentPage < totalPages && (
        <>
          <button
            onClick={() => onPageChange(currentPage + 1)}
            className={`px-3 py-2 ${darkMode ? 'bg-gray-800 text-white' : 'bg-foreground text-white'} rounded hover:bg-foreground focus:outline-none focus:ring-2 focus:ring-blue-600 focus:ring-opacity-50`}
          >
            Next
          </button>
          <button
            onClick={() => onPageChange(totalPages)}
            className={`px-3 py-2 ${darkMode ? 'bg-gray-800 text-white' : 'bg-foreground text-white'} rounded hover:bg-foreground focus:outline-none focus:ring-2 focus:ring-opacity-50`}
          >
            Last
          </button>
        </>
      )}
    </div>
  );
};

export default Pagination;
