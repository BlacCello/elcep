package plugin

import (
	"github.com/MaibornWolff/elcep/main/config"
	"io/ioutil"
	"log"
	"path/filepath"
	"plugin"
)

const factoryMethodName = "NewPlugin"

// this is simply a type alias. This is not a new type!
type factoryMethodType = func(config.Options, interface{}) Plugin

// Provider loads the plugin files and scans for available plugins
type Provider struct {
	Plugins map[string]factoryMethodType
}

// NewPluginProvider returns an instance with loaded Plugins from plugin Files
func NewPluginProvider(pluginFolder string, configuration *config.Configuration) *Provider {
	provider := &Provider{}
	files := findPlugins(pluginFolder)
	log.Printf("found plugins in folder: %v", files)
	activefiles := getActivatedPlugins(files, configuration)
	log.Printf("activated plugins: %v", activefiles)

	provider.initializePlugins(activefiles)
	return provider
}

//TODO write a test for this
func getActivatedPlugins(files []string, configuration *config.Configuration) []string {
	var activatedPlugins []string
	for _, file := range files {
		log.Printf("check if there is a pluginconfig for file %s ", file)
		// TODO its frustrating that we can only differ via thus methods.
		pluginConfig := configuration.ForPlugin(getLogicalPluginName(file))
		if nil != pluginConfig {
			log.Printf("there is a pluginconfig: %v ", pluginConfig)
			// TODO this just does not seem right... can we not just introduce a valid model? this drives me crazy and to disable plugin mechanism as a whole
			options := pluginConfig.Options.(map[interface{}]interface{})
			log.Printf("and we have even options: %v ", configuration)
			enabled := options["enabled"]
			if enabled.(bool) {
				activatedPlugins = append(activatedPlugins, file)
			}
		} else {
			log.Printf("sadly the pluginconfiguration is nil: %v ", pluginConfig)
		}
	}
	return activatedPlugins
}

// GetPluginNames returns a list of logical plugin names
func (provider *Provider) GetPluginNames() []string {
	keys := make([]string, 0, len(provider.Plugins))
	for k := range provider.Plugins {
		keys = append(keys, k)
	}
	return keys
}

func findPlugins(pluginFolder string) []string {
	var foundFileNames []string

	if files, err := ioutil.ReadDir(pluginFolder); err != nil {
		log.Fatal(err)
	} else {
		for _, f := range files {
			foundFileNames = append(foundFileNames, filepath.Join(pluginFolder, f.Name()))
		}
	}

	return foundFileNames
}

func (provider *Provider) initializePlugins(fileNames []string) {
	provider.Plugins = make(map[string]factoryMethodType)
	for _, file := range fileNames {
		plug, err := plugin.Open(file)
		if err != nil {
			log.Fatalf("%s: os.Open(): %s\n", file, err)
		}

		sym, err := plug.Lookup(factoryMethodName)
		if err != nil {
			log.Fatalf("%s: Could not find symbol '%s': %s\n", file, factoryMethodName, err)
		}

		m, ok := sym.(factoryMethodType)
		if !ok {
			var expected factoryMethodType
			log.Fatalf("%s: unexpected type from module symbol %s. Expected `%T`", file, factoryMethodName, expected)
		}

		pluginName := getLogicalPluginName(file)
		provider.Plugins[pluginName] = m

		log.Printf("Plugin detected: %s\n", pluginName)
	}
}

func getLogicalPluginName(file string) string {
	name := filepath.Base(file)
	// i.e. counter.so -> counter
	return name[0 : len(name)-3]
}
