# set base image (host os)
FROM golang:1.18.3-alpine AS builder

#set the working directory in the container
WORKDIR /app

#copy the dependecies file to the working directory
COPY go.mod go.sum ./

#install dependencies
RUN go mod download

# copying the entire content to working directory
COPY . .

#command to run on container start
RUN go build -o main main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 84
CMD [ "/app/main" ]


