package util

// AppendError only if the error is not nil
func AppendError(es []error, e error) []error {
	if es != nil && e != nil {
		es = append(es, e)
	}

	return es
}
