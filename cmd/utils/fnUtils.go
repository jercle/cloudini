package utils

import "github.com/jercle/cloudini/lib"

func GeneratePassword(length int, includeUpper bool, includeNumbers bool, includeSpecial bool) string {
	pwd, err := lib.GenerateRandomString(length, includeUpper, includeNumbers, includeSpecial)
	lib.CheckFatalError(err)
	return pwd
}
