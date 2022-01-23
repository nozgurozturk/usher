<p align="center"><a href="#"><img src="docs/logo.png" alt="usher" width="390" height="135" /></a></p>

## About Usher

Usher is a seating allocation program that can find the most efficient arrangement for a group of people.

*About the name*
> [usher](https://www.oxfordlearnersdictionaries.com/definition/english/usher_1): a person who shows people where to sit in a church, public hall, etc.


### Install

You need to [download](https://golang.org/dl/) and install Go. `1.17` or higher is required for drone navigation
service.

When installation is completed, clone the repo with command:

    git clone https://github.com/nozgurozturk/usher

After the installation, download missing dependencies with command:

    go mod download

### Runing

    make start

### Testing

*integration:*

    make test_integration

*unit:*

    make test_unit


### Build

    make build

### From Dockerfile

    docker build --rm -f ./build/Dockerfile -t usher .

    docker run -it -p 8081:8081 usher
