run:
	make src
	air -- --env $(env)
populate:
	go run cmd/populator/main.go -file population.json -env $(env)
build:
	make src
	go build -o ./tmp/main.exe ./cmd/server/
src:
	tailwind -i ./internal/web/src/input.css -o ./dist/output.css