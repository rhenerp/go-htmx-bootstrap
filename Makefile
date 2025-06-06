APP_NAME := myapp

.PHONY: generate build run clean

generate:
	@echo "🔧 Generating templ files..."
	templ generate

build: generate
	@echo "📦 Building the Go app..."
	go build -o $(APP_NAME) .

run: build
	@echo "🚀 Running the app..."
	./$(APP_NAME)

clean:
	@echo "🧹 Cleaning up..."
	rm -f $(APP_NAME)