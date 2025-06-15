#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$SCRIPT_DIR/common.sh"

INPUT_NAME="$1"

print_header() {
  echo -e "${BLUE}================================${NC}"
  echo -e "${BLUE}  DSA Run All Modules${NC}"
  echo -e "${BLUE}================================${NC}"
  echo ""
}

run_module() {
  local dir_name="$1"

  if [ ! -d "$dir_name" ]; then
    echo -e "${RED}Error: Directory $dir_name does not exist${NC}"
    return 1
  fi

  local package_name=$(echo "$dir_name" | sed 's/^[0-9][0-9][0-9][0-9]-//' | tr '-' '_')
  local temp_main="temp_main_${dir_name}.go"

  echo -e "${YELLOW}Running: $dir_name${NC}"
  echo "----------------------------------------"

  cat >"$temp_main" <<EOF
package main

import (
	"fmt"
	$package_name "github.com/celj/dsa/$dir_name"
)

func main() {
	fmt.Println("Running $dir_name...")
	
	result := $package_name.Run()
	fmt.Printf("Result: %v\n", result)
}
EOF

  if go run "$temp_main"; then
    echo -e "${GREEN}✓ Successfully ran $dir_name${NC}"
    rm "$temp_main"
    return 0
  else
    echo -e "${RED}✗ Failed to run $dir_name${NC}"
    rm "$temp_main"
    return 1
  fi
}

main() {
  local input_selections="$1"
  local total_dirs=0
  local passed_dirs=0
  local failed_dirs=0

  if [ -n "$input_selections" ]; then
    local target_dirs
    if ! target_dirs=$(resolve_target_directories "$input_selections"); then
      exit 1
    fi

    echo -e "${BLUE}Running selected modules: $target_dirs${NC}"
    echo ""

    for dir in $target_dirs; do
      total_dirs=$((total_dirs + 1))
      if run_module "$dir"; then
        passed_dirs=$((passed_dirs + 1))
      else
        failed_dirs=$((failed_dirs + 1))
      fi
      echo ""
    done

    echo -e "${BLUE}================================${NC}"
    echo -e "${BLUE}  Run Summary${NC}"
    echo -e "${BLUE}================================${NC}"
    echo -e "Total modules run: $total_dirs"
    echo -e "${GREEN}Successful: $passed_dirs${NC}"

    if [ $failed_dirs -gt 0 ]; then
      echo -e "${RED}Failed: $failed_dirs${NC}"
      exit 1
    else
      echo -e "${GREEN}All modules ran successfully!${NC}"
      exit 0
    fi
  else
    print_header
    echo -e "${BLUE}Running all algorithm modules...${NC}"
    echo ""

    local dirs=$(get_all_dirs)

    if [ -z "$dirs" ]; then
      echo -e "${YELLOW}No algorithm directories found.${NC}"
      echo "Create one with: make new n=algorithm-name"
      exit 0
    fi

    for dir in $dirs; do
      dir=$(basename "$dir")
      total_dirs=$((total_dirs + 1))

      if run_module "$dir"; then
        passed_dirs=$((passed_dirs + 1))
      else
        failed_dirs=$((failed_dirs + 1))
      fi

      echo ""
    done

    echo -e "${BLUE}================================${NC}"
    echo -e "${BLUE}  Run Summary${NC}"
    echo -e "${BLUE}================================${NC}"
    echo -e "Total modules run: $total_dirs"
    echo -e "${GREEN}Successful: $passed_dirs${NC}"

    if [ $failed_dirs -gt 0 ]; then
      echo -e "${RED}Failed: $failed_dirs${NC}"
      exit 1
    else
      echo -e "${GREEN}All modules ran successfully!${NC}"
      exit 0
    fi
  fi
}

main "$@"
