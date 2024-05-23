provider "aws" {
  region = "us-east-1"
}

terraform {
  backend "s3" {
    bucket  = "email-searcher-tfstates"
    key     = "email-searcher-terraform.tfstate"
    region  = "us-east-1"
    encrypt = true
  }
}
