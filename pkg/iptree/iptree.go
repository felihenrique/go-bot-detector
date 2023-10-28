package iptree

import "github.com/asergeyev/nradix"

type IpTree struct {
	ips *nradix.Tree
}

func New(length int) *IpTree {
	tree := &IpTree{}
	tree.ips = nradix.NewTree(length)

	return tree
}

func (it *IpTree) Add(ipRange string) (bool, error) {
	found, err := it.Has(ipRange)
	if err != nil || found {
		return false, err
	}

	if errAdd := it.ips.AddCIDR(ipRange, 0); errAdd != nil {
		return false, errAdd
	}

	return true, nil
}

func (it *IpTree) Has(ip string) (bool, error) {
	found, err := it.ips.FindCIDR(ip)

	if err != nil || found == nil {
		return false, err
	}

	return true, nil
}
