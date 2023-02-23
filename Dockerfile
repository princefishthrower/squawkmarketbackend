FROM golang:alpine AS build-env

RUN apk add build-base

WORKDIR /app

COPY . .

RUN CGO_ENABLED=1 go build -o main .

FROM alpine

WORKDIR /app

# Copy this app's .env file to the container's /app directory
COPY .env /app/

# also copy in the DB!
COPY squawkmarketbackend.db /app/

# also the google cloud credentials
COPY squawk-market-credentials.json /app/

COPY --from=build-env /app/main /app/

# also add the timezone data
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip

ENV ZONEINFO /zoneinfo.zip

CMD ["./main"]
