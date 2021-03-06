/*
Copyright 2016 The Doctl Authors All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"fmt"
	"testing"

	"github.com/digitalocean/doctl"
	"github.com/stretchr/testify/assert"
)

func TestVolumeActionCommand(t *testing.T) {
	cmd := VolumeAction()
	assert.NotNil(t, cmd)
	assertCommandNames(t, cmd, "attach", "detach", "detach-by-droplet-id", "resize")
}

func TestVolumeActionsAttach(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		tm.volumeActions.On("Attach", testVolume.ID, testDroplet.ID).Return(&testAction, nil)
		config.Args = append(config.Args, testVolume.ID)
		config.Args = append(config.Args, fmt.Sprintf("%d", testDroplet.ID))

		err := RunVolumeAttach(config)
		assert.NoError(t, err)
	})
}

func TestVolumeActionsDetach(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		tm.volumeActions.On("Detach", testVolume.ID).Return(&testAction, nil)
		config.Args = append(config.Args, testVolume.ID)

		err := RunVolumeDetach(config)
		assert.NoError(t, err)
	})
}

func TestVolumeDetachByDropletID(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		tm.volumeActions.On("DetachByDropletID", testVolume.ID, testDroplet.ID).Return(&testAction, nil)
		config.Args = append(config.Args, testVolume.ID)
		config.Args = append(config.Args, fmt.Sprintf("%d", testDroplet.ID))

		err := RunVolumeDetachByDropletID(config)
		assert.NoError(t, err)
	})
}

func TestVolumeResize(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		tm.volumeActions.On("Resize", testVolume.ID, 150, "dev0").Return(&testAction, nil)
		config.Args = append(config.Args, testVolume.ID)

		config.Doit.Set(config.NS, doctl.ArgSizeSlug, 150)
		config.Doit.Set(config.NS, doctl.ArgRegionSlug, "dev0")

		err := RunVolumeResize(config)
		assert.NoError(t, err)
	})
}
