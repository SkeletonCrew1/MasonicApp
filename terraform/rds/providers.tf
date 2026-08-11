terraform {
    required_version = ">= 1.15.7"

    required_providers {
        aws = {
            source  = "hashicorp/aws"
            version = ">= 6.0"
        }
    }

    backend "s3" {
        bucket       = "masons-infra-tfstate"
        region       = "eu-north-1"
        encrypt      = true
        use_lockfile = true
    }
}

provider "aws" {
    region = var.aws_region
}