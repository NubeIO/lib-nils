package nils

// StringIsNil check if value is nil
func StringIsNil(b *string) bool {
	if b == nil {
		return true
	} else {
		return false
	}
}

// Float64IsNil check if value is nil
func Float64IsNil(b *float64) bool {
	if b == nil {
		return true
	} else {
		return false
	}
}

// Float32IsNil check if value is nil
func Float32IsNil(b *float32) bool {
	if b == nil {
		return true
	} else {
		return false
	}
}

// BoolIsNil check if value is nil
func BoolIsNil(b *bool) bool {
	if b == nil {
		return true
	} else {
		return false
	}
}
