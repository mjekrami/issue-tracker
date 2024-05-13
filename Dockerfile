FROM golang:1.22.3-bullseye

# Copy all the necessary files
COPY main.go /app/
COPY go.mod /app/
COPY .env /app/
COPY config/ /app/config
COPY cmd/ /app/cmd

WORKDIR /app/
RUN go get && go build -o issue-tracker
CMD ["./issue-tracker"]



