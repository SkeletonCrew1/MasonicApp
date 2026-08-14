# Creating an RDS instance using Terraform

This documentation describes how to provision an AWS RDS PostgreSQL database, networking, and secure credentials using Terraform.

## Prerequisites

- Git installed
- AWS account
- Terraform installed

## Folder structure

- `envs/` contains environment-specific variable files (`dev.tfvars`, `test.tfvars`, `prod.tfvars`).
- `providers.tf` defines the cloud provider and backend state configuration.
- `variables.tf` defines the input variables for the module.
- `terraform.tfvars` contains the active variable values for deployment.
- `vpc.tf` defines the private network.
- `subnets.tf` defines database subnets and the DB subnet group. 
- `securitygroup.tf` defines the security group and port rules for database access.
- `rds.tf` defines the PostgreSQL RDS cluster configuration.
- `secrets.tf` generates a random password and stores it in AWS Secrets Manager.

## Usage

```bash
1. Configure AWS credentials
2. Clone the MasonicApp repository
3. Change working directory to MasonicApp/terraform/rds/
4. Run terraform init
5. Run terraform plan -var-file="envs/dev.tfvars"
6. Run terraform apply -var-file="envs/dev.tfvars" -auto-approve