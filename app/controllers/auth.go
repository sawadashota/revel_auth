package controllers

import (
  "github.com/revel/revel"
  "api_player/app/models"
)

type Auth struct {
  *revel.Controller
}

func (c Auth) SignOut() revel.Result {
  for key := range c.Session {
    delete(c.Session, key)
  }

  return c.Redirect(SignIn.SignIn)
}

func CheckAuthentication(c *revel.Controller) revel.Result {
  token, hasToken := c.Session["token"]

  if !hasToken {
    c.Flash.Error("Please sign in at first.")
    return c.Redirect(SignIn.SignIn)
  }

  if !models.ExistsSessionUser(token) {
    return c.Redirect(SignIn.SignIn)
  }

  return nil
}

func CheckNotAuthentication(c *revel.Controller) revel.Result {
  _, hasToken := c.Session["token"]

  if !hasToken {
    return nil
  }

  return c.Redirect(Dashboard.Users)
}