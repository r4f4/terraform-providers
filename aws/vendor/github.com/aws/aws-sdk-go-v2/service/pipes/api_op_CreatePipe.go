// Code generated by smithy-go-codegen DO NOT EDIT.

package pipes

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pipes/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Create a pipe. Amazon EventBridge Pipes connect event sources to targets and
// reduces the need for specialized knowledge and integration code.
func (c *Client) CreatePipe(ctx context.Context, params *CreatePipeInput, optFns ...func(*Options)) (*CreatePipeOutput, error) {
	if params == nil {
		params = &CreatePipeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePipe", params, optFns, c.addOperationCreatePipeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePipeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePipeInput struct {

	// The name of the pipe.
	//
	// This member is required.
	Name *string

	// The ARN of the role that allows the pipe to send data to the target.
	//
	// This member is required.
	RoleArn *string

	// The ARN of the source resource.
	//
	// This member is required.
	Source *string

	// The ARN of the target resource.
	//
	// This member is required.
	Target *string

	// A description of the pipe.
	Description *string

	// The state the pipe should be in.
	DesiredState types.RequestedPipeState

	// The ARN of the enrichment resource.
	Enrichment *string

	// The parameters required to set up enrichment on your pipe.
	EnrichmentParameters *types.PipeEnrichmentParameters

	// The parameters required to set up a source for your pipe.
	SourceParameters *types.PipeSourceParameters

	// The list of key-value pairs to associate with the pipe.
	Tags map[string]string

	// The parameters required to set up a target for your pipe.
	TargetParameters *types.PipeTargetParameters

	noSmithyDocumentSerde
}

type CreatePipeOutput struct {

	// The ARN of the pipe.
	Arn *string

	// The time the pipe was created.
	CreationTime *time.Time

	// The state the pipe is in.
	CurrentState types.PipeState

	// The state the pipe should be in.
	DesiredState types.RequestedPipeState

	// When the pipe was last updated, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime)
	// (YYYY-MM-DDThh:mm:ss.sTZD).
	LastModifiedTime *time.Time

	// The name of the pipe.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePipeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePipe{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePipe{}, middleware.After)
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
	if err = addOpCreatePipeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePipe(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePipe(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "pipes",
		OperationName: "CreatePipe",
	}
}