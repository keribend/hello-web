.PHONY: run
run: templ-generate
	go run cmd/hello-web/*.go

.PHONY: templ-generate
templ-generate:
	@templ generate

.PHONY: templ-generate-watch
templ-generate-watch:
	@templ generate -watch -proxyport=7332 -proxy="http://localhost:8080" -open-browser=false -cmd="go run cmd/hello-web/main.go"

.PHONY: tailwind-watch
tailwind-watch:
	@npx @tailwindcss/cli -i input.css -o ./public/static/css/tw.css --watch

.PHONY: sql-generate
sql-generate:
	sqlc generate