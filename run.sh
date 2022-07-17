#!/usr/bin/bash
function checksum(){
  goFiles=$(find . -name "*.go" | sort)
  for i in ${goFiles[@]}
  do
    tmp=$(sha256sum $i | awk -F' ' '{print $1}')

    if [[ -z $finalHash ]]
    then
      finalHash=$tmp
    else
      finalHash=$(echo $finalHash $tmp | sha256sum | awk -F' ' '{print $1}')
    fi
  done
  echo $finalHash
}
port=$PORT
if [[ -z $port ]]
then
  port=8080
fi

file_name=$1
file_sign=$(checksum)




go run $file_name &

trap 'fuser '$port/tcp' -k' EXIT

while true
do
	new_sign=$(checksum)

	if [[ "$new_sign" != "$file_sign" ]]
	then
		echo "Changed"
		fuser "$port/tcp" -k
		go run $file_name &

		file_sign=$new_sign
	fi
	sleep 1


done
