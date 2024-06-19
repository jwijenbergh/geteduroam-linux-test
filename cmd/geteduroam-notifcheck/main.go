package main

import (
	"fmt"
	"os/exec"
	"time"

	"golang.org/x/exp/slog"

	"github.com/geteduroam/linux-app/internal/config"
	"github.com/geteduroam/linux-app/internal/log"
	"github.com/geteduroam/linux-app/internal/nm"
	"github.com/geteduroam/linux-app/internal/notification"
)

func sendnotif(notif string) error {
	_, err := exec.Command("notify-send", "geteduroam", notif).Output()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	log.Initialize("geteduroam-notifcheck", false)
	cfg, err := config.Load()
	if err != nil {
		slog.Error("no previous state", "error", err)
		return
	}
	con, err := nm.PreviousCon(cfg.UUID)
	if err != nil {
		slog.Error("no connection with uuid", "uuid", cfg.UUID, "error", err)
		return
	}
	if con == nil {
		slog.Error("connection is nil")
		return
	}

	if cfg.Validity == nil {
		slog.Info("validity is nil")
		return
	}
	valid := *cfg.Validity
	now := time.Now()
	diff := valid.Sub(now)
	days := int(diff.Hours() / 24)

	var text string
	if days > 10 {
		slog.Info("It is still more than 10 days", "days", days)
		return
	}
	if days < 0 {
		text = "connection is expired"
	}
	if days == 0 {
		text = "connection expires today"
	}
	if days > 0 {
		text = fmt.Sprintf("connection expires in %d days", days)
	}
	msg := fmt.Sprintf("Your eduroam %s. Re-run geteduroam to renew the connection", text)
	err = notification.Send(msg)
	if err != nil {
		slog.Error("failed to send notification", "error", err)
		return
	}
}
