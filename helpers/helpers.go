package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

/**
* HandleError function check error
 */
func HandleError(err error) {

	if err != nil {
		panic(err.Error())
	}
}

/* HashAndSalt ...
*  is test hashed function
 */
func HashAndSalt(pass []byte) string {

	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)

	HandleError(err)
	return string(hashed)
}
