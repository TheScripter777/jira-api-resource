package helpers

func IsStringPtrNilOrEmtpy(ptr *string) bool {
	return ptr == nil || *ptr == ""
}

func IsBoolPtrTrue(ptr *bool) bool {
	return ptr != nil && *ptr == true
}
