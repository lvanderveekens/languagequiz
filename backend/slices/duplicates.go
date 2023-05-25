package slices

func FindDuplicate[T string | int](sliceList []T) *T {
	allKeys := make(map[T]bool)
	for _, item := range sliceList {
		_, value := allKeys[item]
		if value {
			return &item
		} else {
			allKeys[item] = true
		}
	}
	return nil
}
