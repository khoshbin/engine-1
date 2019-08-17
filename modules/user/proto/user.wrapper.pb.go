// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/user/proto/user.proto

package userpb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/fzerorubigd/protobuf/extra"
	_ "github.com/fzerorubigd/protobuf/types"
	_ "github.com/gogo/protobuf/gogoproto"
	elbix_dev_engine_pkg_grpcgw "elbix.dev/engine/pkg/grpcgw"
	gopkg_in_go_playground_validator_v9 "gopkg.in/go-playground/validator.v9"
	golang_org_x_net_context "golang.org/x/net/context"
	github_com_grpc_ecosystem_grpc_gateway_runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	elbix_dev_engine_pkg_assert "elbix.dev/engine/pkg/assert"
	elbix_dev_engine_pkg_resources "elbix.dev/engine/pkg/resources"
	google_golang_org_grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WrappedUserSystemController interface {
	UserSystemServer
	elbix_dev_engine_pkg_grpcgw.Controller
}

type wrappedUserSystemServer struct {
	original UserSystemServer
	v        *gopkg_in_go_playground_validator_v9.Validate
}

func (w *wrappedUserSystemServer) Init(ctx golang_org_x_net_context.Context, conn *google_golang_org_grpc.ClientConn, mux *github_com_grpc_ecosystem_grpc_gateway_runtime.ServeMux) {
	cl := NewUserSystemClient(conn)

	elbix_dev_engine_pkg_assert.Nil(RegisterUserSystemHandlerClient(ctx, mux, cl))
}

func (w *wrappedUserSystemServer) InitGRPC(ctx golang_org_x_net_context.Context, s *google_golang_org_grpc.Server) {
	RegisterUserSystemServer(s, w)
}

func (w *wrappedUserSystemServer) Login(ctx golang_org_x_net_context.Context, req *LoginRequest) (res *UserResponse, err error) {
	ctx, err = elbix_dev_engine_pkg_grpcgw.ExecuteMiddleware(ctx, w.original)
	if err != nil {
		return nil, err
	}
	if err = w.v.StructCtx(ctx, req); err != nil {
		return nil, elbix_dev_engine_pkg_grpcgw.NewBadRequest(err, "validation failed")
	}

	res, err = w.original.Login(ctx, req)
	return
}

func (w *wrappedUserSystemServer) Logout(ctx golang_org_x_net_context.Context, req *LogoutRequest) (res *LogoutResponse, err error) {
	ctx, err = elbix_dev_engine_pkg_grpcgw.ExecuteMiddleware(ctx, w.original)
	if err != nil {
		return nil, err
	}
	if err = w.v.StructCtx(ctx, req); err != nil {
		return nil, elbix_dev_engine_pkg_grpcgw.NewBadRequest(err, "validation failed")
	}

	res, err = w.original.Logout(ctx, req)
	return
}

func (w *wrappedUserSystemServer) Register(ctx golang_org_x_net_context.Context, req *RegisterRequest) (res *UserResponse, err error) {
	ctx, err = elbix_dev_engine_pkg_grpcgw.ExecuteMiddleware(ctx, w.original)
	if err != nil {
		return nil, err
	}
	if err = w.v.StructCtx(ctx, req); err != nil {
		return nil, elbix_dev_engine_pkg_grpcgw.NewBadRequest(err, "validation failed")
	}

	res, err = w.original.Register(ctx, req)
	return
}

func (w *wrappedUserSystemServer) Ping(ctx golang_org_x_net_context.Context, req *PingRequest) (res *UserResponse, err error) {
	ctx, err = elbix_dev_engine_pkg_grpcgw.ExecuteMiddleware(ctx, w.original)
	if err != nil {
		return nil, err
	}
	if err = w.v.StructCtx(ctx, req); err != nil {
		return nil, elbix_dev_engine_pkg_grpcgw.NewBadRequest(err, "validation failed")
	}

	res, err = w.original.Ping(ctx, req)
	return
}

func (w *wrappedUserSystemServer) VerifyToken(ctx golang_org_x_net_context.Context, req *VerifyTokenRequest) (res *UserResponse, err error) {
	ctx, err = elbix_dev_engine_pkg_grpcgw.ExecuteMiddleware(ctx, w.original)
	if err != nil {
		return nil, err
	}
	if err = w.v.StructCtx(ctx, req); err != nil {
		return nil, elbix_dev_engine_pkg_grpcgw.NewBadRequest(err, "validation failed")
	}

	res, err = w.original.VerifyToken(ctx, req)
	return
}

func (w *wrappedUserSystemServer) ChangePassword(ctx golang_org_x_net_context.Context, req *ChangePasswordRequest) (res *ChangePasswordResponse, err error) {
	ctx, err = elbix_dev_engine_pkg_grpcgw.ExecuteMiddleware(ctx, w.original)
	if err != nil {
		return nil, err
	}
	if err = w.v.StructCtx(ctx, req); err != nil {
		return nil, elbix_dev_engine_pkg_grpcgw.NewBadRequest(err, "validation failed")
	}

	res, err = w.original.ChangePassword(ctx, req)
	return
}

func (w *wrappedUserSystemServer) ChangeDisplayName(ctx golang_org_x_net_context.Context, req *ChangeDisplayNameRequest) (res *ChangeDisplayNameResponse, err error) {
	ctx, err = elbix_dev_engine_pkg_grpcgw.ExecuteMiddleware(ctx, w.original)
	if err != nil {
		return nil, err
	}
	if err = w.v.StructCtx(ctx, req); err != nil {
		return nil, elbix_dev_engine_pkg_grpcgw.NewBadRequest(err, "validation failed")
	}

	res, err = w.original.ChangeDisplayName(ctx, req)
	return
}

func (w *wrappedUserSystemServer) ForgotPassword(ctx golang_org_x_net_context.Context, req *ForgotPasswordRequest) (res *ForgotPasswordResponse, err error) {
	ctx, err = elbix_dev_engine_pkg_grpcgw.ExecuteMiddleware(ctx, w.original)
	if err != nil {
		return nil, err
	}
	if err = w.v.StructCtx(ctx, req); err != nil {
		return nil, elbix_dev_engine_pkg_grpcgw.NewBadRequest(err, "validation failed")
	}

	res, err = w.original.ForgotPassword(ctx, req)
	return
}

func NewWrappedUserSystemServer(server UserSystemServer) WrappedUserSystemController {
	return &wrappedUserSystemServer{
		original: server,
		v:        gopkg_in_go_playground_validator_v9.New(),
	}
}
func init() {
	elbix_dev_engine_pkg_resources.RegisterResource("/user.UserSystem/Logout", "")
	elbix_dev_engine_pkg_resources.RegisterResource("/user.UserSystem/Ping", "")
	elbix_dev_engine_pkg_resources.RegisterResource("/user.UserSystem/ChangePassword", "")
	elbix_dev_engine_pkg_resources.RegisterResource("/user.UserSystem/ChangeDisplayName", "")
}
