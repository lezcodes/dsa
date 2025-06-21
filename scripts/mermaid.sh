#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m'

OUTPUT_FILE="output.log"

>"$OUTPUT_FILE"

echo -e "${YELLOW}Checking Mermaid diagrams in Markdown files (excluding root README.md) in numerical order...${NC}"
echo ""

ERRORS_FOUND=0

find . -type f -name "*.md" -print0 | sort -zV | while IFS= read -r -d $'\0' file; do

  relative_file="${file#./}"

  if [[ "$relative_file" == "README.md" ]]; then
    echo -e "${YELLOW}--- Skipping root README.md ---${NC}"
    continue
  fi

  echo -e "${YELLOW}--- Processing $file ---${NC}"

  TEMP_SVG_FILE=/tmp/output.svg

  ALL_OUTPUT=$(mmdc -i "$file" -o "$TEMP_SVG_FILE" 2>&1 || true)

  rm -f "$TEMP_SVG_FILE"

  FILTERED_OUTPUT=$(echo "$ALL_OUTPUT" |
    grep -v "output saved to" |
    grep -v "^\s*✅\s*\./output-[0-9]\+\.svg" |
    grep -v "Parser[0-9]\+\.parseError" |
    grep -v "^[[:space:]]*at ")

  if echo "$FILTERED_OUTPUT" | grep -q "Error: Parse error"; then

    echo -e "${RED}$file${NC}"
    echo "$FILTERED_OUTPUT"
    echo ""

    echo "$file" >>"$OUTPUT_FILE"
    echo "$FILTERED_OUTPUT" >>"$OUTPUT_FILE"
    echo "" >>"$OUTPUT_FILE"
    ERRORS_FOUND=1
  elif [ -n "$FILTERED_OUTPUT" ] && echo "$FILTERED_OUTPUT" | grep -q "Found [0-9]\+ mermaid charts"; then

    echo -e "${GREEN}No Mermaid syntax errors.${NC}"
    echo ""
  else
    echo -e "${GREEN}No Mermaid syntax errors.${NC}"
    echo ""
  fi
done

echo "-------------------------------------"
if [ $ERRORS_FOUND -eq 1 ]; then
  echo -e "${RED}Mermaid syntax check completed with errors. See $OUTPUT_FILE for details.${NC}"
  exit 1
else
  echo -e "${GREEN}Mermaid syntax check completed successfully. No errors found.${NC}"
  exit 0
fi
