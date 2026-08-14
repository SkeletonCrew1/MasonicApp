# K8S
## What is this folder?
This folder specifically is fodler for our application Kubernetes implementation and all of its dependencies.
Folders breakdown:
* secretsmanager - is a folder for our application secrets manager dependency installation and cleanup commands and a policy file used for it.

## How to run?
* Clone this repository.
* Log into AWS using `aws configure` command and providing your credentials and you must have enough permissions to perform these actions.
* If EKS Cluster is online use the next commands inside your terminal of choice to connect to EKS Cluster:
```
aws sts assume-role   --role-arn arn:aws:iam::<account id>:role/eks-admin   --role-session-name session
aws eks update-kubeconfig   --region eu-north-1   --name eks-cluster   --role-arn arn:aws:iam::<account id>:role/eks-admin
```
Do not forget to change "account id" to actual value.
* Given you are inside project repostitory, run `cd k8s` and then run `helm install masonic ./masonic-chart/masonic/ --namespace application --create-namespace --values ./masonic-chart/masonic/values-stage.yaml` to start our whole application.
NOTE! if you make a deployment to production or dev stage change `./masonic-chart/masonic/values-stage.yaml` to respective values file.

## How to clean up the application?
Stopping our application is really easy, just run `helm uninstall masonic -n application`.