package notification

import (
	"os/exec"

	"github.com/geteduroam/linux-app/internal/notification/systemd"
)

// Send sends a single notification with notify-send
func Send(msg string) error {
	_, err := exec.Command("notify-send", "geteduroam", msg).Output()
	return err
}

// HasDaemonSupport returns whether or not notifications can be enabled globally
func HasDaemonSupport() bool {
	// currently we only support systemd
	return systemd.HasDaemonSupport()
}

// EnableDaemon enables the notification using SystemD's user daemon
func EnableDameon() error {
	// currently we only support systemd
	return systemd.EnableDaemon()
}
