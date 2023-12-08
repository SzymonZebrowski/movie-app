FROM node:18-alpine as builder

WORKDIR /app
COPY . .

RUN npm ci --silent

EXPOSE 3000
CMD ["npm", "start"]
