# Build and Test stage
FROM node:20-alpine AS builder

WORKDIR /usr/src/app

COPY package*.json ./

# Install ALL dependencies (including devDependencies like Jest)
RUN npm install

COPY . .

# Run unit and integration tests
# If this fails, the Docker build will stop and fail
RUN npm test

# Production stage
FROM node:20-alpine

WORKDIR /usr/src/app

COPY package*.json ./

# Install ONLY production dependencies
RUN npm install --omit=dev

# Copy only the necessary source code from the builder stage
COPY --from=builder /usr/src/app/src ./src

EXPOSE 4000

CMD ["npm", "start"]
