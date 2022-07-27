# simple-go-container
A simple go application built on a SLEBCI Golang container image

## How to build

1. Clone this repo from github
2. Go into the `simple-go-container` directory
3. `docker build -t hello-from-slebci .`

## How to test

1. `docker images`
2. To start the container in the background, `docker run --rm -d -p 8080:8080/tcp --name hello-from-slebci hello-from-slebci:latest`
3. `docker ps`
4. `w3m -dump http://localhost:8080/hello`
5. To stop the running container, `docker stop hello-from-slebci`

## How to deploy

### To Docker Hub

NOTES:
- USERNAME is Docker Hub user account
- X is the version of the s390x-hello container

1. `docker login --username USERNAME`
2. `docker tag hello-from-slebci:latest USERNAME/s390x-hello:vX`
3. `docker push USERNAME/s390x-hello:vX`
4. `docker logout`
