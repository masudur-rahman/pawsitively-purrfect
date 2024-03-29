FROM ghcr.io/masudur-rahman/golang:1.20

WORKDIR /pawsitively-purrfect

COPY . .
RUN go mod tidy && go mod vendor
RUN go build -o pawsitively-purrfect

#USER nobody:nobody
USER 65535:65535

ENTRYPOINT ["./pawsitively-purrfect"]
