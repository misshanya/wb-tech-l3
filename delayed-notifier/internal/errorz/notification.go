package errorz

import "errors"

var (
	NotificationNotFound         = errors.New("notification not found")
	ChannelNotFoundError         = errors.New("channel not found")
	NotificationIsNotCancellable = errors.New("notification is not cancellable")
)
