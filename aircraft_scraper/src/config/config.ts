import dotenv from 'dotenv';

dotenv.config();

const config = {
  db: {
    host: process.env.MYSQL_HOST || 'localhost',
    user: process.env.MYSQL_USER || 'user',
    password: process.env.MYSQL_PASSWORD || 'password',
    database: process.env.MYSQL_DATABASE || 'airacccidentdata',
    port: parseInt(process.env.MYSQL_PORT || '3306', 10),
  },
  aws: {
    region: process.env.AWS_REGION || '',
    accessKeyId: process.env.AWS_ACCESS_KEY_ID || '',
    secretAccessKey: process.env.AWS_SECRET_ACCESS_KEY || '',
    // Ensure the S3 bucket name matches the subdomain if using AWS for static assets.
    // This allows for proper routing and caching when integrating with Cloudflare.
    s3Bucket: process.env.AWS_S3_BUCKET || '',
  },
  nodeEnv: process.env.NEXT_PUBLIC_ENV || 'development',
  aircraftTypeToImageMapPath: 'src/database/aircraftMapping.json',
};

export default config;
