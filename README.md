# IPMI Exporter for Prometheus

A Prometheus exporter for IPMI metrics that uses `ipmitool` to collect data from IPMI-enabled devices. The exporter provides metrics for sensor data (temperature, fan speeds, voltage) and System Event Log (SEL) entries.

## Metrics

- `ipmi_sensor_status` - Status of IPMI sensors (1 = OK, 0 = Critical)
- `ipmi_fan_speed_rpm` - Fan speed in RPM
- `ipmi_temperature_celsius` - Temperature in Celsius
- `ipmi_voltage` - Voltage readings
- `ipmi_sel_status` - Status of System Event Log entries (1 = OK, 0 = Critical)
- `ipmi_sel_total_entries` - Total number of SEL entries

## Installation

### Docker

```bash
docker run -d \
  --name ipmitool-exporter \
  --privileged \
  -p 9290:9290 \
  greatmedivack/ipmitool-exporter:latest
```

### Systemd 

Copy binary:
```bash
wget https://github.com/GreatMedivack/ipmitool-exporter/releases/download/v1.0.0/ipmitool-exporter-linux-amd64 -o ipmitool-exporter 
sudo cp ipmitool-exporter /usr/local/bin/
sudo chmod 755 /usr/local/bin/ipmitool-exporter
```
Install service:
```bash
sudo cp etc/systemd/system/ipmitool-exporter.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable ipmitool-exporter
sudo systemctl start ipmitool-exporter
```
Check service status:
```bash
sudo systemctl status ipmitool-exporter
```

### Usage

The exporter listens on :9290 by default. Metrics are available at /metrics endpoint:
```bash
curl http://localhost:9290/metrics
```
- Root/sudo access for IPMI operations
