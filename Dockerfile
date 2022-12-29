FROM golang:1.19.4-alpine3.16 AS build

# Create app directory
WORKDIR /app

# Gather local files and Install app dependencies
COPY . .
RUN go mod download

# Build the binary
RUN go build -o ./restep .

# Stage 2 - Copy only binary file
FROM alpine:latest

WORKDIR /

COPY --from=build /app/restep /app/restep

EXPOSE 8000

CMD ["/app/restep"]
