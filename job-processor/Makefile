build:
	@docker run --rm -v `pwd`/..:/app/prometheus-by-example -w /app/prometheus-by-example/job-processor -e CGO_ENABLED=0 -e GOOS=linux golang:1.11 go build  -a -tags netgo -ldflags '-w'
	@docker build -t peti2001/prometheus-by-example .
	@docker push peti2001/prometheus-by-example
	@rm -f job-processor