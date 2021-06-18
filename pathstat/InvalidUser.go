package pathstat

func InvalidUser() *User {
	return &User{
		IntIdNameValidation: *InvalidIntIdNameValidation(),
	}
}
