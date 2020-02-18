package helpers

func CopyMapString(source, destination map[string]string, overwrite bool) map[string]string {
	if source != nil {
		if destination == nil {
			destination = make(map[string]string, 0)
		}

		for k, v := range source {
			if (destination[k] != "" && overwrite) || destination[k] == "" {
				destination[k] = v
			}
		}

		return destination
	}

	return nil
}
