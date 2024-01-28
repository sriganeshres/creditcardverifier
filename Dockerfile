FROM golang:1.21.6-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN go build -o /godocker ./cmd

EXPOSE 8000

CMD [ "/godocker" ]


