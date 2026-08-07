provider "aws" {
  region = "us-east-1"
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">=6.52.0"
    }
  }

  backend "s3" {
    bucket       = "masonicapp-s3bucket"
    key          = "state/users/terraform.tfstate"
    use_lockfile = true
    region       = "us-east-1"
    encrypt      = true
  }
}
