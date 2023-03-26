# bukalawak-api
API for Bukalawak

Cross Compile
```
GOOS=linux GOARCH=amd64 go build -o bin/app-amd64-linux main.go
```

Build Docker image
```
docker build -t bukalawak-api .
```

Docker run
```
docker run -dp 8000:8000 bukalawak-api
```