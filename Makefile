run:
	air
build:
	make tailwind
	go build -o ./tmp/main.exe
tailwind:
	tailwind -i ./resources/css/input.css -o ./dist/output.css
dev_setup:
	go install github.com/cosmtrek/air@latest