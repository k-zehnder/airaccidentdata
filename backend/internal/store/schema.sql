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

-- Insert into Aircrafts Table
INSERT INTO Aircrafts (registration_number, aircraft_make_name, aircraft_model_name, aircraft_operator) VALUES
('ABC123', 'Boeing', '737', 'American Airlines'),
('DEF456', 'Airbus', 'A320', 'Delta Air Lines'),
('GHI789', 'Embraer', 'E175', 'United Airlines');

-- Insert into AircraftImages Table
INSERT INTO AircraftImages (aircraft_id, image_url, s3_url) VALUES
(1, 'https://example.com/image1.jpg', 's3://bucket/image1.jpg'),
(2, 'https://example.com/image2.jpg', 's3://bucket/image2.jpg'),
(3, 'https://example.com/image3.jpg', 's3://bucket/image3.jpg');

-- Insert into Accidents Table
INSERT INTO Accidents (
    updated, entry_date, event_local_date, event_local_time, location_city_name, location_state_name, 
    location_country_name, remark_text, event_type_description, fsdo_description, flight_number, 
    aircraft_missing_flag, aircraft_damage_description, flight_activity, flight_phase, far_part, 
    max_injury_level, fatal_flag, flight_crew_injury_none, flight_crew_injury_minor, 
    flight_crew_injury_serious, flight_crew_injury_fatal, flight_crew_injury_unknown, 
    cabin_crew_injury_none, cabin_crew_injury_minor, cabin_crew_injury_serious, cabin_crew_injury_fatal, 
    cabin_crew_injury_unknown, passenger_injury_none, passenger_injury_minor, passenger_injury_serious, 
    passenger_injury_fatal, passenger_injury_unknown, ground_injury_none, ground_injury_minor, 
    ground_injury_serious, ground_injury_fatal, ground_injury_unknown, aircraft_id
) VALUES (
    'No', '2023-01-15', '2023-01-14', '13:30:00', 'New York', 'New York', 'USA', 
    'Inclement weather conditions.', 'Crash landing', 'FAA', 'AA123', 'N', 'Minor damage to fuselage.', 
    'Scheduled passenger service', 'Landing', 'Part 121', 'Minor', 'N', 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 
    0, 0, 0, 0, 0, 0, 0, 0, 0, 1
);

-- Retrieve Aircrafts with their Images:
SELECT 
    * 
FROM 
    Aircrafts
JOIN 
    AircraftImages ON Aircrafts.id = AircraftImages.aircraft_id;

-- Check for missing foreign key relationships in Accidents table
SELECT *
FROM Accidents
WHERE aircraft_id NOT IN (SELECT id FROM Aircrafts);
