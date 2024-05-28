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
   MYSQL_ROOT_USER=root
   MYSQL_ROOT_PASSWORD=password
   MYSQL_DATABASE=airaccidentdata
   MYSQL_USER=user
   MYSQL_PASSWORD=password
   MYSQL_HOST=mysql
   MYSQL_PORT=3306

   # Backend Configuration
   GO_ENV=development
   SERVER_ADDRESS=0.0.0.0:8080

   # AWS Configuration (production only, for Cloudflare caching with S3 bucket)
   AWS_REGION=your-region
   AWS_ACCESS_KEY_ID=your-access-key-id
   AWS_SECRET_ACCESS_KEY=your-secret-access-key
   AWS_S3_BUCKET=your-s3-bucket

   # Frontend Configuration
   NEXT_PUBLIC_ENV=development

   # Google Maps API Configuration
   GOOGLE_MAPS_API_KEY=your-google-maps-api-key
   ```

3. **Obtain Google Maps Geocoding API Key:**

   To get the coordinates for accidents, you need a Google Maps Geocoding API key. Follow the instructions [here](https://developers.google.com/maps/documentation/geocoding/get-api-key) to obtain and configure your API key. Then, add it to your `.env` file:

   ```dotenv
   GOOGLE_MAPS_API_KEY=your-google-maps-api-key
   ```

4. **Ensure Docker is Installed and Running:**

   Make sure Docker is installed and running on your host machine. You can download Docker Desktop from [here](https://www.docker.com/products/docker-desktop).

   Alternatively, you can install Docker via the command line:

   For **Ubuntu**:

   ```bash
   sudo apt-get update
   sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common
   curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
   sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
   sudo apt-get update
   sudo apt-get install -y docker-ce
   sudo systemctl status docker
   ```

   For **Mac**:

   ```bash
   brew install docker
   brew install docker-compose
   ```

   For **Windows**:

   You can download Docker Desktop for Windows from [here](https://www.docker.com/products/docker-desktop) and follow the installation instructions provided on the website.

5. **Launch Development Environment with Docker:**

   This will build and start all necessary services:

   ```bash
   make dev
   ```

6. **Populate the Database with Accident Data:**

   ```bash
   cd backend
   make data
   cd ..
   ```

7. **Populate the Database with Aircraft Images:**

   ```bash
   cd aircraft_scraper
   make images
   cd ..
   ```

   Your development environment should now be running.

## Accessing the Application

- **Frontend:** Visit `http://localhost:3000` to view the frontend.
- **Swagger UI:** Access the API documentation at `http://localhost:8080/swagger/index.html`
