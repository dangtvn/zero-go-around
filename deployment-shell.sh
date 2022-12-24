#!/bin/bash

docker stop api

docker rm api

docker rmi api:v1

docker build -t api:v1 -f api/Dockerfile .

docker run -itd --net=host --name=api api:v1