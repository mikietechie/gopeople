export PATH=$(go env GOPATH)/bin:$PATH
swag init  -g api/app/app.go --parseDependency --parseInternal

go run main.go