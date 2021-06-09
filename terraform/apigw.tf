resource aws_apigatewayv2_api websocket_gw {
  name         = replace("${var.system_name_prefix}_websocket_gw", "_", "-")
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
}

resource aws_apigatewayv2_stage stage {
  api_id = aws_apigatewayv2_api.websocket_gw.id
  name   = var.aws_env
  auto_deploy = true
}

resource aws_apigatewayv2_integration connection {
  api_id           = aws_apigatewayv2_api.websocket_gw.id
  integration_type = "AWS_PROXY"

  connection_type           = "INTERNET"
  content_handling_strategy = "CONVERT_TO_TEXT"
  integration_method        = "POST"
  integration_uri           = aws_lambda_function.connection.invoke_arn
  passthrough_behavior      = "WHEN_NO_MATCH"
}

resource aws_apigatewayv2_route connection {
  api_id    = aws_apigatewayv2_api.websocket_gw.id
  route_key = "$connect"

  target = "integrations/${aws_apigatewayv2_integration.connection.id}"
}

resource aws_apigatewayv2_integration disconnection {
  api_id           = aws_apigatewayv2_api.websocket_gw.id
  integration_type = "AWS_PROXY"

  connection_type           = "INTERNET"
  content_handling_strategy = "CONVERT_TO_TEXT"
  integration_method        = "POST"
  integration_uri           = aws_lambda_function.disconnection.invoke_arn
  passthrough_behavior      = "WHEN_NO_MATCH"
}

resource aws_apigatewayv2_route disconnection {
  api_id    = aws_apigatewayv2_api.websocket_gw.id
  route_key = "$disconnect"

  target = "integrations/${aws_apigatewayv2_integration.disconnection.id}"
}


resource aws_apigatewayv2_integration send_message {
  api_id           = aws_apigatewayv2_api.websocket_gw.id
  integration_type = "AWS_PROXY"

  connection_type           = "INTERNET"
  content_handling_strategy = "CONVERT_TO_TEXT"
  integration_method        = "POST"
  integration_uri           = aws_lambda_function.send_message.invoke_arn
  passthrough_behavior      = "WHEN_NO_MATCH"
}

resource aws_apigatewayv2_route send_message {
  api_id    = aws_apigatewayv2_api.websocket_gw.id
  route_key = "send_message"

  target = "integrations/${aws_apigatewayv2_integration.send_message.id}"
}
