package alarm

func SendAllAlarm(message string) {
	alarm := NewAlarm()
	emailNotifier := &EmailNotifier{Recipient: "majunhong@gmail.com"}
	logNotifier := &LogNotifier{}
	alarm.RegisterObservers(emailNotifier, logNotifier)
	go alarm.NotifyObservers(message)
}
