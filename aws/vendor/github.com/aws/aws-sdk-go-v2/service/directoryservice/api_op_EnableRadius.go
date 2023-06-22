// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables multi-factor authentication (MFA) with the Remote Authentication Dial
// In User Service (RADIUS) server for an AD Connector or Microsoft AD directory.
func (c *Client) EnableRadius(ctx context.Context, params *EnableRadiusInput, optFns ...func(*Options)) (*EnableRadiusOutput, error) {
	if params == nil {
		params = &EnableRadiusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableRadius", params, optFns, c.addOperationEnableRadiusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableRadiusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the EnableRadius operation.
type EnableRadiusInput struct {

	// The identifier of the directory for which to enable MFA.
	//
	// This member is required.
	DirectoryId *string

	// A RadiusSettings object that contains information about the RADIUS server.
	//
	// This member is required.
	RadiusSettings *types.RadiusSettings

	noSmithyDocumentSerde
}

// Contains the results of the EnableRadius operation.
type EnableRadiusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableRadiusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableRadius{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableRadius{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpEnableRadiusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableRadius(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opEnableRadius(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "EnableRadius",
	}
}