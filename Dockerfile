FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o apiserver ./cmd/apiserver
CMD ["/app/apiserver"]

# COPY . .
# ENTRYPOINT [ "apiserver" ]