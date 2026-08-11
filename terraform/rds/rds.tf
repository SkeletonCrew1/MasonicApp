resource "aws_db_instance" "main" {
    identifier                  = "paranormal-${local.env}"
    engine                      = "postgres"
    engine_version              = "16"
    instance_class              = "db.t4g.micro"
    allocated_storage           = 20

    db_name                     = local.db_name
    username                    = local.db_user

    manage_master_user_password = true

    db_subnet_group_name   = aws_db_subnet_group.main.name
    vpc_security_group_ids = [aws_security_group.rds.id]

    skip_final_snapshot = true
    multi_az            = false
}