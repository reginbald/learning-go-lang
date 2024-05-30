FROM golang:1.22-alpine

WORKDIR /app

ARG some=shit
# RUN go install github.com/cosmtrek/air@latest
RUN apk add git
RUN git clone https://github.com/reginbald/cosmtrek-air.git
RUN cd cosmtrek-air; go install .
RUN go install github.com/a-h/templ/cmd/templ@latest  

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]