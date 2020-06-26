# GO GRPC CHAT

### RUN WITH DOCKER

Start server
```
docker-compose up
```

Run client as much as you want. Specify name via CHATNAME environment variable
```
CHATNAME=Joe docker-compose -f docker-compose-client.yml run client
```

