## Technical task for Flugel.it

### Requirements

To run this code you'll need:

- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli?in=terraform/aws-get-started)
- [go](https://golang.org/doc/install)
- [aws CLI](https://docs.aws.amazon.com/cli/latest/userguide/install-linux.html)

**versions used**

`Terraform v1.0.1
on linux_amd64`

to check run: `terraform version`

`go version go1.13.8 linux/amd64`

to check run: `go version`

`aws-cli/2.2.11 Python/3.8.8 Linux/5.8.0-59-generic exe/x86_64.ubuntu.20 prompt/off`

to check run: `aws --version`


### Setup

Code requires access to S3 instance, thus AWS credentials with sufficient access right must be present in `~/.aws/credentials`

### Running the code

To run the Terraform code:

First you'll need to run `terraform init` in directory where code is located

You could optionally validate using `terraform validate`, but code provided in repo should be valid already

To create infrastructure run `terraform apply`, this will create desired infrastructure

*To destroy the infrastructure run:* `terraform destroy`

### Running Terratest tests

To run tests, issue the command `go test -v` in a directory where the test file is located.

if there are problems with modules this should help: `go get -v -t -d && go mod tidy`

### Automation

Merging or committing to main branch will run the tests automatically


