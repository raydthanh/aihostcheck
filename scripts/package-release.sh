#!/usr/bin/env bash

set -euo pipefail

if [[ $# -ne 4 ]]; then
  echo "usage: $0 <goos> <goarch> <version> <output-directory>" >&2
  exit 2
fi

target_os="$1"
target_arch="$2"
version="$3"
output_directory="$4"

case "$target_os" in
  linux | darwin | windows) ;;
  *)
    echo "unsupported operating system: $target_os" >&2
    exit 2
    ;;
esac

case "$target_arch" in
  amd64 | arm64) ;;
  *)
    echo "unsupported architecture: $target_arch" >&2
    exit 2
    ;;
esac

if [[ ! "$version" =~ ^(v)?[0-9A-Za-z][0-9A-Za-z.-]*$ ]]; then
  echo "invalid version: $version" >&2
  exit 2
fi

display_version="$version"
archive_version="${version#v}"
package_name="aihostcheck_${archive_version}_${target_os}_${target_arch}"
package_directory="$output_directory/$package_name"

mkdir -p "$package_directory/schema"
cp LICENSE README.md "$package_directory/"
cp docs/INSTALL.md "$package_directory/INSTALL.md"
cp schema/report.schema.json "$package_directory/schema/"

binary_name="aihostcheck"
if [[ "$target_os" == "windows" ]]; then
  binary_name+=".exe"
fi

CGO_ENABLED=0 GOOS="$target_os" GOARCH="$target_arch" \
  go build -trimpath -ldflags "-s -w -X main.version=$display_version" \
  -o "$package_directory/$binary_name" ./cmd/aihostcheck

if [[ "$target_os" == "windows" ]]; then
  (
    cd "$output_directory"
    zip -qr "$package_name.zip" "$package_name"
  )
  archive="$output_directory/$package_name.zip"
else
  tar -C "$output_directory" -czf "$output_directory/$package_name.tar.gz" "$package_name"
  archive="$output_directory/$package_name.tar.gz"
fi

rm -r "$package_directory"
sha256sum "$archive"
