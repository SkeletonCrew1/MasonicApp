output "user_passwords" {
  value = {
    for user_key, profile in aws_iam_user_login_profile.users :
    user_key => profile.password
  }
  description = "A map of usernames to their initial temporary passwords"
  sensitive   = true
}
