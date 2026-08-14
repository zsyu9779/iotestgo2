package solution

func FindByName(name string) (string, []any) {
	return "SELECT id, name FROM m04_l06_users WHERE name = ?", []any{name}
}
