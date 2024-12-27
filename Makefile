gen-model-migrations:
	# Generate model migrations
	go run ./gen
	# Tidy up file
	go fmt