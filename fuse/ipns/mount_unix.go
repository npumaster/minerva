//go:build (linux || darwin || freebsd || netbsd || openbsd) && !nofuse
// +build linux darwin freebsd netbsd openbsd
// +build !nofuse

package ipns

import (
	core "github.com/npumaster/minerva/core"
	coreapi "github.com/npumaster/minerva/core/coreapi"
	mount "github.com/npumaster/minerva/fuse/mount"
)

// Mount mounts ipns at a given location, and returns a mount.Mount instance.
func Mount(ipfs *core.IpfsNode, ipnsmp, ipfsmp string) (mount.Mount, error) {
	coreAPI, err := coreapi.NewCoreAPI(ipfs)
	if err != nil {
		return nil, err
	}

	cfg, err := ipfs.Repo.Config()
	if err != nil {
		return nil, err
	}

	allowOther := cfg.Mounts.FuseAllowOther

	fsys, err := NewFileSystem(ipfs.Context(), coreAPI, ipfsmp, ipnsmp)
	if err != nil {
		return nil, err
	}

	return mount.NewMount(ipfs.Process, fsys, ipnsmp, allowOther)
}
