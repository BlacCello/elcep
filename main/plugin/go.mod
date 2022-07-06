module github.com/MaibornWolff/elcep/main/plugin

go 1.17

require (
	github.com/MaibornWolff/elcep/main/config v1.2.0
	github.com/golang/mock v1.6.0
	github.com/olivere/elastic v6.2.19+incompatible
	github.com/prometheus/client_golang v1.0.0
	gopkg.in/go-playground/assert.v1 v1.2.1
)

require (
	github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc // indirect
	github.com/alecthomas/units v0.0.0-20151022065526-2efee857e7cf // indirect
	github.com/beorn7/perks v1.0.0 // indirect
	github.com/fortytw2/leaktest v1.3.0 // indirect
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/mailru/easyjson v0.0.0-20190614124828-94de47d64c63 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/pkg/errors v0.8.0 // indirect
	github.com/prometheus/client_model v0.0.0-20190129233127-fd36f4220a90 // indirect
	github.com/prometheus/common v0.4.1 // indirect
	github.com/prometheus/procfs v0.0.2 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/MaibornWolff/elcep/main/config => ../config
