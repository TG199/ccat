# variables
APP_NAME = ccat
SRC_FILES = ccat.go
BUILD_DIR = ./bin
OUTPUT = $(BUILD_DIR)/$(APP_NAME)

# Build the program
.PHONY: build
build:
	@mkdir -p $(BUILD_DIR)
	@echo "Building the $(APP_NAME) program..."
	go build -o $(OUTPUT) $(SRC_FILES)
	@echo "Build complete!. Binary available at $(OUTPUT)."

# Run the program
.PHONY: run
run:
	@echo "Running the $(APP_NAME) program..."
	$(OUTPUT)

# Clean the build
.PHONY: clean
clean:
	@echo "Cleaning the build..."
	rm -rf $(BUILD_DIR)
	@echo "Clean complete."

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting the code..."
	go fmt $(SRC_FILES)
	@echo "Format complete."

# Lint code
.PHONY: lint
lint:
	@echo "Linting the code..."
	go vet $(SRC_FILES)
	@echo "Linting complete."

# Test run with a file
.PHONY: test
test:
	@echo "Testing $(APP_NAME)..."
	$(OUTPUT) test.txt

