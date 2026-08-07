locals {
  env                  = "development"
  region               = "eu-west-1"
  zone1                = "eu-west-1a"
  zone2                = "eu-west-1b"
  eks_name             = "eks-cluster"
  eks_version          = "1.36"
  vpc_cidr             = "10.0.0.0/16"
  prv_subnet_cidr1     = "10.0.0.0/19"
  prv_subnet_cidr2     = "10.0.32.0/19"
  pub_subnet_cidr1     = "10.0.64.0/19"
  pub_subnet_cidr2     = "10.0.96.0/19"
  all_users_cird_block = "0.0.0.0/0"
}
