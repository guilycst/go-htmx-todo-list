run:
	docker compose up -d db
	air
	docker compose down
run_l:
	make src
	go run ./cmd/server/main.go
build:
	make src
	go build -o ./tmp/main.exe ./cmd/server/
src:
	tailwind -i ./internal/web/src/input.css -o ./dist/output.css
	cp -R ./internal/web/src/*.js ./dist/
populate:
	docker compose up -d db
	go run cmd/populator/main.go -file population.json .env
	docker compose down
populate_l:
	go run cmd/populator/main.go -file population.json .env