# Creating an EKS cluster using Terraform

This documentation describes how to create an EKS cluster and manage policies using Terraform.

## Preriqusites

- Git installed
- AWS account
- Terraform installed

## Folder structure
- `provider.tf` defines the cloud provider and backend.
- `locals.tf` defines local variables.
- `vpc.tf` defines the private network.
- `subnets.tf` defines two private and two public subnets. 
- `routes.tf` defines routes for subnets.
- `igw.tf` defines an internet gateway for public subnets.
- `nat.tf` defines a NAT for private subnets.
- `eks.tf` defines an EKS cluster.
- `nodes.tf` defines nodes configuration. 
- `eks-admin-role.tf` defines an admin role for managing the EKS cluster.
- `aws-lb-controller.tf` sets up AWS Load Balancer Controller.
- `secrets-csi-driver.f` sets up secrets manager driver for our EKS cluster.

## Usage

```bash
1. Configure AWS credentials
2. Clone the MasonicApp repository
3. Change branch to SKEL2-49-implementation-eks-cluster
4. Change working directory to MasonicApp/terraform/eks-cluster/
5. Run terraform init
6. Run terraform apply 
7. Run the following commands and paste the AWS user ID there.
aws eks update-kubeconfig --region eu-north-1 --name eks-cluster
aws sts assume-role --role-arn arn:aws:iam::<USER-ID>:role/eks-admin --role-session-name session
aws eks update-kubeconfig --region eu-north-1 --name eks-cluster --role-arn arn:aws:iam::<USER-ID>:role/eks-admin
```

## Additional:
Jenkins, ArgoCD and Grafana passwords are stored in AWS Secrets Manager.
To expose grafana use: `kubectl port-forward svc/grafana -n monitoring 3000:80`
To expose ArgoCD use: `kubectl port-forward svc/argocd-server -n argocd 443:443`
To expose Jenkins use: `kubectl --namespace jenkins port-forward svc/jenkins 8080:8080`