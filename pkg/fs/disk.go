package fs

import (
	"bufio"
	"os"
	"strings"
)

// IsMounted rchecks if a mount point is mounted
func IsMounted(mountPoint string) (bool, error) {
	mntF, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return false, err
	}

	scanner := bufio.NewScanner(mntF)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		//fspath should be 4th field
		if len(fields) >= 5 {
			if mountPoint == fields[4] {
				return true, nil
			}
		}
	}

	return false, nil
}
