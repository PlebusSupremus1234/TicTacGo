package main

func bold(text string) string { return "\x1b[1m" + text + "\x1b[0m" }
func red(text string) string  { return "\x1b[31m" + text + "\x1b[0m" }
func blue(text string) string { return "\x1b[34m" + text + "\x1b[0m" }

func inArray(array []string, value string) bool {
	for _, i := range array {
        if i == value {
            return true
        }
    }
    return false
}

func indexOf(array []string, value string) int {
	for k, v := range array {
		if value == v {
			return k
		}
	}

	return -1
 }