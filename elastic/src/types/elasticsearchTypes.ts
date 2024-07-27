import { Accident } from './aviationTypes';

// Interface for the accident service
export interface AccidentService {
  fetchAccidents: (page: number, pageSize: number) => Promise<Accident[]>;
  fetchAircraftDetails: (aircraftId: number) => Promise<any>;
  fetchAircraftImages: (aircraftId: number) => Promise<string[]>;
  fetchAccidentInjuries: (accidentId: number) => Promise<any>;
  fetchLocation: (accidentId: number) => Promise<any>;
}

// Interface for the elastic operations
export interface ElasticOperations {
  indexAllAccidentsToElastic: () => Promise<void>;
  searchAccidentsInElastic: (query: string) => Promise<void>;
  fetchRecentAccidentsFromElastic: (size: number) => Promise<void>;
  clearIndex: (index: string) => Promise<void>;
}

// Interface for the Elasticsearch service
export interface ElasticService {
  indexAccident: (accident: Accident) => Promise<void>;
  indexBulkAccidents: (accidents: Accident[]) => Promise<void>;
  searchAccidents: (query: string) => Promise<Accident[]>;
  fetchRecentAccidents: (size: number) => Promise<Accident[]>;
  clearIndex: (index: string) => Promise<void>;
}

// Define the data structure for an Elasticsearch response
export interface ElasticsearchResponse<T> {
  hits: {
    total: { value: number };
    hits: Array<{
      _source: T;
    }>;
  };
}
