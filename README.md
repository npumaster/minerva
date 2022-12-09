
![kubo, an IPFS node in Go](https://ipfs.io/ipfs/bafykbzacecaesuqmivkauix25v6i6xxxsvsrtxknhgb5zak3xxsg2nb4dhs2u/ipfs.go.png)

## What is Minerva?

Minerva was the IPFS implementation and is the most widely used one today. Implementing the *Interplanetary Filesystem* - the Web3 standard for content-addressing, interoperable with HTTP. Thus powered by IPLD's data models and the libp2p for network communication. Kubo is written in Go.

Featureset
- Runs an IPFS-Node as a network service
- [Command Line Interface](https://docs.ipfs.tech/reference/kubo/cli/) to IPFS-Nodes
- Local [Web2-to-Web3 HTTP Gateway functionality](https://github.com/ipfs/specs/tree/main/http-gateways#readme) 
- HTTP RPC API (`/api/v0`) to access and control the daemon
- IPFS's internal Webgui can be used to manage the Kubo nodes

## What is IPFS?

IPFS is a global, versioned, peer-to-peer filesystem. It combines good ideas from previous systems such as Git, BitTorrent, Kademlia, SFS, and the Web. It is like a single BitTorrent swarm, exchanging git objects. IPFS provides an interface as simple as the HTTP web, but with permanence built-in. You can also mount the world at /ipfs.

For more info see: https://docs.ipfs.tech/concepts/what-is-ipfs/
