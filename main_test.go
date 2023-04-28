package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testcaseRoom = []Room{
	{
		Id:       1,
		Duration: 60,
	},
	{
		Id:       2,
		Duration: 180,
	},
}

func TestHasRooms(t *testing.T) {

	want1 := "server4"
	want2 := "server4"
	want3 := "server4"

	// add rooms to server4
	actual1 := newRoom(testcaseRoom[0])
	actual2 := newRoom(testcaseRoom[0])
	actual3 := newRoom(testcaseRoom[0])

	require.Equal(t, want1, actual1)
	require.Equal(t, want2, actual2)
	require.Equal(t, want3, actual3)

}
func TestServer4Full(t *testing.T) {

	want1 := "server4"
	want2 := "server3"

	// filling server4 manually
	actual1 := newRoom(testcaseRoom[0])
	require.Equal(t, want1, actual1)
	for i := 0; i < 5; i++ {
		actual2 := newRoom(testcaseRoom[1])
		require.Equal(t, want1, actual2)
	}

	// TEST WE WANT
	// added room to server3 becasue server4 is full
	actual2 := newRoom(testcaseRoom[1])
	require.Equal(t, want2, actual2)

}

// When rooms filled in server3,4 , adds into working server
func TestNoRooms(t *testing.T) {

	want1 := "server4"
	want2 := "server3"
	want3 := "server1"

	// filling server4 manually
	actual1 := newRoom(testcaseRoom[0])
	require.Equal(t, want1, actual1)
	for i := 0; i < 5; i++ {
		actual1 := newRoom(testcaseRoom[1])
		require.Equal(t, want1, actual1)
	}

	// filling server3 manually
	for i := 0; i < 6; i++ {
		actual2 := newRoom(testcaseRoom[1])
		require.Equal(t, want2, actual2)
	}

	// TEST WE WANT
	actual3 := newRoom(testcaseRoom[1])
	require.Equal(t, want3, actual3)

}

func TestNoAnyRunningServer(t *testing.T) {

	want1 := "server4"
	want2 := "server3"
	want3 := "server1"
	want4 := "server2"

	// filling server4 manually
	actual1 := newRoom(testcaseRoom[0])
	require.Equal(t, want1, actual1)
	for i := 0; i < 5; i++ {
		actual1 := newRoom(testcaseRoom[1])
		require.Equal(t, want1, actual1)
	}

	// filling server3 manually
	for i := 0; i < 6; i++ {
		actual2 := newRoom(testcaseRoom[1])
		require.Equal(t, want2, actual2)
	}

	//filling server1
	actual3 := newRoom(testcaseRoom[0])
	require.Equal(t, want3, actual3)
	for i := 0; i < 6; i++ {
		actual3 := newRoom(testcaseRoom[1])
		require.Equal(t, want3, actual3)
	}

	// TEST WE WANT
	actual4 := newRoom(testcaseRoom[1])
	require.Equal(t, want4, actual4)

}

func TestNoAnyServer(t *testing.T) {

	want1 := "server4"
	want2 := "server3"
	want3 := "server1"
	want4 := "server2"
	want5 := "newserver"

	// filling server4 manually
	actual1 := newRoom(testcaseRoom[0])
	require.Equal(t, want1, actual1)
	for i := 0; i < 5; i++ {
		actual1 := newRoom(testcaseRoom[1])
		require.Equal(t, want1, actual1)
	}

	// filling server3 manually
	for i := 0; i < 6; i++ {
		actual2 := newRoom(testcaseRoom[1])
		require.Equal(t, want2, actual2)
	}

	//filling server1
	actual3 := newRoom(testcaseRoom[0])
	require.Equal(t, want3, actual3)
	for i := 0; i < 6; i++ {
		actual3 := newRoom(testcaseRoom[1])
		require.Equal(t, want3, actual3)
	}

	// filling server2
	actual4 := newRoom(testcaseRoom[0])
	require.Equal(t, want4, actual4)
	for i := 0; i < 6; i++ {
		actual4 := newRoom(testcaseRoom[1])
		require.Equal(t, want4, actual4)
	}

	// TEST WE WANT
	actual5 := newRoom(testcaseRoom[0])
	require.Equal(t, want5, actual5)

}
