SHELL := /bin/bash

.PHONY: new test run help

help:
	@echo "Available targets:"
	@echo "  new NAME=<algorithm-name>   - Create a new algorithm/data structure package"
	@echo "  test [NAME=<directory>]     - Test specific directory or all directories"
	@echo "  run [NAME=<directory>]      - Run algorithm in specific directory or all directories"
	@echo "  help                        - Show this help message"

new:
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is required. Usage: make new NAME=algorithm-name"; \
		exit 1; \
	fi
	@./scripts/new.sh "$(NAME)"

test:
	@./scripts/test.sh "$(NAME)"

run:
	@./scripts/run.sh "$(NAME)"

modernize:
	@go tool modernize ./...
