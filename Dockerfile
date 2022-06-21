FROM alpine:latest
RUN mkdir /app
COPY ./app /app
WORKDIR /app
ENTRYPOINT ["./stop_horny"]
