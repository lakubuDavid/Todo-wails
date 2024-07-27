# justfile
default:
    just --list

run:
    wails dev
dev:
    wails dev

generate:
    sqlc generate
    wails generate module
migrate:
    dbmate migrate

# alias for dbmate
db +args="":
    dbmate {{args}}
