# Dockbroker

> Peer-to-peer computing with Docker. In Go.

Jobs ask dockbrokers for an "offer" for how much money they can execute the job and then pick the cheapest. Nodes that have data locally available or already cached part of the Docker image are cheaper and therefore preferred. Estimated time to completion will affect the price. Jobs can build reputation in dockbrokers when they deliver on time.
