package watermark

import (
	"context"
)

type Service interface {
	Get(ctx context.Context, filters ...Filter) ([]Document, error)
	Status(ctx context.Context, ticketID string) (Status, error)
	Watermark(ctx context.Context, ticketID, mark string) (int, error)
	AddDocument(ctx context.Context, doc *Document) (string, error)
	ServiceStatus(ctx context.Context) (int, error)
}
