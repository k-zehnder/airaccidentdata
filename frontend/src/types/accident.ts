export interface Accident {
    id: number;
    registrationNumber: string;
    remarkText: string;
    aircraftMakeName: string;
    aircraftModelName: string;
    entryDate: string;
    fatalFlag: string;
    flightCrewInjuryNone: number;
    flightCrewInjuryMinor: number;
    flightCrewInjurySerious: number;
    flightCrewInjuryFatal: number;
    flightCrewInjuryUnknown: number;
    cabinCrewInjuryNone: number;
    cabinCrewInjuryMinor: number;
    cabinCrewInjurySerious: number;
    cabinCrewInjuryFatal: number;
    cabinCrewInjuryUnknown: number;
    passengerInjuryNone: number;
    passengerInjuryMinor: number;
    passengerInjurySerious: number;
    passengerInjuryFatal: number;
    passengerInjuryUnknown: number;
    groundInjuryNone: number;
    groundInjuryMinor: number;
    groundInjurySerious: number;
    groundInjuryFatal: number;
    groundInjuryUnknown: number;
}
  
export interface AccidentDetails {
    entryDate: string;
    aircraftMakeName: string;
    location: string;
    summary: string;
    recommendations: string[];
    locationCityName: string;
    locationStateName: string;
    locationCountryName: string;
    remarkText: string;
    eventTypeDescription: string;
}