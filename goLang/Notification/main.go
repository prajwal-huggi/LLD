package main

import (
	"fmt"
	"sync"
)

// model
type Message struct{
	Message string
	SenderId string
	ReceiverId string
}
// Strategy design pattern
//interface
type INotifier interface{
	Send(msg Message)
}

type Sms struct{}
func(s Sms) Send(msg Message){
	fmt.Println("Message: ", msg.Message," is sent via SMS")
}

type Gmail struct{}
func(s Gmail) Send(msg Message){
	fmt.Println("Message: ", msg.Message," is sent via Gmail")
}

type Logging struct{
	notifier INotifier
}
func(l Logging) Send(msg Message){
	fmt.Println("Logged in prometheus")
	l.notifier.Send(msg)
}

// Observer Design Pattern

type IObserver interface{
	update(msg Message)
	GetId() string
	GetFactory() INotifierFactory
}
type Subscriber struct{
	Name string
	SubscriberId string
	Channel string //"sms", "gmail"
}
func(s Subscriber) update(msg Message){
	fmt.Println(s.Name, "received:", msg.Message)

	factory:= s.GetFactory()
	if factory == nil {
		fmt.Println("unknown channel:", s.Channel)
		return
	}

	nm:= GetNotificationManager()
	nm.Send(msg, factory)
}
func (s Subscriber) GetId() string{
	return s.SubscriberId
}
func(s Subscriber) GetFactory()INotifierFactory{
	switch s.Channel {
    case "sms":   return SmsFactory{}
    case "gmail": return GmailFactory{}
    default:      return nil
    }
}

type IObservable interface{
	add(obs IObserver)
	remove(obs IObserver)

	notify(msg Message)
}

type Youtuber struct{
	YoutubeChannel string
	Name string
	Subscribers []IObserver
}
func(y *Youtuber) add(sub IObserver){
	y.Subscribers= append(y.Subscribers, sub)
}
func(y *Youtuber) remove(sub IObserver){
	subs:= y.Subscribers
	subId:= sub.GetId()

	idx:= -1
	for i, val:= range subs{
		subsId:= val.GetId()
		if subsId== subId{
			idx= i
			break;
		}
	}

	if idx== -1{
		fmt.Println("Subscriber haven't subscribed the youtuber")
		return
	}

	y.Subscribers= append(subs[:idx], subs[idx+ 1:]...)
	fmt.Println(sub.GetId()," has unsubscribed ", y.Name)
}

func(y *Youtuber) notify(msg Message){
	for _, val:= range y.Subscribers{
		val.update(msg)
	}
}

// Factory Design Pattern
type INotifierFactory interface{
	Create() INotifier
}

type SmsFactory struct{}
func(s SmsFactory) Create() INotifier{
	return Sms{}
}

type GmailFactory struct{}
func(g GmailFactory) Create() INotifier{
	return Gmail{}
}

// Sigleton Design Pttern
var (
    instance *NotificationManager // starts as nil
    once     sync.Once  // a lock that fires exactly once
)

type NotificationManager struct{}

func GetNotificationManager() *NotificationManager {
    once.Do(func() {
        instance = &NotificationManager{}
    })
    return instance
}

func (manager *NotificationManager) Send(msg Message, factory INotifierFactory) {
	channel := factory.Create()
	notifier := Logging{notifier: channel}  // local, not manager.notifier
	notifier.Send(msg)
}



func main(){

	s1:= Subscriber{
		Name:"Prajwal",
		SubscriberId:"1st",
		Channel:"sms",
	}

	s2:= Subscriber{
		Name:"Jay",
		SubscriberId:"2nd",
		Channel:"gmail",
	}
	
	flyingBeast:= Youtuber{
		YoutubeChannel:"Flying Beast",
		Name:"Gaurav",
	}

	flyingBeast.add(s1)
	flyingBeast.add(s2)

	flyingBeast.notify(Message{
        Message:    " My first youtube video!",
        SenderId:   "yt1",
        ReceiverId: "all",
    })

	fmt.Println("Hello world")
}