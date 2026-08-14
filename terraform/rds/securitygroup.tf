resource "aws_security_group" "rds" {
    name   = "${var.env}-rds-paranormal-sg"
    vpc_id = aws_vpc.main.id

    ingress {
        from_port = 5432
        to_port   = 5432
        protocol  = "tcp"
        cidr_blocks = [var.vpc_cidr]
    }
}