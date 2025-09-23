FROM node:20-alpine AS builder

WORKDIR /app

COPY package*.json ./
# Если используете pnpm/yarn: замените на соответствующие команды
RUN npm ci --silent

COPY . .

RUN npm run build

FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80