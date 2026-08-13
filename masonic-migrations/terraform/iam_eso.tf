resource "aws_iam_policy" "eso_db_policy" {
  name        = "${var.env}-eso-flyway-policy"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = ["secretsmanager:GetSecretValue", "secretsmanager:DescribeSecret"]
        Resource = [var.db_secret_arn]
      }
    ]
  })
}

module "irsa_eso" {
  source = "terraform-aws-modules/iam/aws//modules/iam-assumable-role-with-oidc"
  version = "5.48.0"

  create_role      = true
  role_description = "Role for External Secrets Operator to read DB creds"
  role_name        = "${var.env}-eso-secrets-role"
  provider_url     = var.eks_oidc_issuer_url
  role_policy_arns = [aws_iam_policy.eso_db_policy.arn]

  oidc_fully_qualified_subjects = ["system:serviceaccount:external-secrets:external-secrets"]
}
