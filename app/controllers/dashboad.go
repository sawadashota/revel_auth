package controllers

import (
  "github.com/revel/revel"
  "api_player/app/models"
)

type Dashboard struct {
  *revel.Controller
}

func init() {
  revel.InterceptFunc(CheckAuthentication, revel.BEFORE, &Dashboard{})
}

func (c Dashboard) Users() revel.Result {
  auth := models.GetUserByToken(c.Session["token"])
  users := models.AllUsers()

  return c.Render(users, auth)
}