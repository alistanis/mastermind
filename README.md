# Mastermind

My company wouldn't let me use alternative configuration management, so I decided to roll my own.

Crazy, right?

# Design

Mastermind is designed to be simple. You define files, packages, commands, and services to run as part of a Role.

Each agent/node in your infrastructure may have multiple Roles. When you launch that agent on a machine, you specify what its Role(s) are, and Mastermind will handle the rest.

Mastermind runs over TLS, which requires you to provide a valid Certificate and Private Key in order to use it in anything that resembles a production environment. You may run the server/agent without TLS, but this is certainly not recommended.

Currently, if the Mastermind server were to fail (or if the server existed yet), you'd basically be screwed unless you configure the server to restart itself on failure. High availability is forthcoming as time allows. Clustering and leader election is planned.