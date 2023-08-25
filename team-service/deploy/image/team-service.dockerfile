FROM alpine:latest

RUN mkdir /app

COPY ../../bin/teamApp /app

CMD ["/app/teamApp"]