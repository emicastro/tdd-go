package main

import "fmt"

type INotificationFactory interface {
  SendNotification()
  GetSender() ISender
}

type ISender interface {
  GetSenderMethod() string
  GetSenderChannel() string
}

type SMSNotification struct {
}

type EmailNotification struct {
}

type SMSNotificationSender struct {
}

type EmailNotificationSender struct {}

func (sn SMSNotification) SendNotification() {
  fmt.Println("Sending Notification via SMS")
}

func (sn SMSNotification) GetSender() ISender {
  return SMSNotificationSender{}
}

func (en EmailNotification) SendNotification() {
  fmt.Println("Sending Notification via Email")
}

func (en EmailNotification) GetSender() ISender {
  return EmailNotificationSender{}
}

func (sender SMSNotificationSender) GetSenderMethod() string {
  return "SMS"
}

func (sender SMSNotificationSender) GetSenderChannel() string {
  return "Twillio"
}

func (sender EmailNotificationSender) GetSenderMethod() string {
  return "Email"
}

func (sender EmailNotificationSender) GetSenderChannel() string {
  return "SES"
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
  if notificationType == "SMS" {
    return &SMSNotification{}, nil
  }

  if notificationType == "Email" {
    return &EmailNotification{}, nil
  }

  return nil, fmt.Errorf("No notification type")
}

func sendNotification(f INotificationFactory) {
  f.SendNotification()
}

func getMethod(f INotificationFactory) {
  fmt.Println(f.GetSender().GetSenderMethod())
}

func main () {
  smsFactory, _ := getNotificationFactory("SMS")
  emailFactory, _ := getNotificationFactory("Email")

  sendNotification(smsFactory)
  sendNotification(emailFactory)
  getMethod(smsFactory)
  getMethod(emailFactory)
}
