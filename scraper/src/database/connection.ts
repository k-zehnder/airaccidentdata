import mysql, { RowDataPacket } from 'mysql2/promise';

// Define the Database interface
export interface Database {
    getAircraftTypes(): Promise<{ make: string, model: string, registrationNumber: string }[]>;
    insertAircraftImage(aircraftId: number, imageUrl: string): Promise<void>;
    getAircraftIdByType(make: string, model: string, registrationNumber: string): Promise<number | null>;
    getAllAircraftImages(): Promise<{ aircraftId: number; imageUrl: string }[]>;
    updateAircraftImageWithS3Url(aircraftId: number, imageUrl: string, s3Url: string): Promise<void>;
    close(): Promise<void>;
}

// Factory function to create a database connection
export const createDatabaseConnection = async (config: any): Promise<Database> => {
    // Create a connection to the MySQL database
    const connection = await mysql.createConnection({
        host: config.db.host,
        user: config.db.user,
        password: config.db.password,
        database: config.db.database,
    });

    // Function to get aircraft types from the database
    const getAircraftTypes = async (): Promise<{ make: string, model: string, registrationNumber: string }[]> => {
        // Execute the query to select all aircraft make, model, and registration number
        const [rows] = await connection.query<RowDataPacket[]>(
            'SELECT aircraft_make_name, aircraft_model_name, registration_number FROM Aircrafts',
        );

        return rows.map(row => ({
            make: row.aircraft_make_name,
            model: row.aircraft_model_name,
            registrationNumber: row.registration_number,
        }));
    };

    const updateAircraftImageWithS3Url = async (aircraftId: number, imageUrl: string, s3Url: string): Promise<void> => {
        console.log(`Updating image for aircraft ID ${aircraftId} with S3 URL...`);
        
        try {
            // Execute the query to update the image URL with the S3 URL
            await connection.execute(
                'UPDATE AircraftImages SET s3_url = ? WHERE aircraft_id = ? AND image_url = ?',
                [s3Url, aircraftId, imageUrl],
            );
            
            console.log('Image updated with S3 URL successfully.');
        } catch (error) {
            console.error('Error updating image with S3 URL:', error);
            throw error;
        }
    };
    
    // Function to insert an image URL and its corresponding S3 URL for an aircraft into the database
    const insertAircraftImage = async (
        aircraftId: number,
        imageUrl: string,
    ): Promise<void> => {
        console.log('Inserting image into database:');
        console.log('Aircraft ID:', aircraftId);
        console.log('Image URL:', imageUrl);
    
        try {
            // Execute the query to insert the image URLs
            await connection.execute(
                'INSERT INTO AircraftImages (aircraft_id, image_url) VALUES (?, ?)',
                [aircraftId, imageUrl],
            );
            console.log('Image inserted into the database successfully.');
        } catch (error) {
            console.error('Error inserting image into database:', error);
            throw error;
        }

    };

    const getAllAircraftImages = async (): Promise<{ aircraftId: number; imageUrl: string }[]> => {
        try {
            const [rows] = await connection.query<RowDataPacket[]>(
                'SELECT aircraft_id, image_url FROM AircraftImages'
            );
    
            return rows.map(row => ({
                aircraftId: row.aircraft_id,
                imageUrl: row.image_url,
            }));
        } catch (error) {
            console.error('Error fetching all aircraft images:', error);
            throw error;
        }
    };
    
    // Function to get the aircraft ID by type
    const getAircraftIdByType = async (make: string, model: string, registrationNumber: string): Promise<number | null> => {
        const query = 'SELECT id FROM Aircrafts WHERE aircraft_make_name = ? AND aircraft_model_name = ? AND registration_number = ? LIMIT 1';
        console.log(`Executing query: ${query} with make: ${make}, model: ${model}, registrationNumber: ${registrationNumber}`);
        const [rows] = await connection.query<RowDataPacket[]>(query, [make, model, registrationNumber]);
        
        if (rows.length > 0) {
            console.log(`Found aircraft ID: ${rows[0].id} for make: ${make}, model: ${model}, registrationNumber: ${registrationNumber}`);
            return rows[0].id;
        } else {
            console.log(`No aircraft found for make: ${make}, model: ${model}, registrationNumber: ${registrationNumber}`);
            return null;
        }
    };

    // Function to close the database connection
    const close = async (): Promise<void> => {
        await connection.end();
    };

    // Return an object with the database functions
    return {
        getAircraftTypes,
        insertAircraftImage,
        getAircraftIdByType,
        getAllAircraftImages,
        updateAircraftImageWithS3Url,
        close,
    };
};
