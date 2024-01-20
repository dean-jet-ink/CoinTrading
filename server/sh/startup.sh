#!/bin/sh

go run ./app/migrate/migrate.go

air -c .air.toml