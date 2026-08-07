output "db_endpoint" {
    description = "Database connection endpoint"
    value       = module.db.db_instance_endpoint
}
