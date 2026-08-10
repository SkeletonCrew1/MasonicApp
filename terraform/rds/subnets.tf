resource "aws_subnet" "db_zone1" {
    vpc_id            = aws_vpc.main.id
    cidr_block        = local.subnet_cidr1
    availability_zone = local.zone1

    tags = {
        "Name" = "${local.env}-db-subnet-${local.zone1}"
    }
}

resource "aws_subnet" "db_zone2" {
    vpc_id            = aws_vpc.main.id
    cidr_block        = local.subnet_cidr2
    availability_zone = local.zone2

    tags = {
        "Name" = "${local.env}-db-subnet-${local.zone2}"
    }
}

resource "aws_db_subnet_group" "main" {
    name       = "${local.env}-db-subnet-group"
    subnet_ids = [aws_subnet.db_zone1.id, aws_subnet.db_zone2.id]
}