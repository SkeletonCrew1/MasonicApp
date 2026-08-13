# K8S
## What is this folder?
This folder specifically is fodler for our application Kubernetes implementation and all of its dependencies.
Folders breakdown:
* secretsmanager - is a folder for our application secrets manager dependency installation and cleanup commands and a policy file used for it.
* ingress - is a folder for our application AWS Load Balancer Controller (type of ingress controller) dependency installation and cleanup together with a policy used for it.
* masonic-chart - is a folder for a helm chart of our application, it includes everything: application deployment, application services, ingress and our secrets provider class (used to dynamically pass secrets from AWS Secrets Manager to our application).

## How to run?
* Clone this repository.
* Log into AWS using `aws configure` command and providing your credentials and you must have enough permissions to perform these actions.
* If EKS Cluster is online use the next commands inside your terminal of choice to connect to EKS Cluster:
```
aws sts assume-role   --role-arn arn:aws:iam::<account id>:role/eks-admin   --role-session-name session
aws eks update-kubeconfig   --region eu-north-1   --name eks-cluster   --role-arn arn:aws:iam::<account id>:role/eks-admin
```
Do not forget to change "account id" to actual value.
* Given you are inside project repostitory, run `cd k8s\secretsmanager` and run all the commands inside `setup.txt` in exact order.
* Given you are inside project repostitory, run `cd k8s\ingress`and run all the commands inside `deployment_commands.txt` in order.
* Given you are inside project repostitory, run `cd k8s` and then run `helm install masonic ./masonic-chart/masonic/ --namespace application --create-namespace` to start our whole application.

## How to clean up cluster and application?
Stopping our application is really easy, just run `helm uninstall masonic -n application`.
To clean up the cluster follow the next steps:
* Given you are inside project repostitory, run `cd k8s\secretsmanager` and run all the commands inside `cleanup-secrets.txt` in exact order.
* Given you are inside project repostitory, run `cd k8s\ingress` and run all the commands inside `cleanup-ingress.txt` in exact order.