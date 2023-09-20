FROM node:20-alpine As builder

WORKDIR /usr/src/app

COPY package.json package-lock.json ./

RUN npm install

COPY . .

RUN npm run build -- --configuration production

FROM nginx:1.17.1-alpine

COPY --from=builder /usr/src/app/dist/frontend/ /usr/share/nginx/html
COPY --from=builder /usr/src/app/nginx/default.conf /etc/nginx/conf.d/default.conf
