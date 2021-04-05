FROM ubuntu
WORKDIR /app
COPY main /app/main
COPY data/data.json /app/data/data.json
CMD ["/app/main"]