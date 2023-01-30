package utils

func FilterOutString(s []string, substr string) []string {
	newStr := []string{}
	for _, v := range s {
		if v != substr {
			newStr = append(newStr, v)
		}
	}
	return newStr
}
