docker compose up consul -d


# curl -T data.json http://192.168.49.2:30682/v1/kv/grpc-svc
curl -T data.json http://127.0.0.1:8500/v1/kv/grpc-svc