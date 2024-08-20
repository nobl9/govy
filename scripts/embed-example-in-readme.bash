#!/usr/bin/env bash

set -e

if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <README_PATH>"
  exit 1
fi

README_PATH="$1"

# sed 's/\\/\\\\/g' is used to add extra backslashes which are eaten up by awk.
for file_embed in $(awk '/^\[\/\/\]: # \(embed: .*\)$/ {sub(/\)/, "", $4); print $4}' "$README_PATH"); do
  echo "Found embed directive for: $file_embed" >&2
  awk \
    -v file_embed="$file_embed" \
    -v file_embed_contents="$(sed 's/\\/\\\\/g' "$file_embed")" \
    'q==1 && /^```go$/ {q=2};
    q<2 {print}
    $0 ~ file_embed {q=1}
    q==2 && !/go/ && /^```$/ {q=0; printf "```go\n%s\n```\n", file_embed_contents};' \
    "$README_PATH" >"$README_PATH.tmp"
  mv "$README_PATH.tmp" "$README_PATH"
done
