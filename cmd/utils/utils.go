package utils

import (
	"os"
	"strings"
)

type DistroPackageManager int

func (dpm DistroPackageManager) String() string {
	switch dpm {
	case APT:
		return "APT"
	case DNF:
		return "DNF"
	case PACMAN:
		return "Pacman"
	default:
		return "Unknown distro package manager"
	}
}

const (
	UNKNOWN = iota
	APT
	DNF
	PACMAN
)

func DeterminePackageManager() (DistroPackageManager, error) {
	// Open /etc/os-release file and parse for ID_LIKE
	// Eventually also load ID and version more specifically
	osRelease, err := os.ReadFile("/etc/os-release")

	if err != nil {
		return UNKNOWN, err
	}

	var idLike string

	// Loop over all lines from os-release
	for line := range strings.SplitSeq(string(osRelease), "\n") {
		keyVal := strings.Split(line, "=")
		if len(keyVal) < 2 {
			continue
		}
		if keyVal[0] != "ID_LIKE" {
			continue
		}

		idLike = keyVal[1]
	}

	// log.Info(strings.Split(string(osRelease), "\n"))

	switch {
	case strings.Contains(idLike, "debian"), strings.Contains(idLike, "ubuntu"):
		return APT, nil
	case strings.Contains(idLike, "fedora"), strings.Contains(idLike, "rhel"), strings.Contains(idLike, "rhel"):
		return DNF, nil
	case strings.Contains(idLike, "arch"), strings.Contains(idLike, "manjaro"):
		return PACMAN, nil
	}

	return UNKNOWN, nil
}
