FROM scratch

ADD ca-certificates.crt /etc/ssl/certs/

#CGO_ENABLED=0 GOOS=linux go build -a --ldflags="-s" --installsuffix cgo -o webapp

ADD webapp /
EXPOSE 80
ENTRYPOINT ["/webapp"]