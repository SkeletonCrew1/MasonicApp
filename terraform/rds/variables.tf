variable "aws_region" {
    description = "AWS region for deployment"
    type        = string
    default     = "eu-north-1"
}

variable "env" {
    description = "Environment name (e.g., dev, test, prod)"
    type        = string
}

variable "vpc_cidr" {
    description = "CIDR block for the VPC"
    type        = string
}

variable "subnet_cidr1" {
    description = "CIDR block for the first subnet"
    type        = string
}

variable "subnet_cidr2" {
    description = "CIDR block for the second subnet"
    type        = string
}