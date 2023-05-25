package grpc

import (
	"context"
	"go-base/proto"
	"go-base/usecase"
	"io"
	"log"
	"strconv"
)

type ServerGRPC struct {
	proto.UnimplementedHelloServiceServer
	UseCase *usecase.UseCase
}

// Hello implements HelloServiceServer
func (s *ServerGRPC) Hello(ctx context.Context, in *proto.String) (*proto.String, error) {
	reply := &proto.String{Value: "Message:" + in.GetValue()}
	return reply, nil
}

// Hello implements HelloServiceServer
func (s *ServerGRPC) Hello1(ctx context.Context, in *proto.String) (*proto.String, error) {
	reply := &proto.String{Value: "Message1:" + in.GetValue()}
	return reply, nil
}
func (s *ServerGRPC) GetUser(ctx context.Context, in *proto.GetUserReq) (*proto.User, error) {
	user, err := s.UseCase.UserUseCase.GetOneByEmail(in.GetEmail())
	if err != nil {
		return nil, err
	}
	reply := &proto.User{
		UserName: user.UserName,
		Password: user.Password,
		Email:    user.Email,
	}
	return reply, nil
}
func (s *ServerGRPC) Channel(stream proto.HelloService_ChannelServer) error {
	go func() {
		for i := 0; i < 10; i++ {
			log.Println(i)
			err := stream.Send(&proto.String{Value: "hello" + strconv.Itoa(i)})
			if err != nil {
			}
		}
	}()
	for {
		// Server nhận dữ liệu được gửi từ client
		// trong vòng lặp.
		args, err := stream.Recv()
		if err != nil {
			// Nếu gặp `io.EOF`, client stream sẽ đóng.
			if err == io.EOF {
				return nil
			}
			return err
		}
		log.Println(args.GetValue())
	}

}
