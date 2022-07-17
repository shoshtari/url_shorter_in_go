#!/usr/bin/bash

file_name=$1
port=8080
file_sign="$(sha256sum main.go | awk -F' ' '{print $1 ;}')"

echo $(jobs -p)

go run $file_name &

trap 'fuser '$port/tcp' -k' EXIT

while true
do
	new_sign="$(sha256sum main.go | awk -F' ' '{print $1 ;}')"

	if [[ "$new_sign" != "$file_sign" ]]
	then
		echo "Changed"
		fuser "$port/tcp" -k
		go run $file_name &

		file_sign=$new_sign
	fi
	sleep 1


done
