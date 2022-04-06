API_PORT=5001

## restart: builds all binaries
restart: stop build start
	@echo "Explorer API restarted!"

## clean: cleans all binaries and runs go clean
clean:
	@echo "Cleaning..."
	@- rm -f dist/*
	@go clean
	@echo "Cleaned!"

## build_front: builds the Explorer API
build:
	@echo "Building Explorer API..."
	@go build -o dist/explorer_api .
	@echo "Explorer API built!"

## start: starts the Explorer API
start: build
	@echo "Starting the Explorer API..."
	@./dist/explorer_api -port=${API_PORT} &
	@echo "Explorer API running!"

## stop_invoice: stops the Explorer API
stop:
	@echo "Stopping the Explorer API..."
	@-pkill -SIGTERM -f "explorer_api -port=${API_PORT}"
	@echo "Stopped Explorer API"