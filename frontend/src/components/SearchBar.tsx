import React, { useState, ChangeEvent, FormEvent } from 'react';

interface SearchBarProps {
  onSearch: (query: string) => void;
  searchQuery: string;
}

const SearchBar: React.FC<SearchBarProps> = ({ onSearch, searchQuery }) => {
  const [query, setQuery] = useState<string>(searchQuery);

  const handleSearch = (e: FormEvent) => {
    e.preventDefault();
    onSearch(query);
  };

  return (
    <form onSubmit={handleSearch} className="flex justify-center mb-8">
      <input
        type="text"
        value={query}
        onChange={(e: ChangeEvent<HTMLInputElement>) =>
          setQuery(e.target.value)
        }
        placeholder="Search accidents..."
        className="w-full max-w-lg p-2 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:border-blue-500"
      />
      <button
        type="submit"
        className="ml-2 px-4 py-2 bg-blue-500 text-white rounded-lg shadow-md hover:bg-blue-600"
      >
        Search
      </button>
    </form>
  );
};

export default SearchBar;
