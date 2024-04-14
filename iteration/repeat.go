package iteration

func Repeat(a string, count int) string {
	repeater := ""
	for i := 0; i < count; i++ {
		repeater += a
	}
	return repeater
}
