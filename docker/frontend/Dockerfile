FROM node

WORKDIR /projects/frontend

COPY frontend/package.json ./
COPY frontend/package-lock.json ./

# RUN npm install @angular-devkit/build-angular
RUN npm install -g @angular/cli
RUN npm install
