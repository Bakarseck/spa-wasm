# Codes de couleur ANSI
GREEN := \033[32m
RED := \033[31m
RESET := \033[0m
EXE := forum
DB_PATH := ./db/clash.db
SQL_FILE := ./db/clash.sql
EXEC_FILE := ./db/exec.sql
POST_FILE := ./db/posts.sql
CATEGORY_FILE := ./db/categories.sql

build:
	cd ./App && GOOS=js GOARCH=wasm go build -o ../assets/main.wasm

start: build
	cd ./server && go run .

create-db:
	@sqlite3 $(DB_PATH) < $(SQL_FILE)

auth:
	cd ./auth && go build -o auth && ./auth