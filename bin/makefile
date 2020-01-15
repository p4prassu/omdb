#########################################################################
# Docker Makefile for omdb module.                          			#
#                                                           	   		#
# Author: Prasad Potipireddi             		  	                    #
# Date: Jan 13th, 2020.                                               	#
# Since: CX Cloud Release.                                            	#
# Copyright (c) 2018 Cisco Systems. All rights reserved.              	#
#                                                                      	#
#########################################################################


# Usage: 
	# LOCAL_LIB_PATH is helping to build from the customer path
# build: 
	# below make build example command is used for building the docker. 
	# make build -f ./bin/makefile -e "LOCAL_LIB_PATH=$GOPATH/intro_go/omdb"
# run: 
	# make run example command is used for running docker
	# make run -f ./bin/makefile -e "title=frozen"


# General.
RELEASE_VERSION:=0.0.1

# Docker build-related.
DOCKER_REPO_AWS:=docker_repo/cxc/omdb
DOCKER_IMAGE:=pprassu_omdb
DOCKER_VERSION:=latest

all: build

.GO: build

build:
	docker build -t $(DOCKER_REPO_AWS)/$(DOCKER_IMAGE):$(DOCKER_VERSION) -f ${LOCAL_LIB_PATH}/docker/dockerfile ${LOCAL_LIB_PATH}
		#docker build -t omdb_docker -f ${LOCAL_LIB_PATH}/docker/Dockerfile ${LOCAL_LIB_PATH}
push:
	#docker push --disable-content-trust $(DOCKER_REPO_AWS)/$(DOCKER_IMAGE):$(DOCKER_VERSION)
		#docker push containers.cisco.com/ppotipir/omdb_docker:latest
run:
		docker run --privileged -it $(DOCKER_REPO_AWS)/$(DOCKER_IMAGE):$(DOCKER_VERSION) ${title}

