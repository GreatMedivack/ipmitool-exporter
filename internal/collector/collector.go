package collector

import (
    "log"
    "os/exec"
    "strconv"     // Add this import
    "strings"
    "github.com/prometheus/client_golang/prometheus"
)

type IPMICollector struct {
    sensorStatus *prometheus.GaugeVec
    fanSpeed     *prometheus.GaugeVec
    temperature  *prometheus.GaugeVec
    voltage      *prometheus.GaugeVec
    selStatus    *prometheus.GaugeVec
    selTotal     prometheus.Gauge
}

func NewIPMICollector() *IPMICollector {
    return &IPMICollector{
        sensorStatus: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Namespace: "ipmi",
                Name:      "sensor_status",
                Help:      "Status of IPMI sensors (1 = OK, 0 = Critical)",
            },
            []string{"sensor", "type"},
        ),
        fanSpeed: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Namespace: "ipmi",
                Name:      "fan_speed_rpm",
                Help:      "Fan speed in RPM",
            },
            []string{"fan"},
        ),
        temperature: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Namespace: "ipmi",
                Name:      "temperature_celsius",
                Help:      "Temperature in Celsius",
            },
            []string{"sensor"},
        ),
        voltage: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Namespace: "ipmi",
                Name:      "voltage",
                Help:      "Voltage readings",
            },
            []string{"sensor"},
        ),
        selStatus: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Namespace: "ipmi",
                Name:      "sel_status",
                Help:      "Status of System Event Log entries (1 = OK, 0 = Critical)",
            },
            []string{"event_type", "sensor"},
        ),
        selTotal: prometheus.NewGauge(
            prometheus.GaugeOpts{
                Namespace: "ipmi",
                Name:      "sel_total_entries",
                Help:      "Total number of SEL entries",
            },
        ),
    }
}

func (c *IPMICollector) Collect(ch chan<- prometheus.Metric) {
    c.collectSensorData()
    c.collectSELData()
    c.sensorStatus.Collect(ch)
    c.fanSpeed.Collect(ch)
    c.temperature.Collect(ch)
    c.voltage.Collect(ch)
    c.selStatus.Collect(ch)
    c.selTotal.Collect(ch)
}

func (c *IPMICollector) Describe(ch chan<- *prometheus.Desc) {
    c.sensorStatus.Describe(ch)
    c.fanSpeed.Describe(ch)
    c.temperature.Describe(ch)
    c.voltage.Describe(ch)
    c.selStatus.Describe(ch)
    c.selTotal.Describe(ch)
}

func (c *IPMICollector) collectSensorData() {
    out, err := exec.Command("ipmitool", "sdr").Output()
    if err != nil {
        log.Printf("Error executing ipmitool: %v", err)
        return
    }

    for _, line := range strings.Split(string(out), "\n") {
        if line == "" {
            continue
        }

        fields := strings.Split(line, "|")
        if len(fields) < 3 {
            continue
        }

        sensorName := strings.TrimSpace(fields[0])
        reading := strings.TrimSpace(fields[1])
        status := strings.TrimSpace(fields[2])

        switch {
        case strings.Contains(line, "Fan"):
            if speed, err := strconv.ParseFloat(strings.TrimRight(reading, " RPM"), 64); err == nil {
                c.fanSpeed.WithLabelValues(sensorName).Set(speed)
            }
            c.sensorStatus.WithLabelValues(sensorName, "fan").Set(statusToMetric(status))

        case strings.Contains(line, "Temp"):
            if temp, err := strconv.ParseFloat(strings.TrimRight(reading, " degrees C"), 64); err == nil {
                c.temperature.WithLabelValues(sensorName).Set(temp)
            }
            c.sensorStatus.WithLabelValues(sensorName, "temperature").Set(statusToMetric(status))

        case strings.Contains(line, "Volt"):
            if volt, err := strconv.ParseFloat(strings.TrimRight(reading, " Volts"), 64); err == nil {
                c.voltage.WithLabelValues(sensorName).Set(volt)
            }
            c.sensorStatus.WithLabelValues(sensorName, "voltage").Set(statusToMetric(status))
        }
    }
}

func statusToMetric(status string) float64 {
    status = strings.ToLower(strings.TrimSpace(status))
    if status == "ok" || status == "nominal" || status == "ns" {
        return 1
    }
    return 0
}

func (c *IPMICollector) collectSELData() {
    out, err := exec.Command("ipmitool", "sel", "list").Output()
    if err != nil {
        log.Printf("Error executing ipmitool sel list: %v", err)
        return
    }

    lines := strings.Split(string(out), "\n")
    c.selTotal.Set(float64(len(lines) - 1))
    c.selStatus.Reset()

    for _, line := range lines {
        if line == "" {
            continue
        }

        fields := strings.Split(line, "|")
        if len(fields) < 5 {
            continue
        }

        eventType := strings.TrimSpace(fields[3])
        state := strings.TrimSpace(fields[4])
        description := ""
        if len(fields) > 5 {
            description = strings.TrimSpace(fields[5])
        }

        value := 1.0
        if isCriticalCondition(eventType, description) && strings.Contains(strings.ToLower(state), "deasserted") {
            value = 0.0
        }
        c.selStatus.WithLabelValues(eventType, description).Set(value)
    }
}

func isCriticalCondition(eventType, description string) bool {
    eventLower := strings.ToLower(eventType)
    descLower := strings.ToLower(description)

    criticalConditions := []string{
        "power off/down",
        "critical",
        "failure",
        "error",
        "temperature out of range",
        "voltage out of range",
    }

    for _, condition := range criticalConditions {
        if strings.Contains(eventLower+" "+descLower, condition) {
            return true
        }
    }

    return false
}