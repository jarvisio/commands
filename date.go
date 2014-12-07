package commands

// date returns the current date as a string
func Date() string {
	const layout = "January 2, 2006"
	timeStamp := currentTime()
	return timeStamp.UTC().Format(layout)
}
