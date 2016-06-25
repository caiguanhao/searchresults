package searchresults_test

import (
	"testing"

	"github.com/caiguanhao/searchresults"
)

func TestGetBaiduCount(t *testing.T) {
	bdcount, err := searchresults.GetBaiduCount("万圣节")
	if err != nil {
		t.Error(err)
	} else if bdcount < 10000 {
		t.Error("too few results")
	} else {
		t.Logf("baidu has %d results", bdcount)
	}
}

func TestGetSogouCount(t *testing.T) {
	sgcount, err := searchresults.GetSogouCount("万圣节")
	if err != nil {
		t.Error(err)
	} else if sgcount < 10000 {
		t.Error("too few results")
	} else {
		t.Logf("sogou has %d results", sgcount)
	}
}

func TestGetSoDotComCount(t *testing.T) {
	socount, err := searchresults.GetSoDotComCount("万圣节")
	if err != nil {
		t.Error(err)
	} else if socount < 10000 {
		t.Error("too few results")
	} else {
		t.Logf("so.com has %d results", socount)
	}
}
