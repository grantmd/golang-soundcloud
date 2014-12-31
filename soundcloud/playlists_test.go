package soundcloud

import (
	"net/url"
	"testing"
)

func TestPlaylists(t *testing.T) {
	ret, err := api.Playlists(url.Values{"q": []string{"25 Nights For Nujabes"}})
	if err != nil {
		t.Error(err)
	} else if len(ret) == 0 {
		t.Error("query returned no values")
	} else if ret[0].User.Username != "Ta-ku" {
		t.Error("Ta-ku wasn't the creator first response?!", ret[0].User.Username)
	}
}

func TestPlaylistGet(t *testing.T) {
	ret, err := api.Playlist(nightsfornujabes_id).Get(nil)
	if err != nil {
		t.Error(err)
	} else if ret.Id != nightsfornujabes_id {
		t.Error("id didn't come back as requested", ret.Id)
	} else if ret.Permalink != "25-nights-for-nujabes" {
		t.Error("25-nights-for-nujabes's permalink changed?", ret.Permalink)
	} else if ret.Title != "25 Nights For Nujabes" {
		t.Error("25 Nights For Nujabes title changed?", ret.Title)
	} else if ret.User.Permalink != "takugotbeats" || ret.User.Username != "Ta-ku" {
		t.Error("user object for takugotbeats changed?", ret.User.Permalink, ret.User.Username)
	}
}
