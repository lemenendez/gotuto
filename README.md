# gotuto

A simple Golang tutorial

## Branchs list

1. basic-01

## Checkout branch

´git checkout basic-01´

## Build doocker image

´docker build -t gotuto . ´

## Running the container

´docker run -it --rm -v "$PWD":/go/gotuto  --name got gotuto bash´

## Running a the main go file

´go run github.com/lemenendez/gotuto/tuto/main.go github.com/lemenendez/gotuto/tuto/vars.go ´

## Build the app

´go run github.com/lemenendez/gotuto/tuto/main.go github.com/lemenendez/gotuto/tuto/swap.go ´

## Build the package

´go build ./...´

## Run the binary

´./tuto´
