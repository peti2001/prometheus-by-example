# Mock service with Prometheus exporter

This is a mock service that can be used to test your Prometheus configuration. It contains a service that generates random events to provide mock data for Prometheus

# How to use it?
```
$ git clone https://github.com/peti2001/prometheus-by-example.git
$ cd prometheus-by-example/job-processor
$ docker-compose up
```
Here you can see the mock service's Prometheus exporter: http://localhost:9009/metrics.

1. Open http://localhost:3000 and you can login with admin:admin.
1. Add Prometheus as new source
1. Start adding metrics. E.g: `http_request_duration_microseconds`
