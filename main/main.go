package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/MaibornWolff/elcep/main/config"
	"github.com/MaibornWolff/elcep/main/plugin"
	"github.com/olivere/elastic"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	configuration := config.ReadConfig()
	executor := initExecutor(&configuration)

	go executor.RunPlugins(configuration.Options.Freq)

	http.Handle(configuration.Options.Path, promhttp.Handler())
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(configuration.Options.Port), nil))
}

func initExecutor(configuration *config.Configuration) *plugin.Executor {
	pluginProvider := plugin.NewPluginProvider(configuration.Options.PluginDir, configuration)

	// TODO should be configurable - false for opensearch right now
	client, err := elastic.NewClient(elastic.SetHealthcheckTimeoutStartup(30*time.Second),
		elastic.SetURL(configuration.Options.ElasticsearchURL.String()),
		elastic.SetHealthcheck(false),
		elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}
	executor := &plugin.Executor{
		ElasticClient: client,
	}

	for name, newMon := range pluginProvider.Plugins {
		//TODO this is just wrong, its not even used in counter and bucket
		conf := configuration.ForPlugin(name)
		if conf == nil {
			log.Fatalf("Missing config for plugin %s\n", name)
		}
		executor.BuildPlugins(*configuration, *conf, newMon)
	}

	return executor
}
