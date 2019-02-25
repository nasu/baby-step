package diff

func IsEmptyUsingLen(s string) bool {
	return len(s) == 0
}

func IsEmptyUsingEqual(s string) bool {
	return s == ""
}
