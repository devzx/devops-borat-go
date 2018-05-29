# devops-borat-go
Rewrite of https://github.com/devzx/devops-borat in [Golang](https://golang.org/) .  
A fun little app that posts a random [DevOps Borat](https://twitter.com/DEVOPS_BORAT) tweet to a given chat channel.

## Dependencies
```docker```
## Usage
``` sh
$ git clone git@github.com:devzx/devops-borat-go.git
$ cd devops-borat-go
```
Edit the Dockerfile and add a valid Slack Webhook or Discord Webhook. Requires at least one of the two in order to work.

The application by default will post to the channel twice a day at 8AM and 3PM. You can customise this to your liking by editing the cronjob which also resides within the Dockerfile.
```
$ docker build -t borat .
$ docker run -d --name borat borat
```
### Supports
- Slack
- Discord
