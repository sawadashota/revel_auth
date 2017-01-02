package models

import (
  "github.com/revel/revel"
  "github.com/jinzhu/gorm"
)

type Team struct {
  gorm.Model
  Name  string
  Users []User
}

type uniqueTeamName struct{}

func (team *Team) Validate(v *revel.Validation) {
  v.Required(team.Name).Message("Name is required")
  v.MinSize(team.Name, 3).Message("Name is not enough long")
  v.Check(team.Name, uniqueTeamName{}).Message("The name has been used.")
}

func (u uniqueTeamName) IsSatisfied(value interface{}) bool {
  team := Team{}

  (*DB).Where("name = ?", value).FirstOrInit(&team)

  return team.Name == ""
}

func (u uniqueTeamName) DefaultMessage() string {
  return "Name Should be unique"
}
