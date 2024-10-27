package iterate

func Map[A any, B any](collection []A, function func(item A) B) []B {
	var result = make([]B, len(collection))

	for i := range collection {
		result[i] = function(collection[i])
	}

	return result
}
