package integration

func F64(value float64) *float64 {
	if value == 0 {
		return nil
	}
	return &value
}
