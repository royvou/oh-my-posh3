package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type rubyArgs struct {
	hasRbenv       bool
	hasRvmprompt   bool
	hasChruby      bool
	hasAsdf        bool
	version        string
	hasRB          bool
	hasRakeFile    bool
	hasGemFile     bool
	displayVersion bool
}

func bootStrapRubyTest(args *rubyArgs) *ruby {
	env := new(MockedEnvironment)
	env.On("hasCommand", "rbenv").Return(args.hasRbenv)
	env.On("runCommand", "rbenv", []string{"version-name"}).Return(args.version, nil)
	env.On("hasCommand", "rvm-prompt").Return(args.hasRvmprompt)
	env.On("runCommand", "rvm-prompt", []string{"i", "v", "g"}).Return(args.version, nil)
	env.On("hasCommand", "chruby").Return(args.hasChruby)
	env.On("runCommand", "chruby", []string(nil)).Return(args.version, nil)
	env.On("hasCommand", "asdf").Return(args.hasAsdf)
	env.On("runCommand", "asdf", []string{"current", "ruby"}).Return(args.version, nil)
	env.On("hasFiles", "*.rb").Return(args.hasRB)
	env.On("hasFiles", "Rakefile").Return(args.hasRakeFile)
	env.On("hasFiles", "Gemfile").Return(args.hasGemFile)
	props := &properties{
		values: map[Property]interface{}{
			DisplayVersion: args.displayVersion,
		},
	}
	r := &ruby{
		env:   env,
		props: props,
	}
	return r
}

func TestRubyDisabled(t *testing.T) {
	args := &rubyArgs{}
	ruby := bootStrapRubyTest(args)
	assert.False(t, ruby.enabled(), "ruby is not enabled")
}

func TestRubyChrubyEnabled(t *testing.T) {
	args := &rubyArgs{
		hasRB:     true,
		hasChruby: true,
		version: ` * ruby-2.6.3
		ruby-1.9.3-p392
		jruby-1.7.0
		rubinius-2.0.0-rc1`,
	}
	ruby := bootStrapRubyTest(args)
	assert.True(t, ruby.enabled(), "ruby is enabled")
	assert.Equal(t, "ruby-2.6.3", ruby.version)
}

func TestRubyChrubyEnabledSecondLine(t *testing.T) {
	args := &rubyArgs{
		hasRB:     true,
		hasChruby: true,
		version: ` ruby-2.6.3
		* ruby-1.9.3-p392
		jruby-1.7.0
		rubinius-2.0.0-rc1`,
	}
	ruby := bootStrapRubyTest(args)
	assert.True(t, ruby.enabled(), "ruby is enabled")
	assert.Equal(t, "ruby-1.9.3-p392", ruby.version)
}
