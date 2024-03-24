# Introduction

This module will fetch data from the `Australian government bureau of meteorology` and write it to a point.

## How to build (local)

```
go build -o module-contrib-demo
```

## How to build and run (docker)

### Pre-requisite

- Running BIOS:
    - https://nubeio.github.io/rubix-ce-docs/docs/rubix-ce/setup/docker/#install-and-run-the-rubix-stack
    - Use `ghcr.io/nubeio/rubix-bios-legacy:1.16.0` image for latest
- Install the latest ROS from rubix-ce (RCE)

### Steps to install this module

```bash
bash docker-build.bash
```

# See naming rules

- https://nubeio.github.io/rubix-ce-docs/docs/api-docs/modules/

# How to add APIs

- [api.go](pkg/router.go)
