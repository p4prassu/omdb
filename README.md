# OMDB Movie Ratings

omdb package is a GO package which provides simple and fast way of getting moving ratings from OMDB web site.

## Getting Started
To start using omdb package, install 
* go 
* docker

Run  go get command

    $ go get github.com/<>

## Build Steps

This project structure as follows 

        omdm
        +----bin
        |    +----makefile 
        |    +----omdbRun.sh
        +----conf
        |    +----config.yml
        +----docker
        |    +----dockerfile
        +----utils
        |    +----utils.go
        |    +----utils_test.go
        +----main.go
        +----go.mod
        +----go.sum
        +----ReadMe.md

## Building & Run docker:

 1) Export environemtn variable for local project path

        $export LOCAL_LIB_PATH = <project local path> // Ex. ${HOME}/omdb

 2) Create docker:

        $make build -f ${LOCAL_LIB_PATH}/bin/makefile -e "${LOCAL_LIB_PATH}"

 3) Run docker:

        $make run -f ${LOCAL_LIB_PATH}/bin/makefile -e "title=<movie title>" // ex. frozen

## Compile from source code and Run test

   1) Build using go build command from project directory

            $ cd ${LOCAL_LIB_PATH}
            $ go build -o main .
   2) Run 

            $ ./main -title=<title> // Ex. frozen  

   3) Test

            $ go test ./utils -v

   4) Code Coverage

            $ go test ./utils -v -cover

    

        
