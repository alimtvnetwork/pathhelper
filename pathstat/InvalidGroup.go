package pathstat

func InvalidGroup() *Group {
	return &Group{
		IntIdNameValidation: *InvalidIntIdNameValidation(),
	}
}
