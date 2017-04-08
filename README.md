# Octopus

Octopus is a command line utility written in Go. Octopus is a replacement for Octo.exe and mono on a nix system.

* [Running](#running)
* [Commands](#commands)


## Running

### Install
You can install Go and build Octopus.

### Flags
* --help - Returns help information for Octopus
* --debug - Turns on debug logging
* --server - Base server url for example https://octopus.builds.com
* --apikey - API Key to connect to the server
* --formvalues=KEY:VALUE - Form values that need passed in for a deployment. This flag can be passed in multiple times for multiple values

## Commands

### Promote Release
```
octopus --server={server} --apikey={apikey} promote-release {ProjectName} {FromEnvironment} {ToEnvironment}
```

### Get Project
```
octopus --server={server} --apikey={apikey} get-project {ProjectName}
```
