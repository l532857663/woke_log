#!/bin/bash

set -e

declare -a arr=('./server' './service' './dao' './model' './utils')
for i in "${arr[@]}"                                                                                                                                                                           
do
	OUTPUT="$(gofmt -w $i)"
	if [[ $OUTPUT ]]; then
		echo "The following files contain goimports errors"
		echo $OUTPUT
		echo "The gofmt command must be run for these files"
		exit 1
	fi
done
