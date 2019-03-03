var DEBUG = false

/*
	Append the debug logs
*/
func Debugf(templates string, args ...interface{}) {
	if !DEBUG {
		return
	}

	original_message := fmt.Sprintf((templates), args...)
	fmt.Printf("[Debug] " + original_message)
}

/*
	Append the error logs
*/
func Errorf(templates string, args ...interface{}) {
	original_message := fmt.Sprintf((templates), args...)
	fmt.Printf("[Error] " + original_message)
}
