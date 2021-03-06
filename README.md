# Mountain Backup
File backup tool.

# Table Of Contents
- [Overview](#overview)
- [Configure](#configure)
    - [Upload Configuration](#upload-configuration)
    - [Metrics Configuration](#metrics-configuration)
    - [Backup Configuration](#backup-configuration)
        - [Files](#files)
		- [Prometheus](#prometheus)
- [Develop](#develop)

# Overview
Creates a GZip-ed tar ball and uploads it to a S3 compatible object 
storage service.

**Why "Mountain Backup"?**  
A reference to 
["Steel Mountain"](https://mrrobot.fandom.com/wiki/Steel_Mountain) in the show
Mr. Robot.  

# Configure
The tool's behavior is specified in a TOML configuration file.  

## Upload Configuration
The `Upload` section of the file defines where backups will be stored.  

Configuration:

```toml
[Upload]
# Storage API host
Endpoint = "..."

# Storage API Key ID
KeyID = "..."

# Storage API Secret Access Key
SecretAccessKey = "..."

# Name of bucket in which to upload backup
Bucket = "backups"

# (Optional) Backup name format without file extension, can use strftime symbols
# Defaults to value below
Format = "backup-%Y-%m-%d-%H:%M:%S"
```

## Metrics Configuration
The tool can push metrics to Prometheus about the backup process. To push metrics 
[Prometheus Push Gateway](https://github.com/prometheus/pushgateway) must be accessible to mountain backup.

You can disable metrics by setting `Metrics.Enabled = false`.

Metrics:

- `backup_success`
    - Indicates if backup succeeded
- `backup_number_files`
    - Number of files backed up

Configuration:

```toml
[Metrics]
# (Optional) Set to false if you do not wish to publish Prometheus metrics.
# Defaults to false
Enabled = true

# (Optional) Host which Prometheus Push Gateway can be accessed. Must 
# include scheme
# Defaults to value below
PushGatewayHost = "http://localhost:9091"

# (Optional) Value of `host` label in metrics
# Defaults to value below
LabelHost = "mountain-backup"
```

## Backup Configuration
The mountain backup tools provides different modules to handle unique 
backup scenarios. These modules are configured by creating sub sections in the 
configuration file under the modules name's.  

The names of these sub-sections do not matter. The only constraint is that they 
should be unique in that section.  

For example to configure a module named `ExampleModule` one could create a 
configuration section named `ExampleModule.Foo` or `ExampleModule.Bar`.

### Files
The `Files` module backs up normal files.  

All configuration parameters can include shell globs.

Configuration:

```toml
[Files.XXXXX]
# List of files / directories to backup
Files = [ "..." ]

# Files / directories to exclude from backup
Exclude = [ "..." ]
```

### Prometheus
The `Prometheus` module makes a snapshot of a Prometheus database via the 
admin API and backs it up.  

The snapshot files (`${DataDirectory}/data/snapshots/xxxx`) will be backed up 
as if they were located in the main data directory (`${DataDirectory}/data`).

This way Prometheus will use the data from the backed up snapshot. Instead of 
simply placing the snapshot files in the snapshot directory but starting with 
an empty database.

Configuration:

```toml
[Prometheus.XXXXX]
# Admin API host. Must include scheme
AdminAPIHost = "http://localhost:9090"

# Directory in which Prometheus data is stored
DataDirectory = "/var/lib/prometheus"
```

# Develop
To develop mountain backup you must run Prometheus and Prometheus Push Gateway locally.  

The `dev` make target starts both Prometheus and Prometheus Push Gateway using Docker.
