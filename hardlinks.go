package diffcopy

import (
	"os"

	"github.com/docker/containerd/fs"
)

type Hardlinks struct {
	seenFiles map[string]struct{}
}

func (v *Hardlinks) HandleChange(kind fs.ChangeKind, p string, fi os.FileInfo, err error) error {

	return nil
}
