resource "aws_vpc" "main" {
    cidr_block           = var.vpc_cidr
    enable_dns_support   = true
    enable_dns_hostnames = true
}

resource "aws_subnet" "private_1" {
    vpc_id            = aws_vpc.main.id
    cidr_block        = var.subnet_1_cidr
    availability_zone = data.aws_availability_zones.available.names[0]
}

resource "aws_subnet" "private_2" {
    vpc_id            = aws_vpc.main.id
    cidr_block        = var.subnet_2_cidr
    availability_zone = data.aws_availability_zones.available.names[1]
}

resource "aws_db_subnet_group" "rds_sng" {
    name       = "${var.project}-sng"
    subnet_ids = [aws_subnet.private_1.id, aws_subnet.private_2.id]
}

resource "aws_security_group" "rds_sg" {
    name   = "${var.project}-rds-sg"
    vpc_id = aws_vpc.main.id

    ingress {
        from_port   = 5432
        to_port     = 5432
        protocol    = "tcp"
        cidr_blocks = [aws_vpc.main.cidr_block]
    }
}