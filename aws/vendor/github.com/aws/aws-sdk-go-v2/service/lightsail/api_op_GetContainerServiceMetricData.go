// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the data points of a specific metric of your Amazon Lightsail container
// service. Metrics report the utilization of your resources. Monitor and collect
// metric data regularly to maintain the reliability, availability, and performance
// of your resources.
func (c *Client) GetContainerServiceMetricData(ctx context.Context, params *GetContainerServiceMetricDataInput, optFns ...func(*Options)) (*GetContainerServiceMetricDataOutput, error) {
	if params == nil {
		params = &GetContainerServiceMetricDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetContainerServiceMetricData", params, optFns, c.addOperationGetContainerServiceMetricDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetContainerServiceMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetContainerServiceMetricDataInput struct {

	// The end time of the time period.
	//
	// This member is required.
	EndTime *time.Time

	// The metric for which you want to return information. Valid container service
	// metric names are listed below, along with the most useful statistics to include
	// in your request, and the published unit value.
	//   - CPUUtilization - The average percentage of compute units that are currently
	//   in use across all nodes of the container service. This metric identifies the
	//   processing power required to run containers on each node of the container
	//   service. Statistics: The most useful statistics are Maximum and Average .
	//   Unit: The published unit is Percent .
	//   - MemoryUtilization - The average percentage of available memory that is
	//   currently in use across all nodes of the container service. This metric
	//   identifies the memory required to run containers on each node of the container
	//   service. Statistics: The most useful statistics are Maximum and Average .
	//   Unit: The published unit is Percent .
	//
	// This member is required.
	MetricName types.ContainerServiceMetricName

	// The granularity, in seconds, of the returned data points. All container service
	// metric data is available in 5-minute (300 seconds) granularity.
	//
	// This member is required.
	Period int32

	// The name of the container service for which to get metric data.
	//
	// This member is required.
	ServiceName *string

	// The start time of the time period.
	//
	// This member is required.
	StartTime *time.Time

	// The statistic for the metric. The following statistics are available:
	//   - Minimum - The lowest value observed during the specified period. Use this
	//   value to determine low volumes of activity for your application.
	//   - Maximum - The highest value observed during the specified period. Use this
	//   value to determine high volumes of activity for your application.
	//   - Sum - All values submitted for the matching metric added together. You can
	//   use this statistic to determine the total volume of a metric.
	//   - Average - The value of Sum / SampleCount during the specified period. By
	//   comparing this statistic with the Minimum and Maximum values, you can
	//   determine the full scope of a metric and how close the average use is to the
	//   Minimum and Maximum values. This comparison helps you to know when to increase
	//   or decrease your resources.
	//   - SampleCount - The count, or number, of data points used for the statistical
	//   calculation.
	//
	// This member is required.
	Statistics []types.MetricStatistic

	noSmithyDocumentSerde
}

type GetContainerServiceMetricDataOutput struct {

	// An array of objects that describe the metric data returned.
	MetricData []types.MetricDatapoint

	// The name of the metric returned.
	MetricName types.ContainerServiceMetricName

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetContainerServiceMetricDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetContainerServiceMetricData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetContainerServiceMetricData{}, middleware.After)
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
	if err = addOpGetContainerServiceMetricDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetContainerServiceMetricData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetContainerServiceMetricData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetContainerServiceMetricData",
	}
}