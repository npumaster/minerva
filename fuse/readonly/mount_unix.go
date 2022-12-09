//go:build (linux || darwin || freebsd) && !nofuse
// +build linux darwin freebsd
// +build !nofuse

package readonly

import (
	core "github.com/npumaster/minerva/core"
	mount "github.com/npumaster/minerva/fuse/mount"
)

// Mount mounts IPFS at a given location, and returns a mount.Mount instance.
func Mount(ipfs *core.IpfsNode, mountpoint string) (mount.Mount, error) {
	cfg, err := ipfs.Repo.Config()
	if err != nil {
		return nil, err
	}
	allowOther := cfg.Mounts.FuseAllowOther
	fsys := NewFileSystem(ipfs)
	return mount.NewMount(ipfs.Process, fsys, mountpoint, allowOther)
}
