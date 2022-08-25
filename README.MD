# GO SIMPLE SERVICE
This is a playground repository to explore and experiment the best folder structure to create golang service for my workflow. I will push each iteration of my folder structure and tag i by the iteration number so that we can see the difference and improvement between iteration.

## Folder explanation
### /cmd
This folder contains the main application entry point files for the project. I may have multiple entry points for the project. For example `cmd/simple-service` will contain the entry point for the api service, while `cmd/simple-service-cli` will contain the entry point for the cli service.

### pkg
This folder contain package/module that can be shared across the service. I personally like to keep all the packages in the `pkg` folder until I find the need to relocate them to other folders.

## GO Installation
golang 1.19 +
For further information, you can follow the procedure in https://golang.org/doc/install

## Environment Variable
Environment variable is included in the .env file

## Databases
Currently have not use any database

## Running DRM Server
Install dependencies and creating vendor directory
```
go mod vendor
```
Run the program
```
go run main.go
```

## Credits
[Mark Wolfe's Blog] (https://www.wolfe.id.au/2020/03/10/how-do-i-structure-my-go-project/)