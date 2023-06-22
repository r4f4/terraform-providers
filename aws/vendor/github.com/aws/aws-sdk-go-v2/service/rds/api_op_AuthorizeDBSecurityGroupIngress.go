// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables ingress to a DBSecurityGroup using one of two forms of authorization.
// First, EC2 or VPC security groups can be added to the DBSecurityGroup if the
// application using the database is running on EC2 or VPC instances. Second, IP
// ranges are available if the application accessing your database is running on
// the internet. Required parameters for this API are one of CIDR range,
// EC2SecurityGroupId for VPC, or (EC2SecurityGroupOwnerId and either
// EC2SecurityGroupName or EC2SecurityGroupId for non-VPC). You can't authorize
// ingress from an EC2 security group in one Amazon Web Services Region to an
// Amazon RDS DB instance in another. You can't authorize ingress from a VPC
// security group in one VPC to an Amazon RDS DB instance in another. For an
// overview of CIDR ranges, go to the Wikipedia Tutorial (http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
// . EC2-Classic was retired on August 15, 2022. If you haven't migrated from
// EC2-Classic to a VPC, we recommend that you migrate as soon as possible. For
// more information, see Migrate from EC2-Classic to a VPC (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html)
// in the Amazon EC2 User Guide, the blog EC2-Classic Networking is Retiring –
// Here’s How to Prepare (http://aws.amazon.com/blogs/aws/ec2-classic-is-retiring-heres-how-to-prepare/)
// , and Moving a DB instance not in a VPC into a VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.Non-VPC2VPC.html)
// in the Amazon RDS User Guide.
func (c *Client) AuthorizeDBSecurityGroupIngress(ctx context.Context, params *AuthorizeDBSecurityGroupIngressInput, optFns ...func(*Options)) (*AuthorizeDBSecurityGroupIngressOutput, error) {
	if params == nil {
		params = &AuthorizeDBSecurityGroupIngressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AuthorizeDBSecurityGroupIngress", params, optFns, c.addOperationAuthorizeDBSecurityGroupIngressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AuthorizeDBSecurityGroupIngressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AuthorizeDBSecurityGroupIngressInput struct {

	// The name of the DB security group to add authorization to.
	//
	// This member is required.
	DBSecurityGroupName *string

	// The IP range to authorize.
	CIDRIP *string

	// Id of the EC2 security group to authorize. For VPC DB security groups,
	// EC2SecurityGroupId must be provided. Otherwise, EC2SecurityGroupOwnerId and
	// either EC2SecurityGroupName or EC2SecurityGroupId must be provided.
	EC2SecurityGroupId *string

	// Name of the EC2 security group to authorize. For VPC DB security groups,
	// EC2SecurityGroupId must be provided. Otherwise, EC2SecurityGroupOwnerId and
	// either EC2SecurityGroupName or EC2SecurityGroupId must be provided.
	EC2SecurityGroupName *string

	// Amazon Web Services account number of the owner of the EC2 security group
	// specified in the EC2SecurityGroupName parameter. The Amazon Web Services access
	// key ID isn't an acceptable value. For VPC DB security groups, EC2SecurityGroupId
	// must be provided. Otherwise, EC2SecurityGroupOwnerId and either
	// EC2SecurityGroupName or EC2SecurityGroupId must be provided.
	EC2SecurityGroupOwnerId *string

	noSmithyDocumentSerde
}

type AuthorizeDBSecurityGroupIngressOutput struct {

	// Contains the details for an Amazon RDS DB security group. This data type is
	// used as a response element in the DescribeDBSecurityGroups action.
	DBSecurityGroup *types.DBSecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAuthorizeDBSecurityGroupIngressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAuthorizeDBSecurityGroupIngress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAuthorizeDBSecurityGroupIngress{}, middleware.After)
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
	if err = addOpAuthorizeDBSecurityGroupIngressValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAuthorizeDBSecurityGroupIngress(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAuthorizeDBSecurityGroupIngress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "AuthorizeDBSecurityGroupIngress",
	}
}