# fargate-sumo-forwarder

Forward AWS ECS/Fargate container logs directly into Sumologic collector.

## Overview

There are a few ways to get the logs shipped from ECS/Fargate to Sumologic:

- Use `awslogs` driver to send logs to AWS CLoudWatch Logs and then ship to Sumo via lambda function
- Use `awsfirelens` driver which is a sidecar container running `fluentbit`. 
- Use `splunk` driver in combination with this forwarder software.

## Configuration

Make sure to first set the environment variables:

| Variable      | Description                    | Default |
|---------------|--------------------------------|---------|
| PORT          | Server port                    | 5000    |
| WORKERS       | Number of log delivery workers | 1       |
| AUTH_TOKEN    | Authentication token           |         |
| COLLECTOR_URL | Sumologic collector URL        |         |

Start the server by running: 

```
  PORT=8080 \
  WORKERS=5 \
  AUTH_TOKEN=token \
  COLLECTOR_URL=https://your-collector-url \
  fargate-sumo-forwarder
```

## ECS Configuration

```json
{
  "logConfiguration": {
    "logDriver": "splunk",
    "options": {
      "splunk-url": "https://your-forwarder-url:8080",
      "splunk-token": "token",
      "splunk-source": "source",
      "splunk-sourcetype": "sourcetype",
      "splunk-index": "index",
      "splunk-format": "inline"
    }
  }
}
```

## License

MIT