resource "aws_subnet" "db_zone1" {
    vpc_id            = data.terraform_remote_state.eks.outputs.vpc_id
    cidr_block        = var.subnet_cidr1
    availability_zone = "${var.aws_region}a"

    tags = {
        "Name" = "${var.env}-db-subnet-${var.aws_region}a"
    }
}

resource "aws_subnet" "db_zone2" {
    vpc_id            = data.terraform_remote_state.eks.outputs.vpc_id
    cidr_block        = var.subnet_cidr2
    availability_zone = "${var.aws_region}b"

    tags = {
        "Name" = "${var.env}-db-subnet-${var.aws_region}b"
    }
}

resource "aws_db_subnet_group" "main" {
    name       = "${var.env}-db-subnet-group"
    subnet_ids = [aws_subnet.db_zone1.id, aws_subnet.db_zone2.id]
}