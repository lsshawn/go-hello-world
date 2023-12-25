.PHONY: run css

run:
	@templ generate
	@go run cmd/main.go

.PHONY: css
css:
	tailwindcss -i css/input.css -o css/output.css --minify

.PHONY: css-watch
css-watch:
	tailwindcss -i css/input.css -o css/output.css --watch
