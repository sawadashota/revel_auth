package models

import (
  "github.com/revel/revel"
  "github.com/jinzhu/gorm"
  "strings"
  "fmt"
  _"github.com/go-sql-driver/mysql"
  "time"
)

var DB **gorm.DB

func InitDB() {
  db, err := gorm.Open("mysql", getConnectionString())

  if err != nil {
    revel.ERROR.Println("FATAL", err)
    panic(err)
  }

  db.DB()
  DB = &db
}

type Model struct {
  ID        uint `gorm:"primary_key"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt *time.Time
}

type Validator interface {
  IsSatisfied(interface{}) bool
  DefaultMessage() string
}

func getParamString(param string, defaultValue string) string {
  p, found := revel.Config.String(param)
  if !found {
    if defaultValue == "" {
      revel.ERROR.Fatal("Cound not find parameter: " + param)
    } else {
      return defaultValue
    }
  }
  return p
}

func getConnectionString() string {
  host := getParamString("db.host", "localhost")
  port := getParamString("db.port", "3306")
  user := getParamString("db.user", "Username")
  pass := getParamString("db.password", "Password")
  dbname := getParamString("db.name", "DatabaseName")
  protocol := getParamString("db.protocol", "tcp")
  dbargs := getParamString("dbargs", " ")
  timezone := getParamString("db.timezone", "parseTime=true&loc=Asia%2FTokyo")

  if strings.Trim(dbargs, " ") != "" {
    dbargs = "?" + dbargs
  } else {
    dbargs = ""
  }
  return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s?%s", user, pass, protocol, host, port, dbname, dbargs, timezone)
}
