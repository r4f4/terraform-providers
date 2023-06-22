// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the address blocks that you have added to a directory.
func (c *Client) ListIpRoutes(ctx context.Context, params *ListIpRoutesInput, optFns ...func(*Options)) (*ListIpRoutesOutput, error) {
	if params == nil {
		params = &ListIpRoutesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIpRoutes", params, optFns, c.addOperationListIpRoutesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIpRoutesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIpRoutesInput struct {

	// Identifier (ID) of the directory for which you want to retrieve the IP
	// addresses.
	//
	// This member is required.
	DirectoryId *string

	// Maximum number of items to return. If this value is zero, the maximum number of
	// items is specified by the limitations of the operation.
	Limit *int32

	// The ListIpRoutes.NextToken value from a previous call to ListIpRoutes . Pass
	// null if this is the first call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListIpRoutesOutput struct {

	// A list of IpRoute s.
	IpRoutesInfo []types.IpRouteInfo

	// If not null, more results are available. Pass this value for the NextToken
	// parameter in a subsequent call to ListIpRoutes to retrieve the next set of
	// items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListIpRoutesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListIpRoutes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListIpRoutes{}, middleware.After)
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
	if err = addOpListIpRoutesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIpRoutes(options.Region), middleware.Before); err != nil {
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

// ListIpRoutesAPIClient is a client that implements the ListIpRoutes operation.
type ListIpRoutesAPIClient interface {
	ListIpRoutes(context.Context, *ListIpRoutesInput, ...func(*Options)) (*ListIpRoutesOutput, error)
}

var _ ListIpRoutesAPIClient = (*Client)(nil)

// ListIpRoutesPaginatorOptions is the paginator options for ListIpRoutes
type ListIpRoutesPaginatorOptions struct {
	// Maximum number of items to return. If this value is zero, the maximum number of
	// items is specified by the limitations of the operation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListIpRoutesPaginator is a paginator for ListIpRoutes
type ListIpRoutesPaginator struct {
	options   ListIpRoutesPaginatorOptions
	client    ListIpRoutesAPIClient
	params    *ListIpRoutesInput
	nextToken *string
	firstPage bool
}

// NewListIpRoutesPaginator returns a new ListIpRoutesPaginator
func NewListIpRoutesPaginator(client ListIpRoutesAPIClient, params *ListIpRoutesInput, optFns ...func(*ListIpRoutesPaginatorOptions)) *ListIpRoutesPaginator {
	if params == nil {
		params = &ListIpRoutesInput{}
	}

	options := ListIpRoutesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListIpRoutesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListIpRoutesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListIpRoutes page.
func (p *ListIpRoutesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListIpRoutesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListIpRoutes(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListIpRoutes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "ListIpRoutes",
	}
}