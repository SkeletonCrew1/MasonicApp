terraform {
    required_version = ">= 1.15.7"

    required_providers {
        aws = {
            source  = "hashicorp/aws"
            version = ">= 6.0"
        }
    }

    backend "s3" {
        bucket       = "masonicapp-terraform-state"
        key          = "rds/terraform.tfstate"
        region       = "us-east-1"
        encrypt      = true
        use_lockfile = true
    }
}

provider "aws" {
    region = var.aws_region
}