#!/bin/bash

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

print_header() {
  echo -e "${BLUE}================================${NC}"
  echo -e "${BLUE}  DSA Test Runner${NC}"
  echo -e "${BLUE}================================${NC}"
  echo ""
}

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

run_tests_in_dir() {
  local dir="$1"
  local success=true

  echo -e "${YELLOW}Testing: $dir${NC}"
  echo "----------------------------------------"

  if [ ! -d "$dir" ]; then
    echo -e "${RED}Directory $dir does not exist${NC}"
    return 1
  fi

  echo "Running go test..."
  if ! go test -v "./$dir"; then
    echo -e "${RED}Tests failed in $dir${NC}"
    success=false
  fi

  echo ""
  echo "Running benchmarks..."
  if ! go test -bench=. -benchmem "./$dir"; then
    echo -e "${RED}Benchmarks failed in $dir${NC}"
    success=false
  fi

  echo ""
  echo "Running go vet..."
  if ! go vet "./$dir"; then
    echo -e "${RED}Go vet failed in $dir${NC}"
    success=false
  fi

  echo ""
  echo "Checking gofmt..."
  if [ -n "$(gofmt -l $dir)" ]; then
    echo -e "${RED}Code formatting issues found:${NC}"
    gofmt -l "$dir"
    success=false
  else
    echo -e "${GREEN}Code formatting is correct${NC}"
  fi

  if [ "$success" = true ]; then
    echo -e "${GREEN}✓ All checks passed for $dir${NC}"
    return 0
  else
    echo -e "${RED}✗ Some checks failed for $dir${NC}"
    return 1
  fi
}

get_all_dirs() {
  find . -maxdepth 1 -type d -name "[0-9][0-9][0-9][0-9]-*" | sort
}

main() {
  print_header

  local input_name="$1"
  local total_dirs=0
  local passed_dirs=0
  local failed_dirs=0

  if [ -n "$input_name" ]; then
    local target_dir=$(resolve_directory "$input_name")
    echo -e "${BLUE}Testing specific directory: $target_dir${NC}"
    echo ""

    if run_tests_in_dir "$target_dir"; then
      passed_dirs=1
    else
      failed_dirs=1
    fi
    total_dirs=1
  else
    echo -e "${BLUE}Testing all algorithm directories...${NC}"
    echo ""

    local dirs=$(get_all_dirs)

    if [ -z "$dirs" ]; then
      echo -e "${YELLOW}No algorithm directories found.${NC}"
      echo "Create one with: make new NAME=algorithm-name"
      exit 0
    fi

    for dir in $dirs; do
      dir=$(basename "$dir")
      total_dirs=$((total_dirs + 1))

      if run_tests_in_dir "$dir"; then
        passed_dirs=$((passed_dirs + 1))
      else
        failed_dirs=$((failed_dirs + 1))
      fi

      echo ""
    done
  fi

  echo -e "${BLUE}================================${NC}"
  echo -e "${BLUE}  Test Summary${NC}"
  echo -e "${BLUE}================================${NC}"
  echo -e "Total directories tested: $total_dirs"
  echo -e "${GREEN}Passed: $passed_dirs${NC}"

  if [ $failed_dirs -gt 0 ]; then
    echo -e "${RED}Failed: $failed_dirs${NC}"
    exit 1
  else
    echo -e "${GREEN}All tests passed!${NC}"
    exit 0
  fi
}

main "$@"
