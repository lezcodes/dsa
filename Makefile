SHELL := /bin/bash

.PHONY: new test bench run help

help:
	@echo "Available targets:"
	@echo "  new n=<algorithm-name>      - Create a new algorithm/data structure package"
	@echo "  test [n=<selection>]        - Test specific directories or all directories"
	@echo "  bench [n=<selection>]       - Run benchmarks for specific directories or all directories"
	@echo "  run [n=<selection>]         - Run algorithms in specific directories or all directories"
	@echo "  help                        - Show this help message"
	@echo ""
	@echo "Selection formats:"
	@echo "  n=linear-search,binary-search               - Run by algorithm names"
	@echo "  n=1,2,15,20                                 - Run by directory indices"
	@echo "  n=linear-search,8,0002-binary-search,15     - Mix names and indices"

new:
	@if [ -z "$(n)" ]; then \
		echo "Error: n is required. Usage: make new n=algorithm-name"; \
		exit 1; \
	fi
	@./scripts/new.sh "$(n)"

test:
	@./scripts/test.sh "$(n)"

bench:
	@./scripts/bench.sh "$(n)"

run:
	@./scripts/run.sh "$(n)"

modern:
	@go tool modernize -fix ./...
