.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: run
run:
	make templ-generate
	go run cmd/hello-web/main.go

