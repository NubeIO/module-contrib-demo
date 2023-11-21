# Module example to show the use case of rubix-os (ROS)

This module will fetch data from the `Australian government bureau of meteorology` and write it to a point.

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
