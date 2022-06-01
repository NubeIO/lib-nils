package nils

// Bool returns a pointer to a variable holding the supplied bool constant
func Bool(x bool) *bool {
	return &x
}

// BoolValue returns the bool value pointed to by p
func BoolValue(p *bool) bool {
	return *p
}

// Float32 returns a pointer to a variable holding the supplied float32 constant
func Float32(x float32) *float32 {
	return &x
}

// Float32Value returns the float32 value pointed to by p
func Float32Value(p *float32) float32 {
	return *p
}

// Float64 returns a pointer to a variable holding the supplied float64 constant
func Float64(x float64) *float64 {
	return &x
}

// Float64Value returns the float64 value pointed to by p
func Float64Value(p *float64) float64 {
	return *p
}

// Int returns a pointer to a variable holding the supplied int constant
func Int(x int) *int {
	return &x
}

// IntValue returns the int value pointed to by p
func IntValue(p *int) int {
	return *p
}

// Int8 returns a pointer to a variable holding the supplied int8 constant
func Int8(x int8) *int8 {
	return &x
}

// Int8Value returns the int8 value pointed to by p
func Int8Value(p *int8) int8 {
	return *p
}

// Int16 returns a pointer to a variable holding the supplied int16 constant
func Int16(x int16) *int16 {
	return &x
}

// Int16Value returns the int16 value pointed to by p
func Int16Value(p *int16) int16 {

	return *p
}

// Int32 returns a pointer to a variable holding the supplied int32 constant
func Int32(x int32) *int32 {
	return &x
}

// Int32Value returns the int32 value pointed to by p
func Int32Value(p *int32) int32 {
	return *p
}

// Int64 returns a pointer to a variable holding the supplied int64 constant
func Int64(x int64) *int64 {
	return &x
}

// Int64Value returns the int64 value pointed to by p
func Int64Value(p *int64) int64 {
	return *p
}

// Uint returns a pointer to a variable holding the supplied uint constant
func Uint(x uint) *uint {
	return &x
}

// UintValue returns the uint value pointed to by p
func UintValue(p *uint) uint {
	return *p
}

// Uint8 returns a pointer to a variable holding the supplied uint8 constant
func Uint8(x uint8) *uint8 {
	return &x
}

// Uint8Value returns the uint8 value pointed to by p
func Uint8Value(p *uint8) uint8 {
	return *p
}

// Uint16 returns a pointer to a variable holding the supplied uint16 constant
func Uint16(x uint16) *uint16 {
	return &x
}

// Uint16Value returns the uint16 value pointed to by p
func Uint16Value(p *uint16) uint16 {
	return *p
}

// Uint32 returns a pointer to a variable holding the supplied uint32 constant
func Uint32(x uint32) *uint32 {
	return &x
}

// Uint32Value returns the uint32 value pointed to by p
func Uint32Value(p *uint32) uint32 {

	return *p
}

// Uint64 returns a pointer to a variable holding the supplied uint64 constant
func Uint64(x uint64) *uint64 {
	return &x
}

// Uint64Value returns the uint64 value pointed to by p
func Uint64Value(p *uint64) uint64 {
	return *p
}

// String returns a pointer to a variable holding the supplied string constant
func String(x string) *string {
	return &x
}

// StringValue returns the string value pointed to by p
func StringValue(p *string) string {
	return *p
}
