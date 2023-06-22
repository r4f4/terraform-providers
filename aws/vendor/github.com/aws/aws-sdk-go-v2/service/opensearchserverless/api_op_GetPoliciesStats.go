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

// Returns statistical information about your OpenSearch Serverless access
// policies, security configurations, and security policies.
func (c *Client) GetPoliciesStats(ctx context.Context, params *GetPoliciesStatsInput, optFns ...func(*Options)) (*GetPoliciesStatsOutput, error) {
	if params == nil {
		params = &GetPoliciesStatsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPoliciesStats", params, optFns, c.addOperationGetPoliciesStatsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPoliciesStatsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPoliciesStatsInput struct {
	noSmithyDocumentSerde
}

type GetPoliciesStatsOutput struct {

	// Information about the data access policies in your account.
	AccessPolicyStats *types.AccessPolicyStats

	// Information about the security configurations in your account.
	SecurityConfigStats *types.SecurityConfigStats

	// Information about the security policies in your account.
	SecurityPolicyStats *types.SecurityPolicyStats

	// The total number of OpenSearch Serverless security policies and configurations
	// in your account.
	TotalPolicyCount *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPoliciesStatsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetPoliciesStats{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetPoliciesStats{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPoliciesStats(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPoliciesStats(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aoss",
		OperationName: "GetPoliciesStats",
	}
}