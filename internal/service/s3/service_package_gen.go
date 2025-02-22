// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package s3

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []func(context.Context) (datasource.DataSourceWithConfigure, error) {
	return []func(context.Context) (datasource.DataSourceWithConfigure, error){}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []func(context.Context) (resource.ResourceWithConfigure, error) {
	return []func(context.Context) (resource.ResourceWithConfigure, error){}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_canonical_user_id": DataSourceCanonicalUserID,
		"aws_s3_bucket":         DataSourceBucket,
		"aws_s3_bucket_object":  DataSourceBucketObject,
		"aws_s3_bucket_objects": DataSourceBucketObjects,
		"aws_s3_bucket_policy":  DataSourceBucketPolicy,
		"aws_s3_object":         DataSourceObject,
		"aws_s3_objects":        DataSourceObjects,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_s3_bucket":                                      ResourceBucket,
		"aws_s3_bucket_accelerate_configuration":             ResourceBucketAccelerateConfiguration,
		"aws_s3_bucket_acl":                                  ResourceBucketACL,
		"aws_s3_bucket_analytics_configuration":              ResourceBucketAnalyticsConfiguration,
		"aws_s3_bucket_cors_configuration":                   ResourceBucketCorsConfiguration,
		"aws_s3_bucket_intelligent_tiering_configuration":    ResourceBucketIntelligentTieringConfiguration,
		"aws_s3_bucket_inventory":                            ResourceBucketInventory,
		"aws_s3_bucket_lifecycle_configuration":              ResourceBucketLifecycleConfiguration,
		"aws_s3_bucket_logging":                              ResourceBucketLogging,
		"aws_s3_bucket_metric":                               ResourceBucketMetric,
		"aws_s3_bucket_notification":                         ResourceBucketNotification,
		"aws_s3_bucket_object":                               ResourceBucketObject,
		"aws_s3_bucket_object_lock_configuration":            ResourceBucketObjectLockConfiguration,
		"aws_s3_bucket_ownership_controls":                   ResourceBucketOwnershipControls,
		"aws_s3_bucket_policy":                               ResourceBucketPolicy,
		"aws_s3_bucket_public_access_block":                  ResourceBucketPublicAccessBlock,
		"aws_s3_bucket_replication_configuration":            ResourceBucketReplicationConfiguration,
		"aws_s3_bucket_request_payment_configuration":        ResourceBucketRequestPaymentConfiguration,
		"aws_s3_bucket_server_side_encryption_configuration": ResourceBucketServerSideEncryptionConfiguration,
		"aws_s3_bucket_versioning":                           ResourceBucketVersioning,
		"aws_s3_bucket_website_configuration":                ResourceBucketWebsiteConfiguration,
		"aws_s3_object":                                      ResourceObject,
		"aws_s3_object_copy":                                 ResourceObjectCopy,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.S3
}

var ServicePackage = &servicePackage{}
