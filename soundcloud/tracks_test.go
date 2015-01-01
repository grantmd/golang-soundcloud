package soundcloud

import (
	"net/url"
	"testing"
)

func TestTracks(t *testing.T) {
	ret, err := api.Tracks(url.Values{"q": []string{"welcome to night vale"}})
	if err != nil {
		t.Error(err)
	} else if len(ret) == 0 {
		t.Error("query returned no values")
	} else if ret[0].User.Username != "planetoffinks" {
		t.Error("planetoffinks wasn't the creator first response?!")
	}
}

func TestTrackGet(t *testing.T) {
	ret, err := api.Track(regain_control_id).Get(nil)
	if err != nil {
		t.Error(err)
	} else if ret.Id != regain_control_id {
		t.Error("id didn't come back as requested", ret.Id)
	} else if ret.Permalink != "shirobon-regain-control" {
		t.Error("shirobon-regain-control's permalink changed?", ret.Permalink)
	} else if ret.Title != "Regain Control" {
		t.Error("Regain Control title changed?", ret.Title)
	} else if ret.UserId != shiroban_id || ret.UserId != ret.User.Id {
		t.Error("user object is messed up", ret.UserId, ret.User)
	} else if ret.User.Permalink != "shirobon" || ret.User.Username != "Shirobon" {
		t.Error("user object for shirobon changed?", ret.User.Permalink, ret.User.Username)
	} else if ret.PlaybackCount < 1000 || ret.FavoritingsCount < 1000 || ret.CommentCount < 100 {
		t.Error("Some counts are wrong?", ret)
	}
}

func TestTrackStream(t *testing.T) {
	track, err := api.Track(regain_control_id).Get(nil)
	if err != nil {
		t.Error(err)
	}

	streamUrl, err := track.Stream()
	if err != nil {
		t.Error(err)
	} else if streamUrl == nil {
		t.Error("streamUrl is empty?")
	}
}
