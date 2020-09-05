#go build -o main app/*
docker build . -t awssns_go
kind load docker-image awssns_go:latest
