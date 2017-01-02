package models


func GetUserByEmail(email string) (user User) {
  (*DB).Limit(1).Find(&user, "email = ?", email).Related(&user.Team)
  return
}

func GetUserByToken(token string) (user User) {
  (*DB).Limit(1).Find(&user, "remember_token = ?", token).Related(&user.Team)
  return
}


func AllUsers() (users []User) {
  (*DB).Find(&users)
  return
}

func CreateUser(user User) {
  (*DB).Create(&user)
}

func SeedRememberToken(tokenUser TokenUser) {
  (*DB).Table("users").Where("email = ?", tokenUser.Email).Update("remember_token", tokenUser.RememberToken)
}

func ExistsSessionUser(token string) bool {
  var count int
  (*DB).Model(&User{}).Where("remember_token = ?", token).Count(&count)
  return count == 1
}