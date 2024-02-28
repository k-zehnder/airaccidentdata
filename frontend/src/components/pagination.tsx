// Pagination component
import React from 'react';

interface PaginationProps {
  currentPage: number;
  totalPages: number;
  onPageChange: (page: number) => void;
  darkMode?: boolean; 
}

const Pagination: React.FC<PaginationProps> = ({ currentPage, totalPages, onPageChange, darkMode }) => {
  // Function to create page numbers
  const getPageNumbers = () => {
    const pages = [];
    for (let i = 1; i <= totalPages; i++) {
      pages.push(i);
    }
    return pages;
  };

  return (
    <div className="flex justify-center my-8">
      {currentPage > 1 && (
        <button
          onClick={() => onPageChange(currentPage - 1)}
          className={`mx-2 px-4 py-2 ${darkMode ? 'bg-gray-800 text-white' : 'bg-foreground text-white'} rounded hover:bg-foreground focus:outline-none focus:ring-2 focus:bg-foreground focus:ring-opacity-50`}
        >
          Previous
        </button>
      )}

      {getPageNumbers().map((page) => (
        <button
          key={page}
          onClick={() => onPageChange(page)}
          className={`mx-1 px-3 py-2 text-sm font-medium ${currentPage === page ? darkMode ? 'bg-gray-800 text-white' : 'bg-foreground text-white' : 'bg-white text-gray-700'} rounded hover:bg-foreground hover:text-white focus:outline-none focus:ring-2 focus:ring-foreground focus:ring-opacity-50`}
        >
          {page}
        </button>
      ))}

      {currentPage < totalPages && (
        <button
          onClick={() => onPageChange(currentPage + 1)}
          className={`mx-2 px-4 py-2 ${darkMode ? 'bg-gray-800 text-white' : 'bg-foreground text-white'} rounded hover:bg-foreground focus:outline-none focus:ring-2 focus:ring-blue-600 focus:ring-opacity-50`}
        >
          Next
        </button>
      )}
    </div>
  );
};

export default Pagination;
