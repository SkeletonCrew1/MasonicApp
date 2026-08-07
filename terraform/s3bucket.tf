resource "aws_s3_bucket" "main" {
  bucket = "masonicapp-s3bucket"

  tags = {
    ManagedBy = "Terraform"
    Name      = "masonicapp-s3bucket"
  }
}

resource "aws_s3_bucket_versioning" "main" {
  bucket = aws_s3_bucket.main.id
  versioning_configuration {
    status = "Enabled"
  }
}
