resource aws_dynamodb_table connection {
  name         = replace("${var.system_name_prefix}_connection", "_", "-")
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "connectionId"

  attribute {
    name = "connectionId"
    type = "S"
  }
  point_in_time_recovery {
    enabled = true
  }
}