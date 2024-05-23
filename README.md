# airaccidentdata

## Quickstart

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/k-zehnder/airaccidentdata.git
   cd airaccidentdata
   ```

2. **Configure Environment:**

   Copy the example environment file:

   ```bash
   cp .env.example .env
   ```

   **Edit the `.env` file:**

   Open the `.env` file and set your environment variables:

   ```dotenv
   # MySQL Configuration
   MYSQL_ROOT_PASSWORD=password
   MYSQL_DATABASE=airaccidentdata
   MYSQL_USER=user
   MYSQL_PASSWORD=password
   MYSQL_HOST=mysql
   MYSQL_PORT=3306

   # Backend Configuration
   GO_ENV=development
   SERVER_ADDRESS=0.0.0.0:8080

   # AWS Configuration (for aircraft_scraper service, needed for production environment only)
   AWS_REGION=your-region
   AWS_ACCESS_KEY_ID=your-access-key-id
   AWS_SECRET_ACCESS_KEY=your-secret-access-key
   AWS_S3_BUCKET=your-s3-bucket

   # Frontend Configuration
   NEXT_PUBLIC_ENV=development

   # Google Maps API Configuration
   GOOGLE_MAPS_API_KEY=your-google-maps-api-key

   # FAA Accident Data Configuration
   CSV_FILE_PATH=downloaded_file.csv
   ```

3. **Ensure Docker is Installed and Running:**

   Make sure Docker is installed and running on your host machine. You can download Docker from [here](https://www.docker.com/products/docker-desktop).

4. **Launch Development Environment with Docker:**

   This will build and start all necessary services:

   ```bash
   make dev
   ```

5. **Populate the Database with Accident Data:**

   ```bash
   cd backend
   make fetch-data
   cd ..
   ```

6. **Populate the Database with Aircraft Images:**

   ```bash
   cd aircraft_scraper
   make fetch-images
   cd ..
   ```

   Your development environment should now be running.

## Accessing the Application

- **Frontend:** Visit `http://localhost:3000` to view the frontend.
- **Swagger UI:** Access the API documentation at `http://localhost:8080/swagger/index.html`.
