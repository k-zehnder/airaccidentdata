import dotenv from 'dotenv';

dotenv.config();

const config = {
  nodeEnv: process.env.NEXT_PUBLIC_ENV || 'development',
};

export default config;
