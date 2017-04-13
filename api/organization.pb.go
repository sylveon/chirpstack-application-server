// Code generated by protoc-gen-go.
// source: organization.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Request the organizations defined in the system.
type ListOrganizationRequest struct {
	// Max number of organizations to return in the result-set.
	Limit int32 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int32 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *ListOrganizationRequest) Reset()                    { *m = ListOrganizationRequest{} }
func (m *ListOrganizationRequest) String() string            { return proto.CompactTextString(m) }
func (*ListOrganizationRequest) ProtoMessage()               {}
func (*ListOrganizationRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *ListOrganizationRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListOrganizationRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

// Request the user information.
type OrganizationRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *OrganizationRequest) Reset()                    { *m = OrganizationRequest{} }
func (m *OrganizationRequest) String() string            { return proto.CompactTextString(m) }
func (*OrganizationRequest) ProtoMessage()               {}
func (*OrganizationRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *OrganizationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetOrganizationResponse struct {
	// ID of the user.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Organization name.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Organization display name.
	DisplayName string `protobuf:"bytes,3,opt,name=displayName" json:"displayName,omitempty"`
	// Can the organization create and "own" Gateways?
	CanHaveGateways bool `protobuf:"varint,4,opt,name=canHaveGateways" json:"canHaveGateways,omitempty"`
	// When the user was created.
	CreatedAt string `protobuf:"bytes,5,opt,name=createdAt" json:"createdAt,omitempty"`
	// When the user was last updated (excludes changes in application access).
	UpdatedAt string `protobuf:"bytes,6,opt,name=updatedAt" json:"updatedAt,omitempty"`
}

func (m *GetOrganizationResponse) Reset()                    { *m = GetOrganizationResponse{} }
func (m *GetOrganizationResponse) String() string            { return proto.CompactTextString(m) }
func (*GetOrganizationResponse) ProtoMessage()               {}
func (*GetOrganizationResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *GetOrganizationResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetOrganizationResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetOrganizationResponse) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *GetOrganizationResponse) GetCanHaveGateways() bool {
	if m != nil {
		return m.CanHaveGateways
	}
	return false
}

func (m *GetOrganizationResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *GetOrganizationResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

// Add a new organization.
type CreateOrganizationRequest struct {
	// Organization name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Organization display name.
	DisplayName string `protobuf:"bytes,2,opt,name=displayName" json:"displayName,omitempty"`
	// Can the organization create and "own" Gateways?
	CanHaveGateways bool `protobuf:"varint,3,opt,name=canHaveGateways" json:"canHaveGateways,omitempty"`
}

func (m *CreateOrganizationRequest) Reset()                    { *m = CreateOrganizationRequest{} }
func (m *CreateOrganizationRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateOrganizationRequest) ProtoMessage()               {}
func (*CreateOrganizationRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *CreateOrganizationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateOrganizationRequest) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *CreateOrganizationRequest) GetCanHaveGateways() bool {
	if m != nil {
		return m.CanHaveGateways
	}
	return false
}

type CreateOrganizationResponse struct {
	// ID of the organization.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateOrganizationResponse) Reset()                    { *m = CreateOrganizationResponse{} }
func (m *CreateOrganizationResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateOrganizationResponse) ProtoMessage()               {}
func (*CreateOrganizationResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

func (m *CreateOrganizationResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Not quite the AddOrganizationRequest.
type UpdateOrganizationRequest struct {
	// The ID of the organization to be updated.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The new name.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// The new display name.
	DisplayName string `protobuf:"bytes,3,opt,name=displayName" json:"displayName,omitempty"`
	// Can the organization create and "own" Gateways?
	CanHaveGateways bool `protobuf:"varint,4,opt,name=canHaveGateways" json:"canHaveGateways,omitempty"`
}

func (m *UpdateOrganizationRequest) Reset()                    { *m = UpdateOrganizationRequest{} }
func (m *UpdateOrganizationRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateOrganizationRequest) ProtoMessage()               {}
func (*UpdateOrganizationRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

func (m *UpdateOrganizationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateOrganizationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateOrganizationRequest) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *UpdateOrganizationRequest) GetCanHaveGateways() bool {
	if m != nil {
		return m.CanHaveGateways
	}
	return false
}

type ListOrganizationResponse struct {
	TotalCount int32                      `protobuf:"varint,1,opt,name=totalCount" json:"totalCount,omitempty"`
	Result     []*GetOrganizationResponse `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
}

func (m *ListOrganizationResponse) Reset()                    { *m = ListOrganizationResponse{} }
func (m *ListOrganizationResponse) String() string            { return proto.CompactTextString(m) }
func (*ListOrganizationResponse) ProtoMessage()               {}
func (*ListOrganizationResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *ListOrganizationResponse) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListOrganizationResponse) GetResult() []*GetOrganizationResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

type OrganizationEmptyResponse struct {
}

func (m *OrganizationEmptyResponse) Reset()                    { *m = OrganizationEmptyResponse{} }
func (m *OrganizationEmptyResponse) String() string            { return proto.CompactTextString(m) }
func (*OrganizationEmptyResponse) ProtoMessage()               {}
func (*OrganizationEmptyResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{7} }

type OrganizationUserRequest struct {
	// The organization id.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The user's id.
	UserID int64 `protobuf:"varint,2,opt,name=userID" json:"userID,omitempty"`
	// The user's admin status for the organization
	IsAdmin bool `protobuf:"varint,3,opt,name=isAdmin" json:"isAdmin,omitempty"`
}

func (m *OrganizationUserRequest) Reset()                    { *m = OrganizationUserRequest{} }
func (m *OrganizationUserRequest) String() string            { return proto.CompactTextString(m) }
func (*OrganizationUserRequest) ProtoMessage()               {}
func (*OrganizationUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{8} }

func (m *OrganizationUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrganizationUserRequest) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *OrganizationUserRequest) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

type DeleteOrganizationUserRequest struct {
	// The organization id.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The user's id.
	UserID int64 `protobuf:"varint,2,opt,name=userID" json:"userID,omitempty"`
}

func (m *DeleteOrganizationUserRequest) Reset()                    { *m = DeleteOrganizationUserRequest{} }
func (m *DeleteOrganizationUserRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteOrganizationUserRequest) ProtoMessage()               {}
func (*DeleteOrganizationUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{9} }

func (m *DeleteOrganizationUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeleteOrganizationUserRequest) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

// Request the users in an organization.
type ListOrganizationUsersRequest struct {
	// The organization id.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Max number of users to return in the result-set.
	Limit int32 `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int32 `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
}

func (m *ListOrganizationUsersRequest) Reset()                    { *m = ListOrganizationUsersRequest{} }
func (m *ListOrganizationUsersRequest) String() string            { return proto.CompactTextString(m) }
func (*ListOrganizationUsersRequest) ProtoMessage()               {}
func (*ListOrganizationUsersRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{10} }

func (m *ListOrganizationUsersRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ListOrganizationUsersRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListOrganizationUsersRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type GetOrganizationUserRequest struct {
	// ID of the organization.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// ID of the user.
	UserID int64 `protobuf:"varint,2,opt,name=userID" json:"userID,omitempty"`
}

func (m *GetOrganizationUserRequest) Reset()                    { *m = GetOrganizationUserRequest{} }
func (m *GetOrganizationUserRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOrganizationUserRequest) ProtoMessage()               {}
func (*GetOrganizationUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{11} }

func (m *GetOrganizationUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetOrganizationUserRequest) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

// Response for a user in the organization
type GetOrganizationUserResponse struct {
	// ID of the user.
	UserID int64 `protobuf:"varint,1,opt,name=userID" json:"userID,omitempty"`
	// Username of the user.
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	// If the user is a system admin, capable of creating other users.
	IsAdmin bool `protobuf:"varint,3,opt,name=isAdmin" json:"isAdmin,omitempty"`
	// When the user was created.
	CreatedAt string `protobuf:"bytes,4,opt,name=createdAt" json:"createdAt,omitempty"`
	// When the user was last updated (excludes changes in application access).
	UpdatedAt string `protobuf:"bytes,5,opt,name=updatedAt" json:"updatedAt,omitempty"`
}

func (m *GetOrganizationUserResponse) Reset()                    { *m = GetOrganizationUserResponse{} }
func (m *GetOrganizationUserResponse) String() string            { return proto.CompactTextString(m) }
func (*GetOrganizationUserResponse) ProtoMessage()               {}
func (*GetOrganizationUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{12} }

func (m *GetOrganizationUserResponse) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *GetOrganizationUserResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetOrganizationUserResponse) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *GetOrganizationUserResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *GetOrganizationUserResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

// Response for the users in an organization.
type ListOrganizationUsersResponse struct {
	// The total number of users in the organization.
	TotalCount int32 `protobuf:"varint,1,opt,name=totalCount" json:"totalCount,omitempty"`
	// The users in the requested limit, offset range.
	Result []*GetOrganizationUserResponse `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
}

func (m *ListOrganizationUsersResponse) Reset()                    { *m = ListOrganizationUsersResponse{} }
func (m *ListOrganizationUsersResponse) String() string            { return proto.CompactTextString(m) }
func (*ListOrganizationUsersResponse) ProtoMessage()               {}
func (*ListOrganizationUsersResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{13} }

func (m *ListOrganizationUsersResponse) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListOrganizationUsersResponse) GetResult() []*GetOrganizationUserResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*ListOrganizationRequest)(nil), "api.ListOrganizationRequest")
	proto.RegisterType((*OrganizationRequest)(nil), "api.OrganizationRequest")
	proto.RegisterType((*GetOrganizationResponse)(nil), "api.GetOrganizationResponse")
	proto.RegisterType((*CreateOrganizationRequest)(nil), "api.CreateOrganizationRequest")
	proto.RegisterType((*CreateOrganizationResponse)(nil), "api.CreateOrganizationResponse")
	proto.RegisterType((*UpdateOrganizationRequest)(nil), "api.UpdateOrganizationRequest")
	proto.RegisterType((*ListOrganizationResponse)(nil), "api.ListOrganizationResponse")
	proto.RegisterType((*OrganizationEmptyResponse)(nil), "api.OrganizationEmptyResponse")
	proto.RegisterType((*OrganizationUserRequest)(nil), "api.OrganizationUserRequest")
	proto.RegisterType((*DeleteOrganizationUserRequest)(nil), "api.DeleteOrganizationUserRequest")
	proto.RegisterType((*ListOrganizationUsersRequest)(nil), "api.ListOrganizationUsersRequest")
	proto.RegisterType((*GetOrganizationUserRequest)(nil), "api.GetOrganizationUserRequest")
	proto.RegisterType((*GetOrganizationUserResponse)(nil), "api.GetOrganizationUserResponse")
	proto.RegisterType((*ListOrganizationUsersResponse)(nil), "api.ListOrganizationUsersResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Organization service

type OrganizationClient interface {
	// Get organization list.
	List(ctx context.Context, in *ListOrganizationRequest, opts ...grpc.CallOption) (*ListOrganizationResponse, error)
	// Get data for a particular organization.
	Get(ctx context.Context, in *OrganizationRequest, opts ...grpc.CallOption) (*GetOrganizationResponse, error)
	// Create a new organization.
	Create(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*CreateOrganizationResponse, error)
	// Update an existing organization.
	Update(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*OrganizationEmptyResponse, error)
	// Delete an organization.
	Delete(ctx context.Context, in *OrganizationRequest, opts ...grpc.CallOption) (*OrganizationEmptyResponse, error)
	// Get organization's user list.
	ListUsers(ctx context.Context, in *ListOrganizationUsersRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
	// Get data for a particular organization user.
	GetUser(ctx context.Context, in *GetOrganizationUserRequest, opts ...grpc.CallOption) (*GetOrganizationUserResponse, error)
	// Add a new user to an organization.
	AddUser(ctx context.Context, in *OrganizationUserRequest, opts ...grpc.CallOption) (*OrganizationEmptyResponse, error)
	// Update a user in an organization.
	UpdateUser(ctx context.Context, in *OrganizationUserRequest, opts ...grpc.CallOption) (*OrganizationEmptyResponse, error)
	// Delete a user from an organization.
	DeleteUser(ctx context.Context, in *DeleteOrganizationUserRequest, opts ...grpc.CallOption) (*OrganizationEmptyResponse, error)
}

type organizationClient struct {
	cc *grpc.ClientConn
}

func NewOrganizationClient(cc *grpc.ClientConn) OrganizationClient {
	return &organizationClient{cc}
}

func (c *organizationClient) List(ctx context.Context, in *ListOrganizationRequest, opts ...grpc.CallOption) (*ListOrganizationResponse, error) {
	out := new(ListOrganizationResponse)
	err := grpc.Invoke(ctx, "/api.Organization/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) Get(ctx context.Context, in *OrganizationRequest, opts ...grpc.CallOption) (*GetOrganizationResponse, error) {
	out := new(GetOrganizationResponse)
	err := grpc.Invoke(ctx, "/api.Organization/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) Create(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*CreateOrganizationResponse, error) {
	out := new(CreateOrganizationResponse)
	err := grpc.Invoke(ctx, "/api.Organization/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) Update(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*OrganizationEmptyResponse, error) {
	out := new(OrganizationEmptyResponse)
	err := grpc.Invoke(ctx, "/api.Organization/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) Delete(ctx context.Context, in *OrganizationRequest, opts ...grpc.CallOption) (*OrganizationEmptyResponse, error) {
	out := new(OrganizationEmptyResponse)
	err := grpc.Invoke(ctx, "/api.Organization/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) ListUsers(ctx context.Context, in *ListOrganizationUsersRequest, opts ...grpc.CallOption) (*ListUserResponse, error) {
	out := new(ListUserResponse)
	err := grpc.Invoke(ctx, "/api.Organization/ListUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) GetUser(ctx context.Context, in *GetOrganizationUserRequest, opts ...grpc.CallOption) (*GetOrganizationUserResponse, error) {
	out := new(GetOrganizationUserResponse)
	err := grpc.Invoke(ctx, "/api.Organization/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) AddUser(ctx context.Context, in *OrganizationUserRequest, opts ...grpc.CallOption) (*OrganizationEmptyResponse, error) {
	out := new(OrganizationEmptyResponse)
	err := grpc.Invoke(ctx, "/api.Organization/AddUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) UpdateUser(ctx context.Context, in *OrganizationUserRequest, opts ...grpc.CallOption) (*OrganizationEmptyResponse, error) {
	out := new(OrganizationEmptyResponse)
	err := grpc.Invoke(ctx, "/api.Organization/UpdateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) DeleteUser(ctx context.Context, in *DeleteOrganizationUserRequest, opts ...grpc.CallOption) (*OrganizationEmptyResponse, error) {
	out := new(OrganizationEmptyResponse)
	err := grpc.Invoke(ctx, "/api.Organization/DeleteUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Organization service

type OrganizationServer interface {
	// Get organization list.
	List(context.Context, *ListOrganizationRequest) (*ListOrganizationResponse, error)
	// Get data for a particular organization.
	Get(context.Context, *OrganizationRequest) (*GetOrganizationResponse, error)
	// Create a new organization.
	Create(context.Context, *CreateOrganizationRequest) (*CreateOrganizationResponse, error)
	// Update an existing organization.
	Update(context.Context, *UpdateOrganizationRequest) (*OrganizationEmptyResponse, error)
	// Delete an organization.
	Delete(context.Context, *OrganizationRequest) (*OrganizationEmptyResponse, error)
	// Get organization's user list.
	ListUsers(context.Context, *ListOrganizationUsersRequest) (*ListUserResponse, error)
	// Get data for a particular organization user.
	GetUser(context.Context, *GetOrganizationUserRequest) (*GetOrganizationUserResponse, error)
	// Add a new user to an organization.
	AddUser(context.Context, *OrganizationUserRequest) (*OrganizationEmptyResponse, error)
	// Update a user in an organization.
	UpdateUser(context.Context, *OrganizationUserRequest) (*OrganizationEmptyResponse, error)
	// Delete a user from an organization.
	DeleteUser(context.Context, *DeleteOrganizationUserRequest) (*OrganizationEmptyResponse, error)
}

func RegisterOrganizationServer(s *grpc.Server, srv OrganizationServer) {
	s.RegisterService(&_Organization_serviceDesc, srv)
}

func _Organization_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Organization/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).List(ctx, req.(*ListOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Organization/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).Get(ctx, req.(*OrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Organization/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).Create(ctx, req.(*CreateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Organization/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).Update(ctx, req.(*UpdateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Organization/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).Delete(ctx, req.(*OrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Organization/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).ListUsers(ctx, req.(*ListOrganizationUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Organization/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).GetUser(ctx, req.(*GetOrganizationUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Organization/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).AddUser(ctx, req.(*OrganizationUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Organization/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).UpdateUser(ctx, req.(*OrganizationUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrganizationUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Organization/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).DeleteUser(ctx, req.(*DeleteOrganizationUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Organization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Organization",
	HandlerType: (*OrganizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Organization_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Organization_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Organization_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Organization_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Organization_Delete_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Organization_ListUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Organization_GetUser_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _Organization_AddUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Organization_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Organization_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "organization.proto",
}

func init() { proto.RegisterFile("organization.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 741 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0xe3, 0xc4, 0x6d, 0xa7, 0x08, 0xa4, 0xa5, 0x34, 0x8e, 0x9b, 0x34, 0x61, 0xa5, 0xa2,
	0xa8, 0xa0, 0x44, 0x14, 0x0e, 0x88, 0x5b, 0xd5, 0xa2, 0x80, 0x84, 0x40, 0xb2, 0xd4, 0x13, 0x08,
	0xb4, 0xd4, 0xdb, 0xb2, 0x92, 0x63, 0xbb, 0xd9, 0x0d, 0x28, 0x94, 0x4a, 0x88, 0x2b, 0x47, 0x3e,
	0x82, 0x2f, 0xe1, 0x0b, 0xf8, 0x05, 0xae, 0xfc, 0x03, 0xf2, 0xd8, 0x49, 0x1d, 0xc7, 0x9b, 0x44,
	0x15, 0xdc, 0xbc, 0x33, 0xe3, 0x79, 0x33, 0x6f, 0xdf, 0x8c, 0x0d, 0x24, 0x1c, 0x9c, 0xb2, 0x40,
	0x7c, 0x62, 0x4a, 0x84, 0x41, 0x27, 0x1a, 0x84, 0x2a, 0x24, 0x26, 0x8b, 0x84, 0x03, 0x43, 0xc9,
	0x07, 0x89, 0xc1, 0xa9, 0x9f, 0x86, 0xe1, 0xa9, 0xcf, 0xbb, 0x2c, 0x12, 0x5d, 0x16, 0x04, 0xa1,
	0xc2, 0x68, 0x99, 0x78, 0x69, 0x0f, 0xaa, 0xcf, 0x85, 0x54, 0x2f, 0x33, 0x89, 0x5c, 0x7e, 0x36,
	0xe4, 0x52, 0x91, 0x0d, 0xa8, 0xf8, 0xa2, 0x2f, 0x94, 0x6d, 0xb4, 0x8c, 0x76, 0xc5, 0x4d, 0x0e,
	0x64, 0x13, 0xac, 0xf0, 0xe4, 0x44, 0x72, 0x65, 0x97, 0xd0, 0x9c, 0x9e, 0xe8, 0x0e, 0xdc, 0x2c,
	0x4a, 0x72, 0x1d, 0x4a, 0xc2, 0xc3, 0x0c, 0xa6, 0x5b, 0x12, 0x1e, 0xfd, 0x69, 0x40, 0xb5, 0xc7,
	0x73, 0x78, 0x32, 0x0a, 0x03, 0xc9, 0xf3, 0xb1, 0x84, 0x40, 0x39, 0x60, 0x7d, 0x8e, 0x40, 0x6b,
	0x2e, 0x3e, 0x93, 0x16, 0xac, 0x7b, 0x42, 0x46, 0x3e, 0x1b, 0xbd, 0x88, 0x5d, 0x26, 0xba, 0xb2,
	0x26, 0xd2, 0x86, 0x1b, 0xc7, 0x2c, 0x78, 0xca, 0x3e, 0xf0, 0x1e, 0x53, 0xfc, 0x23, 0x1b, 0x49,
	0xbb, 0xdc, 0x32, 0xda, 0xab, 0x6e, 0xde, 0x4c, 0xea, 0xb0, 0x76, 0x3c, 0xe0, 0x4c, 0x71, 0x6f,
	0x5f, 0xd9, 0x15, 0xcc, 0x74, 0x69, 0x88, 0xbd, 0xc3, 0xc8, 0x4b, 0xbd, 0x56, 0xe2, 0x9d, 0x18,
	0xe8, 0x39, 0xd4, 0x0e, 0x30, 0xb4, 0xa8, 0xe9, 0x71, 0xe1, 0x86, 0xbe, 0xf0, 0xd2, 0x52, 0x85,
	0x9b, 0x85, 0x85, 0xd3, 0x7b, 0xe0, 0x14, 0x81, 0x17, 0xd3, 0x48, 0xbf, 0x19, 0x50, 0x3b, 0xc2,
	0xc2, 0x97, 0xb8, 0xa0, 0xff, 0x4d, 0x3a, 0x8d, 0xc0, 0x9e, 0x15, 0x5c, 0x5a, 0xf9, 0x36, 0x80,
	0x0a, 0x15, 0xf3, 0x0f, 0xc2, 0x61, 0x30, 0x96, 0x5d, 0xc6, 0x42, 0x1e, 0x82, 0x35, 0xe0, 0x72,
	0xe8, 0xc7, 0xda, 0x33, 0xdb, 0xeb, 0x7b, 0xf5, 0x0e, 0x8b, 0x44, 0x47, 0x23, 0x27, 0x37, 0x8d,
	0xa5, 0x5b, 0x50, 0xcb, 0xfa, 0x9f, 0xf4, 0x23, 0x35, 0x1a, 0x07, 0xd1, 0x57, 0x50, 0xcd, 0x3a,
	0x8f, 0x24, 0x1f, 0xe8, 0x98, 0xd9, 0x04, 0x2b, 0x1e, 0xab, 0x67, 0x87, 0xc8, 0x8d, 0xe9, 0xa6,
	0x27, 0x62, 0xc3, 0x8a, 0x90, 0xfb, 0x5e, 0x5f, 0x04, 0xe9, 0x7d, 0x8d, 0x8f, 0xb4, 0x07, 0x8d,
	0x43, 0xee, 0xf3, 0x69, 0xe2, 0xaf, 0x00, 0x41, 0x5f, 0x43, 0x3d, 0x4f, 0x5a, 0x9c, 0x46, 0xea,
	0xf2, 0x4c, 0x46, 0xb7, 0x54, 0x3c, 0xba, 0xe6, 0xd4, 0xe8, 0x1e, 0x82, 0x93, 0xe3, 0xf0, 0x2a,
	0x35, 0xfe, 0x30, 0x60, 0xab, 0x30, 0x4d, 0x7a, 0xb9, 0x97, 0xef, 0x19, 0x53, 0xf4, 0x39, 0xb0,
	0x1a, 0x3f, 0x65, 0x44, 0x37, 0x39, 0xeb, 0xa9, 0x9d, 0x9e, 0xdd, 0xf2, 0xdc, 0xd9, 0xad, 0xe4,
	0x67, 0x77, 0x04, 0x0d, 0x0d, 0x9b, 0x4b, 0xea, 0xf0, 0x51, 0x4e, 0x87, 0xad, 0x22, 0x1d, 0x66,
	0x9b, 0x1f, 0x6b, 0x71, 0xef, 0xcf, 0x2a, 0x5c, 0xcb, 0x06, 0x91, 0xb7, 0x50, 0x8e, 0x6b, 0x21,
	0x89, 0x94, 0x35, 0xab, 0xd8, 0x69, 0x68, 0xbc, 0xa9, 0x88, 0x9d, 0xaf, 0xbf, 0x7e, 0x7f, 0x2f,
	0x6d, 0x10, 0x82, 0x4b, 0x3e, 0xfb, 0x51, 0x90, 0xe4, 0x0d, 0x98, 0x3d, 0xae, 0x88, 0x8d, 0x19,
	0x8a, 0x72, 0xcf, 0x1d, 0x22, 0xda, 0xc4, 0xd4, 0x35, 0x52, 0x9d, 0x4d, 0xdd, 0x3d, 0x17, 0xde,
	0x05, 0x79, 0x0f, 0x56, 0xb2, 0x8b, 0xc8, 0x36, 0x26, 0xd2, 0x6e, 0x45, 0xa7, 0xa9, 0xf5, 0xa7,
	0x58, 0x0d, 0xc4, 0xaa, 0xd2, 0x82, 0x36, 0x1e, 0x1b, 0xbb, 0xc4, 0x07, 0x2b, 0x59, 0x63, 0x29,
	0x92, 0x76, 0xa7, 0x39, 0xdb, 0x33, 0xcd, 0x4e, 0x0f, 0x3d, 0x45, 0xa0, 0xba, 0xa3, 0x6b, 0x2a,
	0x46, 0x3b, 0x06, 0x2b, 0x99, 0xdd, 0x39, 0xd4, 0x2d, 0xc2, 0x49, 0xc9, 0xdb, 0xd5, 0x92, 0xe7,
	0xc3, 0x5a, 0x7c, 0xa9, 0xa8, 0x3e, 0x72, 0xbb, 0xf0, 0x92, 0xb3, 0x73, 0xee, 0xdc, 0x9a, 0x84,
	0x64, 0xd5, 0x45, 0x77, 0x10, 0xa7, 0x49, 0x1a, 0x1a, 0x9c, 0xee, 0x10, 0x01, 0x3e, 0xc3, 0x4a,
	0x8f, 0xe3, 0x9b, 0xa4, 0xa9, 0x57, 0x6c, 0x82, 0xb4, 0x50, 0xd2, 0xb4, 0x83, 0xa0, 0x6d, 0x72,
	0x67, 0x2e, 0x68, 0xf7, 0x3c, 0x19, 0xf3, 0x0b, 0x72, 0x06, 0x2b, 0xfb, 0x9e, 0x87, 0xe8, 0xf5,
	0x19, 0xde, 0xb2, 0xd0, 0x8b, 0x58, 0x6d, 0x23, 0x30, 0xa5, 0xf3, 0xbb, 0x8d, 0xef, 0xf0, 0x02,
	0x20, 0x11, 0xc9, 0x3f, 0x40, 0xbd, 0x8f, 0xa8, 0x77, 0x9d, 0x25, 0xdb, 0x8d, 0xe1, 0xbf, 0x18,
	0x00, 0x89, 0x86, 0x10, 0x9f, 0x22, 0xc2, 0xdc, 0x0f, 0xc2, 0xc2, 0x2a, 0x52, 0xd2, 0x77, 0x97,
	0xac, 0xe2, 0x9d, 0x85, 0x7f, 0x79, 0x0f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x45, 0x82, 0xf7,
	0x66, 0x2a, 0x0a, 0x00, 0x00,
}
