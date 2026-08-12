provider "aws" {
  region = "eu-north-1"
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">=6.57.0"
    }
  }

  backend "s3" {
    bucket       = "masonicapp-terraform-state"
    key          = "state/users/terraform.tfstate"
    use_lockfile = true
    region       = "us-east-1"
    encrypt      = true
  }
}
