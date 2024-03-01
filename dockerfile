FROM alpine:3.14 

RUN apk add --no-cache libc6-compat 
WORKDIR /app

COPY ./output/ .

ENTRYPOINT /app/todo.app || (ls && pwd)
