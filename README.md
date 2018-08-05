# golang-http           

This example golang application reads variables from the environment and based
on it displays a greeting.

# Variables
The variables are stored in a ConfigMaps that are exposed to the container using a Deployment.
The ConfigMaps are provided as terraform resources in staging and productien.
For preview and devpods the configmaps are provided by Charts.

## Running locally
If you need to run the container locally you can use it like this:
```shell
$ docker run -d -e PLAYER_NAME="seb" -e ENV_NAME=dev -e COLOR="red" -p 8080:8080 golang-http:0.0.1
```

