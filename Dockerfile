FROM golang:1.19.4-alpine3.16 AS build

# Create app directory
WORKDIR /app

# Gather local files, install app dependencies, and build the binary
COPY . .
RUN go mod download && go build -o ./restep .

# Stage 2 - Copy only binary file and run
FROM alpine:3.16

WORKDIR /

COPY --from=build /app/restep /app/restep

EXPOSE 8000

CMD ["/app/restep"]
