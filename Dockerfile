FROM golang:latest as build

WORKDIR /go/src/app/
COPY . .
RUN go mod download
RUN go install -mod=mod github.com/githubnemo/CompileDaemon@latest
EXPOSE 8080
ENTRYPOINT CompileDaemon --build="go build -o main ." --command=./main

#RUN go mod download
#RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/web-app ./main.go
#
#FROM alpine:latest
#RUN apk --no-cache add ca-certificates
#WORKDIR /usr/bin
#COPY --from=build /go/src/app/bin /go/bin
#EXPOSE 80
#ENTRYPOINT /go/bin/web-app --port 80