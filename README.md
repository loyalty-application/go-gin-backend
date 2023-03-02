# go-gin-backend
The `go-gin-backend` repository is a microservice of the `loyalty-application`
It serves as the REST API Server, handling request from the users and applications

## Dependencies
This repository has the follow dependencies:
- [loyalty-application]()
- [Docker](https://docs.docker.com/engine/install/)
- Docker Compose

## Setup
### Cloning the repo
You are advised not to clone this repository by itself, rather you should head over to the main repository and clone this repository recursively as a submodule instead.

### Environment variables
You will need to rename the `example.env` file to `.env` and replace the environment variables with the ones you want to use.

Please take note that all submodules associate with this repo will also have their own `example.env` files that need to be copied, renamed and modified.

## Contributing
To contribute to the repo, refer to the [Development Guide]() located in the `/docs` folder of this repo to understand the structure and development practices 
