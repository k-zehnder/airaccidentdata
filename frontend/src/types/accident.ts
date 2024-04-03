export interface Aircraft {
    aircraft_make_name: string;
    aircraft_model_name: string;
    aircraft_operator: string;
    id: number;
    registration_number: string;
}

export interface Accident {
    id: number;
    entry_date: string;
    event_local_date: string;
    event_local_time: string;
    location_city_name: string;
    location_state_name: string;
    location_country_name: string;
    latitude?: number; 
    longitude?: number; 
    remark_text: string;
    event_type_description: string;
    fatal_flag: string;
    flight_crew_injury_none: number;
    flight_crew_injury_minor: number;
    flight_crew_injury_serious: number;
    flight_crew_injury_fatal: number;
    flight_crew_injury_unknown: number;
    cabin_crew_injury_none: number;
    cabin_crew_injury_minor: number;
    cabin_crew_injury_serious: number;
    cabin_crew_injury_fatal: number;
    cabin_crew_injury_unknown: number;
    passenger_injury_none: number;
    passenger_injury_minor: number;
    passenger_injury_serious: number;
    passenger_injury_fatal: number;
    passenger_injury_unknown: number;
    ground_injury_none: number;
    ground_injury_minor: number;
    ground_injury_serious: number;
    ground_injury_fatal: number;
    ground_injury_unknown: number;
    aircraft_id: number;
    imageUrl: string;
    aircraftDetails?: Aircraft;
}
