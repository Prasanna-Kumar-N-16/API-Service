# ==================================================================================== #
# HELP
# ==================================================================================== #

.PHONY: help

## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## run_all: will execute all make commands
run_all: remove clean build test test_coverage

## remove: will delete the binary
remove:
	if [ -e API-Service ]; then \
        rm API-Service; \
        echo "File removed."; \
    else \
        echo "File does not exist."; \
    fi

## clean: will clean all the modules   
clean: remove
	go mod tidy -v

## build: generates a go binary
build:
	go build .

## test: will run test on all the files
test:
	go test ./...

## test_coverage: will say the code coverage
test_coverage:
	go test ./... -coverprofile=coverage.out


