package infrastructure

import (
	"reflect"
	"testing"
	"time"

	"fmt"

	"github.com/bkatrenko/ports/proto"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	testMongoAddr = "mongodb://localhost:27017"
)

func TestUpsertPort(t *testing.T) {
	Convey("test insert port", t, func() {
		storage, err := NewMongoStorage(testMongoAddr, time.Minute)
		if err != nil {
			t.Fatal(err)
		}

		keys := []string{"TEST1", "TEST2"}
		ports := []map[string]*proto.Port{
			map[string]*proto.Port{
				keys[0]: &proto.Port{
					Name:        "PORT1",
					City:        "CITY1",
					Country:     "UK",
					Alias:       []string{"UK1", "BR1"},
					Regions:     []string{"London1"},
					Coordinates: []float32{1.1, 1.1},
					Province:    "LND1",
					Timezone:    "London/UK1",
					Unlocks:     []string{"UNLOCK1", "UNLOCK1"},
					Code:        "UK1",
				},
			},

			map[string]*proto.Port{
				keys[1]: &proto.Port{
					Name:        "PORT2",
					City:        "CITY2",
					Country:     "US",
					Alias:       []string{"UK2", "BR2"},
					Regions:     []string{"London2"},
					Coordinates: []float32{2.2, 2.2},
					Province:    "LND2",
					Timezone:    "London/UK2",
					Unlocks:     []string{"UNLOCK2", "UNLOCK2"},
					Code:        "UK2",
				},
			},
		}

		retrievedPort := putAndRetrieve(t, storage, ports[0], keys[0])
		if !reflect.DeepEqual(retrievedPort, ports[0][keys[0]]) {
			fmt.Println(ports[0][keys[0]], retrievedPort)
			t.Fatal("port [0] isn't equal after retrieve")
		}

		ports[0][keys[0]].City = "Kiyv"

		retrievedPort = putAndRetrieve(t, storage, ports[0], keys[0])
		if !reflect.DeepEqual(retrievedPort, ports[0][keys[0]]) {
			t.Fatal("port [0] isn't equal after retrieve and update")
		}

		retrievedPort = putAndRetrieve(t, storage, ports[1], keys[1])
		if !reflect.DeepEqual(retrievedPort, ports[1][keys[1]]) {
			t.Fatal("port [1] isn't equal after retrieve")
		}
	})
}

func putAndRetrieve(t *testing.T, storage *MongoStorage, port map[string]*proto.Port, key string) *proto.Port {
	if err := storage.PutPorts(port); err != nil {
		t.Fatal(err)
	}

	foundPorts, err := storage.GetPorts(0, 100)
	if err != nil {
		t.Fatal(err)
	}

	return foundPorts[key]
}
