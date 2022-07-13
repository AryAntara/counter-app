# APP

this an api for counting words,characters,characters without space,syllables,sentences, and paragraph

## INSTALL

first you must be clone this repository

```	bash
git clone https://github.com/aryarte/counter_application
```

## USAGE
### using go 

you must be have `go` is installation on your computer.
download dependencies are usage for this API

```bash 
go mod download 
```

build the application

```bash
go run build .
```

then run the application

```bash
./counter 
```

now your application runing on port `8080`

### using docker and kubernetes

### docker setup
you must be have `docker` installation on your computer.
to run with `kubernetes` or `kubectl` the `API` must be push on `docker hub`.
you have to setup an account, to do it you can read official docs at https://docs.docker.com/docker-hub.

first you need to do is to build this application.

```bash
docker build -t <your_username>/<your_repository> .
```

you can run test after building.

```bash
docker run <your_username>/<your_repository>
```

login with your docker ID to push and pull images from docker hub.

```bash 
docker login
```

then you can push your image to docker.

```bash
docker push <your_username>/<your_repository>
```

### kubernetes setup

you need to change the `counterApplication.yaml` file on line 17.
change the `<your_docker_username>` and `<your_docker_repository>` with your docker username and repository
are used for this application.

finally you can run it 

```bash 
kubectl apply -f counterApplication.yaml
```

To usage this application if has been runing.
you can acces route `/count` with `POST` method.
and text to count in body.

```bash
# test using curl
curl -X POST localhost:8080/count -d "text"
```

it will count the string `text`
