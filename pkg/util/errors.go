package util

import "errors"

// AppendError only if the error is not nil
func AppendError(es []error, e error) []error {
	if es != nil && e != nil {
		es = append(es, e)
	}

	return es
}

// MergeErrors into a single error
func MergeErrors(es []error) error {
	if es == nil || (es != nil && len(es) == 0) {
		return nil
	}

	s := ""
	for _, e := range es {
		if e != nil {
			if len(s) > 0 {
				s += ". "
			}
			s += e.Error()
		}
	}

	return errors.New(s)
}
