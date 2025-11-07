package adapters

import (
	"log"

	"github.com/eliastarin/ssp2-southpark/go-api/domain"
	"github.com/eliastarin/ssp2-southpark/go-api/ports"
)

type MemoryPublisher struct {
	Buffer []domain.Message
}

var _ ports.MessagePublisher = (*MemoryPublisher)(nil)

func NewMemoryPublisher() *MemoryPublisher {
	return &MemoryPublisher{Buffer: make([]domain.Message, 0)}
}

func (m *MemoryPublisher) Publish(msg domain.Message) error {
	m.Buffer = append(m.Buffer, msg)
	log.Printf("[memory-publisher] stored message: %+v", msg)
	return nil
}
