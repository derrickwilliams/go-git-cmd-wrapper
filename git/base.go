package git

import (
	"log"
	"os/exec"
	"strings"

	"github.com/ldez/go-git-cmd-wrapper/types"
)

var cmdExecutor = commander

// Init https://git-scm.com/docs/git-init
func Init(options ...types.Option) (string, error) {
	return command("init", options...)
}

// Push https://git-scm.com/docs/git-push
func Push(options ...types.Option) (string, error) {
	return command("push", options...)
}

// Pull https://git-scm.com/docs/git-pull
func Pull(options ...types.Option) (string, error) {
	return command("pull", options...)
}

// Clone https://git-scm.com/docs/git-clone
func Clone(options ...types.Option) (string, error) {
	return command("clone", options...)
}

// Remote https://git-scm.com/docs/git-remote
func Remote(options ...types.Option) (string, error) {
	return command("remote", options...)
}

// Fetch https://git-scm.com/docs/git-fetch
func Fetch(options ...types.Option) (string, error) {
	return command("fetch", options...)
}

// Rebase https://git-scm.com/docs/git-rebase
func Rebase(options ...types.Option) (string, error) {
	return command("rebase", options...)
}

// Checkout https://git-scm.com/docs/git-checkout
func Checkout(options ...types.Option) (string, error) {
	return command("checkout", options...)
}

// Config https://git-scm.com/docs/git-config
func Config(options ...types.Option) (string, error) {
	return command("config", options...)
}

// Branch https://git-scm.com/docs/git-branch
func Branch(options ...types.Option) (string, error) {
	return command("branch", options...)
}

// RevParse https://git-scm.com/docs/git-rev-parse
func RevParse(options ...types.Option) (string, error) {
	return command("rev-parse", options...)
}

// Reset https://git-scm.com/docs/git-reset
func Reset(options ...types.Option) (string, error) {
	return command("reset", options...)
}

// Commit https://git-scm.com/docs/git-commit
func Commit(options ...types.Option) (string, error) {
	return command("commit", options...)
}

// Add https://git-scm.com/docs/git-add
func Add(options ...types.Option) (string, error) {
	return command("add", options...)
}

// Merge https://git-scm.com/docs/git-merge
func Merge(options ...types.Option) (string, error) {
	return command("merge", options...)
}

// Raw use to execute arbitrary commands.
func Raw(cmd string, options ...types.Option) (string, error) {
	return command(cmd, options...)
}

// Debug display command line
func Debug(g *types.Cmd) {
	g.Debug = true
}

// Debugger display command line
func Debugger(debug bool) types.Option {
	return func(g *types.Cmd) {
		g.Debug = debug
	}
}

// Cond apply conditionally some options
func Cond(apply bool, options ...types.Option) types.Option {
	if apply {
		return func(g *types.Cmd) {
			g.ApplyOptions(options...)
		}
	}
	return NoOp
}

// NoOp do nothing.
func NoOp(g *types.Cmd) {}

func command(name string, options ...types.Option) (string, error) {
	g := &types.Cmd{Base: "git", Options: []string{name}}

	g.ApplyOptions(options...)

	return cmdExecutor(g.Base, g.Debug, g.Options...)
}

func commander(name string, debug bool, args ...string) (string, error) {
	if debug {
		log.Println(name, strings.Join(args, " "))
	}

	cmd := exec.Command(name, args...)

	output, err := cmd.CombinedOutput()

	return string(output), err
}
