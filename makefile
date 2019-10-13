swagger-check:
	which swagger

swagger-generate: swagger-check
	swagger generate spec -o ./swagger.yaml --scan-models

swagger-serve: swagger-check
	swagger serve -F=swagger swagger.yaml --port 80