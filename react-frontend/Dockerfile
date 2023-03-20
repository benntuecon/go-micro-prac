FROM node:alpine
RUN npm install -g create-react-app
WORKDIR /app
COPY package.json .
RUN npm install react-scripts --force 
RUN npm install

COPY . .

EXPOSE 3000
CMD ["npm", "start"]











