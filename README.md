# Golang Goban

An Goban webapp written in Golang

## Getting Started

This is a Docker app all you need to do is install the prereqs below then start the app as instructed

### Prerequisites

Docker
Docker-Compose

### Starting

Right now the app starts before the db is healthy, I need to add a health check to the docker-compose. In the meantime do:

docker-compose up -d db

docker-compose logs -f db //checking for "server is ready for connections"

docker-compose up app