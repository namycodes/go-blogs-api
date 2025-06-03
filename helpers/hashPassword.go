package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error){
  hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)

  if err != nil{
     return nil, err
  }

  return hash, nil
}

func VerifyPassword(hashedPassword []byte ,password string) (bool, error){
  err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))

  if err != nil {
    return false, err
  }

  return true, nil
}

