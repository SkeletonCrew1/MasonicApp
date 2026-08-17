resource "random_password" "db_password" {
    length           = 16
    special          = true
    override_special = "!#$%&*()-_=+[]{}<>:?"
}

locals {
    db_secrets = {
        "USERS_DB_HOST"     = aws_db_instance.main.address
        "USERS_DB_PORT"     = "5432"
        "USERS_DB_USER"     = var.db_user
        "USERS_DB_PASSWORD" = random_password.db_password.result
        "USERS_DB_NAME"     = "main_db"

        "MAP_DB_HOST"       = aws_db_instance.main.address
        "MAP_DB_PORT"       = "5432"
        "MAP_DB_USER"       = var.db_user
        "MAP_DB_PASSWORD"   = random_password.db_password.result
        "MAP_DB_NAME"       = "map_db"

        "VOTING_DB_HOST"       = aws_db_instance.main.address
        "VOTING_DB_PORT"       = "5432"
        "VOTING_DB_USER"       = var.db_user
        "VOTING_DB_PASSWORD"   = random_password.db_password.result
        "VOTING_DB_NAME"       = "voting_db"
    }
}

resource "aws_secretsmanager_secret" "db_secrets" {
    for_each                = local.db_secrets
    name                    = each.key
    description             = "Pass for Paranormal RDS for ${each.key}"
    recovery_window_in_days = 0
}

resource "aws_secretsmanager_secret_version" "db_password" {
    for_each      = local.db_secrets
    secret_id     = aws_secretsmanager_secret.db_secrets[each.key].id
    secret_string = jsonencode({
        (each.key) = each.value
     })
}