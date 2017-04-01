package custom

import "strconv"

// MaxCustomTypeTagFn check if param max return
func MaxCustomTypeTagFn(i interface{}, ctx interface{}, params ...string) bool {
	if len(params) == 0 {
		return true
	}

	max, err := strconv.Atoi(params[0])
	if err != nil {
		return true
	}

	switch v := i.(type) {
	case int:
		return v <= max
	case string:
		return len(v) <= max
	default:
	}

	return true
}

// MinCustomTypeTagFn check if param min return true, else false
func MinCustomTypeTagFn(i interface{}, ctx interface{}, params ...string) bool {
	if len(params) == 0 {
		return true
	}

	min, err := strconv.Atoi(params[0])
	if err != nil {
		return true
	}

	switch v := i.(type) {
	case int:
		return v >= min
	case string:
		return len(v) >= min
	default:
	}

	return true
}

// GtCustomTypeTagFn check if grate equal return true, else false
func GtCustomTypeTagFn(i interface{}, ctx interface{}, params ...string) bool {
	if len(params) == 0 {
		return true
	}

	min, err := strconv.Atoi(params[0])
	if err != nil {
		return true
	}

	switch v := i.(type) {
	case int:
		return v > min
	case string:
		return len(v) > min
	default:
	}

	return true
}

// LtCustomTypeTagFn check if less than return true, else false
func LtCustomTypeTagFn(i interface{}, ctx interface{}, params ...string) bool {
	if len(params) == 0 {
		return true
	}

	max, err := strconv.Atoi(params[0])
	if err != nil {
		return true
	}

	switch v := i.(type) {
	case int:
		return v < max
	case string:
		return len(v) < max
	default:
	}

	return true
}
