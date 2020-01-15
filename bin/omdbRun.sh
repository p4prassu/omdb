#!/bin/bash

#########################################################################
# Shell script for omdb module.                          			#
#                                                           	   		#
# Author: Prasad Potipireddi <ppotipir@cisco.com>              		  	#
# Date: Jan 13th, 2020.                                               	#
# Since: CX Cloud Release.                                            	#
# Copyright (c) 2018 Cisco Systems. All rights reserved.              	#
#                                                                      	#
#########################################################################

# Usage:
    # ondbRun.sh <title>

# helpFunction will display the usage of script if user hasn't provided 
# right set of arguments and exit the script
helpFunction()
{
   echo ""
   echo "Usage: $0 title"
   exit 1 # Exit script after printing help
}

# Print helpFunction in case parameters are empty
if [ -z "$1" ]  
then
   echo "Some or all of the parameters are empty";
   helpFunction
fi

# Begin script in case all parameters are correct
echo "title=$1"

# calling actual go main with given title
./bin/main -title=$1