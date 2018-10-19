package git

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"

	"github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4/plumbing/object"

	GIT "gopkg.in/src-d/go-git.v4"
)

type Git interface {
	Commit(versionFile, version string) string
	Push(username, password, commit, branch string)
}

type git struct {
	path     string
	repo     *GIT.Repository
	worktree *GIT.Worktree
	gitName  string
	gitEmail string
}

func New(path, name, email string) Git {
	path, err := os.Getwd()
	if err != nil {
		logrus.Error(err.Error())
	}
	logrus.Infof("Set git path %s", path)
	g := git{
		path:     path,
		repo:     nil,
		worktree: nil,
		gitName:  name,
		gitEmail: email,
	}
	g.openRepo()
	g.setWorktree()
	return &g
}

func (g *git) Push(username, password, commit, branch string) {
	logrus.Infof("git push --set-upstream origin %s:%v", commit, branch)
	err := g.repo.Push(&GIT.PushOptions{
		//RefSpecs: []config.RefSpec{config.RefSpec(fmt.Sprintf("%s:%s", commit, branch))},
		Auth: &http.BasicAuth{
			username,
			password,
		},
	})
	if err != nil {
		logrus.Panicf("Cannot push to %v\n %v", branch, err)
	}
}

func (g *git) Commit(versionFile, version string) string {
	logrus.Infof("git status", versionFile)
	status, err := g.worktree.Status()
	if err != nil {
		logrus.Errorf("Cannot get status of current git repo\n %s", err.Error())
	}
	logrus.Infoln(status.String())

	logrus.Infof("git add %v", versionFile)
	_, err = g.worktree.Add(versionFile)
	if err != nil {
		logrus.Panicf("Cannot add %v\n %v", versionFile, err)
	}

	logrus.Infof("git commit -m \"[CI SKIP] ci upgrade to version %s", version)
	commit, err := g.worktree.Commit(fmt.Sprintf("[CI SKIP] ci upgrade to version %s", version), &GIT.CommitOptions{
		Author: &object.Signature{
			Name:  g.gitName,
			Email: g.gitEmail,
			When:  time.Now(),
		},
	})
	if err != nil {
		logrus.Panicf("Cannot commit %v\n %v", versionFile, err)
	}

	logrus.Infof("git show -s")
	obj, err := g.repo.CommitObject(commit)
	if err != nil {
		logrus.Panicf("Cannot show commit\n %v", versionFile, err)
	}
	logrus.Info(obj)
	return obj.Hash.String()
}

func (g *git) openRepo() {
	r, err := GIT.PlainOpen(g.path)
	if err != nil {
		log.Panicf("Cannot open git repository\n %v", err)
	}
	g.repo = r
}

func (g *git) setWorktree() {
	w, err := g.repo.Worktree()
	if err != nil {
		log.Panicf("Cannot set git worktree\n %v", err)
	}
	g.worktree = w
}
