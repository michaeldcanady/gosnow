package gosnow

import(
  "golang.org/x/crypto/bcrypt"
  "log"
)

//Used to hash password stored in client struct
func hashPassword(password string) (string){
  pwd := []byte(password)
  // Use GenerateFromPassword to hash & salt pwd.
    // MinCost is just an integer constant provided by the bcrypt
    // package along with DefaultCost & MaxCost.
    // The cost can be any value you want provided it isn't lower
    // than the MinCost (4)
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
    // GenerateFromPassword returns a byte slice so we need to
    // convert the bytes to a string and return it
    return string(hash)
}

//Used to compare re-entered password to stores
//Used if session times out
func comparePasswords(hashedPwd string, pwd string) (bool, string) {
  plainPwd := []byte(pwd)
  // Since we'll be getting the hashed password from the DB it
    // will be a string so we'll need to convert it to a byte slice
    byteHash := []byte(hashedPwd)
    err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
    if err != nil {
        log.Println(err)
        return false, ""
    }

    return true, pwd
}
