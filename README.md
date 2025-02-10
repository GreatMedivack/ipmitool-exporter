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
  --name ipmi-exporter \
  --privileged \
  -p 9290:9290 \
  greatmedivack/ipmi-exporter:latest
```

### Systemd 

Copy binary:
```bash 
sudo cp ipmitool-exporter /usr/local/bin/
sudo chmod 755 /usr/local/bin/ipmitool-exporter
```
Install service:
```bash
sudo cp etc/systemd/system/ipmi-exporter.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable ipmi-exporter
sudo systemctl start ipmi-exporter
```
Check service status:
```bash
sudo systemctl status ipmi-exporter
```

### Usage

The exporter listens on :9290 by default. Metrics are available at /metrics endpoint:
```bash
curl http://localhost:9290/metrics
```

### ## Trae AI
Used 1 Reference

Let me add the systemd installation instructions and complete the README:

```markdown
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
# Build image
docker build -t ipmi-exporter .

# Run container
docker run -d \
  --name ipmi-exporter \
  --privileged \
  -p 9290:9290 \
  greatmedivack/ipmi-exporter:latest
 ```
```

### Systemd
1. Build from source:
```bash
go build -o ipmitool-exporter cmd/main.go
 ```

2. Copy binary:
```bash
sudo cp ipmitool-exporter /usr/local/bin/
sudo chmod 755 /usr/local/bin/ipmitool-exporter
 ```
```

3. Install service:
```bash
sudo cp etc/systemd/system/ipmi-exporter.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable ipmi-exporter
sudo systemctl start ipmi-exporter
 ```
```

4. Check service status:
```bash
sudo systemctl status ipmi-exporter
 ```

## Usage
The exporter listens on :9290 by default. Metrics are available at /metrics endpoint:

```bash
# Check metrics
curl http://localhost:9290/metrics

# Check service logs
sudo journalctl -u ipmi-exporter -f
 ```

## Requirements
- ipmitool must be installed on the system
- Root/sudo access for IPMI operations
## Building from source
```bash
# Clone repository
git clone https://github.com/GreatMedivack/ipmitool-exporter.git
cd ipmitool-exporter

# Build binary
go build -o ipmitool-exporter cmd/main.go

# Run locally
sudo ./ipmitool-exporter
 ```
```

## GreatMedivack
Delete

Да ты заебал, я хочу чтобы инструкция по установке в systemd тоже была в файле README.md

## Trae AI
Used 1 Reference

Извините за путаницу. Давайте добавим инструкцию по установке через systemd в существующий README.md:

```markdown
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
  --name ipmi-exporter \
  --privileged \
  -p 9290:9290 \
  greatmedivack/ipmi-exporter:latest
 ```
```

### Systemd
1. Copy binary:
```bash
sudo cp ipmitool-exporter /usr/local/bin/
sudo chmod 755 /usr/local/bin/ipmitool-exporter
 ```
```

2. Install service:
```bash
sudo cp etc/systemd/system/ipmi-exporter.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable ipmi-exporter
sudo systemctl start ipmi-exporter
 ```
```

3. Check service status:
```bash
sudo systemctl status ipmi-exporter
 ```

## Usage
The exporter listens on :9290 by default. Metrics are available at /metrics endpoint:

```bash
curl http://localhost:9290/metrics
 ```

## Requirements
- ipmitool must be installed on the system
- Root/sudo access for IPMI operations