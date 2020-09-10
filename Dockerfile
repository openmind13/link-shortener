FROM golang:latest
WORKDIR /home/app/
COPY ./ /home/app/
RUN go mod download && go build -o apiserver ./cmd/apiserver
EXPOSE 8080
CMD ./apiserver