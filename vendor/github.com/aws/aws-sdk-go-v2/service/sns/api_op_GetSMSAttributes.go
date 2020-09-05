// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the GetSMSAttributes request.
type GetSMSAttributesInput struct {
	_ struct{} `type:"structure"`

	// A list of the individual attribute names, such as MonthlySpendLimit, for
	// which you want values.
	//
	// For all attribute names, see SetSMSAttributes (https://docs.aws.amazon.com/sns/latest/api/API_SetSMSAttributes.html).
	//
	// If you don't use this parameter, Amazon SNS returns all SMS attributes.
	Attributes []string `locationName:"attributes" type:"list"`
}

// String returns the string representation
func (s GetSMSAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// The response from the GetSMSAttributes request.
type GetSMSAttributesOutput struct {
	_ struct{} `type:"structure"`

	// The SMS attribute names and their values.
	Attributes map[string]string `locationName:"attributes" type:"map"`
}

// String returns the string representation
func (s GetSMSAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSMSAttributes = "GetSMSAttributes"

// GetSMSAttributesRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Returns the settings for sending SMS messages from your account.
//
// These settings are set with the SetSMSAttributes action.
//
//    // Example sending a request using GetSMSAttributesRequest.
//    req := client.GetSMSAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/GetSMSAttributes
func (c *Client) GetSMSAttributesRequest(input *GetSMSAttributesInput) GetSMSAttributesRequest {
	op := &aws.Operation{
		Name:       opGetSMSAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSMSAttributesInput{}
	}

	req := c.newRequest(op, input, &GetSMSAttributesOutput{})

	return GetSMSAttributesRequest{Request: req, Input: input, Copy: c.GetSMSAttributesRequest}
}

// GetSMSAttributesRequest is the request type for the
// GetSMSAttributes API operation.
type GetSMSAttributesRequest struct {
	*aws.Request
	Input *GetSMSAttributesInput
	Copy  func(*GetSMSAttributesInput) GetSMSAttributesRequest
}

// Send marshals and sends the GetSMSAttributes API request.
func (r GetSMSAttributesRequest) Send(ctx context.Context) (*GetSMSAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSMSAttributesResponse{
		GetSMSAttributesOutput: r.Request.Data.(*GetSMSAttributesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSMSAttributesResponse is the response type for the
// GetSMSAttributes API operation.
type GetSMSAttributesResponse struct {
	*GetSMSAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSMSAttributes request.
func (r *GetSMSAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
