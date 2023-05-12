FROM golang:1.20.1-alpine

WORKDIR /app

COPY go.mod ./
#COPY go.sum ./
#RUN go mod download

COPY *.go ./

#Build application and set specific name for the app
RUN go build -o /go-ci-test

EXPOSE 8088

CMD [ "/go-ci-test" ]