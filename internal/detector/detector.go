package detector

import "github.com/felihenrique/go-botdetector/pkg/iptree"

type Detector struct {
	tree *iptree.IpTree
}

func New() Detector {
	d := Detector{}
	d.tree = iptree.New(0)

	return d
}

func (d *Detector) IncludeRanges(ranges []string) {
	for _, v := range ranges {
		d.tree.Add(v)
	}
}

func (d *Detector) LoadDefaultRanges() {

}
