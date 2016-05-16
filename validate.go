package Validate

import "fmt"

//Defigning a vlidate function which return a isValid and message
type  Validate func(value string) (message string, isValid bool)

//Accept as value from clid and validate func which must return true to proceed..
/*
    Example
    GetInput(&id, "Enter user Id: ", func(value string)(message string, isValid bool){
	return "User Id must be a number", IsNumeric(value)

    })
 */
func GetInput(value *string, message string,  validFn Validate){
	isValid := false
	var validMsg string
	for !isValid{
		fmt.Println(message)
		fmt.Scanln(value)
		validMsg, isValid = validFn(*value)
		if !isValid{
			fmt.Println(validMsg)
		}
	}
}
