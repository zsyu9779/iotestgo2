package definedtypeassignment

type UserID int

func demonstrate() {
	var raw int = 1
	var id UserID = raw
	_ = id
}
