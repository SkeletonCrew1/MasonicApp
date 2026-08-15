data "http" "alb_iam_policy" {
  url = "https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/v2.14.1/docs/install/iam_policy.json"
}

resource "aws_iam_policy" "aws_load_balancer_controller_policy" {
  name        = "AWSLoadBalancerControllerIAMPolicy"
  description = "IAM policy for AWS Load Balancer Controller"
  policy      = data.http.alb_iam_policy.response_body
}

data "aws_iam_policy_document" "alb_controller_assume_role" {
  statement {
    actions = ["sts:AssumeRoleWithWebIdentity"]
    effect  = "Allow"

    principals {
      type        = "Federated"
      identifiers = [aws_iam_openid_connect_provider.eks.arn]
    }

    condition {
      test     = "StringEquals"
      variable = "${replace(aws_iam_openid_connect_provider.eks.url, "https://", "")}:sub"
      values   = ["system:serviceaccount:kube-system:aws-load-balancer-controller"]
    }
  }
}

resource "aws_iam_role" "aws_load_balancer_controller" {
  name               = "AmazonEKSLoadBalancerControllerRole"
  assume_role_policy = data.aws_iam_policy_document.alb_controller_assume_role.json
}

resource "aws_iam_role_policy_attachment" "alb_controller_attachment" {
  policy_arn = aws_iam_policy.aws_load_balancer_controller_policy.arn
  role       = aws_iam_role.aws_load_balancer_controller.name
}

resource "helm_release" "aws_load_balancer_controller" {
  name       = "aws-load-balancer-controller"
  repository = "https://aws.github.io/eks-charts"
  chart      = "aws-load-balancer-controller"
  namespace  = "kube-system"
  version    = "1.14.0"

  values = [
    yamlencode({
      clusterName = aws_eks_cluster.eks.name
      region      = "eu-north-1"
      vpcId       = aws_vpc.main.id
      serviceAccount = {
        create = true
        name   = "aws-load-balancer-controller"
        annotations = {
          "eks.amazonaws.com/role-arn" = aws_iam_role.aws_load_balancer_controller.arn
        }
      }
      subnetTags = {
        "kubernetes.io/role/elb" = "1"
      }
    })
  ]

  depends_on = [
    aws_iam_role_policy_attachment.alb_controller_attachment,
    aws_eks_node_group.general
  ]
}

resource "aws_ec2_tag" "public_subnet_tags_1" {
  resource_id = aws_subnet.public_zone1.id
  key         = "kubernetes.io/cluster/${aws_eks_cluster.eks.name}"
  value       = "shared"
}

resource "aws_ec2_tag" "public_subnet_tags_2" {
  resource_id = aws_subnet.public_zone2.id
  key         = "kubernetes.io/cluster/${aws_eks_cluster.eks.name}"
  value       = "shared"
}

resource "aws_ec2_tag" "public_subnet_role_1" {
  resource_id = aws_subnet.public_zone1.id
  key         = "kubernetes.io/role/elb"
  value       = "1"
}

resource "aws_ec2_tag" "public_subnet_role_2" {
  resource_id = aws_subnet.public_zone2.id
  key         = "kubernetes.io/role/elb"
  value       = "1"
}