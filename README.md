# Sensor Collector
> Executable Binary

## Build for Rasperry Pi
`$ make rasperrypi`

## Raspbian Configuration Systemd

1. Create a configuration file and edit it. This file will tell systemd which program needs to be executed:
`$ sudo nano /lib/systemd/system/sensor_collector.service`

    1.1. Add the following lines in the file:
```bash
[Unit]

Description=Sensor Collector

After=multi-user.target

[Service]

Type=idle

ExecStart=/home/sensor-collector/sensor_collector

[Install]

WantedBy=multi-user.target
```

2. Save and exit the nano file using Ctrl+x,Y and ENTER.

3. Change the permissions on the configuration file to 644:

`$ sudo chmod 644 /lib/systemd/system/sensor_collector.service`

4. Now all the tell the systemd to start the process on boot up :

`$ sudo systemctl daemon-reload`

`$ sudo systemctl enable myscript.service`

5. Now reboot your Pi and the process should run:

`$ sudo reboot`

