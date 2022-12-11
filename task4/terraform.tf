terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "terra_server" {
  ami           = "ami-0a6b2839d44d781b2"
  instance_type = "t2.micro"

  tags = {
    Name = "Testserver"
  }
}
