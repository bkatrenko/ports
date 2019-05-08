package testing

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bkatrenko/ports/proto"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	testURL = "http://localhost:8080/v1/ports?offset=0&limit=100"
)

// TestIntegration is very simple but cute test - after start all components it's will request
// ports data and do some assertions to validate it.
// Will be cool to add more checks - but the main reason of creating this test is just give some
// understanding of how to test that this apps works fine
func TestIntegration(t *testing.T) {

	var expectedPortsLen = 100

	Convey("check all components in work - just make ports request", t, func() {
		resp, err := http.Get(testURL)
		if err != nil {
			t.Fatal(err)
		}

		body := map[string]*proto.Port{}

		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}

		// TODO: add more assertions
		So(len(body), ShouldEqual, expectedPortsLen)
		So(body["AEAJM"].Name, ShouldEqual, "Ajman")
		So(body["AEAJM"].Timezone, ShouldEqual, "Asia/Dubai")
	})
}
