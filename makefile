run:
	@go run ./cmd/url-shortener/main.go

build:
	@go build -o ./bin/url-shortener ./cmd/url-shortener/main.go

clean:
	@rm ./bin/url-shortener && rm -d ./bin/

full: build
	@./bin/url-shortener