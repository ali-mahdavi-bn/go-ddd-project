package event_dispatcher

import (
	"github.com/rozturac/rmqc"
	"go-e-s/src/backbone/application/consts"
	"go-e-s/src/backbone/service_layer"
	"reflect"
)

type RabbitMQEventDispatcher struct {
	rbt     *rmqc.RabbitMQ
	appName string
}

func NewRabbitMQEventDispatcher(rbt *rmqc.RabbitMQ) service_layer.IEventDispatcher {
	return &RabbitMQEventDispatcher{rbt: rbt}
}

func (handler RabbitMQEventDispatcher) Dispatch(events []service_layer.IBaseEvent) {
	for _, event := range events {
		t := reflect.TypeOf(event)
		eventName := t.Elem().Name()
		handler.rbt.Publish(consts.AppName, eventName, event)
	}
}
