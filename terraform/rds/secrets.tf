resource "random_password" "db_password" {
    length           = 16
    special          = true
    override_special = "!#$%&*()-_=+[]{}<>:?"
}

resource "aws_secretsmanager_secret" "db_password" {
    name                    = "${local.env}/paranormal/db-password"
    description             = "Pass for Paranormal RDS"
    recovery_window_in_days = 0
}

resource "aws_secretsmanager_secret_version" "db_password" {
    secret_id     = aws_secretsmanager_secret.db_password.id
    secret_string = jsonencode({
        username = local.db_user
        password = random_password.db_password.result
        engine   = "postgres"
        port     = 5432
     })
}