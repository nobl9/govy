#!/usr/bin/env bash

if [ "$#" -ne 1 ] || [ -z "$1" ]; then
	# Display an error message and exit
	echo "Usage: $0 <repo_name>"
	exit 1
fi

REPO="$1"

i=0
while IFS= read -r line; do
	((i++))
	if [ "$i" -eq 1 ]; then
		continue
	fi
	IFS=',' read -ra values <<<"$line"
	name="${values[0]}"
	description="${values[1]}"
	color="${values[2]}"
	echo "Creating '$name' label..."
	gh label --repo nobl9/"$REPO" create "$name" --description "$description" --color "$color"
done <./bootstrap/labels.csv
