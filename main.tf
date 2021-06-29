terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }

  required_version = ">= 0.14.9"
}

provider "aws" {
  profile = "default"
  region  = "eu-north-1"
}

resource "aws_s3_bucket" "flugel-bucket" {
  bucket = "flugel-task"
  acl    = "public-read"
}
resource "aws_s3_bucket_object" "test1" {
  bucket  = aws_s3_bucket.flugel-bucket.id
  key     = "test1.txt"
  content = timestamp()
}

resource "aws_s3_bucket_object" "test2" {
  bucket  = aws_s3_bucket.flugel-bucket.id
  key     = "test2.txt"
  content = timestamp()
}
