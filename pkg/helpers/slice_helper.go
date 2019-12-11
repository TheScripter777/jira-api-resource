package helpers

func CleanStringSlice(slice []string) []string {
	var r = make([]string, 0)
	for _, str := range slice {
		if str != "" {
			r = append(r, str)
		}
	}

	return r
}
