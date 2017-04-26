package main

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"github.com/signifai/snap-plugin-processor-metric-repeat-filter/metric-repeat-filter"
)

func main() {
	plugin.StartProcessor(MetricRepeatFilter.MetricRepeatFilterPlugin{}, MetricRepeatFilter.Name, MetricRepeatFilter.Version)
}

