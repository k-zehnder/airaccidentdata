# Stage 1: Build the Next.js/React application
FROM node:latest as build
WORKDIR /app
COPY ./package*.json ./
RUN npm install
COPY ./ .
RUN npm run build

# Stage 2: Run the Next.js app
FROM node:latest
WORKDIR /app
COPY --from=build /app .
EXPOSE 3000
CMD ["npm", "start"]
