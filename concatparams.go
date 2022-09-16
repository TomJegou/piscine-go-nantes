package piscine

func ConcatParams(args []string) string {
	var result string
	for i := 0; i < len(args)-1; i++ {
		result += args[i]
		result += "\n"
	}
	result += args[len(args)-1]
	return result
}
