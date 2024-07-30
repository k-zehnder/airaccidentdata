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
      // Search by query with pagination
      response = (await client.search<ElasticSearchResponse<Accident>>({
        index: 'accidents',
        body: {
          query: {
            multi_match: {
              query,
              fields: ['remark_text', 'aircraftDetails.aircraft_make_name'],
            },
          },
          from: (page - 1) * size,
          size,
        },
      })) as unknown as ElasticSearchResponse<Accident>;
    } else {
      // Fetch most recent accidents with pagination
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

      // Set Cache-Control headers only for homepage search
      res.setHeader('Cache-Control', 'public, max-age=7200, s-maxage=7200');
    }

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
