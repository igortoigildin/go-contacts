package safecall

// To be called if panic happens
type RecoverFunc func(interface{}) error

// Standard handler of panic, which calls user custom func.
func Catch(onRecover RecoverFunc) {
	if r := recover(); r != nil {
		onRecover(r)
	}
}
