FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /loyalty-application/go-gin-backend 

EXPOSE 8080

CMD [ "/loyalty-application/go-gin-backend" ]
