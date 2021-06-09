terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

provider "aws" {
  region = "ap-northeast-1"
}

data aws_region current {}

variable aws_env {
  default = "dev"
}

variable system_name_prefix {
  default = "message-service-dev"
}


output websocket_url {
  value = aws_apigatewayv2_api.websocket_gw.api_endpoint
}

