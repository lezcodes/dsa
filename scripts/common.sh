#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

get_directory_by_index() {
  local index="$1"
  local padded_index=$(printf "%04d" "$index")
  local dir=$(find . -maxdepth 1 -type d -name "${padded_index}-*" | head -1)
  if [ -n "$dir" ]; then
    echo "$dir" | sed 's|^\./||'
  else
    echo ""
  fi
}

resolve_directory() {
  local input="$1"

  if [[ "$input" =~ ^[0-9]+$ ]]; then
    local dir=$(get_directory_by_index "$input")
    if [ -z "$dir" ]; then
      echo -e "${RED}Error: No directory found for index '$input'${NC}" >&2
      echo "Available directories:" >&2
      find . -maxdepth 1 -type d -name "[0-9][0-9][0-9][0-9]-*" | sort | sed 's|^\./||' >&2
      return 1
    fi
    echo "$dir"
    return 0
  fi

  if [ -d "$input" ]; then
    echo "$input"
    return 0
  fi

  local matches=$(find . -maxdepth 1 -type d -name "*-${input}" | sort)

  if [ -z "$matches" ]; then
    echo -e "${RED}Error: No directory found matching '$input'${NC}" >&2
    echo "Available directories:" >&2
    find . -maxdepth 1 -type d -name "[0-9][0-9][0-9][0-9]-*" | sort | sed 's|^\./||' >&2
    return 1
  fi

  local count=$(echo "$matches" | wc -l)

  if [ "$count" -eq 1 ]; then
    echo "$matches" | sed 's|^\./||'
    return 0
  else
    echo -e "${RED}Error: Multiple directories found matching '$input':${NC}" >&2
    echo "$matches" | sed 's|^\./||' >&2
    return 1
  fi
}

parse_selections() {
  local input="$1"
  if [ -z "$input" ]; then
    echo ""
    return 0
  fi

  echo "$input" | tr ',' '\n'
}

get_all_dirs() {
  find . -maxdepth 1 -type d -name "[0-9][0-9][0-9][0-9]-*" | sort
}

resolve_target_directories() {
  local input_selections="$1"
  local target_dirs=""

  if [ -n "$input_selections" ]; then
    local selections=$(parse_selections "$input_selections")

    while IFS= read -r selection; do
      if [ -n "$selection" ]; then
        local target_dir
        if ! target_dir=$(resolve_directory "$selection"); then
          return 1
        fi
        target_dirs="$target_dirs $target_dir"
      fi
    done <<<"$selections"
  fi

  echo "$target_dirs"
}
