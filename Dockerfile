FROM golang:1.19 AS build

WORKDIR /app

COPY / ./

RUN go build -o /main

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /main /main

USER nonroot:nonroot

ENTRYPOINT ["./main"]