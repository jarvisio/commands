package commands

// Now returns the current time as a string
// you can choose between 12 and 24 hour time format
// using the use24 parameter.
func Clock(use24 bool) string {
	layout := "3:04PM"
	if use24 {
		layout = "15:04"
	}
	t := currentTime()
	return t.Format(layout)
}
