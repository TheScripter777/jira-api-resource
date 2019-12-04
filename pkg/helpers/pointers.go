package helpers

func IsStringPtrNilOrEmtpy(ptr *string) bool {
	if ptr == nil || *ptr == "" {
		return true
	}

	return false
}
