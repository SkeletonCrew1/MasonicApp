resource "aws_iam_user" "users" {
  for_each = toset(local.users)

  name = each.value
  path = "/"

  tags = {
    ManagedBy = "Terraform"
  }
}

resource "aws_iam_user_login_profile" "users" {
  for_each = aws_iam_user.users

  user                    = each.value.name
  password_length         = 20
  password_reset_required = true

}

resource "aws_iam_user_policy_attachment" "users" {
  for_each = toset(local.users)

  user       = aws_iam_user.users[each.value].name
  policy_arn = "arn:aws:iam::aws:policy/AdministratorAccess"
}
