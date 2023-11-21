# Introduction

This module will fetch data from the `Australian government bureau of meteorology` and write it to a point.

## How to build

#### Download private repo dependencies

```
export GITHUB_TOKEN=<YOUR_GITHUB_TOKEN>
git config --global url."https://$GITHUB_TOKEN:x-oauth-basic@github.com/NubeIO".insteadOf "https://github.com/NubeIO"
```

#### Download private repo dependencies if above commands doesn't work

```
export GITHUB_TOKEN=<YOUR_GITHUB_TOKEN>
export GOPRIVATE=github.com/NubeIO
go get -v
```

#### If both of them doesn't work

- Check `~/.gitconfig` file & remove lines with `git config --global url."https://*` & then retry again

```
go build -o module-contrib-demo
```

## How to build and run (docker)

### Pre-requisite

- Running BIOS:
    - https://nubeio.github.io/rubix-ce-docs/docs/rubix-ce/setup/docker/#install-and-run-the-rubix-stack
- Install the latest ROS from rubix-ce (RCE)

```
bash docker-build.bash <GITHUB_TOKEN>
```

## How to build and run (local)

### Prerequisite

- Folder `<YOUR_ROS_PATH>` needs to have `app-amd64`

To build and run rubix-os you can use the bash script.

```
bash build.bash <YOUR_ROS_PATH>
```

Example

```
bash build.bash ~/nube/rubix-os
```

### See naming rules

- https://nubeio.github.io/rubix-ce-docs/docs/api-docs/modules/

### How to add APIs

- [api.go](pkg/api.go)
