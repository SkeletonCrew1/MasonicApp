terraform {
    required_version = ">= 1.15.7"

    required_providers {
        aws = {
            source  = "hashicorp/aws"
            version = ">= 6.0"
        }
    }

    backend "s3" {
        bucket       = "masonicapp-terraform-state-dev"
        key          = "rds/terraform.tfstate"
        region       = "eu-north-1"
        encrypt      = true
        use_lockfile = true
    }
}

provider "aws" {
    region = var.aws_region
}

data "terraform_remote_state" "eks" {
    backend = "s3"

    config = {
        bucket = "masonicapp-terraform-state-dev"
        key    = "state/eks-cluster/terraform.tfstate"
        region = "eu-north-1"
    }
}
