package xcontrib

func IN(find string, ss ...string) bool {
	for _, s := range ss {
		if find == s {
			return true
		}
	}
	return false
}
