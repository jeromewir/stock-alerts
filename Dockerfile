FROM golang:1.15.6-alpine AS build
RUN apk add ca-certificates
WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -o /bin/stocksalerts
ENTRYPOINT ["/bin/stocksalerts"]
