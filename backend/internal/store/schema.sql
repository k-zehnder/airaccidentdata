-- Create Database if not exists
CREATE DATABASE IF NOT EXISTS airaccidentdata;
USE airaccidentdata;

-- Create Aircrafts Table
CREATE TABLE IF NOT EXISTS Aircrafts (
    id INT AUTO_INCREMENT UNIQUE PRIMARY KEY,
    registration_number VARCHAR(255) UNIQUE,
    aircraft_make_name VARCHAR(255),
    aircraft_model_name VARCHAR(255),
    aircraft_operator VARCHAR(255)
);

-- Create Aircraft Images Table
CREATE TABLE IF NOT EXISTS AircraftImages (
    id INT AUTO_INCREMENT UNIQUE PRIMARY KEY,
    aircraft_id INT,
    image_url TEXT,
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

-- Create Accident Images Table
CREATE TABLE IF NOT EXISTS AccidentImages (
    id INT AUTO_INCREMENT UNIQUE PRIMARY KEY,
    accident_id INT,
    image_url VARCHAR(255),
    FOREIGN KEY (accident_id) REFERENCES Accidents(id)
);

-- Insert a new aircraft
INSERT INTO Aircrafts (registration_number, aircraft_make_name, aircraft_model_name, aircraft_operator)
VALUES ('REG123', 'Make', 'Model', 'Operator')
ON DUPLICATE KEY UPDATE
    aircraft_make_name = VALUES(aircraft_make_name),
    aircraft_model_name = VALUES(aircraft_model_name),
    aircraft_operator = VALUES(aircraft_operator);

-- Insert a new accident
INSERT INTO Accidents (
    updated, 
    entry_date, 
    event_local_date, 
    event_local_time, 
    location_city_name, 
    location_state_name, 
    location_country_name, 
    remark_text, 
    event_type_description, 
    fsdo_description, 
    flight_number, 
    aircraft_missing_flag, 
    aircraft_damage_description, 
    flight_activity, 
    flight_phase, 
    far_part, 
    max_injury_level, 
    fatal_flag, 
    flight_crew_injury_none, 
    flight_crew_injury_minor, 
    flight_crew_injury_serious, 
    flight_crew_injury_fatal, 
    flight_crew_injury_unknown, 
    cabin_crew_injury_none, 
    cabin_crew_injury_minor, 
    cabin_crew_injury_serious, 
    cabin_crew_injury_fatal, 
    cabin_crew_injury_unknown, 
    passenger_injury_none, 
    passenger_injury_minor, 
    passenger_injury_serious, 
    passenger_injury_fatal, 
    passenger_injury_unknown, 
    ground_injury_none, 
    ground_injury_minor, 
    ground_injury_serious, 
    ground_injury_fatal, 
    ground_injury_unknown,
    aircraft_id
) VALUES (
    'no',
    '2024-02-28', 
    '2024-02-28', 
    '14:45:00', 
    'Springfield', 
    'Illinois', 
    'USA', 
    'Another accident occurred', 
    'Accident', 
    'Springfield FSDO', 
    'FL124', 
    'No', 
    'Major', 
    'Commercial Air Transport', 
    'Takeoff', 
    'Part 121', 
    'Serious', 
    'Yes', 
    0, 
    1, 
    0, 
    0, 
    0, 
    0, 
    0, 
    1, 
    0, 
    0, 
    0, 
    0, 
    0, 
    1, 
    0, 
    0, 
    0, 
    0, 
    0, 
    0, 
    1  
);

-- Insert a new image for the aircraft
INSERT INTO AircraftImages (aircraft_id, image_url)
VALUES (1, 'https://example.com/aircraft_image.jpg');

-- Insert a new image for the accident
INSERT INTO AccidentImages (accident_id, image_url)
VALUES (1, 'https://example.com/accident_image.jpg');

-- Query for accidents of a specific aircraft registration number
SELECT * 
FROM Accidents 
INNER JOIN Aircrafts ON Aircrafts.id = Accidents.aircraft_id
WHERE Aircrafts.registration_number = 'REG123';
