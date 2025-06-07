# Build the application for Linux
build-lin:
	@echo "Building for linux..."	
	
	@GOOS=linux GOARCH=amd64 go build -o remote cmd/api/main.go

# Build the application for Windows
build-win:
	@echo "Building for windows..."	
	
	@GOOS=windows GOARCH=amd64 go build -o remote cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f remote
