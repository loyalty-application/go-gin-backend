FROM golang:1.20

WORKDIR /app

ENV GIN_MODE=release

ENV CGO_ENABLED 1

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . ./

#RUN wget -O - https://truststore.pki.rds.amazonaws.com/global/global-bundle.pem > ./rds-combined-ca-bundle.pem

RUN go build -o /loyalty-application/go-gin-backend

EXPOSE 8080

CMD [ "/loyalty-application/go-gin-backend" ]
