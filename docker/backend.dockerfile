FROM golang:1.13.1-alpine3.10 as build
ENV GO111MODULE on
RUN apk add --update --no-cache git 

RUN mkdir -p /go/src/backend
WORKDIR /go/src/backend
COPY . .
RUN go build

FROM alpine:3.10
COPY --from=build /go/src/backend/backend /bin/backend

ENTRYPOINT ["/bin/backend"]

