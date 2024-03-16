package adapter

import (
	"go-e-s/src/backbone/service_layer"
	"time"
)

type IBaseEntity interface {
	Table() string
}

type BaseEntity struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
	domainEvents []service_layer.IBaseEvent `json:"domain_events"`
}

func (u *BaseEntity) ClearDomainEvents() {
	u.domainEvents = nil
}
func (u *BaseEntity) GetDomainEvents() []service_layer.IBaseEvent {
	return u.domainEvents
}

func (u *BaseEntity) AddEvent(event service_layer.IBaseEvent) {
	u.domainEvents = append(u.domainEvents, event)
}
