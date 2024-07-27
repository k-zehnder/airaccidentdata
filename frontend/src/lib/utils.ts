import { type ClassValue, clsx } from 'clsx';
import { twMerge } from 'tailwind-merge';

// Merge Tailwind CSS class names with conditional class names
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

// Format a date string into a more readable format
export function formatDate(dateString: string) {
  const options: Intl.DateTimeFormatOptions = {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  };
  const date = new Date(dateString);
  return date.toLocaleDateString('en-US', options);
}

// Generate page numbers for pagination, including ellipsis for large gaps
export const getPageNumbers = (currentPage: number, totalPages: number) => {
  const pageNumbers = [];
  // Determine the start and end pages for the range
  let startPage = Math.max(currentPage - 2, 1);
  let endPage = Math.min(currentPage + 2, totalPages);

  // Always show the first page
  if (startPage > 1) pageNumbers.push(1);
  // Add ellipsis if there's a gap between the first page and the start page
  if (startPage > 2) pageNumbers.push(-1);

  // Add page numbers within the range
  for (let i = startPage; i <= endPage; i++) {
    pageNumbers.push(i);
  }

  // Add ellipsis if there's a gap between the end page and the last page
  if (endPage < totalPages - 1) pageNumbers.push(-2);
  // Always show the last page
  if (endPage < totalPages) pageNumbers.push(totalPages);

  return pageNumbers;
};
