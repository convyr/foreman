package que

import (
	"flag"
	"time"

	nats "github.com/nats-io/go-nats"
)

var (
	NatsURL string
)

type Que struct {
	NatsClient *nats.Conn
	Channel    string
	Timeout    time.Duration
	Group      string
}

func Flags() {
	flag.StringVar(&NatsURL, "nats-url", nats.DefaultURL, "Url for Nats")
}

func New(ch string) (*Que, error) {
	q := Que{
		Channel: ch,
		Group:   fmt.Sprintf("%v-convyr", ch),
		Timeout: time.Second,
	}
	var err error
	q.NatsClient, err = nats.Connect(NatsURL)
	return &q, err
}

func (q *Que) Watch() error {
	go func() {
		_, err := nc.QueueSubscribe(q.Channel, q.Group, func(m *nats.Msg) {
			// To belt
			fmt.Prntln(string(m.Data))
			if m.Reply != "" {
				//belt results to go here
				nc.Publish(m.Reply, []byte("received"))
			}
		})
		if err != nil {
			return err
		}
	}()
}

func (q *Que) Close() {
	q.NatsClient.Close()
}
