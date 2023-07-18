# Go + HTMX To-Do List
A modest todo list app built with Go, HTMX and Tailwind CSS. The primary objective behind creating this app was to explore the capabilities of HTMX and gain insights into its practical implementation.

![Todo app screenshot](./docs/todo.png)

# What it can currently do

- Create, Read, Update and Delete tasks
- Complete tasks
- Data is stored in PostgreSQL

# Dependencies
- Go (1.19)
- Make
- Tailwind CSS Standalone CLI
- Air (live-reloading)
- Docker

# Implementation details

The majority of Go code is dedicated to an HTTP server that manages requests from the HTMX library. When these requests are successful, they trigger the execution of an HTML template using Go's [html/template package](https://pkg.go.dev/html/template) and return an HTML document as a response.

HTMX leverages these HTML responses to dynamically replace parts of the page without requiring a full page reload. This approach enables the application to provide interactivity comparable to popular JavaScript frameworks/libraries, but with reduced reliance on actual JavaScript code.

The visual appearance of the page is managed using the Tailwind CSS framework. The Standalone CLI tool is responsible for compiling the *.html files found in the [./internal/web/templates/](./internal/web/templates/) directory and producing a CSS file [./dist/output.css](./dist/output.css). This generated CSS file is then served by the Go server.

For storage, a PostgreSQL database is utilized within a container.

# How to run

## Server
Serves the HTMX app.

In a browser visit  ***http://localhost:8080***

### Without Docker

 - Download the [Tailwind CSS Standalone CLI](https://tailwindcss.com/blog/standalone-cli) and setup in your PATH as **tailwind**
 - Setup a PosgreSQL server: The connection string is passed as a ENV variable named **PG_CONN_STR**, ENV variables are defined in the [.env](./.env), look it up for reference

Then run the Make file:

```$ make run_l```

### Docker Compose

#### Source on host machine with live-reload, database on Docker

```$ make run```

#### Everithing on Docker

```$ docker compose up```

## Populator
    
Inserts the contents of [population.json](./population.json) file into the database

### Without Docker

 - Setup a PosgreSQL server: The connection string is passed as a ENV variable named **PG_CONN_STR**, ENV variables are defined in the [.env](./.env), look it up for reference

 Then run the Make file:

```$ make populate_l```


### Docker Compose

```$ make populate```
