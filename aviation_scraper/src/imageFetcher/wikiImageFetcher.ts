import axios from 'axios';

export interface Fetcher {
    fetchHtmlFromUrl(url: string): Promise<string>;
    fetchImageFromUrl(url: string): Promise<Buffer>;
}

export const createAxiosFetcher = (): Fetcher => {
    const fetchHtmlFromUrl = async (url: string): Promise<string> => {
        const response = await axios.get(url);
        return response.data;
    };

    const fetchImageFromUrl = async (url: string): Promise<Buffer> => {
        const response = await axios.get(url, { responseType: 'arraybuffer' });
        return response.data;
    };

    return { fetchHtmlFromUrl, fetchImageFromUrl };
};
