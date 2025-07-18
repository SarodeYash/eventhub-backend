package util

import "golang.org/x/crypto/bcrypt"

func ToHahshPass(pass string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	return string(bytes), err

}

/*bcrypt.GenerateFromPassword used to provide as an hash password
which can be stored in Database instead of original password
So if the database is compromised then user's original password
should not be there.*/

func CompareHashPassword(password string, retrivedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(retrivedPassword), []byte(password))
	return err == nil
}

/*This function gets two inputs as first a plain password and the
another is the hashed or retrived from the database password
and then it compares it and check if they are same or not
if its same then err is nil if its not same err is nul nil*/
