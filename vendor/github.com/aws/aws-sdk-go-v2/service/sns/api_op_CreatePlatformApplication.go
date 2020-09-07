// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input for CreatePlatformApplication action.
type CreatePlatformApplicationInput struct {
	_ struct{} `type:"structure"`

	// For a list of attributes, see SetPlatformApplicationAttributes (https://docs.aws.amazon.com/sns/latest/api/API_SetPlatformApplicationAttributes.html)
	//
	// Attributes is a required field
	Attributes map[string]string `type:"map" required:"true"`

	// Application names must be made up of only uppercase and lowercase ASCII letters,
	// numbers, underscores, hyphens, and periods, and must be between 1 and 256
	// characters long.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The following platforms are supported: ADM (Amazon Device Messaging), APNS
	// (Apple Push Notification Service), APNS_SANDBOX, and GCM (Firebase Cloud
	// Messaging).
	//
	// Platform is a required field
	Platform *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreatePlatformApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePlatformApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePlatformApplicationInput"}

	if s.Attributes == nil {
		invalidParams.Add(aws.NewErrParamRequired("Attributes"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Platform == nil {
		invalidParams.Add(aws.NewErrParamRequired("Platform"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response from CreatePlatformApplication action.
type CreatePlatformApplicationOutput struct {
	_ struct{} `type:"structure"`

	// PlatformApplicationArn is returned.
	PlatformApplicationArn *string `type:"string"`
}

// String returns the string representation
func (s CreatePlatformApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreatePlatformApplication = "CreatePlatformApplication"

// CreatePlatformApplicationRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Creates a platform application object for one of the supported push notification
// services, such as APNS and GCM (Firebase Cloud Messaging), to which devices
// and mobile apps may register. You must specify PlatformPrincipal and PlatformCredential
// attributes when using the CreatePlatformApplication action.
//
// PlatformPrincipal and PlatformCredential are received from the notification
// service.
//
//    * For ADM, PlatformPrincipal is client id and PlatformCredential is client
//    secret.
//
//    * For Baidu, PlatformPrincipal is API key and PlatformCredential is secret
//    key.
//
//    * For APNS and APNS_SANDBOX, PlatformPrincipal is SSL certificate and
//    PlatformCredential is private key.
//
//    * For GCM (Firebase Cloud Messaging), there is no PlatformPrincipal and
//    the PlatformCredential is API key.
//
//    * For MPNS, PlatformPrincipal is TLS certificate and PlatformCredential
//    is private key.
//
//    * For WNS, PlatformPrincipal is Package Security Identifier and PlatformCredential
//    is secret key.
//
// You can use the returned PlatformApplicationArn as an attribute for the CreatePlatformEndpoint
// action.
//
//    // Example sending a request using CreatePlatformApplicationRequest.
//    req := client.CreatePlatformApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/CreatePlatformApplication
func (c *Client) CreatePlatformApplicationRequest(input *CreatePlatformApplicationInput) CreatePlatformApplicationRequest {
	op := &aws.Operation{
		Name:       opCreatePlatformApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreatePlatformApplicationInput{}
	}

	req := c.newRequest(op, input, &CreatePlatformApplicationOutput{})

	return CreatePlatformApplicationRequest{Request: req, Input: input, Copy: c.CreatePlatformApplicationRequest}
}

// CreatePlatformApplicationRequest is the request type for the
// CreatePlatformApplication API operation.
type CreatePlatformApplicationRequest struct {
	*aws.Request
	Input *CreatePlatformApplicationInput
	Copy  func(*CreatePlatformApplicationInput) CreatePlatformApplicationRequest
}

// Send marshals and sends the CreatePlatformApplication API request.
func (r CreatePlatformApplicationRequest) Send(ctx context.Context) (*CreatePlatformApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePlatformApplicationResponse{
		CreatePlatformApplicationOutput: r.Request.Data.(*CreatePlatformApplicationOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePlatformApplicationResponse is the response type for the
// CreatePlatformApplication API operation.
type CreatePlatformApplicationResponse struct {
	*CreatePlatformApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePlatformApplication request.
func (r *CreatePlatformApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
