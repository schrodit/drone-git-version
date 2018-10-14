package git

import (
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing/object"

	GIT "gopkg.in/src-d/go-git.v4"
)

type Git interface {
	Commit(versionFile string)
	Push(branch string)
}

type git struct {
	path     string
	repo     *GIT.Repository
	worktree *GIT.Worktree
	gitName  string
	gitEmail string
}

func New(path, name, email string) Git {
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

func (g *git) Push(branch string) {
	logrus.Infof("git push --set-upstream origin HEAD:%v", branch)
	err := g.repo.Push(&GIT.PushOptions{
		RefSpecs: []config.RefSpec{config.RefSpec(fmt.Sprintf("HEAD:%s", branch))},
	})
	if err != nil {
		logrus.Panicf("Cannot push to %v\n %v", branch, err)
	}
}

func (g *git) Commit(versionFile string) {
	logrus.Infof("git add %v", versionFile)
	_, err := g.worktree.Add(versionFile)
	if err != nil {
		logrus.Panicf("Cannot add %v\n %v", versionFile, err)
	}

	logrus.Infof("git commit -m \"[CI SKIP] ci upgrade to version $VERSION\" ")
	commit, err := g.worktree.Commit("[CI SKIP] ci upgrade to version $VERSION", &GIT.CommitOptions{
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
