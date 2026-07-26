package main

import (
	"fmt"
	"sync"
)

// =============================================================
// MODEL
// S → Message has one job: carry data between components
// =============================================================

type Message struct {
	Content    string
	SenderID   string
	ReceiverID string
}

// =============================================================
// STRATEGY PATTERN — INotifier + concrete channels
//
// S → each channel has one job: deliver via its medium
// O → add new channel = new struct, zero edits to existing code
// L → any INotifier can substitute any other
// D → callers depend on INotifier, not SMS/Gmail directly
// =============================================================

type INotifier interface {
	Send(msg Message)
}

type SMS struct{}

func (s SMS) Send(msg Message) {
	fmt.Println("  [SMS] →", msg.Content, "| to:", msg.ReceiverID)
}

type Gmail struct{}

func (g Gmail) Send(msg Message) {
	fmt.Println("  [Gmail] →", msg.Content, "| to:", msg.ReceiverID)
}

type WhatsApp struct{}

func (w WhatsApp) Send(msg Message) {
	fmt.Println("  [WhatsApp] →", msg.Content, "| to:", msg.ReceiverID)
}

// =============================================================
// DECORATOR PATTERN — Logging + Timestamp wrap INotifier
//
// S → Logging has one job: log before delivering
// O → add new decorator = new struct, nothing else changes
// L → LoggingDecorator IS-A INotifier, substitutable anywhere
// D → depends on INotifier abstraction, not concrete channel
// =============================================================

// LoggingDecorator — logs to prometheus before delivery
type LoggingDecorator struct {
	notifier INotifier // HAS-A INotifier (wraps it)
}

func (l LoggingDecorator) Send(msg Message) {
	fmt.Println("  [Prometheus Log] message:", msg.Content) // log first
	l.notifier.Send(msg)                                    // then deliver
}

// TimestampDecorator — adds timestamp to message before delivery
type TimestampDecorator struct {
	notifier INotifier
}

func (t TimestampDecorator) Send(msg Message) {
	msg.Content = "[2025-04-13 14:22:00] " + msg.Content // mutate content
	t.notifier.Send(msg)                                  // pass down chain
}

// =============================================================
// FACTORY PATTERN — INotifierFactory + concrete factories
//
// S → each factory has one job: create one type of notifier
// O → registry is open for extension, closed for modification
// L → any factory substitutes any other via interface
// D → callers depend on INotifierFactory, not SmsFactory directly
// =============================================================

type INotifierFactory interface {
	Create() INotifier
}

type SmsFactory struct{}

func (s SmsFactory) Create() INotifier { return SMS{} }

type GmailFactory struct{}

func (g GmailFactory) Create() INotifier { return Gmail{} }

type WhatsAppFactory struct{}

func (w WhatsAppFactory) Create() INotifier { return WhatsApp{} }

// O → registry: add new channel = one new line here, zero edits elsewhere
var channelRegistry = map[string]INotifierFactory{
	"sms":      SmsFactory{},
	"gmail":    GmailFactory{},
	"whatsapp": WhatsAppFactory{},
}

func GetFactory(channel string) (INotifierFactory, bool) {
	f, ok := channelRegistry[channel]
	return f, ok
}

// =============================================================
// SINGLETON PATTERN — NotificationManager
//
// S → one job: orchestrate delivery through factory + decorator
// L → returns INotificationManager interface, not concrete type
// D → Subscriber depends on INotificationManager abstraction
// =============================================================

// INotificationManager — abstraction callers depend on (D)
type INotificationManager interface {
	Send(msg Message, factories []INotifierFactory)
}

type NotificationManager struct {
	notifier INotifier
}

// Send — creates channel via factory, wraps in decorators, delivers
// O → adding a new decorator here is the only change needed
func (nm *NotificationManager) Send(msg Message, factories []INotifierFactory) {
	for _, factory := range factories {
		channel := factory.Create() // get concrete channel (SMS/Gmail/WhatsApp)

		// stack decorators — timestamp first, then logging, then deliver
		// read bottom-up: TimestampDecorator → LoggingDecorator → channel
		notifier := LoggingDecorator{
			notifier: TimestampDecorator{
				notifier: channel,
			},
		}
		notifier.Send(msg) // fires: log → timestamp → deliver
	}
}

// Singleton plumbing — sync.Once guarantees thread safety
var (
	nmInstance *NotificationManager
	nmOnce     sync.Once
)

// GetNotificationManager — returns interface not concrete type (L, D)
func GetNotificationManager() INotificationManager {
	nmOnce.Do(func() {
		nmInstance = &NotificationManager{}
	})
	return nmInstance
}

// =============================================================
// OBSERVER PATTERN — IObserver, INotificationObserver,
//                    IObservable, Subscriber, Youtuber
//
// I → split IObserver into base + notification-specific (ISP)
// S → Youtuber's job: manage subscribers and broadcast
// S → Subscriber's job: receive update and trigger delivery
// D → Subscriber depends on INotificationManager, not concrete
// =============================================================

// IObserver — base contract every observer must satisfy (I)
type IObserver interface {
	Update(msg Message)
	GetID() string
}

// INotificationObserver — extended contract for observers
// that also send notifications (I — only who needs it implements it)
type INotificationObserver interface {
	IObserver                             // embeds base
	GetFactories() []INotifierFactory     // only notification observers need this
}

// Subscriber — implements INotificationObserver
type Subscriber struct {
	Name         string
	SubscriberID string
	Channels     []string             // ["sms", "gmail"] — just names, not factories
	manager      INotificationManager // D → injected abstraction, not concrete type
}

func (s Subscriber) GetID() string { return s.SubscriberID }

// GetFactories — resolves channel names to factories via registry (O)
// Adding new channel = new entry in channelRegistry, nothing here changes
func (s Subscriber) GetFactories() []INotifierFactory {
	var factories []INotifierFactory
	for _, ch := range s.Channels {
		if f, ok := GetFactory(ch); ok {
			factories = append(factories, f)
		} else {
			fmt.Println("  [Warning] unknown channel:", ch)
		}
	}
	return factories
}

// Update — one job: receive notification and trigger delivery (S)
func (s Subscriber) Update(msg Message) {
	fmt.Printf("\n%s notified via %v:\n", s.Name, s.Channels)
	s.manager.Send(msg, s.GetFactories()) // depends on interface (D)
}

// Logger — implements base IObserver only (I — doesn't need GetFactories)
// S → one job: log that a notification was broadcast
type Logger struct {
	ID string
}

func (l Logger) GetID() string { return l.ID }
func (l Logger) Update(msg Message) {
	fmt.Println("\n[Logger] broadcast detected:", msg.Content)
}

// IObservable — contract for the subject (YouTuber)
type IObservable interface {
	Add(obs IObserver)
	Remove(obs IObserver)
	Notify(msg Message)
}

// Youtuber — implements IObservable
// S → one job: manage observer list and broadcast
type Youtuber struct {
	ChannelName string
	Name        string
	subscribers []IObserver // stores base interface (I — not INotificationObserver)
}

func (y *Youtuber) Add(obs IObserver) {
	y.subscribers = append(y.subscribers, obs)
	fmt.Println(obs.GetID(), "subscribed to", y.ChannelName)
}

func (y *Youtuber) Remove(obs IObserver) {
	idx := -1
	for i, val := range y.subscribers {
		if val.GetID() == obs.GetID() {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println("subscriber not found:", obs.GetID())
		return
	}
	y.subscribers = append(y.subscribers[:idx], y.subscribers[idx+1:]...)
	fmt.Println(obs.GetID(), "unsubscribed from", y.ChannelName)
}

// Notify — broadcasts to ALL observers (S — just loops and delegates)
func (y *Youtuber) Notify(msg Message) {
	fmt.Println("\n--- YouTuber", y.Name, "broadcasting ---")
	for _, obs := range y.subscribers {
		obs.Update(msg)
	}
}

// =============================================================
// MAIN — wires everything together
// D → concrete types only created here, at the top level
//      everything below depends on interfaces
// =============================================================

func main() {
	// singleton notification manager — injected into subscribers
	nm := GetNotificationManager()

	// subscribers — each with their own channel preferences
	s1 := Subscriber{
		Name:         "Prajwal",
		SubscriberID: "sub-001",
		Channels:     []string{"sms", "gmail"},   // two channels
		manager:      nm,                          // D → injected
	}

	s2 := Subscriber{
		Name:         "Jay",
		SubscriberID: "sub-002",
		Channels:     []string{"whatsapp"},        // one channel
		manager:      nm,
	}

	s3 := Subscriber{
		Name:         "Rahul",
		SubscriberID: "sub-003",
		Channels:     []string{"sms", "whatsapp", "gmail"}, // all three
		manager:      nm,
	}

	// logger — base observer, no notification delivery (I)
	logger := Logger{ID: "logger-001"}

	// youtuber — the subject
	flyingBeast := &Youtuber{
		ChannelName: "Flying Beast",
		Name:        "Gaurav",
	}

	// add all observers
	flyingBeast.Add(s1)
	flyingBeast.Add(s2)
	flyingBeast.Add(s3)
	flyingBeast.Add(logger) // Logger satisfies IObserver ✅

	// broadcast — one call fans out to all subscribers
	flyingBeast.Notify(Message{
		Content:    "New video uploaded — India Trip Vlog!",
		SenderID:   "flying-beast",
		ReceiverID: "all-subscribers",
	})

	fmt.Println("\n--- Prajwal unsubscribes ---")
	flyingBeast.Remove(s1)

	// second broadcast — s1 no longer receives
	flyingBeast.Notify(Message{
		Content:    "Second video — Goa Trip!",
		SenderID:   "flying-beast",
		ReceiverID: "all-subscribers",
	})
}