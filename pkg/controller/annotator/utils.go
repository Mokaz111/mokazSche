package annotator

import (
	"fmt"
	"mokazSche/pkg/plugins/apis/policy"
	"strings"
	"time"
)

func getMaxHotVauleTimeRange(hotValues []policy.HotValuePolicy) time.Duration {
	var max time.Duration

	if hotValues == nil {
		return max
	}

	for _, tr := range hotValues {
		if max < tr.TimeRange.Duration {
			max = tr.TimeRange.Duration
		}
	}

	return max
}
func handlingMetaKeyWithMetricName(nodeName, metricName string) string {
	return nodeName + "/" + metricName
}
func splitMetaKeyWithMetricName(key string) (string, string, error) {
	parts := strings.Split(key, "/")

	if len(parts) != 2 {
		return "", "", fmt.Errorf("unexpected key format: %q", key)
	}

	return parts[0], parts[1], nil
}
