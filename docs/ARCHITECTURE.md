# Architecture

## Features

_No expertise_ infra management for small to medium size teams which

1. Dev workflow story ⛵
   1. Devcontainer integration on local
   2. local joins a namespace's network to access complete dev environment
   3. access namespaces services / webpages from local
   4. AI agents can access complex development environment ⛵ 🤖
   5. Provide context to AI coding agent about the environment ⛵ 🤖
2. Easy to use and opinionated. As a result many advance features will not be supported
3. Management Console Support out of the box ⛵
4. Out of the box ⛵
   1. tracing 🌟
   2. log management 🌟
   3. monitoring - VictoriaMetrics 🌟
   4. backup and recovery 🌟
   5. operator installation management like helm chart CLI 🌟
   6. Traffik, SSL and domain management using cloudflare
   7. simple persistent volume support
   8. Cloud provider support 🌟
      1. Hetzner
      2. AWS
      3. GCP
      4. Azure
      5. Digitalocean
5. Single binary which auto scaffolds env using nix
   1. install dependencies
6. Declarative
7. Operator support
8. Clustering support
   1. Equal nodes and automatic leader election
   2. Hashicorp RAFT
   3. bbolt for storage
   4. snapshoting of architecture
   5. resurrect from snapshot
   6. Sparse reconciliation cycles
9. Stack support - higher level constructs to describe the whole architecture in terms of lower level constructs like services ⛵
