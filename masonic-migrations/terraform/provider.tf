terraform {
    required_version = ">= 1.15.7"

    required_providers {
        aws = {
            source  = "hashicorp/aws"
            version = ">= 6.0"
        }
    }
}

provider "aws" {
    region = var.aws_region
}

data "aws_caller_identity" "current" {}
