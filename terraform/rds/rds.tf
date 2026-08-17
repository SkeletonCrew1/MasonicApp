resource "aws_db_instance" "main" {
    identifier                  = "paranormal-${var.env}"
    engine                      = "postgres"
    engine_version              = "18"
    allow_major_version_upgrade = true
    instance_class              = "db.t4g.micro"
    allocated_storage           = 20

    db_name                     = var.db_name
    username                    = var.db_user

    password                    = random_password.db_password.result

    db_subnet_group_name   = aws_db_subnet_group.main.name
    vpc_security_group_ids = [aws_security_group.rds.id]

    skip_final_snapshot = true
    multi_az            = false

    publicly_accessible  = true
}
