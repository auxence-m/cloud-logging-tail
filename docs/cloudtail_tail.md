## cloudtail tail

Display and stream Google Cloud Logging entries matching the specified filters

### Synopsis

The tail command will fetch and list all Google Cloud Logging entries from the last 24 hours by default unless specified otherwise with the available flags

```
cloudtail tail [projectID] [flags]
```

### Examples

```

The following examples demonstrate common usage patterns for tail.

# Stream all logs in real time
cloudtail tail projectID --follow

# Stream logs from a specific resource type
cloudtail tail projectID --resource-type=gce_instance --follow

# Stream only ERROR severity logs
cloudtail tail projectID --severity=ERROR --follow

# Display the most recent 100 log entries
cloudtail tail projectID --limit=100

# Display logs from the last 30 minutes
cloudtail tail projectID --since=30m

# Display logs from the last hour and continue streaming
cloudtail tail projectID -since=1h --follow

# Display logs newer than a specific point in time
cloudtail tail projectID --since-time=2026-02-12T12:30:00Z

# Filter logs by log name and resource type
cloudtail tail projectID \
	--log-name=projects/projectID/logs/cloudbuild \
	--resource-type=k8s_container

# Combine severity and time-based filtering
cloudtail tail projectID --severity=WARNING --since=1h

# Use an advanced filter expression for complex queries
cloudtail tail projectID \
	--filter='severity>="ERROR" AND timestamp>="2026-01-01T00:00:00Z" AND timestamp<="2023-01-31T12:00:00Z"'

# Combine advanced filtering with a result limit
cloudtail tail projectID \
	--filter='severity>="ERROR" AND timestamp>="2026-01-01T00:00:00Z" AND timestamp<="2023-01-31T12:00:00Z"' \
	--limit=100

# Stream logs using an advanced filter expression
cloudtail tail projectID --filter='severity>="CRITICAL"' --follow

# Write log output to a file instead of stdout
cloudtail tail projectID --severity=INFO --output=logs.txt

# Stream logs and write them to a file
cloudtail tail projectID --follow --output=logs.txt

# Combine multiple filters for a focused query
cloudtail tail projectID \
	--log-name=projects/projectID/logs/cloudbuild \
	--resource-type=k8s_container \
	--severity=ERROR \
	--since=15m

# Retrieve recent logs using a fixed timestamp and save them
cloudtail tail projectID \
	--since-time=2026-01-13T12:30:00Z \
	--limit=200 \
	--output=incident.log

Notes:
  - To include historical logs, use --since or --since-time. 
    A timestamp in --filter alone does not include past entries.
  - --follow streams logs in real-time. Without --since or --since-time, 
    only new entries from the time of execution are shown.
  - --limit applies only to the initial historical fetch. Streaming ignores it.

```

### Options

```
      --filter string          Apply a raw filter expression for advanced queries (e.g. severity>="WARNING" AND severity<="ERROR" AND timestamp>="2026-05-18T12:00:00Z")
  -f, --follow                 Stream new log entries as they are generated
  -h, --help                   help for tail
  -n, --limit int              Maximum number of logs to display (defaults to -1, showing all logs). (default -1)
      --log-name string        Filter logs by log name
  -o, --output string          Write logs to the specified file (defaults to stdout).
      --resource-type string   Filter logs by resource type
      --severity string        Filter logs by severity level (DEFAULT, DEBUG, INFO, NOTICE, WARNING, ERROR, CRITICAL, ALERT, EMERGENCY)
      --since string           Show logs newer than a relative duration (e.g. 1h, 30m, 20s, 1h15m30s). Only one of since-time / since may be used
      --since-time string      Show logs newer than an RFC3339 timestamp (e.g. 2026-01-13T12:30:00Z). Only one of since-time / since may be used
```

### SEE ALSO

* [cloudtail](cloudtail.md)	 - cloudtail displays or tail logs from Google Cloud Logging

