FROM golang:latest

WORKDIR /app/

COPY . .

RUN go mod download -x

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

EXPOSE ${APP_PORT}

ENTRYPOINT CompileDaemon --build="go build -o main cmd/main.go" --command="./main"