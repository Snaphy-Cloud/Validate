##Command line tool for scanning value and validating a value.

###Install

    go get github.com/Snaphy-Cloud/Validate


###Example
    //Scanning an Input email and checking its validity
    GetInput(&email, "Enter email address ", func(email string)(message string, isValid bool){
       return "Email must be a valid one", IsEmail(email)
    })
