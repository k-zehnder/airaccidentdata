-- Create Database if not exists
CREATE DATABASE IF NOT EXISTS airaccidentdata;
USE airaccidentdata;

-- Create Aircrafts Table
CREATE TABLE IF NOT EXISTS Aircrafts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    registration_number VARCHAR(255) NOT NULL,
    aircraft_make_name VARCHAR(255) NOT NULL,
    aircraft_model_name VARCHAR(255) NOT NULL,
    aircraft_operator VARCHAR(255),
    unique_identifier VARCHAR(255) AS (CONCAT(aircraft_make_name, ' ', aircraft_model_name, ' ', registration_number)) STORED UNIQUE
);

-- Create Aircraft Images Table
CREATE TABLE IF NOT EXISTS AircraftImages (
    id INT AUTO_INCREMENT UNIQUE PRIMARY KEY,
    aircraft_id INT,
    image_url TEXT,
    s3_url VARCHAR(255),
    FOREIGN KEY (aircraft_id) REFERENCES Aircrafts(id)
);

-- Create Accidents Table
CREATE TABLE IF NOT EXISTS Accidents (
    id INT AUTO_INCREMENT UNIQUE PRIMARY KEY,
    updated VARCHAR(255),
    entry_date DATE,
    event_local_date DATE,
    event_local_time TIME,
    location_city_name VARCHAR(255),
    location_state_name VARCHAR(255),
    location_country_name VARCHAR(255),
    latitude FLOAT,
    longitude FLOAT,
    remark_text TEXT,
    event_type_description VARCHAR(255),
    fsdo_description VARCHAR(255),
    flight_number VARCHAR(255),
    aircraft_missing_flag VARCHAR(50),
    aircraft_damage_description VARCHAR(255),
    flight_activity VARCHAR(255),
    flight_phase VARCHAR(255),
    far_part VARCHAR(50),
    max_injury_level VARCHAR(50),
    fatal_flag VARCHAR(50),
    flight_crew_injury_none INT,
    flight_crew_injury_minor INT,
    flight_crew_injury_serious INT,
    flight_crew_injury_fatal INT,
    flight_crew_injury_unknown INT,
    cabin_crew_injury_none INT,
    cabin_crew_injury_minor INT,
    cabin_crew_injury_serious INT,
    cabin_crew_injury_fatal INT,
    cabin_crew_injury_unknown INT,
    passenger_injury_none INT,
    passenger_injury_minor INT,
    passenger_injury_serious INT,
    passenger_injury_fatal INT,
    passenger_injury_unknown INT,
    ground_injury_none INT,
    ground_injury_minor INT,
    ground_injury_serious INT,
    ground_injury_fatal INT,
    ground_injury_unknown INT,
    aircraft_id INT,
    FOREIGN KEY (aircraft_id) REFERENCES Aircrafts(id)
);
