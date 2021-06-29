package test

import (
	// "fmt"
	"testing"

  "github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestBucket(t *testing.T) {
	t.Parallel()

	region := "eu-north-1"
	s3BucketName := "flugel-task"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
  		// The path to where our Terraform code is located
  		TerraformDir: "./",
    })
	defer terraform.Destroy(t, terraformOptions)
  terraform.InitAndApply(t, terraformOptions)

	aws.AssertS3BucketExists(t, region, s3BucketName)
	aws.GetS3ObjectContents(t, region, s3BucketName, "test1.txt")
	aws.GetS3ObjectContents(t, region, s3BucketName, "test2.txt")
}
