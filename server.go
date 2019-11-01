package main

import (
  "database/sql"
  "fmt"
  "log"
  "net/url"
  "os"

  _ "github.com/go-sql-driver/mysql"
  "github.com/labstack/echo/v4"
  "github.com/yinhr/todo-api/delivery"
  "github.com/yinhr/todo-api/repository"
  "github.com/yinhr/todo-api/usecase"
  "gopkg.in/boj/redistore.v1"
)

func main() {
  e := echo.New()

  dbHost := os.Getenv("DBHOST") 
  dbPort := os.Getenv("DBPORT")
  dbName := os.Getenv("DBNAME")
  dbUser := os.Getenv("DBUSER")
  dbPass := os.Getenv("DBPASS")
  params := url.Values{}
  params.Add("loc", "Asia/Tokyo")
  params.Add("parseTime", "true")
  params.Add("timeout", "10s")
  dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbUser, dbPass, dbHost, dbPort, dbName, params.Encode())

  conn, err := sql.Open(`mysql`, dns)
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }

  defer func() {
    if err := conn.Close(); err != nil {
      log.Fatal(err)
    }
  }()

  redisHost := os.Getenv("REDISHOST")
  redisPort := os.Getenv("REDISPORT")
  redisPass := os.Getenv("REDISPASS")
  store, err := redistore.NewRediStore(10, "tcp", redisHost + ":" + redisPort, "", []byte(redisPass))
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
  defer store.Close()

  repositories := repository.GetInitializedRepositories(conn)

  usecases := usecase.GetInitializedUsecases(&repositories)

  handlers := delivery.GetInitializedHandlers(&usecases)
  delivery.ConfigureHandler(e, &handlers)
  delivery.ConfiguerMiddleware(e, store)
  
  e.Logger.Fatal(e.Start(":4000"))
}
