# Build the application for Linux
build-lin:
	@echo "Building for linux..."	
	
	@GOOS=linux GOARCH=amd64 go build -o remote cmd/api/main.go

run-lin:
	@echo "Building for linux..."	
	
	@GOOS=linux GOARCH=amd64 go build -o remote cmd/api/main.go
	@echo "Running for linux..."	
	
	@./remote

# Build the application for Windows
build-win:
	@echo "Building for windows..."	
	
	@GOOS=windows GOARCH=amd64 go build -o remote cmd/api/main.go

run-win:
	@echo "Building for windows..."	
	
	@GOOS=windows GOARCH=amd64 go build -o remote cmd/api/main.go
	@echo "Running for windows..."	
	
	@./remote

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f remote
