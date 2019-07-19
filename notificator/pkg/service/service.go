package service

import (
	"context"

	"github.com/twinj/uuid"
)

// NotificatorService describes the service.
type NotificatorService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	SendEmail(ctx context.Context, email string, content string) (string, error)
}

type basicNotificatorService struct{}

func (b *basicNotificatorService) SendEmail(ctx context.Context, email string, content string) (string, error) {
	// TODO implement the business logic of SendEmail

	id := uuid.NewV4()
	// if err != nil {
	// 	return "", err
	// }
	return id.String(), nil
}

// NewBasicNotificatorService returns a naive, stateless implementation of NotificatorService.
func NewBasicNotificatorService() NotificatorService {
	return &basicNotificatorService{}
}

// New returns a NotificatorService with all of the expected middleware wired in.
func New(middleware []Middleware) NotificatorService {
	var svc NotificatorService = NewBasicNotificatorService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
