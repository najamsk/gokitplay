package service

import (
	"context"
	"log"

	sdetcd "github.com/go-kit/kit/sd/etcd"
	"github.com/najamsk/kitplay/first/notificator/pkg/grpc/pb"
	"google.golang.org/grpc"
)

// UsersService describes the service.
type UsersService interface {
	// Add your methods here
	Create(ctx context.Context, email string) error
}

type basicUsersService struct {
	notificatorServiceClient pb.NotificatorClient
}

func (b *basicUsersService) Create(ctx context.Context, email string) (e0 error) {
	// TODO implement the business logic of Create
	reply, err := b.notificatorServiceClient.SendEmail(context.Background(), &pb.SendEmailRequest{
		Email:   email,
		Content: "hi, thanks for registeration",
	})

	if reply != nil {
		log.Printf("Email ID: %s", reply.Id)
	}
	return err
}

// NewBasicUsersService returns a naive, stateless implementation of UsersService.
func NewBasicUsersService() UsersService {

	var (
		etcdServer = "http://etcd:2379"
		prefix     = "/services/notificator/"
	)
	client, err := sdetcd.NewClient(context.Background(), []string{etcdServer}, sdetcd.ClientOptions{})
	// client, err = sdectd.NewClient(context.Background(), []string{etcdServer}, sdectd.ClientOptions{})
	if err != nil {
		log.Printf("unable to connect to etcd: %s", err.Error())
		return new(basicUsersService)
	}

	entries, err := client.GetEntries(prefix)
	if err != nil || len(entries) == 0 {
		log.Printf("unable to connect to entries: %s", err.Error())
		return new(basicUsersService)
	}

	conn, err := grpc.Dial(entries[0], grpc.WithInsecure())
	if err != nil {
		log.Printf("unable to connect to notificator: %s", err.Error())
		return new(basicUsersService)
	}
	return &basicUsersService{
		notificatorServiceClient: pb.NewNotificatorClient(conn),
	}
}

// New returns a UsersService with all of the expected middleware wired in.
func New(middleware []Middleware) UsersService {
	var svc UsersService = NewBasicUsersService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
