#!/usr/bin/env bash
set -euo pipefail

if (($# != 3)); then
	printf 'Usage: bash %s CASE_ID OUTPUT_DIRECTORY GOVY_CHECKOUT\n' "$0" >&2
	exit 2
fi

case "$1" in
[1-9]) ;;
10)
	printf 'Case 10 is explanation-only; grade explanation.md against evals.json.\n' >&2
	exit 2
	;;
*)
	printf 'Invalid case ID: %s (expected 1-9)\n' "$1" >&2
	exit 2
	;;
esac

case_id="$1"
if [[ ! -d "$2" || ! -d "$3" ]]; then
	printf 'Output directory and Govy checkout must exist.\n' >&2
	exit 2
fi
output_dir="$(cd -- "$2" && pwd -P)"
govy_checkout="$(cd -- "$3" && pwd -P)"
script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd -P)"
if [[ ! -f "$govy_checkout/go.mod" || ! -f "$govy_checkout/go.sum" ]]; then
	printf 'Govy checkout must contain go.mod and go.sum.\n' >&2
	exit 2
fi

shopt -s nullglob
sources=("$output_dir"/*.go)
if (("${#sources[@]}" == 0)); then
	printf 'No Go output files in %s\n' "$output_dir" >&2
	exit 1
fi
for source in "${sources[@]}"; do
	if [[ ! -f "$source" ]]; then
		printf 'Go output is not a regular file: %s\n' "$source" >&2
		exit 1
	fi
	case "${source##*/}" in
	govy_eval_*.go)
		printf 'Reserved output filename: %s\n' "${source##*/}" >&2
		exit 1
		;;
	esac
done

eval_workdir="$(mktemp -d "${TMPDIR:-/tmp}/govy-skill-eval-$(date -u +%Y%m%dT%H%M%SZ)-XXXXXX")"
cleanup() {
	local status="$?"
	trap - EXIT
	if ! rm -rf -- "$eval_workdir"; then
		printf 'Failed to remove temporary evaluation directory: %s\n' "$eval_workdir" >&2
		if ((status == 0)); then
			status=1
		fi
	fi
	exit "$status"
}
trap cleanup EXIT
trap 'exit 130' INT
trap 'exit 143' TERM

cp -- "$govy_checkout/go.mod" "$govy_checkout/go.sum" "$eval_workdir/"
cp -- "${sources[@]}" "$script_dir/testdata/govy_eval_helpers_test.go" \
	"$script_dir/testdata/govy_eval_case_${case_id}_test.go" "$eval_workdir/"
cd -- "$eval_workdir"

export GOWORK=off GOPROXY=off GOSUMDB=off GOTOOLCHAIN=local GOFLAGS=
go mod edit -module=govy-skill-eval \
	-require=github.com/nobl9/govy@v0.0.0 \
	"-replace=github.com/nobl9/govy=$govy_checkout"
go test -mod=mod -count=1 -json ./...
