# Hell Let Loose Geo-Fences

HLL Geofence is a demo tool using the new RCONv2 protocol of Hell Let Loose to employ geofences for the participating players in a game session.
These geofences will ensure that the player does not get out of them in the same way the game would do when a player moves out of the map.
It will warn the player once going outside a defined fence and punish them if they're still out of the fence after a configurable number of seconds (10 by default).

## Usage

The tool can be built and started with a single simple command:

```bash
go run cmd/cmd.go
```

It requires that go in version 1.24 or higher to be installed on the system.

## Configuration

The tool requires a configuration file with credentials to the servers that should be observed and the fences.
While the config is auto-generated during the first start, you can create it prior to it by copying the `config.example.yml` file of this repository to `config.yml`.

# Showcase

A short showcase video of how the tool works in action can be found on YouTube:

[![Showcase video](http://img.youtube.com/vi/ETN9Y2ROR5s/0.jpg)](http://www.youtube.com/watch?v=ETN9Y2ROR5s "Hell Let Loose Geofences")
