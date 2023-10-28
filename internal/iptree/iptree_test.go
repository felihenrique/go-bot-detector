package iptree_test

import (
	"testing"

	"github.com/felihenrique/go-botdetector/internal/iptree"
	"github.com/stretchr/testify/assert"
)

func TestCreateEmpty(t *testing.T) {
	it := iptree.New(0)
	added, err := it.Add("192.165.1.0")
	assert.True(t, added)
	assert.Nil(t, err)

	found, err := it.Has("192.165.1.0")
	assert.Nil(t, err)
	assert.True(t, found)
}

func TestCreate(t *testing.T) {
	it := iptree.New(10)
	added, err := it.Add("1.1.1.1")
	assert.True(t, added)
	assert.Nil(t, err)

	found, err := it.Has("1.1.1.1")
	assert.Nil(t, err)
	assert.True(t, found)
}

func TestAddParsing(t *testing.T) {
	it := iptree.New(0)

	added, err := it.Add("192.168.1.10/23")
	assert.True(t, added)
	assert.Nil(t, err)

	added2, err := it.Add("0.0.0.0/24")
	assert.True(t, added2)
	assert.Nil(t, err)

	added3, err := it.Add("192.111.15")
	assert.False(t, added3)
	assert.NotNil(t, err)

	added4, err := it.Add("134.111.5.3/ab")
	assert.False(t, added4)
	assert.NotNil(t, err)
}

func TestFind(t *testing.T) {
	it := iptree.New(1)
	added, err := it.Add("10.5.3.1/23")
	assert.True(t, added)
	assert.Nil(t, err)

	found, err := it.Has("10.5.2.10")
	assert.True(t, found)
	assert.Nil(t, err)

	found2, err := it.Has("10.5.2.0/24")
	assert.True(t, found2)
	assert.Nil(t, err)
}
