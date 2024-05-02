CREATE DATABASE IF NOT EXISTS airaccidentdata;
USE airaccidentdata;

CREATE TABLE IF NOT EXISTS Aircrafts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    registration_number VARCHAR(255),
    aircraft_make_name VARCHAR(255),
    aircraft_model_name VARCHAR(255),
    aircraft_operator VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS Locations (
    id INT AUTO_INCREMENT PRIMARY KEY,
    city_name VARCHAR(255),
    state_name VARCHAR(255),
    country_name VARCHAR(255),
    latitude FLOAT,
    longitude FLOAT
);

CREATE TABLE IF NOT EXISTS Accidents (
    id INT AUTO_INCREMENT PRIMARY KEY,
    updated VARCHAR(255),
    entry_date DATE,
    event_local_date DATE,
    event_local_time TIME,
    remark_text VARCHAR(1024),
    event_type_description VARCHAR(255),
    fsdo_description VARCHAR(255),
    flight_number VARCHAR(255),
    aircraft_missing_flag VARCHAR(50),
    aircraft_damage_description VARCHAR(255),
    flight_activity VARCHAR(255),
    flight_phase VARCHAR(255),
    far_part VARCHAR(50),
    fatal_flag VARCHAR(50),
    aircraft_id INT,
    location_id INT,
    FOREIGN KEY (aircraft_id) REFERENCES Aircrafts(id),
    FOREIGN KEY (location_id) REFERENCES Locations(id)
);

CREATE TABLE IF NOT EXISTS Injuries (
    id INT AUTO_INCREMENT PRIMARY KEY,
    person_type VARCHAR(50),
    injury_severity VARCHAR(50),
    count INT,
    accident_id INT,
    FOREIGN KEY (accident_id) REFERENCES Accidents(id)
);

CREATE TABLE IF NOT EXISTS AircraftImages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    image_url VARCHAR(2048),
    s3_url VARCHAR(255),
    aircraft_id INT,
    FOREIGN KEY (aircraft_id) REFERENCES Aircrafts(id)
);
