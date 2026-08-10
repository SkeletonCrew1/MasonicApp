resource "aws_security_group" "rds" {
    name   = "${local.env}-rds-sg"
    vpc_id = aws_vpc.main.id

    ingress {
        from_port = 5432
        to_port   = 5432
        protocol  = "tcp"
        cidr_blocks = [local.vpc_cidr]
    }
}