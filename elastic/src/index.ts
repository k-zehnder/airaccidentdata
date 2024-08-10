import mysql from 'mysql2/promise';
import { Client } from '@elastic/elasticsearch';
import config from './config/config';
import { createAccidentService } from './services/accidentService';
import { createElasticService } from './services/elasticService';
import { Accident } from './types/aviationTypes';
import {
  AccidentService,
  ElasticOperations,
  ElasticService,
} from './types/elasticsearchTypes';

// Factory function to create Elasticsearch operations
const createElasticOperations = (
  accidentService: AccidentService,
  elasticService: ElasticService
): ElasticOperations => ({
  // Index all accidents to Elasticsearch
  indexAllAccidentsToElastic: async (): Promise<void> => {
    try {
      const pageSize = 10;
      let page = 1;
      let totalAccidents = 0;

      while (true) {
        const accidents = await accidentService.fetchAccidents(page, pageSize);
        if (accidents.length === 0) break;

        const detailedAccidents: Accident[] = await Promise.all(
          accidents.map(async (accident) => {
            const aircraftDetails = await accidentService.fetchAircraftDetails(
              accident.aircraft_id
            );
            const images = await accidentService.fetchAircraftImages(
              accident.aircraft_id
            );
            const injuries = await accidentService.fetchAccidentInjuries(
              accident.id
            );
            const location = await accidentService.fetchLocation(accident.id);

            return {
              ...accident,
              aircraftDetails,
              imageUrl:
                images[0] ||
                'https://upload.wikimedia.org/wikipedia/commons/e/e2/BK-117_Polizei-NRW_D-HNWL.jpg',
              injuries,
              location,
            };
          })
        );

        console.log(
          `Indexing ${detailedAccidents.length} accidents to Elasticsearch...`
        );
        await elasticService.indexBulkAccidents(detailedAccidents);
        totalAccidents += detailedAccidents.length;
        console.log(
          `Indexed ${totalAccidents} accidents to Elasticsearch so far.`
        );

        page++;
      }

      console.log('Indexed all accidents to Elasticsearch successfully.');
    } catch (error) {
      console.error('Error indexing accidents:', error);
    }
  },

  // Search for accidents in Elasticsearch
  searchAccidentsInElastic: async (query: string): Promise<void> => {
    try {
      console.log(`Searching for accidents with query: ${query}`);
      const results = await elasticService.searchAccidents(query);
      console.log('Search results:', results);
    } catch (error) {
      console.error('Error searching accidents:', error);
    }
  },

  // Fetch the most recent accidents from Elasticsearch
  fetchRecentAccidentsFromElastic: async (size: number): Promise<void> => {
    try {
      console.log(
        `Fetching ${size} most recent accidents from Elasticsearch...`
      );
      const results = await elasticService.fetchRecentAccidents(size);
      console.log('Recent accidents:', results);
    } catch (error) {
      console.error('Error fetching recent accidents:', error);
    }
  },

  // Clear the Elasticsearch index
  clearIndex: async (index: string): Promise<void> => {
    try {
      await elasticService.clearIndex(index);
      console.log(`Cleared all documents in the index ${index} successfully.`);
    } catch (error) {
      console.error(`Error clearing index ${index}:`, error);
    }
  },
});

// Execute indexing, searching, and fetching recent accidents with optional clearing
const run = async (clear: boolean = false): Promise<void> => {
  // Initialize MySQL connection pool
  const pool = mysql.createPool(config.mysql);

  // Initialize Elasticsearch client
  const esClient = new Client({
    node: config.elasticsearch.host,
    auth: {
      apiKey: config.elasticsearch.apiKey,
    },
  });

  // Create services
  const accidentService = createAccidentService(pool);
  const elasticService = createElasticService(esClient);

  // Create Elasticsearch operations
  const elasticOps = createElasticOperations(accidentService, elasticService);

  // Usage example:
  // To clear the Elasticsearch index before running the script, use:
  // ts-node index.ts --clear
  if (clear) {
    await elasticOps.clearIndex('accidents');
  }

  await elasticOps.indexAllAccidentsToElastic();
  await elasticOps.searchAccidentsInElastic('appleton');
  await elasticOps.fetchRecentAccidentsFromElastic(5);
};

// Parse command line arguments using built-in process.argv
const args = process.argv.slice(2);
const clearFlag = args.includes('--clear');

// Run the script
run(clearFlag);
