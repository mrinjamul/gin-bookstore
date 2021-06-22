# gin-bookstore

Simple CRUD Operations for bookStore.

```qoute
Gin by Example !
```

## Used Tools and Technologies

- [Golang](https://golang.org/)
- [Gin Gonic](https://gin-gonic.com/)
- [Gorm](https://github.com/jinzhu/gorm)
- [go-sqlite](https://github.com/mattn/go-sqlite3)

## Available secrets

For ENV,

    # Golang ENV (.env)
    GIN_MODE=release
    PORT=<your port>

## Build

```shell
go mod download
go build -o gin-bookstore .
```

## Running

```shell
./gin-bookstore
```

## Endpoints

| Methods | Endpoints  | Description       |
| ------- | ---------- | ----------------- |
| GET     | /          | Hello World       |
| GET     | /books     | Get All books     |
| POST    | /books     | Create a book     |
| GET     | /books/:id | Get book by id    |
| PATCH   | /books/:id | Update book by id |
| DELETE  | /books/:id | Delete book by id |

## Author

- Injamul Mohammad Mollah <mrinjamul@gmail.com>

## License

- under [MIT license](LICENSE)
