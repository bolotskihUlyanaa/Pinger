FROM node:13.12.0-alpine
WORKDIR /web
COPY package*.json ./
RUN npm install
COPY . .
CMD ["npm", "start"]