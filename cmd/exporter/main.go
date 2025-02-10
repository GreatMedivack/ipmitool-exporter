package main

import (
    "log"
    "net/http"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "ipmitool-exporter/internal/collector"  // Add this import
)

const (
    namespace = "ipmi"
    port     = ":9290"
)

func main() {
    // Update this line to use the correct package name
    ipmiCollector := collector.NewIPMICollector()
    prometheus.MustRegister(ipmiCollector)

    http.Handle("/metrics", promhttp.Handler())
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(`<html>
            <head><title>IPMI Exporter</title></head>
            <body>
            <h1>IPMI Exporter</h1>
            <p><a href="/metrics">Metrics</a></p>
            </body>
            </html>`))
    })

    log.Printf("Starting IPMI exporter on %s", port)
    log.Fatal(http.ListenAndServe(port, nil))
}