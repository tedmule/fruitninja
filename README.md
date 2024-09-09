Fruitninja is a blah (still working on)

# Environment
* **FRUIT_NINJA_NAME**: Fruit name for HTTP response(apple, banana, ...)
* **FRUIT_NINJA_COUNT**: Number of fruit for HTTP response

# Dev
Use air to reload
```
go install github.com/cosmtrek/air@latest
```

# Dev
```
air go run .
```

# Build
Build for alpine
```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ninja .
```