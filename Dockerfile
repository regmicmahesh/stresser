FROM golang:alpine AS build
WORKDIR /app
COPY . .
RUN go build -o . ./...

FROM alpine
WORKDIR /app
COPY --from=build /app/stresser .
ENTRYPOINT ["/app/stresser"]
CMD ["-h"]