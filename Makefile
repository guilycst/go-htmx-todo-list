run:
	tailwind -i ./resources/css/input.css -o ./dist/output.css
	go run main.go