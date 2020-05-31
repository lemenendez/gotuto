# gotuto

A simple Golang tutorial. It runs using docker

## Branchs list

1. basic-01
2. basic-02
3. basic-03
4. basic-04

## Checkout branch

git checkout branch-name

´git checkout basic-01´

## Build docker image

´docker build -t gotuto . ´

## Running the container

´docker run -it --rm -v "$PWD":/go/gotuto  --name got gotuto bash´

## Running the code

´go run github.com/lemenendez/gotuto/tuto/slice_simple.go´
´go run github.com/lemenendez/gotuto/tuto/slice.go´

## Build the package

´go build ./...´

## Run the binary

´./tuto´
