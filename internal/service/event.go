package service

import (
	"encoding/json"
	"log"
)

type EventService struct {
	producer ProducerInterface
}

func NewEventService(producer ProducerInterface) *EventService {
	return &EventService{producer: producer}
}

func (e *EventService) PublishOrderStatusChanged(orderID int, oldStatus, newStatus string) {
	event := map[string]interface{}{
		"order_id":   orderID,
		"old_status": oldStatus,
		"new_status": newStatus,
	}

	message, err := json.Marshal(event)
	if err != nil {
		log.Printf("Failed to marshal event: %v\n", err)
		return
	}

	err = e.producer.Publish(nil, message)
	if err != nil {
		log.Printf("Failed to publish event: %v\n", err)
	} else {
		log.Printf("Event published: %s\n", message)
	}
}
