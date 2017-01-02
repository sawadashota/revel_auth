package models

import (
  "github.com/revel/revel"
  "github.com/jinzhu/gorm"
)

type User struct {
  gorm.Model
  TeamID        int
  Name          string
  Email         string
  Password      []byte
  Role          string
  RememberToken string

  Team          Team
}

type MockUser struct {
  TeamID          int
  Name            string
  Email           string
  Password        string
  PasswordConfirm string
  Role            string
}

type TokenUser struct {
  Email         string
  RememberToken string
}

type uniqueName struct{}
type uniqueEmail struct{}

func (mockUser *MockUser) Validate(v *revel.Validation) {
  v.Required(mockUser.Name).Message("Name is required")
  v.MinSize(mockUser.Name, 3).Message("Name is not enough long")
  v.Check(mockUser.Name, uniqueName{}).Message("The name has been used.")
  v.Check(mockUser.Email,
    revel.ValidRequired(),
    revel.ValidEmail()).Message("Email is required")
  v.Check(mockUser.Email, uniqueEmail{}).Message("Email should be unique")
  v.Required(mockUser.Password).Message("Password is required")
  v.Required(mockUser.PasswordConfirm).Message("Password confirmation is required")
  v.Required(mockUser.Role).Message("Role is required")
}

func (u uniqueEmail) IsSatisfied(value interface{}) bool {
  user := User{}

  (*DB).Where("email = ?", value).FirstOrInit(&user)

  return user.Email == ""
}

func (u uniqueEmail) DefaultMessage() string {
  return "Email Should be unique"
}

func (u uniqueName) IsSatisfied(value interface{}) bool {
  user := User{}

  (*DB).Where("name = ?", value).FirstOrInit(&user)

  return user.Email == ""
}

func (u uniqueName) DefaultMessage() string {
  return "Name Should be unique"
}
