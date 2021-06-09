# API Gateway Websocket with Golang

## Architecture

![](pics/sys_arch.png)

## Build and Deploy

```
make deploy
```

## Verification

use [wscat](https://github.com/websockets/wscat)

```
wscat -c wss://<api-id>.execute-api.<region>.amazonaws.com/<stage>
```

![](./pics/wscat.gif)
