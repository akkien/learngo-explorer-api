## build_front: builds the Explorer
build:
	@echo "Building Explorer..."
	@go build -o ./explorer .
	@echo "Explorer built!"

## start: starts the Explorer
start: build
	@echo "Starting the Explorer..."
	@./explorer &
	@echo "Explorer running!"

## stop_invoice: stops the Explorer
stop:
	@echo "Stopping the Explorer..."
	@-pkill -SIGTERM -f "explorer"
	@echo "Stopped Explorer"