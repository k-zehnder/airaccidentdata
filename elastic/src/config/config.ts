import dotenv from 'dotenv';

// Load environment variables from .env file
dotenv.config();

const config = {
  mysql: {
    host: process.env.MYSQL_HOST || 'localhost',
    user: process.env.MYSQL_USER || 'root',
    password: process.env.MYSQL_PASSWORD || 'password',
    database: process.env.MYSQL_DATABASE || 'airaccidentdata',
    waitForConnections: true,
    connectionLimit: 10,
    queueLimit: 0,
  },
  elasticsearch: {
    host: process.env.ELASTICSEARCH_HOST || 'http://localhost:9200',
    apiKey: process.env.ELASTICSEARCH_API_KEY || '',
  },
};

export default config;
