package goapifytask

import (
	"context"
	"log"
	"time"

	"github.com/data-harvesters/goapify"
)

type State int

const (
	Initialize State = -1
)

type Scraper interface {
	Next(state State) (State, error)

	Stop()

	Context() context.Context
}

type Base struct {
	actor *goapify.Actor

	context context.Context
	cancel  func()
}

func New(actor *goapify.Actor) *Base {
	ctx, cancel := context.WithCancel(actor.Context())

	return &Base{
		actor:   actor,
		context: ctx,
		cancel:  cancel,
	}
}

func (b *Base) Stop() {
	b.cancel()
}

func (b *Base) Context() context.Context {
	return b.context
}

func (b *Base) Actor() *goapify.Actor {
	return b.actor
}

func Run(s Scraper) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	state := Initialize
	var err error

	for {
		select {
		default:
			state, err = s.Next(state)
			if err != nil {
				return nil
			}
		case <-s.Context().Done():
			return nil
		}

		time.Sleep(100 * time.Millisecond)
	}
}
