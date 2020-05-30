# gotuto

A simple Golang tutorial

## Branchs list

1. basic-01
2. basic-02
3. basic-03
4. basic-04
5. basic-05
6. basic-06
7. basic-07

## Checkout branch

´git checkout basic-06´

## Build doocker image

´docker build -t gotuto . ´

## Running the container

´docker run -it --rm -v "$PWD":/go/gotuto  --name got gotuto bash´

## Running a the main go file

´go run github.com/lemenendez/gotuto/tuto/main.go github.com/lemenendez/gotuto/tuto/trunc.go ´

## Build the app

´go build github.com/lemenendez/gotuto/tuto/main.go github.com/lemenendez/gotuto/tuto/search.got ´

## Build the package

´go build ./...´

## Run the binary

´./tuto´
