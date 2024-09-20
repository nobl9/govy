#!/usr/bin/env bash

set -euo pipefail

if [ "$#" -ne 1 ]; then
  echo "Provide a single argument with a version of the release draft to use." >&2
  echo "Usage: $0 <VERSION>"
  exit 1
fi

VERSION="$1"

RELEASE_NOTES=$(gh release view "$VERSION" --json body --jq .body)

BREAKING_CHANGES_HEADER="Breaking Changes"
RELEASE_NOTES_HEADER="Release Notes"

commit_message_re="-\s.*\s\(#([0-9]+)\)\s@.*"
rls_header_re="^##.*(Features|$BREAKING_CHANGES_HEADER|Bug Fixes|Fixed Vulnerabilities)"

extract_header() {
  local body="$1"
  local header_name="$2"
  awk "
    /^\s?$/ {next};
    /^--+/ {rn=0};
    /^Signed-off-by|Co-authored-by/ {rn=0};
    /^## $header_name/ {rn=1};
    rn && !/^##/ && !/^--+/ {print};
    /^##/ && !/^## $header_name/ {rn=0}" <<<"$body"
}

indent() {
  while IFS= read -r line; do
    printf "  %s\n" "${line%"${line##*[![:space:]]}"}"
  done <<<"$1"
}

new_notes=""
rls_header=""
while IFS= read -r line; do
  new_notes+="$line\n"
  if [[ $line == \##* ]]; then
    if ! [[ $line =~ $rls_header_re ]]; then
      rls_header=""
      continue
    fi
    rls_header="${BASH_REMATCH[1]}"
  fi
  if [[ $rls_header == "" ]] || [[ $line != -* ]] || [[ $line == *"@renovate"* ]]; then
    continue
  fi
  if ! [[ $line =~ $commit_message_re ]]; then
    continue
  fi
  pr_number="${BASH_REMATCH[1]}"
  pr_body=$(gh pr view "$pr_number" --json body --jq '.body')

  add_notes() {
    local notes="$1"
    if [[ $notes != "" ]]; then
      new_notes+=$(indent "> $notes")
      new_notes+="\n"
    fi
  }

  rn=$(extract_header "$pr_body" "$RELEASE_NOTES_HEADER")
  bc=$(extract_header "$pr_body" "$BREAKING_CHANGES_HEADER")

  case $rls_header in
  "$BREAKING_CHANGES_HEADER") add_notes "$bc" ;;
  *) add_notes "$rn" ;;
  esac

done <<<"$RELEASE_NOTES"

echo "Uploading release notes for $VERSION"
# shellcheck disable=2059
printf "$new_notes" | gh release edit "$VERSION" --verify-tag -F -
