SHELL := /bin/bash

.PHONY: new test run help

help:
	@echo "Available targets:"
	@echo "  new NAME=<algorithm-name>   - Create a new algorithm/data structure package"
	@echo "  test [NAME=<directory>]     - Test specific directory or all directories"
	@echo "  run NAME=<directory>        - Run the algorithm in the specified directory"
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
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is required. Usage: make run NAME=directory-name"; \
		exit 1; \
	fi
	@./scripts/run.sh "$(NAME)"
