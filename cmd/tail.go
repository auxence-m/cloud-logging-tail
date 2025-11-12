package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/auxence-m/cloudtail/stream"
	"github.com/spf13/cobra"
)

var (
	logName      string
	resourceType string
	severity     string
	since        string
	sinceTime    string
	follow       bool
	limit        int
	output       string
)

// tailCmd represents the tail command
var tailCmd = &cobra.Command{
	Use:          "tail [projectID]",
	Short:        "Stream Google Cloud Logging entries directly into the terminal in real time",
	Long:         `The tail command will fetch and stream all Google Cloud Logging entries from the last 24 hours by default unless specified otherwise with the available flags`,
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	RunE:         tailRun,
}

func tailRun(cmd *cobra.Command, args []string) error {
	var (
		parseDuration time.Duration
		parseTime     time.Time
		parseSeverity string
	)

	// Trim flags
	trimmedLogName := strings.TrimSpace(logName)
	trimmedResourceType := strings.TrimSpace(resourceType)
	trimmedSeverity := strings.TrimSpace(severity)
	trimmedSince := strings.TrimSpace(since)
	trimmedSinceTime := strings.TrimSpace(sinceTime)

	// Validate severity flag
	if trimmedSeverity != "" {
		s, err := validateSeverityFlag(trimmedSeverity)
		if err != nil {
			return err
		}
		parseSeverity = s
	}

	// Validate since flag
	if trimmedSince != "" {
		d, err := validateSinceFlag(trimmedSince)
		if err != nil {
			return err
		}
		parseDuration = d
	}

	// Validate sinceTime flag
	if trimmedSinceTime != "" {
		t, err := validateSinceTimeFlag(trimmedSinceTime)
		if err != nil {
			return err
		}
		parseTime = t
	}

	filter := stream.Filter{
		LogName:      trimmedLogName,
		ResourceType: trimmedResourceType,
		Severity:     parseSeverity,
		Since:        parseDuration,
		SinceTime:    parseTime,
	}

	filterStr := stream.BuildFilterString(&filter)

	fmt.Println(filterStr)

	return nil
}

// validateSeverityFlag ensures the --severity flag has a valid value
func validateSeverityFlag(severity string) (string, error) {
	upper := strings.ToUpper(severity)

	validSeverities := map[string]struct{}{
		"INFO":    {},
		"DEBUG":   {},
		"WARNING": {},
		"NOTICE":  {},
		"ERROR":   {},
	}

	_, found := validSeverities[upper]
	if !found {
		return "", fmt.Errorf("invalid value for --severity flag: %q. (valid values: INFO, WARNING, ERROR, etc.)", severity)
	}

	return upper, nil
}

// validateSinceFlag validates a --since flag in the form of "1h", "30m", or "20s" and converts it into a time.Duration.
func validateSinceFlag(since string) (time.Duration, error) {
	parseDuration, err := time.ParseDuration(since)
	if err != nil {
		return 0, fmt.Errorf("invalid value for --since flag: %q (valid values: 1h, 30m, 20s, 1h15m30s, etc.): \n%w", since, err)
	}

	if parseDuration < 0 {
		return 0, fmt.Errorf("the --since flag duration must be positive (got %q)", since)
	}

	return parseDuration, nil
}

// validateSinceTimeFlag validates that the --since-time flag is a valid RFC3339 timestamp.
func validateSinceTimeFlag(sinceTime string) (time.Time, error) {
	parsedTime, err := time.Parse(time.RFC3339, sinceTime)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid value for --sinceTime flag: %q (must be RFC3339 format): \n%w", sinceTime, err)
	}

	return parsedTime, nil
}

func init() {
	rootCmd.AddCommand(tailCmd)

	tailCmd.Flags().StringVar(&logName, "logName", "", "Retrives the logs with the specified logName")
	tailCmd.Flags().StringVar(&resourceType, "resource-type", "", "Retrives the logs with the specified resource-type")
	tailCmd.Flags().StringVar(&severity, "severity", "", "Retrives the logs with the specified severity level. (e.g., INFO, WARNING, ERROR)")
	tailCmd.Flags().StringVar(&since, "since", "", "Retrieves logs newer than a specified relative duration (e.g., 1h, 30m, 20s, 1h15m30s). Only one of --since-time or --since may be used")
	tailCmd.Flags().StringVar(&sinceTime, "since-time", "", "Retrieves logs newer than a specific timestamp in RFC3339 format (e.g., YYYY-MM-DDTHH:MM:SSZ). Only one of --since-time or --since may be used")

	tailCmd.MarkFlagsMutuallyExclusive("since", "since-time")

	tailCmd.Flags().BoolVar(&follow, "follow", false, "Specify if the logs should be streamed in real-time as they are generated")
	tailCmd.Flags().IntVar(&limit, "limit", -1, "Number of recent logs to display. Defaults to -1 with no effect, showing all logs")
	tailCmd.Flags().StringVar(&output, "output", "", "Specify the output file to write the logs to")
}
