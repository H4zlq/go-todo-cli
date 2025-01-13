package utils

func FormatStatus(status string) string {
	switch status {
	case "in-progress":
		return "In Progress"
	case "done":
		return "Done"
	default:
		return "Todo"
	}
}