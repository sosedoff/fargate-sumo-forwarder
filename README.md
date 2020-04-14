# fargate-sumo-forwarder

Forward AWS ECS/Fargate container logs directly into Sumologic collector.

## Overview

There are a few ways to get the logs shipped from ECS/Fargate to Sumologic:

- Use `awslogs` driver to send logs to AWS CLoudWatch Logs and then ship to Sumo via lambda function
- Use `awsfirelens` driver which is a sidecar container running `fluentbit`. 
- Use `splunk` driver in combination with this forwarder software.

## Configuration

Make sure to first set the environment variables:

- `PORT` - Server port
- `COLLECTOR_URL` - Sumologic hosted collector URL
- `COLLECTOR_WORKERS` - Number of log shipping workers

Start the server by running: 

```
fargate-sumo-forwarder
```

## ECS Configuration

```json
{
  "logConfiguration": {
    "logDriver": "splunk",
    "options": {
      "splunk-url": "https://your-collector-url",
      "splunk-token": "token",
      "splunk-source": "source",
      "splunk-index": "index",
      "splunk-format": "inline"
    }
  }
}
```

## License

MIT