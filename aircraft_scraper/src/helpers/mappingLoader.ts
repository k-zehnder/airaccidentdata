// Loads and parses an aircraft type mapping file into a Map object.
import fs from 'fs';
import { AircraftMapping } from '../types/aircraft';

// Loads the aircraft type map from a JSON file
export const loadAircraftTypeMap = (
  filePath: string
): Map<string, AircraftMapping> => {
  try {
    const data = fs.readFileSync(filePath);
    const jsonData = JSON.parse(data.toString());
    return new Map(Object.entries(jsonData));
  } catch (error) {
    console.error(
      `Error reading or parsing aircraft type map JSON file: ${error}`
    );
    return new Map();
  }
};
