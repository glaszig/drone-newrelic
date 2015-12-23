# drone-newrelic

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-newrelic/status.svg)](http://beta.drone.io/drone-plugins/drone-newrelic)
[![](https://badge.imagelayers.io/plugins/drone-newrelic:latest.svg)](https://imagelayers.io/?images=plugins/drone-newrelic:latest 'Get your own badge on imagelayers.io')

Drone plugin for deployment via New Relic.

## Usage

```sh
./drone-newrelic <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone-plugins/drone-newrelic",
        "full_name": "drone/drone",
        "owner": "drone",
        "name": "drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/home/user/golang/src/github.com/drone-plugins/drone-newrelic"
    },
    "vargs": {
        "tasks": "production deploy"
    }
}
EOF
```

## Docker

Build the Docker container using `make`:

```sh
make deps build docker
```

### Example

```sh
docker run -i plugins/drone-newrelic <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone-plugins/drone-newrelic",
        "full_name": "drone/drone",
        "owner": "drone",
        "name": "drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/home/user/golang/src/github.com/drone-plugins/drone-newrelic"
    },
    "vargs": {
        "tasks": "production deploy"
    }
}
EOF
```
