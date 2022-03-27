#build stage
FROM golang:alpine 
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .

ENTRYPOINT /app
LABEL Name=dockerfiletest Version=0.0.1
EXPOSE 8080
