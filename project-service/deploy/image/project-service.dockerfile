FROM alpine:latest

RUN mkdir /app

COPY ../../bin/projectApp /app

CMD ["/app/projectApp"]