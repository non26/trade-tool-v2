from golang:1.19.2-bullseye

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o goapp ./cmd/myapp
# RUN go build -o goapp

EXPOSE 8080

CMD [ "./goapp" ]