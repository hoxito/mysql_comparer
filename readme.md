# mysql_comparer

This is a simple mysql databases comparer that shows wich tables are missing from 1 db respect another and returns the sql script needed to sync both dbs.

### Auth

Currently, the auth service is not implemented.

## Requirements

Go 1.18  [golang.org](https://golang.org/doc/install)

## Initial Config
Set Environment Variables

    export GO111MODULE=on
    export GOFLAGS=-mod=vendor

To download the app correctly you must run:

    go get github.com/hoxito/mysql_comparer

Once downloaded you will have the code in folder

    cd $GOPATH/src/github.com/hoxito/mysql_comparer

Or you could just run

    https://github.com/hoxito/mysql_comparer.git

To have the repository

# Installation

To install the application in your local machine:

1- Install required libs

    go mod download
	go mod vendor


2- build and execution

    go install
    mysql_comparer

env file contents:


## API-DOCS
To load api docs you have to run the application and go to

> (http://localhost:3010/swagger/index.html)


## Run Docker Containers
If you wish to run the application in docker containers you should run:

    cd /mysql_comparer
    docker build -t dev-mysqlcomp .
    docker run -it --name dev-mysqlcomp -p 3000:3000 -v $PWD:/go/src/github.com/hoxito/mysql_comparer dev-mysqlcomp

## Testing
To run tests, you must cd into the desired test folder and run:

    go test -v