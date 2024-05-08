# FROM golang:1.22.3-alpine3.19

# WORKDIR /app

# # Download Go modules
# COPY . .
# RUN go mod download


# RUN CGO_ENABLED=0 GOOS=linux go build -o api-server ./cmd/api-server

# EXPOSE 8080

# CMD ["./api-server"]

FROM golang:1.22.3-alpine3.19 as build
WORKDIR /src

RUN apk --update add --no-cache ca-certificates openssl git tzdata && \
update-ca-certificates

COPY . .
# COPY ca-bundle.crt /etc/ssl/certs/ca-bundle.crt
# COPY ca-bundle.trust.crt /etc/ssl/certs/ca-bundle.trust.crt 
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api-server ./cmd/api-server

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /bin/api-server /bin/api-server

EXPOSE 1323
CMD ["/bin/api-server"]