resource "aws_s3_bucket" "tfstate" {
    bucket = "masons-infra-tfstate"
}

resource "aws_s3_bucket_versioning" "tfstate" {
    bucket = aws_s3_bucket.tfstate.id
    
    versioning_configuration {
      status = "Enabled"
    }
}