package bench

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/chrootarchive"
	"github.com/docker/docker/pkg/reexec"
)

func init() {
	reexec.Init()
}

func benchmarkInitialCopy(b *testing.B, fn func(string, string) error, size int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		tmpdir, err := createTestDir(size)
		if err != nil {
			b.Error(err)
		}
		destdir, err := ioutil.TempDir(os.Getenv("DIRRSYNC_BASE_DIR"), "destdir")
		if err != nil {
			os.RemoveAll(tmpdir)
			b.Error(err)
		}
		b.StartTimer()
		err = fn(tmpdir, destdir)
		os.RemoveAll(tmpdir)
		os.RemoveAll(destdir)
		if err != nil {
			b.Error(err)
		}
	}
}

func copyWithTar(src, dest string) error {
	return archive.CopyWithTar(src, dest)
}

func chrootCopyWithTar(src, dest string) error {
	return chrootarchive.CopyWithTar(src, dest)
}

func cpa(src, dest string) error {
	cmd := exec.Command("cp", "-a", src, dest)
	return cmd.Run()
}

func rsync(src, dest string) error {
	cmd := exec.Command("rsync", "-a", "--del", src, dest)
	return cmd.Run()
}

func gnutar(src, dest string) error {
	tar := exec.Command("tar", "-cf", "-", src)
	unpack := exec.Command("tar", "xf", "-", "-C", dest)
	stdout, err := tar.StdoutPipe()
	if err != nil {
		return err
	}
	unpack.Stdin = stdout
	go tar.Run()
	return unpack.Run()
}

func BenchmarkCopyWithTar10(b *testing.B) {
	benchmarkInitialCopy(b, copyWithTar, 10)
}

func BenchmarkCopyWithTar50(b *testing.B) {
	benchmarkInitialCopy(b, copyWithTar, 50)
}

func BenchmarkCopyWithTar200(b *testing.B) {
	benchmarkInitialCopy(b, copyWithTar, 200)
}

func BenchmarkCopyWithTar1000(b *testing.B) {
	benchmarkInitialCopy(b, copyWithTar, 1000)
}

func BenchmarkChrootCopyWithTar10(b *testing.B) {
	benchmarkInitialCopy(b, chrootCopyWithTar, 10)
}

func BenchmarkChrootCopyWithTar50(b *testing.B) {
	benchmarkInitialCopy(b, chrootCopyWithTar, 50)
}

func BenchmarkChrootCopyWithTar200(b *testing.B) {
	benchmarkInitialCopy(b, chrootCopyWithTar, 200)
}

func BenchmarkChrootCopyWithTar1000(b *testing.B) {
	benchmarkInitialCopy(b, chrootCopyWithTar, 1000)
}

func BenchmarkCPA10(b *testing.B) {
	benchmarkInitialCopy(b, cpa, 10)
}

func BenchmarkCPA50(b *testing.B) {
	benchmarkInitialCopy(b, cpa, 50)
}

func BenchmarkCPA200(b *testing.B) {
	benchmarkInitialCopy(b, cpa, 200)
}

func BenchmarkCPA1000(b *testing.B) {
	benchmarkInitialCopy(b, cpa, 1000)
}

func BenchmarkDiffCopy10(b *testing.B) {
	benchmarkInitialCopy(b, diffCopyReg, 10)
}

func BenchmarkDiffCopy50(b *testing.B) {
	benchmarkInitialCopy(b, diffCopyReg, 50)
}

func BenchmarkDiffCopy200(b *testing.B) {
	benchmarkInitialCopy(b, diffCopyReg, 200)
}

func BenchmarkDiffCopy1000(b *testing.B) {
	benchmarkInitialCopy(b, diffCopyReg, 1000)
}

func BenchmarkDiffCopyProto10(b *testing.B) {
	benchmarkInitialCopy(b, diffCopyProto, 10)
}

func BenchmarkDiffCopyProto50(b *testing.B) {
	benchmarkInitialCopy(b, diffCopyProto, 50)
}

func BenchmarkDiffCopyProto200(b *testing.B) {
	benchmarkInitialCopy(b, diffCopyProto, 200)
}

func BenchmarkDiffCopyProto1000(b *testing.B) {
	benchmarkInitialCopy(b, diffCopyProto, 1000)
}

func BenchmarkRsync10(b *testing.B) {
	benchmarkInitialCopy(b, rsync, 10)
}

func BenchmarkRsync50(b *testing.B) {
	benchmarkInitialCopy(b, rsync, 50)
}

func BenchmarkRsync200(b *testing.B) {
	benchmarkInitialCopy(b, rsync, 200)
}

func BenchmarkRsync1000(b *testing.B) {
	benchmarkInitialCopy(b, rsync, 1000)
}

func BenchmarkGnuTar10(b *testing.B) {
	benchmarkInitialCopy(b, gnutar, 10)
}

func BenchmarkGnuTar50(b *testing.B) {
	benchmarkInitialCopy(b, gnutar, 50)
}

func BenchmarkGnuTar200(b *testing.B) {
	benchmarkInitialCopy(b, gnutar, 200)
}

func BenchmarkGnuTar1000(b *testing.B) {
	benchmarkInitialCopy(b, gnutar, 1000)
}
