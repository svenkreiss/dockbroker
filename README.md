# Dockbroker

> Peer-to-peer computing with Docker. In Go.

[![GoDoc](https://godoc.org/github.com/svenkreiss/dockbroker?status.png)](https://godoc.org/github.com/svenkreiss/dockbroker)
[![Build Status](https://travis-ci.org/svenkreiss/dockbroker.png?branch=master)](https://travis-ci.org/svenkreiss/dockbroker)

Jobs ask dockbrokers for an "offer" for how much money they can execute the job and then pick the cheapest. Nodes that have data locally available or already cached part of the Docker image are cheaper and therefore preferred. Estimated time to completion will affect the price. Jobs can build reputation in dockbrokers when they deliver on time.
