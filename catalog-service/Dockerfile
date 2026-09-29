FROM golang:1.24-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /out/catalog-service .

FROM alpine:3.20
COPY --from=build /out/catalog-service /usr/local/bin/catalog-service
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/catalog-service"]
