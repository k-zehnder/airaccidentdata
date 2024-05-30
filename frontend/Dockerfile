# Stage 1: Build the Next.js/React application
FROM node:latest as build
WORKDIR /app

# Copy package files and install dependencies
COPY ./package*.json ./
RUN npm install

# Set environment variable for build
ENV NEXT_PUBLIC_ENV=production

# Copy the rest of the source files and build the application
COPY ./ .
RUN npm run build

# Stage 2: Run the Next.js app
FROM node:latest
WORKDIR /app
COPY --from=build /app .
EXPOSE 3000

# Command to run the app
CMD ["npm", "start"]
