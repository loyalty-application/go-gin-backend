# Development Guide
Before contributing to the project, please read this document to understand how the project is organised.

## Folder Structure 
The project consists of the following folders:
- `/models` Models to represent the REST API requests and database structure
- `/collections` Contains the MongoDB data layer's abstraction
- `/services` Business Logic layer that abstracts the calculations
- `/middlewares` Middleware for the REST API
- `/controllers` Controller layer, only logic for controller should be here, other things belong in other layers
- `/config` Configuration files to load things like env variables and db connection
- `/kafka` Kafka Producers
- `/docs` Swagger API documentation and other docs related to this microservice

## Branches
In this project, the `main` branch contains the production ready and tested code.

To start writing code and contributing to the project, please checkout to the `development` branch with the following command:
```bash
git checkout development
```
After which, you should branch out from `development` into your own branch, following any of the convetions below:
1. `<name>/development`
2. `<feature>/development`

For example, if you're currently working on the authentication feature, you may want to branch out to `auth/development` by doing:
```bash
# you need to be on the development branch first before executing the command
git checkout -b auth/development
```

You may want to merge your branch back into `development` frequently to check if it works or `fast-forward` your branch to the latest `development` commits
```bash
# while on your own branch i.e auth/development, run the following command to fast-forward
git merge development
```

After you've finished development of your feature, merge your feature branch back into `development` by doing:
```bash
# switch back to development branch
git checkout development
# replace auth/development with your own branch
git merge auth/development
```

When you've tested the code sufficiently and ensured its working as intended, merge it from the `development` branch into `main` branch by doing:
```bash
# switch to main branch
git checkout main
# merge the development branch into main branch
git merge development
```
Do note that before you do this, you should check with others working on this branch to ensure that the updates they've merged into `development` are tested and ready to be merged into `main`

## Docker
The project contains two seperate Dockerfile(s), namely `Dockerfile` and `Dockerfile.dev`.
The `Dockerfile` is for production use while the `Dockerfile.dev` is for development use.

In the development phase, we make use of `air` to monitor the files for changes in order to perform live reloading.
After each live reload, it also makes use of the `swag` command to generate new `swagger` documentation based on the annotations written.

You will not need to handle any of these manually as the scripts have already been written.
