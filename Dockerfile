FROM golang:1.21.6

WORKDIR /app

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go run internal/gen.go

EXPOSE 1323
