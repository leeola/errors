package errors

import "bytes"

const (
	defaultErrorsSep = "; "
)

func Join(errs []error) error {
	if len(errs) == 0 {
		return nil
	}

	return multiErr{
		Errs: errs,
		Sep:  defaultErrorsSep,
	}
}

func JoinSep(errs []error, sep string) error {
	if len(errs) == 0 {
		return nil
	}

	return multiErr{
		Errs: errs,
		Sep:  sep,
	}
}

func Split(err error) []error {
	sErr, ok := err.(multiErr)
	if !ok {
		return []error{err}
	}

	return sErr.Errs
}

type multiErr struct {
	Errs []error
	Sep  string
}

func (e multiErr) Error() string {
	var buf bytes.Buffer

	for i, err := range e.Errs {
		if i != 0 {
			buf.WriteString(e.Sep)
		}

		buf.WriteString(err.Error())
	}

	return buf.String()
}
