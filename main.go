package main

import (
	// "encoding/json"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"

	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/ini.v1"
)

const (
	namespace = "marklogic"
)


func metricsHandler(w http.ResponseWriter, r *http.Request, status Status) {

	collectors := map[string]bool {
		"summary": 	true,
		"forests": 	true,
		"hosts": 	true,
		"servers": 	true,
	}

	up := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: namespace + "_up",
		Help: "Was the last query of MarkLogic successful?",
	})

	registry := prometheus.NewRegistry()
	registry.MustRegister(up)

	status.registry = registry
	status.Collect(collectors)

	up.Set(1)

	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}

func main() {
	var (
		listenAddress = kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry.").Default(":9307").String()
		metricsPath   = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Default("/metrics").String() 
		iniFile       = kingpin.Flag("config.ini", "Configuration file").Default("./marklogic.ini").String()
		uri           = kingpin.Flag("marklogic.uri", "HTTP API address of a MarkLogic server").Default("http://localhost:8002").String()
	)

	log.AddFlags(kingpin.CommandLine)
	kingpin.Version(version.Print(namespace + "_exporter"))
	kingpin.HelpFlag.Short('h')	
	kingpin.Parse()

    cfg, err := ini.InsensitiveLoad(*iniFile)
   	if err != nil {
   		log.Fatal("Issue when loading config from " + *iniFile)
   	}

   	status := Status{
   		user: cfg.Section("auth").Key("username").String(),
   		passwd: cfg.Section("auth").Key("password").String(),
   		uri: *uri,
   	}

	log.Infoln("Starting " + namespace + "_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	http.HandleFunc(*metricsPath, func(w http.ResponseWriter, r *http.Request) {
		metricsHandler(w, r, status)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
						 <head><title>MarkLogic Exporter</title></head>
						 <body>
						 <h1>MarkLogic Exporter</h1>
						 <p><a href='`+ *metricsPath + `'>Metrics</a></p>
						 </body>
						 </html>`))		
	})

	log.Infoln("Listening on", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}