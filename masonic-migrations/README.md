# Creating an EKS cluster using Terraform

This documentation describes how to provision an AWS RDS PostgreSQL instance. How to configure ESO and execute shema migrations using Flyway.

## Preriqusites

- Git installed
- AWS account
- Terraform installed
- kubectl configured

## Folder structure
- `provider.tf` defines the cloud provider and backend.
- `locals.tf` defines local variables.
- `terraform/rds/rds.tf` defines the AWS RDS instance.
- `terraform/rds/securitygroup.tf` defines the security group for RDS.
- `terraform/rds/secrets.tf` generates and stores database credentials in AWS Secrets Manager.
- `terraform/rds/providers.tf` defines the cloud provider and backend.
- `masonic-migrations/k8s/03-flyway-job.yaml` defines a Flyway migration job for 3 dbs.
- `masonic-migrations/sql/` contains schema migrations for the microservices architecture.

## Usage

```bash
1. Configure AWS credentials
2. Change working directory to MasonicApp/terraform/rds/
3. Run terraform init
4. Run terraform apply 
5. Retrieve USERS_DB_HOST and USERS_DB_PASSWORD from AWS Secrets Manager
6. Run an ephemeral pod: kubectl run pg-client --rm -i --tty --image=postgres:18-alpine --restart=Never -- /bin/sh
7. Inside the pod, connect to RDS: psql -h <USERS_DB_HOST> -p 5432 -U miraculous -d postgres -W
8. Create databases: CREATE DATABASE main_db; CREATE DATABASE map_db; CREATE DATABASE voting_db;
9. Exit psql (\q) and the pod (exit)
10. Change working directory to MasonicApp/
11. Run migrations: kubectl apply -f masonic-migrations/k8s/03-flyway-job.yaml -n application
```

## Additional:
Temporary solution: RDS is currently configured with publicly_accessible = true and ingress/egress rules are set to 0.0.0.0/0
And database credentials(e. g. ports, hosts, passwords, users) are stored in AWS Secrets Manager. 