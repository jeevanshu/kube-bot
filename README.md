# kube-bot

A Discord bot for interacting with your Kubernetes Cluster and Objects.

Invite `kube-bot` to your server to help you manage your K8s cluster without requiring access to commandline.

It will listen for messages that prefix `!k`. 

## Commands

For a list of commands use `!k help` in your discord channel.

`kube-bot` can be used to list, update, delete, scale Kubernetes objects.

## Build

### Make it yours

To launch your own version of `kube-bot`, make use of the `Dockerfile` in the repo and
pass the variables mentioned in `.env.sample` file.