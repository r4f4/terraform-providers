// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearchserverless

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearchserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns attributes for one or more VPC endpoints associated with the current
// account. For more information, see Access Amazon OpenSearch Serverless using an
// interface endpoint (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-vpc.html)
// .
func (c *Client) BatchGetVpcEndpoint(ctx context.Context, params *BatchGetVpcEndpointInput, optFns ...func(*Options)) (*BatchGetVpcEndpointOutput, error) {
	if params == nil {
		params = &BatchGetVpcEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetVpcEndpoint", params, optFns, c.addOperationBatchGetVpcEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetVpcEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetVpcEndpointInput struct {

	// A list of VPC endpoint identifiers.
	//
	// This member is required.
	Ids []string

	noSmithyDocumentSerde
}

type BatchGetVpcEndpointOutput struct {

	// Details about the specified VPC endpoint.
	VpcEndpointDetails []types.VpcEndpointDetail

	// Error information for a failed request.
	VpcEndpointErrorDetails []types.VpcEndpointErrorDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetVpcEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpBatchGetVpcEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpBatchGetVpcEndpoint{}, middleware.After)
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
	if err = addOpBatchGetVpcEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetVpcEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetVpcEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aoss",
		OperationName: "BatchGetVpcEndpoint",
	}
}