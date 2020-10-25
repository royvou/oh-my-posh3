package main

import "regexp"

type ruby struct {
	props   *properties
	env     environmentInfo
	version string
}

func (r *ruby) string() string {
	if r.props.getBool(DisplayVersion, true) {
		return r.version
	}
	return ""
}

func (r *ruby) init(props *properties, env environmentInfo) {
	r.props = props
	r.env = env
}

func (r *ruby) enabled() bool {
	if !r.env.hasFiles("*.rb") && !r.env.hasFiles("Rakefile") && !r.env.hasFiles("Gemfile") {
		return false
	}
	if r.env.hasCommand("rbenv") {
		r.version, _ = r.env.runCommand("rbenv", "version-name")
	} else if r.env.hasCommand("rvm-prompt") {
		r.version, _ = r.env.runCommand("rvm-prompt", "i", "v", "g")
	} else if r.env.hasCommand("chruby") {
		version, _ := r.env.runCommand("chruby")
		regex := regexp.MustCompile(`\* (?P<version>.*)\n`)
		match := groupDict(regex, version)
		r.version = match["version"]
	} else if r.env.hasCommand("asdf") {
		version, _ := r.env.runCommand("asdf", "current", "ruby")
		//TODO: strip this
		r.version = version
	} else {
		r.version = "system"
	}
	return true
}
