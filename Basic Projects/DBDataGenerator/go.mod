module github.com/dinesh.natukula/myapp

go 1.24.3

require github.com/brianvoe/gofakeit/v6 v6.28.0

require github.com/dinesh.natukula/mylib v0.0.0

replace github.com/dinesh.natukula/mylib => ../go-libs/mylib
