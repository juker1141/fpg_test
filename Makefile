
LOGGER_BINARY=loggerServiceApp


## build_logger: builds the logger binary as a linux executable
build_logger:
	@echo "Building logger binary..."
	cd logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} .
	@echo "Done!"

.PHONY: build_logger