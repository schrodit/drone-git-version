package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "git-version plugin"
	app.Usage = "git-version plugin"
	app.Action = run
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "github_username",
			Usage:  "Username for github account",
			EnvVar: "PLUGIN_GITHUB_USERNAME,GITHUB_USERNAME",
		},
		cli.StringFlag{
			Name:   "github_password",
			Usage:  "password for github account",
			EnvVar: "PLUGIN_GITHUB_PASSWORD,GITHUB_PASSWORD",
		},
		cli.StringFlag{
			Name:   "git_name",
			Usage:  "Username for git config",
			EnvVar: "PLUGIN_GIT_NAME,GIT_NAME",
		},
		cli.StringFlag{
			Name:   "git_email",
			Usage:  "Email for git config",
			EnvVar: "PLUGIN_GIT_EMAIL,GIT_EMAIL",
		},
		cli.StringFlag{
			Name:   "file",
			Usage:  "Kubernetes helm release",
			EnvVar: "PLUGIN_FILE,FILE",
		},
		cli.StringFlag{
			Name:   "repo_owner",
			Usage:  "Current repository owner or org",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo_name",
			Usage:  "Current repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "branch",
			Usage:  "Current branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "deployment_type",
			Usage:  "Deployment type",
			EnvVar: "DRONE_DEPLOY_TO",
		},
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}
	plugin := Plugin{
		Config: Config{
			GitHubUsername: c.String("github_username"),
			GitHubPassword: c.String("github_password"),
			GitName:        c.String("git_name"),
			GitEmail:       c.String("git_email"),
			File:           c.String("file"),
			RepoOwner:      c.String("repo_owner"),
			RepoName:       c.String("repo_name"),
			Branch:         c.String("branch"),
			DeploymentType: c.String("deployment_type"),
		},
	}
	return plugin.Exec()
}
