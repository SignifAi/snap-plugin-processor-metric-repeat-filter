/*
http://www.apache.org/licenses/LICENSE-2.0.txt
Copyright 2017 SignifAI Inc
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
