#!/bin/bash

set -e

if [ -z "$1" ]; then
  echo "Error: Directory name is required"
  echo "Usage: $0 <directory-name>"
  exit 1
fi

INPUT_NAME="$1"

resolve_directory() {
  local input="$1"

  if [ -d "$input" ]; then
    echo "$input"
    return 0
  fi

  local matches=$(find . -maxdepth 1 -type d -name "*-${input}" | sort)

  if [ -z "$matches" ]; then
    echo "Error: No directory found matching '$input'"
    echo "Available directories:"
    find . -maxdepth 1 -type d -name "[0-9][0-9][0-9][0-9]-*" | sort | sed 's|^\./||'
    exit 1
  fi

  local count=$(echo "$matches" | wc -l)

  if [ "$count" -eq 1 ]; then
    echo "$matches" | sed 's|^\./||'
    return 0
  else
    echo "Error: Multiple directories found matching '$input':"
    echo "$matches" | sed 's|^\./||'
    exit 1
  fi
}

DIR_NAME=$(resolve_directory "$INPUT_NAME")

if [ ! -d "$DIR_NAME" ]; then
  echo "Error: Directory $DIR_NAME does not exist"
  exit 1
fi

PACKAGE_NAME=$(echo "$DIR_NAME" | sed 's/^[0-9][0-9][0-9][0-9]-//' | tr '-' '_')

TEMP_MAIN="temp_main_${DIR_NAME}.go"

cat >"$TEMP_MAIN" <<EOF
package main

import (
	"fmt"
	$PACKAGE_NAME "github.com/celj/dsa/$DIR_NAME"
)

func main() {
	fmt.Println("Running $DIR_NAME...")
	
	result := $PACKAGE_NAME.Run()
	fmt.Printf("Result: %v\n", result)
}
EOF

echo "Running $DIR_NAME..."
go run "$TEMP_MAIN"

rm "$TEMP_MAIN"
