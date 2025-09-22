FROM node:20-alpine AS builder

WORKDIR /app

COPY package*.json ./
# Если используете pnpm/yarn: замените на соответствующие команды
RUN npm ci --silent

COPY . .

RUN npm run build