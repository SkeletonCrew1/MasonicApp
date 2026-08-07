variable "aws_region" {
    description = "AWS region for deployment"
    type        = string
    default     = "eu-north-1"
}
variable "db_password" { 
    description = "Password for the database for nov TF_VAR"
    type        = string
    sensitive   = true 
    default     = null
}