# K8S

## What is this folder?
This folder specifically is fodler for our application Kubernetes Helm implementation.

## Breakdown:
* `masonic-chart\masonic` breakdown:
* * `values-stage.yaml` - values used by Helm for stage environment
* `masonic-chart\masonic\templates` breakdown:
* * `application-ingress.yaml` - is an ingress manifest to expose our application to the internet.
* * `auth-deployment.yaml` - is a deployment manifest for auth microservice.
* * `auth-service.yaml` - is a service manifest that looks for auth pods.
* * `backend-deployment.yaml` - is a deployment manifest for backend microservice.
* * `backend-service.yaml` - is a service manifest that looks for backend pods.
* * `frontend-deployment.yaml` - is a deployment manifest for frontend microservice.
* * `frontend-service.yaml` - is a service manifest that looks for frontend pods.
* * `inquisitor-deployment.yaml` - is a deployment manifest for daily inquisitor microservice.
* * `mail-deployment.yaml` - is a deployment manifest for mail microservice.
* * `mail-service.yaml` - is a service manifest that look for mail pods.
* * `password-deployment.yaml` - is a deployment manifest for daily password microservice.
* * `posting-deployment.yaml` - is a deployment manifest for posting microservice.
* * `posting-service.yaml` - is a service manifest that looks for posting pods.
* * `secrets-provider.yaml` - is a secrets provider class manifest that fetches all of application required secrets and adds them to a secrets map.

## How to run our application without Jenkins and/or ArgoCD?
* Clone this repository.
* Log into AWS using `aws configure` command and providing your credentials and you must have enough permissions to perform these actions.
* Log into AWS EKS Cluster using (do not forget to change account id to actual value):
```
aws sts assume-role   --role-arn arn:aws:iam::<account id>:role/eks-admin   --role-session-name session
aws eks update-kubeconfig   --region eu-north-1   --name eks-cluster   --role-arn arn:aws:iam::<account id>:role/eks-admin
```
* Given you are inside project repostitory, run `cd k8s` and then run `helm install masonic ./masonic-chart/masonic/ --namespace application --create-namespace --values ./masonic-chart/masonic/values-stage.yaml` to start our whole application, If you make a deployment to production or dev stage change `./masonic-chart/masonic/values-stage.yaml` to the respective values file.

## How to clean up the application without ArgoCD?
Stopping our application is really easy, just run `helm uninstall masonic -n application`.
