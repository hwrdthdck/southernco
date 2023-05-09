// Code generated by mdatagen. DO NOT EDIT.

package metadata

import "go.opentelemetry.io/collector/confmap"

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for expvarreceiver metrics.
type MetricsConfig struct {
	ProcessRuntimeMemstatsBuckHashSys   MetricConfig `mapstructure:"process.runtime.memstats.buck_hash_sys"`
	ProcessRuntimeMemstatsFrees         MetricConfig `mapstructure:"process.runtime.memstats.frees"`
	ProcessRuntimeMemstatsGcCPUFraction MetricConfig `mapstructure:"process.runtime.memstats.gc_cpu_fraction"`
	ProcessRuntimeMemstatsGcSys         MetricConfig `mapstructure:"process.runtime.memstats.gc_sys"`
	ProcessRuntimeMemstatsHeapAlloc     MetricConfig `mapstructure:"process.runtime.memstats.heap_alloc"`
	ProcessRuntimeMemstatsHeapIdle      MetricConfig `mapstructure:"process.runtime.memstats.heap_idle"`
	ProcessRuntimeMemstatsHeapInuse     MetricConfig `mapstructure:"process.runtime.memstats.heap_inuse"`
	ProcessRuntimeMemstatsHeapObjects   MetricConfig `mapstructure:"process.runtime.memstats.heap_objects"`
	ProcessRuntimeMemstatsHeapReleased  MetricConfig `mapstructure:"process.runtime.memstats.heap_released"`
	ProcessRuntimeMemstatsHeapSys       MetricConfig `mapstructure:"process.runtime.memstats.heap_sys"`
	ProcessRuntimeMemstatsLastPause     MetricConfig `mapstructure:"process.runtime.memstats.last_pause"`
	ProcessRuntimeMemstatsLookups       MetricConfig `mapstructure:"process.runtime.memstats.lookups"`
	ProcessRuntimeMemstatsMallocs       MetricConfig `mapstructure:"process.runtime.memstats.mallocs"`
	ProcessRuntimeMemstatsMcacheInuse   MetricConfig `mapstructure:"process.runtime.memstats.mcache_inuse"`
	ProcessRuntimeMemstatsMcacheSys     MetricConfig `mapstructure:"process.runtime.memstats.mcache_sys"`
	ProcessRuntimeMemstatsMspanInuse    MetricConfig `mapstructure:"process.runtime.memstats.mspan_inuse"`
	ProcessRuntimeMemstatsMspanSys      MetricConfig `mapstructure:"process.runtime.memstats.mspan_sys"`
	ProcessRuntimeMemstatsNextGc        MetricConfig `mapstructure:"process.runtime.memstats.next_gc"`
	ProcessRuntimeMemstatsNumForcedGc   MetricConfig `mapstructure:"process.runtime.memstats.num_forced_gc"`
	ProcessRuntimeMemstatsNumGc         MetricConfig `mapstructure:"process.runtime.memstats.num_gc"`
	ProcessRuntimeMemstatsOtherSys      MetricConfig `mapstructure:"process.runtime.memstats.other_sys"`
	ProcessRuntimeMemstatsPauseTotal    MetricConfig `mapstructure:"process.runtime.memstats.pause_total"`
	ProcessRuntimeMemstatsStackInuse    MetricConfig `mapstructure:"process.runtime.memstats.stack_inuse"`
	ProcessRuntimeMemstatsStackSys      MetricConfig `mapstructure:"process.runtime.memstats.stack_sys"`
	ProcessRuntimeMemstatsSys           MetricConfig `mapstructure:"process.runtime.memstats.sys"`
	ProcessRuntimeMemstatsTotalAlloc    MetricConfig `mapstructure:"process.runtime.memstats.total_alloc"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		ProcessRuntimeMemstatsBuckHashSys: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsFrees: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsGcCPUFraction: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsGcSys: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsHeapAlloc: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsHeapIdle: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsHeapInuse: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsHeapObjects: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsHeapReleased: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsHeapSys: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsLastPause: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsLookups: MetricConfig{
			Enabled: false,
		},
		ProcessRuntimeMemstatsMallocs: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsMcacheInuse: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsMcacheSys: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsMspanInuse: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsMspanSys: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsNextGc: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsNumForcedGc: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsNumGc: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsOtherSys: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsPauseTotal: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsStackInuse: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsStackSys: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsSys: MetricConfig{
			Enabled: true,
		},
		ProcessRuntimeMemstatsTotalAlloc: MetricConfig{
			Enabled: false,
		},
	}
}

// MetricsBuilderConfig is a configuration for expvarreceiver metrics builder.
type MetricsBuilderConfig struct {
	Metrics MetricsConfig `mapstructure:"metrics"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics: DefaultMetricsConfig(),
	}
}
