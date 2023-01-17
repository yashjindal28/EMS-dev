# set base image (host os)
FROM node:14-alpine as build-step

#set the working directory in the container
WORKDIR '/app'

#copy the dependecies file to the working directory
COPY ./package.json ./

#install dependencies
RUN npm install

# copying the entire content to working directory
COPY . .

EXPOSE 4201

#command to run on container start
CMD ["npm","run","start"]
#RUN npm run build --prod



#stage 2
#FROM nginx:stable-alpine
#COPY --from=build-step /app/dist/angular-EMS-frontend /usr/share/nginx/html

#EXPOSE 4201

