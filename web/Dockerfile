# Base stage for common dependencies and environment setup
FROM node:lts-alpine as base
WORKDIR /app
COPY package*.json ./
RUN npm install

# Development stage
FROM base as development

COPY . .

EXPOSE 5173

CMD ["npm", "run", "dev"]

# Build stage for compiling the app
FROM base as builder
COPY . .
RUN npm run build

# Production stage
FROM node:lts-alpine as production
WORKDIR /app
COPY --from=builder /app/dist /app/dist
RUN npm install -g http-server
CMD ["http-server", "dist"]
