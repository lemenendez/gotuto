# gotuto

A simple Golang tutorial

## Branchs list

1. basic-01

## checkout branch

´git checkout basic-01´ 

## build doocker image

´docker build -t gotuto . ´

## running the container

´docker run -it --rm -v /go/gotuto:"$PWD"  --name got gotuto bash´
