package time_diff

import "time"

func After(t1, t2 time.Time) bool {
	return t1.After(t2)
}
func EqualAfter(t1, t2 time.Time) bool {
	return t1.Equal(t2) || t1.After(t2)
}
func Unix(t1, t2 time.Time) bool {
	return t1.Unix() > t2.Unix()
}
func EqualUnix(t1, t2 time.Time) bool {
	return t1.Unix() >= t2.Unix()
}
func UnixNano(t1, t2 time.Time) bool {
	return t1.UnixNano() > t2.UnixNano()
}
func EqualUnixNano(t1, t2 time.Time) bool {
	return t1.UnixNano() >= t2.UnixNano()
}
