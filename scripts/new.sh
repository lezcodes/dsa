#!/bin/bash

set -e

if [ -z "$1" ]; then
  echo "Error: Algorithm name is required"
  echo "Usage: $0 <algorithm-name>"
  exit 1
fi

ALGORITHM_NAME="$1"

get_next_number() {
  local max_num=0
  for dir in [0-9][0-9][0-9][0-9]-*; do
    if [ -d "$dir" ]; then
      local num=$(echo "$dir" | sed 's/^\([0-9][0-9][0-9][0-9]\)-.*/\1/')
      local decimal_num=$((10#$num))
      if [ "$decimal_num" -gt "$max_num" ]; then
        max_num="$decimal_num"
      fi
    fi
  done
  printf "%04d" $((max_num + 1))
}

NEXT_NUM=$(get_next_number)
DIR_NAME="${NEXT_NUM}-${ALGORITHM_NAME}"

if [ -d "$DIR_NAME" ]; then
  echo "Error: Directory $DIR_NAME already exists"
  exit 1
fi

echo "Creating directory: $DIR_NAME"
mkdir "$DIR_NAME"

PACKAGE_NAME=$(echo "$ALGORITHM_NAME" | tr '-' '_')

cat >"$DIR_NAME/${PACKAGE_NAME}.go" <<EOF
package $PACKAGE_NAME

func Run() any {
	return "TODO: Implement $ALGORITHM_NAME"
}
EOF

cat >"$DIR_NAME/${PACKAGE_NAME}_test.go" <<EOF
package $PACKAGE_NAME

import "testing"

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run()
	}
}
EOF

cat >"$DIR_NAME/README.md" <<EOF
# $ALGORITHM_NAME

## Description
TODO: Add description of the algorithm/data structure

## Complexity
- Time Complexity: TODO
- Space Complexity: TODO

## Usage
\`\`\`bash
make run n=$DIR_NAME
\`\`\`

## Testing
\`\`\`bash
make test n=$DIR_NAME
\`\`\`
EOF

echo "Successfully created $DIR_NAME with:"
echo "  - ${PACKAGE_NAME}.go (implementation)"
echo "  - ${PACKAGE_NAME}_test.go (tests and benchmarks)"
echo "  - README.md (documentation)"
echo ""
echo "To get started:"
echo "  make run n=$DIR_NAME"
