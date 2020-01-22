# semver
A command line tool for interacting with SemVer-compliant versions.

This CLI tool follows the [Semantic Versioning](https://semver.org) standard as closely as possible for making determinations regarding version structure and precedence.

# Installation
### Locally
```bash
go get github.com/particledecay/semver
```
### Use via Docker
```bash
docker run particledecay/semver:latest diff <version.1> <version.2>
```

# Usage
### Get detailed diff between versions
```bash
# v1.1.3 is 2 patch versions ahead
$ semver diff v1.1.1 v1.1.3
  MAJOR  MINOR  PATCH  PRERELEASE  BUILD
      0      0      2           0      0
```
```bash
# v1.1.1 is 2 patch versions behind
$ semver diff v1.1.3 v1.1.1
  MAJOR  MINOR  PATCH  PRERELEASE  BUILD
      0      0     -2           0      0
```
```bash
# v2.6.0 is 1 major version ahead and then 6 minor versions ahead from there
$ docker run particledecay/semver diff v1.0.0 v2.6.0
  MAJOR  MINOR  PATCH  PRERELEASE  BUILD
      1      6      0           0      0
```
