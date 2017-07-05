FROM alpine:3.5

COPY gobase /app/

COPY files/ /app/files/
RUN ls -la /app/files/*

WORKDIR app

CMD ["./gobase"]