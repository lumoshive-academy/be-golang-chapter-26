package service

import (
	"be-golang-chapter-26/LA-Chapter-26A-H/config"
	"be-golang-chapter-26/LA-Chapter-26A-H/notification"
)

// Notifier adalah struct yang menggunakan EmailService, SMSService, dan Config
type Notifier struct {
	EmailService *notification.EmailService
	SMSService   *notification.SMSService
	NotifConfig  *config.NotifConfig
}

// SendNotification mengirim notifikasi melalui email dan SMS
func (n *Notifier) SendNotification(message string) {
	n.EmailService.SendEmail(message)
	n.SMSService.SendSMS(message)
}
