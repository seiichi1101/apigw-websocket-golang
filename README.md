# API Gateway Websocket with Golang

## Architecture

![](pics/sys_arch.png)

## init

```
make init
```

## Build and Deploy

```
make deploy
```

## Verification

use [wscat](https://github.com/websockets/wscat)

![](./pics/wscat.gif)

- connection

```
wscat -c wss://<api-id>.execute-api.<region>.amazonaws.com/<stage>
```

- send message

```
{"action": "send_message", "targetId": "<target-connection-id>", "message": "hello world"}
```
