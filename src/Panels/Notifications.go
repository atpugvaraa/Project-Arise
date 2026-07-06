package panels

func NewNotificationPanel(message string) Panel {
	return Panel{
		Kind:  Notification,
		Title: "NOTIFICATION",
		Lines: []string{
			message,
		},
	}
}
