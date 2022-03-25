package valueobject

//newuser describes the value object that should be passed when a user is created
type NewUser struct {
	FirstName		string
	LastName  		string
	Email    		string
	Password        string
}
