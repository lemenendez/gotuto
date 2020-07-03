# gotuto

A simple Golang tutorial. It runs using docker

## Branchs list

1. basic-01
2. basic-02
3. basic-03
4. basic-04
5. basic-05
6. basic-06
7. basic-07
8. basic-08
9. basic-09
10. int-01
11. int-02
12. int-03
13. int-04
14. int-05

## Checkout branch

git checkout branch-name

´git checkout basic-01´

## Build docker image

´docker build -t gotuto . ´

## Running the container

´docker run -it --rm -v "$PWD":/go/gotuto  --name got gotuto bash´

## Running the code

´go run sortv0.go´

## Build the package

´go build ./...´

## Run the binary

´./tuto´
