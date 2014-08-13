# Dockbroker

> Peer-to-peer computing with Docker. Written in Go.
> The concept is explained [here](http://www.svenkreiss.com/blog/data-center-bazaar/).
> At this time, only a skeleton of the code exists.

[![GoDoc](https://godoc.org/github.com/svenkreiss/dockbroker?status.png)](https://godoc.org/github.com/svenkreiss/dockbroker)
[![Build Status](https://travis-ci.org/svenkreiss/dockbroker.png?branch=master)](https://travis-ci.org/svenkreiss/dockbroker)


## Summary

Clients ask dockbrokers for an "offer" for how much money they can execute a job and then pick the cheapest. Brokers that have data locally available or already cached part of the Docker image are cheaper and therefore preferred. Estimated time to completion will affect the price. Clients can build reputation in brokers when they deliver on time and brokers build reputation in clients when their job description and estimated run-time are good.


## License

Dockbroker was written by Sven Kreiss and made available under the [MIT license](https://github.com/svenkreiss/dockbroker/blob/master/LICENSE).
