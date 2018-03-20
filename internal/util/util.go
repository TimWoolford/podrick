package util

func FirstPopulated(vals... string) string{
	for _, val := range vals {
		if len(val) > 0 {
			return val
		}
	}
	return ""
}