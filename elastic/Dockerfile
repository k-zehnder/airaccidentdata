# Use the official Node.js image as the base
FROM node:latest

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Install ts-node globally
RUN npm install -g ts-node typescript

# Keep the container running
CMD ["tail", "-f", "/dev/null"]
