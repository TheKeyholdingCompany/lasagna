package helpers

func RemoveElements[T string](array []T, element T) []T {
	var r []T
	for _, el := range array {
		if el != element {
			r = append(r, el)
		}
	}
	return r
}
