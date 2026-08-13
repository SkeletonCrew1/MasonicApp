variable "aws_region" {
    description = "AWS region"
    type        = string
}

variable "eks_oidc_issuer_url" {
    description = "url OCID of EKS cluster"
    type        = string 
}

variable "db_secret_arn" {
    description = "ARN secret in AWS Secret Manager from RDS"
    type        = string
}