test:
	go test
test_cover:
	go test -cover
test_cover_profile:
	go test -coverprofile=coverage.out
test_tool_cover: test_cover_profile
	go tool cover -html=coverage.out