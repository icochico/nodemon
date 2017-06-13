# NodeMon
A sensor data collector and fusion service for IoT
```
                   _
 _ __    ___    __| |  ___  _ __ ___    ___   _ __
| '_ \  / _ \  / _` | / _ \| '_ ` _ \  / _ \ | '_ \
| | | || (_) || (_| ||  __/| | | | | || (_) || | | |
|_| |_| \___/  \__,_| \___||_| |_| |_| \___/ |_| |_|
```

<b>NodeMon</b> is a sensor fusion service that collects, shapes and distributes data from local sensors.<br/>
NodeMon uses NATS (http://nats.io/) to distribute the aggregated data to potential clients along the network.

<b>Dependencies</b>

Go (>= 1.8) https://golang.org/dl/<br/>
gNATSd (>= 0.9.6) http://nats.io/download/nats-io/gnatsd/<br />
GNU make (suggested) https://www.gnu.org/software/make/<br/>

<b>Build</b>

With GNU make:

```
cd nodemon
make
```

Without make:

```
cd nodemon
go get ihmc.us/nodemon
go install ihmc.us/nodemon
```

<b>Run</b>

Run gNATSd (NATS Server):

Unzip in a directory and launch:
```
  gnatsd (starts the NATS server)
```

Run NodeMon:
```
Usage:
  nodemon [command]

Available Commands:
  help        Help about any command
  run         Run the NodeMon

Flags:
      --config string            config file (default is $HOME/.nodemon.yaml)
      --disservice-debug         Enable detailed debug log level for DisServiceSensor
      --disservice-port uint16   DisService statistics port (default 2400)
  -h, --help                     help for nodemon
      --http-debug               Enable detailed debug log level for HTTP Server
      --http-port uint16         HTTP server port (default 1323)
      --log-level string         Log level. Values: debug, info, warn, error, fatal, panic (default "debug")
      --mockets-debug            Enable detailed debug log level for MocketsSensor
      --mockets-port uint16      Mockets statistics port (default 1400)
      --nats-address string      NATS server address (default "localhost")
      --nats-debug               Enable detailed debug log level for NATS Server (default true)
      --nats-port uint16         NATS server port (default 4222)
      --netsensor-debug          Enable detailed debug log level for NetSensor
      --netsensor-port uint16    NetSensor statistics port (default 7777)
  -t, --toggle                   Help message for toggle
      --viper                    Use Viper for configuration (default true)

Use "nodemon [command] --help" for more information about a command.
```



