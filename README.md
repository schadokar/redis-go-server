# Redis-Go-Server

A prebuild golang server to work with redis db.

## Dependencies

- go-redis github.com/go-redis/redis/v7
- gorilla/mux github.com/gorilla/mux

## Docker Image

The docker image is hosted on [Docker Hub](https://hub.docker.com/r/schadokar/redis-go-server).

## Prerequisite

The Redis DB must be up and running. It can be locally hosted, Redis docker image or cloud-based.

Use docker-compose to pass the Redis DB details to the server.

Set the environment variable of the redis-db:

- REDIS_DB_URL=URL
- REDIS_DB_PASSWORD=Password
- REDIS_DB=db index

**Note** If the password is an empty string `""`, don't send the empty string.

```
REDIS_DB_PASSWORD=
```

It will take that as an empty string.

## Instructions

#### Pull the image

```docker
docker pull schadokar/redis-go-server
```

#### Create a docker-compose file

```
version: "2"

services:
  redis-db:
    image: redis:alpine
    ports:
      - 6379:6379
    container_name: redis-db

  redis-go-server:
    image: schadokar/redis-go-server:latest

    environment:
      - REDIS_DB_URL=redis-db:6379
      - REDIS_DB_PASSWORD=
      - REDIS_DB=0
    ports:
      - 8080:8080
    depends_on:
      - redis-db
    container_name: redis-go-server
```

First, we're pulling the `redis:alpine` image and running it as `redis-db` container.  
The second image is `schadokar/redis-go-server` which is running as `redis-go-server` container.

Check the environment variables.

### Endpoints of the server

There are 3 endpoints of the server:

- POST METHOD **set**  
  It takes a json object with 2 keys: "key" and "value"

```
{
    "key": "strongest-avenger",
    "value": "thor"
}
```

- GET METHOD **get**  
  Return all the keys stored in the Redis DB.
- GET METHOD **get/{key}**  
  Return the value of the key

### Example

- Run the docker-compose file

```
docker-compose -f <file-path> up -d

Output

docker-compose -f ./docker-compose.yaml up -d
Creating network "deployments_default" with the default driver
Creating redis-db ... done
Creating redis-go-server ... done
```

- Save a key-value pair in the Redis DB using `cURL`

```
curl -d '{"key":"strongest-avenger","value":"thor"}' -X POST http://localhost:8080/set

Output

  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    64  100    22  100    42    687   1312 --:--:-- --:--:-- --:--:--  2000  "Successfully Saved!"
```

- Get the value of the key `strongest-avenger`

```
curl -X GET http://localhost:8080/get/strongest-avenger

% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
 100     7  100     7    0     0    148      0 --:--:-- --:--:-- --:--:--   148   "thor"
```

- Get all the keys stored in the Redis DB

```
curl -X GET http://localhost:8080/get

  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    22  100    22    0     0   1466      0 --:--:-- --:--:-- --:--:--  1466   ["strongest-avenger"]
```

# To build the image using the Dockerfile in the deployments

Open the terminal in the redis-go-server directory and run

```docker
docker build -t <tag of the image> -f deployments/Dockerfile .
```

# To run the docker-compose

Open the terminal in the redis-go-server directory and run

```docker
docker-compose -f deployments/docker-compose.yaml up -d
```

# References

[Callicoder](https://www.callicoder.com/docker-golang-image-container-example/)  
[Builder pattern vs. Multi-stage builds in Docker](https://blog.alexellis.io/mutli-stage-docker-builds)
