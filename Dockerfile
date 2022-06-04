FROM  golang:1.18.3-alpine3.16

WORKDIR /classroom-api

COPY . .

EXPOSE 3333

CMD ["go", "run", "./src/main/main.go"]
