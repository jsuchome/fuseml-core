package workflow

import (
	"time"

	"github.com/jonboulle/clockwork"
	"github.com/tektoncd/cli/pkg/formatted"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func formatDesc(desc string) string {
	if len(desc) > 60 {
		return desc[0:59] + "..."
	}
	return desc
}

func formatAge(startTime *string) string {
	if startTime != nil {
		st, _ := time.Parse(time.RFC3339, *startTime)
		return formatted.Age(&v1.Time{Time: st}, clockwork.NewRealClock())
	}
	return "---"
}

func formatDuration(startTime, completionTime *string) string {
	layout := time.RFC3339
	st := time.Time{}
	ct := time.Time{}
	if startTime != nil {
		st, _ = time.Parse(layout, *startTime)
	}
	if completionTime != nil {
		ct, _ = time.Parse(layout, *completionTime)
	}
	return formatted.Duration(&v1.Time{Time: st}, &v1.Time{Time: ct})
}
