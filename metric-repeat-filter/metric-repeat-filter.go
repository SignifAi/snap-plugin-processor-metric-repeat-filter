package MetricRepeatFilter

import (
	"strings"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

const (
	Name    = "metric-repeat-filter"
	Version = 1
)

type MetricRepeatFilterPlugin struct {
	currentMaps map[string]interface{}
}

func New() MetricRepeatFilterPlugin {
	var mrfp MetricRepeatFilterPlugin
	mrfp.currentMaps = make(map[string]interface{})
	return mrfp
}

func (MetricRepeatFilterPlugin) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()

	return *policy, nil
}

func (mrf MetricRepeatFilterPlugin) Process(metrics []plugin.Metric, cfg plugin.Config) ([]plugin.Metric, error) {
	newMetrics := make([]plugin.Metric, 0)
	for _, metric := range metrics {
		metricName := strings.Join(metric.Namespace.Strings(), ".")
		oldValue, present := mrf.currentMaps[metricName]
		if present && oldValue == metric.Data {
			continue
		} else {
			mrf.currentMaps[metricName] = metric.Data
			newMetrics = append(newMetrics, metric)
		}
	}

	return newMetrics, nil
}
