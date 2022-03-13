FROM golang:1.17-alpine AS build
ADD . /ted
ENV CGO_ENABLED=0
WORKDIR /ted
RUN go build -o ted.bin ./cmd/ted

FROM alpine:latest
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

COPY --from=build /ted/ted.bin /ted/ted.bin
#COPY ./keys/private.key /keys/private.key
#COPY ./keys/public.key /keys/public.key
EXPOSE 9999

