package proto

import (
	"context"
	"github.com/erfanmomeniii/user-management/internal/model"
	"github.com/erfanmomeniii/user-management/internal/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion7

type UserClient interface {
	SaveUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*SaveUserReply, error)
	DeleteUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*DeleteUserReply, error)
	UpdateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UpdateUserReply, error)
	GetUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*GetSingleUserReply, error)
	GetUsers(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*GetMultipleUserReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) SaveUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*SaveUserReply, error) {
	out := new(SaveUserReply)
	err := c.cc.Invoke(ctx, "/User/SaveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*DeleteUserReply, error) {
	out := new(DeleteUserReply)
	err := c.cc.Invoke(ctx, "/User/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UpdateUserReply, error) {
	out := new(UpdateUserReply)
	err := c.cc.Invoke(ctx, "/User/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*GetSingleUserReply, error) {
	out := new(GetSingleUserReply)
	err := c.cc.Invoke(ctx, "/User/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUsers(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*GetMultipleUserReply, error) {
	out := new(GetMultipleUserReply)
	err := c.cc.Invoke(ctx, "/User/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type UserServerInterface interface {
	SaveUser(context.Context, *UserRequest) (*SaveUserReply, error)
	DeleteUser(context.Context, *UserId) (*DeleteUserReply, error)
	UpdateUser(context.Context, *UserRequest) (*UpdateUserReply, error)
	GetUser(context.Context, *UserId) (*GetSingleUserReply, error)
	GetUsers(context.Context, *UserId) (*GetMultipleUserReply, error)
	mustEmbedUnimplementedUserServer()
}

// UserServer must be embedded to have forward compatible implementations.
type UserServer struct{}

func (UserServer) SaveUser(ctx context.Context, request *UserRequest) (*SaveUserReply, error) {
	if _, err := repository.User.Save(model.User{FirstName: request.FirstName, LastName: request.LastName,
		Nickname: request.Nickname, Email: request.Email, Password: request.Password, Country: request.Country}); err != nil {
		ctx.Done()
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	var response *SaveUserReply
	response.Message = "User saved successfully"
	ctx.Done()
	return response, nil
}

func (UserServer) DeleteUser(ctx context.Context, user *UserId) (*DeleteUserReply, error) {
	if _, err := repository.User.Delete(user.id); err != nil {
		ctx.Done()
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	var response *DeleteUserReply
	response.Message = "User deleted successfully"
	ctx.Done()
	return response, nil
}

func (UserServer) UpdateUser(context.Context, *UserRequest) (*UpdateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

func (UserServer) GetUser(ctx context.Context, user *UserId) (*GetSingleUserReply, error) {

	u, err := repository.User.Get(user.id)

	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	var response *GetSingleUserReply

	response.Id = u.Id
	response.Password = u.Password
	response.Country = u.Country
	response.Email = u.Email
	response.LastName = u.LastName
	response.FirstName = u.FirstName
	response.Nickname = u.Nickname

	return response, nil
}

func (UserServer) GetUsers(context.Context, *UserId) (*GetMultipleUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServerInterface) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_SaveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SaveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/SaveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerInterface).SaveUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerInterface).DeleteUser(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerInterface).UpdateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerInterface).GetUser(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerInterface).GetUsers(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "User",
	HandlerType: (*UserServerInterface)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveUser",
			Handler:    _User_SaveUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _User_DeleteUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _User_GetUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _User_GetUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
