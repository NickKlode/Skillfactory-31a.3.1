package main

import (
	"gonews/pkg/api"
	"gonews/pkg/storage"
	"gonews/pkg/storage/memdb"
	"gonews/pkg/storage/mongodb"
	"gonews/pkg/storage/postgresql"
	"log"
	"net/http"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	// Создаём объект сервера.
	var srv server

	// Создаём объекты баз данных.
	//
	// БД в памяти.
	db := memdb.New()

	// Реляционная БД PostgreSQL.
	db2, err := postgresql.New("postgres://postgres:ZAQzaqzaq97@localhost:5433/posts")
	if err != nil {
		log.Fatal(err)
	}
	// Документная БД MongoDB.
	db3, err := mongodb.New("mongodb://server.domain:27017/")
	if err != nil {
		log.Fatal(err)
	}
	_, _, _ = db, db2, db3

	// Инициализируем хранилище сервера конкретной БД.
	srv.db = db

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// Запускаем веб-сервер на порту 8080 на всех интерфейсах.
	// Предаём серверу маршрутизатор запросов,
	// поэтому сервер будет все запросы отправлять на маршрутизатор.
	// Маршрутизатор будет выбирать нужный обработчик.
	http.ListenAndServe(":8080", srv.api.Router())
}
