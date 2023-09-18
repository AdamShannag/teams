FROM alpine:latest

RUN mkdir /app

COPY ../../bin/userApp /app

CMD ["/app/userApp"]