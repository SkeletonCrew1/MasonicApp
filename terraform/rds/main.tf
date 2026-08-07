module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 6.0"

  name = "pg-vpc"
  cidr = "10.0.0.0/16"
  azs  = ["${var.aws_region}a", "${var.aws_region}b"]

  database_subnets             = ["10.0.1.0/24", "10.0.2.0/24"]
  create_database_subnet_group = true

  create_database_subnet_route_table = true
}

module "security_group" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 6.0"

  name        = "pg-sg"
  description = "Allow PostgreSQL access within VPC"
  vpc_id      = module.vpc.vpc_id

  ingress_rules = {
    postgresql = {
      from_port   = 5432
      to_port     = 5432
      ip_protocol    = "tcp"
      cidr_ipv4 = module.vpc.vpc_cidr_block
    }
  }
}

module "db" {
  source  = "terraform-aws-modules/rds/aws"
  version = "~> 6.0"

  identifier = "t4g-postgres"

  engine               = "postgres"
  engine_version       = "18"
  family               = "postgres18"
  major_engine_version = "18"

  instance_class    = "db.t4g.micro"
  allocated_storage = 20

  db_name  = "simpledb"
  username = "dbadmin"

  manage_master_user_password = false
  password                    = var.db_password

  db_subnet_group_name   = module.vpc.database_subnet_group_name
  vpc_security_group_ids = [module.security_group.id]
  multi_az               = false
  publicly_accessible    = false

  backup_retention_period  = 0
  skip_final_snapshot      = true
}