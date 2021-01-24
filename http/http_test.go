// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package http

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHTTP(t *testing.T) {
	url := "http://luajit.org/luajit.html"

	client := New()
	assert.NotNil(t, client)

	b, err := client.DownloadIf(context.Background(), url, time.Now())
	assert.Nil(t, b)
	assert.NoError(t, err)
}
