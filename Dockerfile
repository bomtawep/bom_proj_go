FROM golang:latest as build

WORKDIR /go/src/app/
COPY . .
COPY .env.production .env
RUN go mod download
RUN go install -mod=mod github.com/githubnemo/CompileDaemon@latest
EXPOSE 8080
ENTRYPOINT CompileDaemon --build="go build -o main ./cmd" --command=./main

FROM nginx:latest
COPY --from=build /go/src/app/main /usr/share/nginx/html
EXPOSE 80