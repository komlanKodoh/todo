#!/bin/sh

# Clear the output directory if it already exists
rm -rf ./output

# build the front end project
(cd ./frontend && npm run build)


# build the backend project 
(cd ./backend && CGO_ENABLED=0 go build -o ./todo.app) || exit 1
echo "Backend build successful"


# Copy frontend static files to backend folder
cp -r ./frontend/build ./backend/static

mkdir output

# Copy the binary to the output folder
cp -r ./frontend/build/ ./output/static/

# Copy the go binary to the output folder
cp -r ./backend/todo.app ./output/todo.app

echo "Successfuly copied output directory"

# Build dokcer image
docker build -t todo-app . 

docker tag todo-app:latest postrgresslearn/todo-app:latest
docker tag todo-app:latest postrgresslearn/todo-app:1.0.0

docker push postrgresslearn/todo-app --all-tags




