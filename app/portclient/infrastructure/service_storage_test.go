package infrastructure

import (
	"testing"

	"github.com/bkatrenko/ports/proto"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	exampleFilePath = "../../../ports.json"
)

func TestDecodeFile(t *testing.T) {
	Convey("test decode file", t, func() {
		storage := ServiceStorage{}
		reader, err := storage.readFile(exampleFilePath)
		if err != nil {
			t.Fatal(err)
		}

		ch := storage.decodeFile(reader)

		res := []map[string]*proto.Port{}
		for v := range ch {
			res = append(res, v)
		}

		// just check few fields from first and last ports map
		// TODO: add more assertions
		So(len(res), ShouldEqual, 1632)
		So(res[0]["AEAJM"].Name, ShouldEqual, "Ajman")
		So(res[0]["AEAJM"].Timezone, ShouldEqual, "Asia/Dubai")

		So(res[1631]["ZWUTA"].Name, ShouldEqual, "Mutare")
		So(res[1631]["ZWUTA"].Timezone, ShouldEqual, "Africa/Harare")
	})
}
