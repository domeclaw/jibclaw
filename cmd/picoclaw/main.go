// JIBClaw - Ultra-lightweight personal AI agent
// Inspired by and based on nanobot: https://github.com/HKUDS/nanobot
// License: MIT
//
// Copyright (c) 2026 JIBClaw contributors

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/sipeed/picoclaw/cmd/picoclaw/internal"
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/agent"
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/auth"
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/cron"
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/gateway"
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/migrate"
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/model"
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/onboard"
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/skills"
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/status"
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/version"
	"github.com/sipeed/picoclaw/pkg/config"
)

func NewJibclawCommand() *cobra.Command {
	short := fmt.Sprintf("%s jibclaw - Personal AI Assistant v%s\n\n", internal.Logo, config.GetVersion())

	cmd := &cobra.Command{
		Use:     "jibclaw",
		Short:   short,
		Example: "jibclaw version",
	}

	cmd.AddCommand(
		onboard.NewOnboardCommand(),
		agent.NewAgentCommand(),
		auth.NewAuthCommand(),
		gateway.NewGatewayCommand(),
		status.NewStatusCommand(),
		cron.NewCronCommand(),
		migrate.NewMigrateCommand(),
		skills.NewSkillsCommand(),
		model.NewModelCommand(),
		version.NewVersionCommand(),
	)

	return cmd
}

const (
	colorBlue = "\033[1;38;2;62;93;185m"
	colorRed  = "\033[1;38;2;213;70;70m"
	banner    = "\r\n" +
		colorBlue + "     ██╗██╗██████╗  ██████╗" + colorRed + "██╗      █████╗ ██╗    ██╗\n" +
		colorBlue + "     ██║██║██╔══██╗██╔════╝" + colorRed + "██║     ██╔══██╗██║    ██║\n" +
		colorBlue + "     ██║██║██████╔╝██║     " + colorRed + "██║     ███████║██║ █╗ ██║\n" +
		colorBlue + "██   ██║██║██╔══██╗██║     " + colorRed + "██║     ██╔══██║██║███╗██║\n" +
		colorBlue + "╚█████╔╝██║██████╔╝╚██████╗" + colorRed + "███████╗██║  ██║╚███╔███╔╝\n" +
		colorBlue + " ╚════╝ ╚═╝╚═════╝  ╚═════╝" + colorRed + "╚══════╝╚═╝  ╚═╝ ╚══╝╚══╝\n " +
		"\033[0m\r\n"
)

func main() {
	fmt.Printf("%s", banner)
	cmd := NewJibclawCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
