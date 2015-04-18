package assingle

import (
	"github.com/aerospike/aerospike-client-go"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type ASSingleSuite struct {
	suite.Suite
	data   DataTest
	client *aerospike.Client
	as     *ASSingle
}

func TestASSingleSuite(t *testing.T) {
	suite.Run(t, &ASSingleSuite{})
}

func (t *ASSingleSuite) SetupSuite() {
	client, err := aerospike.NewClient("192.168.99.100", 3000)
	if err != nil {
		panic(err)
	}
	t.client = client

	t.data = DataTest{
		Name:      "test@test.com",
		Email:     "test@test.com",
		Age:       90,
		Scope:     []string{"1", "2"},
		CreatedAt: time.Now(),
	}
}

func (t *ASSingleSuite) SetupTest() {
	t.as = New(t.client, "test", 512)
}

func (t *ASSingleSuite) TearDownTest() {
	key1, _ := aerospike.NewKey(t.as.ns, "access_token", "1")
	key2, _ := aerospike.NewKey(t.as.ns, "access_token", "2")
	key3, _ := aerospike.NewKey(t.as.ns, "access_token", "3")
	t.as.client.Delete(nil, key1)
	t.as.client.Delete(nil, key2)
	t.as.client.Delete(nil, key3)
}

func (t *ASSingleSuite) TearDownSuite() {
	t.as.Close()
}

func (t *ASSingleSuite) TestPut() {
	err := t.as.Put(nil, "access_token", "1", &t.data)
	t.NoError(err)

	//put exist key
	policy := aerospike.NewWritePolicy(0, 0)
	policy.RecordExistsAction = aerospike.CREATE_ONLY
	err = t.as.Put(policy, "access_token", "1", &t.data)
	t.Equal(ErrKeyExist, err)
}

func (t *ASSingleSuite) TestDelete() {
	err := t.as.Put(nil, "access_token", "1", &t.data)
	t.NoError(err)

	err = t.as.Delete(nil, "access_token", "1")
	t.NoError(err)

	//delete with not exist
	err = t.as.Delete(nil, "access_token", "1")
	t.Equal(ErrKeyNotExist, err)
}

func (t *ASSingleSuite) TestGet() {
	getData := DataTest{}

	err := t.as.Put(nil, "access_token", "1", &t.data)
	t.NoError(err)

	err = t.as.Get(nil, "access_token", "1", &getData)
	t.NoError(err)
	t.Equal(t.data, getData)

	//get not exist
	err = t.as.Get(nil, "access_token", "2", &getData)
	t.Equal(ErrNotFound, err)
}

func (t *ASSingleSuite) TestExist() {
	getData := DataTest{}

	err := t.as.Put(nil, "access_token", "1", &t.data)
	t.NoError(err)

	err = t.as.Get(nil, "access_token", "1", &getData)
	t.NoError(err)
	t.Equal(t.data, getData)

	//get not exist
	err = t.as.Get(nil, "access_token", "2", &getData)
	t.Equal(ErrNotFound, err)
}

func (t *ASSingleSuite) TestMGet() {
	err := t.as.Put(nil, "access_token", "1", &t.data)
	t.NoError(err)

	err = t.as.Put(nil, "access_token", "2", &t.data)
	t.NoError(err)

	err = t.as.Put(nil, "access_token", "3", &t.data)
	t.NoError(err)

	it, err := t.as.MGet(nil, "access_token", "0", "1", "2", "4", "3")
	t.NoError(err)
	t.Equal(3, it.Size())

	list := make([]DataTest, it.Size())
	for i := 0; i < it.Size(); i++ {
		list[i] = DataTest{}
		it.Scan(&list[i])
		t.Equal(t.data, list[i])
	}

	t.NoError(it.Close())
}
