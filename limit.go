package ff

func Limit(max int) func(str string) string {
	return func(str string) string {
		if len(str) > max {
			str = str[0:max]
		}
		return str
	}
}
