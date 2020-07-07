# Configure Jenkins with Docker Compose

## Prepare network and volumes

Command line as follows

```shell
docker network create jenkins
docker volume create jenkins-docker-certs
docker volume create jenkins-data
```

## Check conifgure with docker-compose

Docker compose command line

```shell
docker-compose config
networks:
  jenkins:
    external:
      name: jenkins
services:
  jenkins:
    container_name: jenkins-blueocean
    depends_on:
    - jenkins-docker
    environment:
      DOCKER_CERT_PATH: /certs/client
      DOCKER_HOST: tcp://docker:2376
      DOCKER_TLS_VERIFY: 1
    image: jenkinsci/blueocean
    networks:
      jenkins: null
    ports:
    - published: 8080
      target: 8080
    - published: 50000
      target: 50000
    volumes:
    - jenkins-data:/var/jenkins_home:rw
    - jenkins-docker-certs:/certs/client:ro
  jenkins-docker:
    container_name: jenkins-bocker
    environment:
      DOCKER_TLS_CERTDIR: /certs
    image: docker:dind
    networks:
      jenkins:
        aliases:
        - docker
    ports:
    - published: 2376
      target: 2376
    privileged: true
    volumes:
    - jenkins-docker-certs:/certs/client:rw
    - jenkins-data:/var/jenkins_home:rw
version: '3.3'
volumes:
  jenkins-data:
    external: true
  jenkins-docker-certs: {}
```

## Run Jenkins service  in daemon mode

docker-compose command line

```shell
docker-compose up -d
Recreating jenkins-bocker ... done
Creating jenkins-blueocean ... done
```

Since container `jenkins-bocker` is a running one, docker compose will use `Recreating` instead of `Creating` with its specail running mechanism.

## Destroy Jenkins container

Docker compose command line

```shell
docker-compose down
Stopping jenkins-blueocean ... done
Stopping jenkins-bocker    ... done
Removing jenkins-blueocean ... done
Removing jenkins-bocker    ... done
Network jenkins is external, skipping
```

The destroy process will skip `Network` stage since `Network` is  extenal source. Please vist docker official document to find more information.

## Jenkins Configure Plugin mirror

Configure orginal configure file from container and modify it as file `hudson.model.UpdateCenter.xml` under the smae folder.

```shell
# copy configure file from source
docker cp jenkins-blueocean:/var/jenkins_home/hudson.model.UpdateCenter.xml hudson.model.UpdateCenter.xml
# copy the modified one to container
docker cp  hudson.model.UpdateCenter.xml jenkins-blueocean:/var/jenkins_home/hudson.model.UpdateCenter.xml


```

## Jenkins service status

Jenkins container IPAddress

```shell
docker inspect --format='{{.NetworkSettings.Networks.jenkins.IPAddress}}'  jenkins-blueocean
```

## Reference

1. [Downloading and running jenkins in docker](https://www.jenkins.io/doc/book/installing/#downloading-and-running-jenkins-in-docker)
2. [Docker compose reference and guidelines](https://docs.docker.com/compose/compose-file/#external)
