data "aws_caller_identity" "current" {}

resource "aws_iam_role" "eks-admin" {
   name = "eks-admin"
   assume_role_policy = jsonencode({
       Version = "2012-10-17"
       Statement = [
           {
               Effect = "Allow"
               Action = "sts:AssumeRole"
               Principal = {
                   AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
               }
           }
       ]
   })  
}

resource "aws_iam_policy" "eks-admin-policy" {
   name = "eks-admin-policy"
   policy = jsonencode({
       Version = "2012-10-17"
       Statement = [
           {
               Effect = "Allow"
               Action = [
                   "eks:*"
               ]
               Resource = "*"
           },
           {
               Effect = "Allow"
               Action = "iam:PassRole"
               Resource = "*"
               Condition = {
                   StringEquals = {
                       "iam:PassedToService" = "eks.amazonaws.com"
                   }
               }
           },
       ]
   })
}

resource "aws_iam_role_policy_attachment" "eks-admin-policy-attachment" {
   role       = aws_iam_role.eks-admin.name
   policy_arn = aws_iam_policy.eks-admin-policy.arn
}

resource "aws_iam_group" "eks-admin-group" {
   name = "eks-admin-group"
}

resource "aws_iam_user_group_membership" "eks-admin-group-membership" {
   for_each = toset(local.users)
   user = each.value

   groups = [
       aws_iam_group.eks-admin-group.name,
   ]
}

resource "aws_iam_group_policy_attachment" "eks-admin-policy-attachment" {
   group      = aws_iam_group.eks-admin-group.name
   policy_arn = aws_iam_policy.eks-admin-policy.arn
}

resource "aws_eks_access_entry" "manager" {
  cluster_name      = aws_eks_cluster.eks.name
  principal_arn     = aws_iam_role.eks-admin.arn
  kubernetes_groups = ["my-admin"]
}
