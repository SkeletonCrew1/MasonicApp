resource "aws_db_instance" "databases" {
  for_each = var.databases

  identifier             = "${var.project}-db-${each.key}"
  engine                 = "postgres"
  engine_version         = "18"
  instance_class         = "db.t3.micro"
  allocated_storage      = 10

  db_name                = "${each.key}_db"
  username               = "postgres"
  password               = var.db_password

  skip_final_snapshot    = true
  publicly_accessible    = false

  db_subnet_group_name   = aws_db_subnet_group.rds_sng.name
  vpc_security_group_ids = [aws_security_group.rds_sg.id]
}