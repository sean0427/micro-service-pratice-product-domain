mock_gen:
	mockgen -package=mock --source=mongodb/mongodb.go > mock/mongo_client_mock.go
	mockgen -package=mock --source=product.go > mock/repo_mock.go

test:
	go test ./...
