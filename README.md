## Variables


Use `const` for variables, that shouldn't change
Use `var` for variables, that can be changed

When asking for an user input, make sure to define, what you're going to use e.g
`var userName string` here we're asking for a username
`var UserTickets int` integrer, since we're asking for a number

#Printing

Use `Printf` to use variables in ur print methods
Use `Println` to print the text on a new line

e.g `fmt.Printf("Welcome %v \n", userAccountName)`

# User Inputs

`fmt.Scan(&<variableName>)`
We're passing the refrence

use `uint` for positive number variables

# Arrays

`var variableName [size]variableType`

e.g `var bookings = [50]string{}`