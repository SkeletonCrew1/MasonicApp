resource "aws_security_group" "rds" {
    name   = "${var.env}-rds-paranormal-sg"
    vpc_id = data.terraform_remote_state.eks.outputs.vpc_id

    ingress {
        from_port   = 5432
        to_port     = 5432
        protocol    = "tcp"
        cidr_blocks = [data.terraform_remote_state.eks.outputs.vpc_cidr_block] 
    }

    egress {
        from_port   = 0
        to_port     = 0
        protocol    = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
}