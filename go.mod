module auth

go 1.12

require (
	github.com/gin-contrib/sessions v0.0.1
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jinzhu/gorm v1.9.11
	github.com/kr/pretty v0.1.0 // indirect
	golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c
)

replace golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c => github.com/golang/crypto v0.0.0-20191002192127-34f69633bfdc
