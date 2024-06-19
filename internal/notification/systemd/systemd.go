package systemd

import (
	"os"
	"os/exec"

	"golang.org/x/exp/slog"
)

func hasSystemD() bool {
	if _, err := os.Stat("/run/systemd/system"); !os.IsNotExist(err) {
		return true
	}
	return false
}

func hasUnit(unit string) bool {
	_, err := exec.Command("systemctl", "--user", "list-unit-files", unit).Output()
	return err == nil
}

func hasUnitFiles() bool {
	if !hasUnit("geteduroam-notifs.service") {
		slog.Error("geteduroam-notifs.service is not installed anywhere")
		return false
	}
	if !hasUnit("geteduroam-notifs.timer") {
		slog.Error("geteduroam-notifs.timer is not installed anywhere")
		return false
	}
	return true
}

// HasDaemonSupport returns whether or not notifications can be enabled globally
// This depends on if systemd is used and if the unit is ready to be enabled
func HasDaemonSupport() bool {
	if !hasSystemD() {
		return false
	}
	if !hasUnitFiles() {
		return false
	}
	return true
}

// EnableDaemon enables the notification daemon using systemctl commands
func EnableDaemon() error {
	_, err := exec.Command("systemctl", "--user", "enable", "--now", "geteduroam-notifs.timer").Output()
	return err
}
