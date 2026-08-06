variable "aws_region" {}
variable "project" {}
variable "databases" { type = set(string) }
variable "vpc_cidr" {}
variable "subnet_1_cidr" {}
variable "subnet_2_cidr" {}
variable "db_password" { sensitive = true }