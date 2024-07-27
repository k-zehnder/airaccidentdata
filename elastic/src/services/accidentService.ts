import mysql from 'mysql2/promise';
import { RowDataPacket } from 'mysql2';
import { Accident, Aircraft, Injury, Location } from '../types/aviationTypes';

// Create the accident service with a MySQL connection pool
export const createAccidentService = (pool: mysql.Pool) => {
  // Fetch a paginated list of accidents from the database
  const fetchAccidents = async (
    page: number,
    pageSize: number
  ): Promise<Accident[]> => {
    const offset = (page - 1) * pageSize;
    const [rows] = await pool.query<Accident[] & RowDataPacket[]>(
      'SELECT * FROM Accidents LIMIT ?, ?',
      [offset, pageSize]
    );
    return rows;
  };

  // Fetch the details of a specific aircraft by its ID
  const fetchAircraftDetails = async (
    aircraftId: number
  ): Promise<Aircraft> => {
    const [rows] = await pool.query<Aircraft[] & RowDataPacket[]>(
      'SELECT * FROM Aircrafts WHERE id = ?',
      [aircraftId]
    );
    return rows[0];
  };

  // Fetch the image URLs of a specific aircraft by its ID
  const fetchAircraftImages = async (aircraftId: number): Promise<string[]> => {
    const [rows] = await pool.query<{ image_url: string } & RowDataPacket[]>(
      'SELECT image_url FROM AircraftImages WHERE aircraft_id = ?',
      [aircraftId]
    );
    return rows.map((row) => row.image_url);
  };

  // Fetch the injuries associated with a specific accident by its ID
  const fetchAccidentInjuries = async (
    accidentId: number
  ): Promise<Injury[]> => {
    const [rows] = await pool.query<Injury[] & RowDataPacket[]>(
      'SELECT * FROM Injuries WHERE accident_id = ?',
      [accidentId]
    );
    return rows;
  };

  // Fetch the location details by location ID
  const fetchLocation = async (locationId: number): Promise<Location> => {
    const [rows] = await pool.query<Location[] & RowDataPacket[]>(
      'SELECT * FROM Locations WHERE id = ?',
      [locationId]
    );
    return rows[0];
  };

  return {
    fetchAccidents,
    fetchAircraftDetails,
    fetchAircraftImages,
    fetchAccidentInjuries,
    fetchLocation,
  };
};
