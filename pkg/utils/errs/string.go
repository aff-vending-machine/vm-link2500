package errs

func ToString(err error) string {
	if err != nil {
		return err.Error()
	}

	return ""
}
