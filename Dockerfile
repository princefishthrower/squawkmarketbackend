FROM golang:alpine AS build-env

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine

WORKDIR /app

# Copy this app's .env file to the container's /app directory
COPY .env /app/

COPY --from=build-env /app/main /app/

# also add the timezone data
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip

ENV ZONEINFO /zoneinfo.zip

CMD ["./main"]
