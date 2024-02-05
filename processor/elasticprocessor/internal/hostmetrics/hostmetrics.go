package hostmetrics

import (
	"fmt"
	"path"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

var scraperToElasticDataset = map[string]string{
	"cpu":        "system.cpu",
	"disk":       "system.diskio",
	"filesystem": "system.filesystem",
	"load":       "system.load",
	"memory":     "system.memory",
	"network":    "system.network",
	"paging":     "system.memory",
	"processes":  "system.process.summary",
	"process":    "system.process",
}

// AddElasticSystemMetrics computes additional metrics for compatibility with the Elastic system integration.
// The `scopeMetrics` input should be metrics generated by a specific hostmetrics scraper.
// `scopeMetrics` are modified in place.
func AddElasticSystemMetrics(scopeMetrics pmetric.ScopeMetrics) error {
	scope := scopeMetrics.Scope()
	scraper := path.Base(scope.Name())

	dataset, ok := scraperToElasticDataset[scraper]
	if !ok {
		return fmt.Errorf("no dataset defined for scaper '%s'", scraper)
	}

	switch scraper {
	case "cpu":
		return addCPUMetrics(scopeMetrics.Metrics(), dataset)
	case "memory":
		return addMemoryMetrics(scopeMetrics.Metrics(), dataset)
	case "load":
		return addLoadMetrics(scopeMetrics.Metrics(), dataset)
	case "process":
		return addProcessMetrics(scopeMetrics.Metrics(), dataset)
	case "processes":
		return addProcessesMetrics(scopeMetrics.Metrics(), dataset)
	default:
		return fmt.Errorf("no matching transform function found for scope '%s'", scope.Name())
	}
}
