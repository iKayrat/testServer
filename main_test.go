package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

var testcaseServers = []Server{
	{
		Address:     "server1",
		CPUcapacity: 800,
		CPUusage:    0,
		Status:      true,
		Rooms:       []Room{},
	},
	{
		Address:     "server2",
		CPUcapacity: 800,
		CPUusage:    0,
		Status:      false,
		Rooms:       []Room{},
	},
	{
		Address:     "server3",
		CPUcapacity: 800,
		CPUusage:    80,
		Status:      true,
		Rooms:       []Room{{Id: rand.Int63n(10), Duration: 60}},
	},
	{
		Address:     "server4",
		CPUcapacity: 800,
		CPUusage:    120,
		Status:      true,
		Rooms:       []Room{{Id: rand.Int63n(10), Duration: 180}},
	},
}

var testcaseRoom = []Room{
	{
		Id:       1,
		Duration: 60,
	},
	// {
	// 	Id:       2,
	// 	Duration: 180,
	// },
	// {
	// 	Id:       3,
	// 	Duration: 60,
	// },
}

func TestNoRooms(t *testing.T) {

	want := "server1"

	tc := testcaseServers[:2]

	for _, tr := range testcaseRoom {
		actual := newRoom(tr, tc)

		require.Equal(t, want, actual)
	}
}

func TestUsageOverCapacity(t *testing.T) {

	want := "server1"

	tc := testcaseServers
	tc[2].CPUusage = 800
	tc[3].CPUusage = 800

	for _, tr := range testcaseRoom {
		actual := newRoom(tr, tc)

		require.Equal(t, want, actual)

		if want != actual {
			t.Error("not equal")
		}

	}

}
