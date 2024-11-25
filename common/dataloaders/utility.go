package dataloaders

func sliceToMap[T Model](objects ...T) map[string]T {
	var (
		result = make(map[string]T, len(objects))
	)

	for _, v := range objects {
		result[v.GetID()] = v
	}

	return result
}
