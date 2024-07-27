export interface Accident {
  id: number;
  updated: string;
  entry_date: string;
  event_local_date: string;
  event_local_time: string;
  aircraft_id: number;
  aircraftDetails?: Aircraft;
  event_type_description: string;
  fatal_flag: string;
  imageUrl?: string;
  remark_text: string;
  injuries?: Injury[];
  location?: Location;
}

export interface Aircraft {
  id: number;
  registration_number: string;
  aircraft_make_name: string;
  aircraft_model_name: string;
  aircraft_operator: string;
}

export interface Injury {
  id: number;
  person_type: string;
  injury_severity: string;
  count: number;
  accident_id: number;
}

export interface Location {
  city_name: string;
  state_name: string;
  country_name: string;
  latitude: number;
  longitude: number;
}
