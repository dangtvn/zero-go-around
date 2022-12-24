#!/bin/bash

docker stop api
docker stop service_user

docker rm api
docker rm service_user

docker rmi api:v1
docker rmi service_user:v1

docker build -t api:v1 -f api/Dockerfile .
docker build -t service_user:v1 -f rpc/service/user/Dockerfile .

docker run -itd --net=host --name=api api:v1
docker run -itd --net=host --name=service_user service_user:v1