# Project remoteMediaControl

Remote Media controler, a simple app that extends your media keyboard buttons to mobile devices.
It exposes a simple js page with 3 buttons prev, play/pause, next. Then broadcasts the command as a keyboard button.
When run, shows you local IP and port (dependant on port env variable, default is 8080).
You should be available to see the app in your local area network devices (like other WiFi devices).

Should work on both Linux and Windows

## Getting Started

Clone the repository and run depending on your OS.
```bash
make run-lin
```
or
```bash
make run-win
```

This should build and run the app.

## Requirements

To build the Go 1.24 is required.