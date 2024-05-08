export interface Accident {
  id: number;
  updated: string;
  entry_date: string;
  event_local_date: string;
  event_local_time: string;
  aircraftDetails?: Aircraft;
  aircraft_damage_description?: string;
  aircraft_id: number;
  aircraft_missing_flag?: string;
  event_type_description: string;
  far_part?: string;
  fatal_flag: string;
  flight_activity?: string;
  flight_number?: string;
  flight_phase?: string;
  fsdo_description?: string;
  imageUrl: string;
  location?: Location;
  location_id?: number;
  remark_text: string;
  injuries?: Injury[];
}

export interface Aircraft {
  id: number;
  registration_number: string;
  aircraft_make_name: string;
  aircraft_model_name: string;
  aircraft_operator?: string;
}

export interface Location {
  id: number;
  city_name: string;
  state_name: string;
  country_name: string;
  latitude: number;
  longitude: number;
}

export interface Injury {
  id: number;
  person_type: string;
  injury_severity: string;
  count: number;
  accident_id: number;
}
