# golang-http

This example golang application reads variables from the environment and based
on it displays a greeting.

# URLs
Not enough time was used for security/whitelisting. Thus the current endpoints
are accesible via nginx reverse proxy'ed locations that target each env:
- [dev/local](https://sebosp.com/golang-http/local/)
- [staging](https://sebosp.com/golang-http/staging/)
- [prod](https://sebosp.com/golang-http/prod/)

# Variables
The variables are stored in ConfigMaps that are exposed to the container
in the Deployment setup: `charts/golang-http/templates/deployment.yaml`
The ConfigMaps are provided by [sebosp/terraform](https://github.com/sebosp/terraform) in staging and production.
For preview and devpods the configmaps are provided by `charts/preview/templates/configmap.yaml`

## Running locally
If you need to run the container locally you can use it like this:
```shell
$ docker run -d -e PLAYER_NAME="seb" -e ENV_NAME=dev -e COLOR="red" -p 8080:8080 sebosp/golang-http:0.0.3
```

# Dev Workflow

## Background
Jenkins X runs on GKE and is used for CI/CD, it is setup via `jx init`.
Requires already working `helm` and `kubectl`.

## Initialization
`jx create quickstart` was used to initialize the repo.

## Monitoring pipeline
Either go to blue ocean via:
```bash
jx console`
```
Or follow it on the terminal:
```bash
jx get activity -f golang-http -w
```

## Pipeline
Following GitOps, the `Jenkinsfile` controls the pipeline setup allowing for
progressive and accountable pipeline evolution.
`Jenkinsfile` contains rules for different branch prefixes and the sequential
operations to run.

`jx env` allows you to select an env, for example
- staging
- prod
- sebosp-golang-http-feature-helloworld

## Dev environment setup
The dev setup is contained within [docker image tvl](https://github.com/sebosp/tvl)
Based on alpine it contains several vim utilities for coding Go such as:
vim-go
YouCompleteMe
UltiSnips

To mirror local code to <-> from kubernetes pods ksync is used:
- `jx created devpod --reuse --sync` (Creates a dev namespace for my user)
- `jx sync` (calls ksync and mirrors changes back and forth)

This allows running skaffold on a pod and allows debugging the pipeline.

## Developing

When a new branch is pushed, a preview env is starting for it.
a Makefile has targets for building/testing the Go code.
Currently staging is set to release automatically.
On merge to master, a versioned release is created on ChartMuseum.
If the pipeline succeeds on master, it is possible to promote to prod.
`jx promote golang-http --env production --version 0.0.3`

These promotions as PRs on the 
- [staging](https://github.com/sebosp/environment-dolphinchisel-staging/)
- [production](https://github.com/sebosp/environment-dolphinchisel-production/)

The ingress urls can be seen with:
```bash
$ jx get urls
golang-http http://golang-http.jx-production.235.138.176.118.nip.io
```

## Gotchas
- Branches map to namespaces. That means the branch name has to be considered
since a long branch name may create records too big for DNS.(63 chars rfc1035/rfc1123).
Thus, prefer concise short branch names to very long names.
- git-flow uses prefixes such as feature/, since branch names map to namespaces
this tends to not work well, the file `~/.git/config` can be modified to
change the prefixes. Also "underscore" as doesn't play too well (TODO: document why)o

# TODO
- Drop privileges to normal user
- Code coverage
- Testing Preview/Staging html.
- anchore for CVEs.
