package ports

import "github.com/eliastarin/ssp2-southpark/go-api/domain"

type MessagePublisher interface {
	Publish(msg domain.Message) error
}
