package domain

type INotifier interface {
	NotifyUser(userID int32, message interface{}) error
	Broadcast(message interface{}) error
}