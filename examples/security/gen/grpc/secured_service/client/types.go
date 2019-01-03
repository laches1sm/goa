// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// secured_service gRPC client types
//
// Command:
// $ goa gen goa.design/goa/examples/security/design -o
// $(GOPATH)/src/goa.design/goa/examples/security

package client

import (
	"goa.design/goa/examples/security/gen/grpc/secured_service/pb"
	securedservice "goa.design/goa/examples/security/gen/secured_service"
)

// NewSigninRequest builds the gRPC request type from the payload of the
// "signin" endpoint of the "secured_service" service.
func NewSigninRequest(payload *securedservice.SigninPayload) *pb.SigninRequest {
	message := &pb.SigninRequest{}
	return message
}

// NewCreds builds the result type of the "signin" endpoint of the
// "secured_service" service from the gRPC response type.
func NewCreds(message *pb.SigninResponse) *securedservice.Creds {
	result := &securedservice.Creds{
		JWT:        message.Jwt,
		APIKey:     message.ApiKey,
		OauthToken: message.OauthToken,
	}
	return result
}

// NewSecureRequest builds the gRPC request type from the payload of the
// "secure" endpoint of the "secured_service" service.
func NewSecureRequest(payload *securedservice.SecurePayload) *pb.SecureRequest {
	message := &pb.SecureRequest{}
	if payload.Fail != nil {
		message.Fail = *payload.Fail
	}
	return message
}

// NewSecureResponse builds the result type of the "secure" endpoint of the
// "secured_service" service from the gRPC response type.
func NewSecureResponse(message *pb.SecureResponse) string {
	result := message.Field
	return result
}

// NewDoublySecureRequest builds the gRPC request type from the payload of the
// "doubly_secure" endpoint of the "secured_service" service.
func NewDoublySecureRequest(payload *securedservice.DoublySecurePayload) *pb.DoublySecureRequest {
	message := &pb.DoublySecureRequest{
		Key: payload.Key,
	}
	return message
}

// NewDoublySecureResponse builds the result type of the "doubly_secure"
// endpoint of the "secured_service" service from the gRPC response type.
func NewDoublySecureResponse(message *pb.DoublySecureResponse) string {
	result := message.Field
	return result
}

// NewAlsoDoublySecureRequest builds the gRPC request type from the payload of
// the "also_doubly_secure" endpoint of the "secured_service" service.
func NewAlsoDoublySecureRequest(payload *securedservice.AlsoDoublySecurePayload) *pb.AlsoDoublySecureRequest {
	message := &pb.AlsoDoublySecureRequest{}
	if payload.Username != nil {
		message.Username = *payload.Username
	}
	if payload.Password != nil {
		message.Password = *payload.Password
	}
	if payload.Key != nil {
		message.Key = *payload.Key
	}
	return message
}

// NewAlsoDoublySecureResponse builds the result type of the
// "also_doubly_secure" endpoint of the "secured_service" service from the gRPC
// response type.
func NewAlsoDoublySecureResponse(message *pb.AlsoDoublySecureResponse) string {
	result := message.Field
	return result
}
