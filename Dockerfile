# build stage
FROM golang:1.18.3-alpine3.16 AS build-stage

WORKDIR /classroom-api

ADD ./src ./src
COPY ./go.mod .
COPY ./go.sum .
 
RUN ["go", "build", "./src/main/main.go"]

# run stage
FROM alpine

WORKDIR /api

COPY --from=build-stage /classroom-api/main /api

EXPOSE 3333
EXPOSE 5003

CMD ["./main"]
