package errors

import "slices"

func Any(err error, targets ...error) bool {
	return slices.ContainsFunc(targets, func(target error) bool {
		return Is(err, target)
	})
}

func All(err error, targets ...error) bool {
	var exists = false

	for _, target := range targets {
		exists = false

		if Is(err, target) {
			exists = true
		}

		if !exists {
			break
		}
	}
	return exists
}
