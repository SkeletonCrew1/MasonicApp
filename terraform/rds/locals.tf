locals {
    env          = var.env
    zone1        = "${var.aws_region}a"
    zone2        = "${var.aws_region}b"

    vpc_cidr     = var.vpc_cidr
    subnet_cidr1 = var.subnet_cidr1
    subnet_cidr2 = var.subnet_cidr2

    db_name      = "paranormal_db"
    db_user      = "miraculous"
}