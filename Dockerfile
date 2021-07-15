##
## Build
##

FROM golang:1.16-buster AS build

WORKDIR /app

COPY . .

RUN go mod download


RUN go build -o /docker-hashGen

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /docker-hashGen /docker-hashGen

USER nonroot:nonroot

ENTRYPOINT ["/docker-hashGen"]