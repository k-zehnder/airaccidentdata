import { NextApiRequest, NextApiResponse } from 'next';
import { Client } from '@elastic/elasticsearch';
import { Accident } from '@/types/aviationTypes';

export interface ElasticSearchHit<T> {
  _index: string;
  _type: string;
  _id: string;
  _score: number;
  _source: T;
}

export interface ElasticSearchResponse<T> {
  hits: {
    total: { value: number; relation: string };
    max_score: number;
    hits: Array<ElasticSearchHit<T>>;
  };
}

// Initialize Elasticsearch client
const client = new Client({
  node: 'http://elasticsearch:9200',
  auth: {
    apiKey: process.env.ELASTICSEARCH_API_KEY!,
  },
});

export default async (req: NextApiRequest, res: NextApiResponse) => {
  const { query, size = 10, page = 1 } = req.body;

  try {
    let response: ElasticSearchResponse<Accident>;

    if (query) {
      // Search by query with pagination, including registration number
      response = (await client.search<ElasticSearchResponse<Accident>>({
        index: 'accidents',
        body: {
          query: {
            bool: {
              should: [
                {
                  // Exact match for the registration number
                  term: {
                    'aircraftDetails.registration_number.keyword': query,
                  },
                },
                {
                  // Fallback to multi-match search across other fields
                  multi_match: {
                    query,
                    fields: ['*'],
                  },
                },
              ],
            },
          },
          from: (page - 1) * size,
          size,
        },
      })) as unknown as ElasticSearchResponse<Accident>;
    } else {
      // Fetch the most recent accidents with pagination if no query is provided
      response = (await client.search<ElasticSearchResponse<Accident>>({
        index: 'accidents',
        body: {
          sort: [{ event_local_date: { order: 'desc' } }],
          from: (page - 1) * size,
          size,
          query: {
            match_all: {},
          },
        },
      })) as unknown as ElasticSearchResponse<Accident>;
    }

    // Map the Elasticsearch hits to the response format
    const results = response.hits.hits.map(
      (hit: ElasticSearchHit<Accident>) => hit._source
    );
    const total = response.hits.total.value;
    res.status(200).json({ results, total });
  } catch (error) {
    console.error('Search service error:', error);
    res.status(500).json({ error: 'Failed to fetch search results' });
  }
};
