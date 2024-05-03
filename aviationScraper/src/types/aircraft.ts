export interface AircraftMapping {
  url: string;
  type: 'exact' | 'direct' | 'indirect';
}

export interface AircraftType {
  id: number;
  make: string;
  model: string;
  registrationNumber: string;
}
