## Append HCL format with hclwrite to a Terraform file

Example of use [hclwrite package](https://pkg.go.dev/github.com/hashicorp/hcl/v2/hclwrite). \
For instance, its only append Terraform resource  [aws_ssm_parameter](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ssm_parameter)

### Usage

```
go mod tidy
go build -o /usr/local/bin/append main.go
append
```
Running "append" an interactive session starts asking for values of aws_ssm_parameter Resource.
