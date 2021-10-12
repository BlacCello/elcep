package main

import (
	"github.com/MaibornWolff/elcep/main/config"
	"github.com/MaibornWolff/elcep/main/plugin"
	"github.com/olivere/elastic"
	"github.com/patrickmn/go-cache"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"time"
)

var bucketCache = cache.New(cache.NoExpiration, 60*time.Minute)

// TODO is the second parameter not the pluginOption? this should matter!!! i.e. the timekey coulkd be interesting for differenz plugins. as an index could be
// The factory method for the plugin
// noinspection GoUnusedExportedFunction
func NewPlugin(options config.Options, _ interface{}) plugin.Plugin {
	return &bucketAggregationPlugin{
		timeKey: options.TimeKey,
	}
}

type bucketAggregationPlugin struct {
	timeKey    string
	monitors   []*bucketAggregationMonitor
	collectors []prometheus.Collector
}

func (plugin *bucketAggregationPlugin) GetMonitors() []*bucketAggregationMonitor {
	return plugin.monitors
}

func (plugin *bucketAggregationPlugin) BuildMetrics(queries []config.Query) []prometheus.Collector {
	for _, query := range queries {
		log.Printf("Query loaded: %#v\n", query)
		monitor := NewAggregationMonitor(Create(query, plugin.timeKey))
		plugin.monitors = append(plugin.monitors, monitor)
		plugin.collectors = append(plugin.collectors, monitor.counter)
	}
	return plugin.collectors
}

func (plugin *bucketAggregationPlugin) Perform(elasticClient *elastic.Client) {
	for _, monitor := range plugin.monitors {
		monitor.Perform(elasticClient)
	}
}

func main() {}
