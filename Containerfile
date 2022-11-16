FROM golang:1.19 as builder

WORKDIR /usr/share/app

COPY . .
RUN go mod download && go mod verify
RUN go build

FROM ubuntu

WORKDIR /usr/share/app

COPY index.gohtml .
COPY --from=builder /usr/share/app/cubiks /usr/local/bin/cubiks

CMD ["cubiks"]