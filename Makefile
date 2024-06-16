default: generate run

run:
	@go run cmd/main.go

generate:
	@templ generate
