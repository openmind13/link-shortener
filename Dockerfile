FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get -d -u ./...
RUN go build -o apiserver ./cmd/apiserver
CMD ["/app/apiserver"]

# COPY . .
# ENTRYPOINT [ "apiserver" ]