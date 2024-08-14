import { Client } from '@elastic/elasticsearch';
import { Accident } from '../types/aviationTypes';
import { ElasticService } from '../types/elasticsearchTypes';
import config from '../config/config';

// Factory function to create an Elasticsearch service
export const createElasticService = (client: Client): ElasticService => {
  // Index a single accident document into Elasticsearch
  const indexAccident = async (accident: Accident): Promise<void> => {
    await client.index({
      index: 'accidents',
      id: String(accident.id),
      body: accident,
    });
  };

  // Index multiple accident documents into Elasticsearch
  const indexBulkAccidents = async (accidents: Accident[]): Promise<void> => {
    const body = accidents.flatMap((accident) => [
      { index: { _index: 'accidents', _id: String(accident.id) } },
      accident,
    ]);
    const bulkResponse = await client.bulk({ body });
    if (bulkResponse.errors) {
      console.error(
        'Errors occurred during bulk indexing:',
        bulkResponse.items
      );
    }
  };

  // Search for accidents in Elasticsearch
  const searchAccidents = async (query: string): Promise<Accident[]> => {
    const response = await client.search({
      index: 'accidents',
      body: {
        query: {
          bool: {
            should: [
              {
                term: {
                  'aircraftDetails.registration_number.keyword': query, // Exact match for registration number
                },
              },
              {
                multi_match: {
                  query,
                  fields: [
                    'remark_text',
                    'event_type_description',
                    'fatal_flag',
                  ],
                },
              },
            ],
          },
        },
      },
    });
    return response.hits.hits.map((hit: any) => hit._source as Accident);
  };

  // Fetch the most recent accidents from Elasticsearch
  const fetchRecentAccidents = async (size: number): Promise<Accident[]> => {
    const response = await client.search({
      index: 'accidents',
      body: {
        sort: [{ event_local_date: { order: 'desc' } }],
        size,
        query: {
          match_all: {},
        },
      },
    });
    return response.hits.hits.map((hit: any) => hit._source as Accident);
  };

  // Clear the Elasticsearch index
  const clearIndex = async (index: string): Promise<void> => {
    await client.deleteByQuery({
      index,
      body: {
        query: {
          match_all: {},
        },
      },
    });
  };

  return {
    indexAccident,
    indexBulkAccidents,
    searchAccidents,
    fetchRecentAccidents,
    clearIndex,
  };
};
