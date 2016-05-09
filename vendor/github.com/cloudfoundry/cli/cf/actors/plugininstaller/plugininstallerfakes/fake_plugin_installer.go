// This file was generated by counterfeiter
package plugininstallerfakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/actors/plugininstaller"
)

type FakePluginInstaller struct {
	InstallStub        func(inputSourceFilepath string) string
	installMutex       sync.RWMutex
	installArgsForCall []struct {
		inputSourceFilepath string
	}
	installReturns struct {
		result1 string
	}
}

func (fake *FakePluginInstaller) Install(inputSourceFilepath string) string {
	fake.installMutex.Lock()
	fake.installArgsForCall = append(fake.installArgsForCall, struct {
		inputSourceFilepath string
	}{inputSourceFilepath})
	fake.installMutex.Unlock()
	if fake.InstallStub != nil {
		return fake.InstallStub(inputSourceFilepath)
	} else {
		return fake.installReturns.result1
	}
}

func (fake *FakePluginInstaller) InstallCallCount() int {
	fake.installMutex.RLock()
	defer fake.installMutex.RUnlock()
	return len(fake.installArgsForCall)
}

func (fake *FakePluginInstaller) InstallArgsForCall(i int) string {
	fake.installMutex.RLock()
	defer fake.installMutex.RUnlock()
	return fake.installArgsForCall[i].inputSourceFilepath
}

func (fake *FakePluginInstaller) InstallReturns(result1 string) {
	fake.InstallStub = nil
	fake.installReturns = struct {
		result1 string
	}{result1}
}

var _ plugininstaller.PluginInstaller = new(FakePluginInstaller)
