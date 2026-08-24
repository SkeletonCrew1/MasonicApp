output "vpc_id" {
    value       = aws_vpc.main.id
    description = "for dev infra" 
}

output "vpc_cidr_block" {
    value       = aws_vpc.main.cidr_block
    description = "boundaries of the primary network" 
}