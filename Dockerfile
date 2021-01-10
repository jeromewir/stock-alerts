FROM golang:1.15.6-alpine AS build

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -o /bin/stocksalerts

FROM scratch
COPY --from=build /bin/stocksalerts /bin/stocksalerts
ENTRYPOINT ["/bin/stocksalerts"]