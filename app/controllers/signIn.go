package controllers

import (
  "github.com/revel/revel"
  "api_player/app/models"
  "golang.org/x/crypto/bcrypt"
  "time"
  "fmt"
  "crypto/sha256"
  "math/rand"
)

type SignIn struct {
  *revel.Controller
}

func init() {
  revel.InterceptFunc(CheckNotAuthentication, revel.BEFORE, &SignIn{})
}

func (c SignIn) Authentication(email, password string) revel.Result {
  user := models.GetUserByEmail(email)

  if user.Email != "" {

    err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))

    if err == nil {

      tokenUser := models.TokenUser{
        Email: email,
        RememberToken: generateToken(),
      }

      c.Session["token"] = tokenUser.RememberToken
      models.SeedRememberToken(tokenUser)

      c.Flash.Success("Welcome! " + user.Name)
      return c.Redirect(Dashboard.Users)
    }
  }

  c.Flash.Out["mail"] = email
  c.Flash.Error("Incorrect email address or password.")
  return c.Redirect(SignIn.SignIn)
}

func (c SignIn) SignIn() revel.Result {
  return c.Render()
}

func (c SignIn) SignUp() revel.Result {
  return c.Render()
}

func (c SignIn) Register(mockUser *models.MockUser) revel.Result {
  mockUser.TeamID = 2
  mockUser.Role = "admin"

  mockUser.Validate(c.Validation)

  if c.Validation.HasErrors() {
    c.Validation.Keep()
    c.FlashParams()
    return c.Redirect(SignIn.SignUp)
  }

  if mockUser.Password != mockUser.PasswordConfirm {
    c.Flash.Error("Passwords are unmatched")
    c.FlashParams()
    return c.Redirect(SignIn.SignUp)
  }

  var user models.User

  user.TeamID = mockUser.TeamID
  user.Name = mockUser.Name
  user.Email = mockUser.Email
  user.Password, _ = bcrypt.GenerateFromPassword([]byte(mockUser.Password), bcrypt.DefaultCost)
  user.Role = mockUser.Role

  models.CreateUser(user);

  return c.Redirect(SignIn.SignIn)
}

func generateToken() string {
  var token string

  for {
    rand.Seed(time.Now().UnixNano())
    token = fmt.Sprintf("%x", sha256.Sum256([]byte(string(rand.Intn(15)))))

    if !models.ExistsSessionUser(token) {
      break
    }
  }

  return token
}