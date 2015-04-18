package assingle

import (
	"github.com/aerospike/aerospike-client-go"
	"github.com/plimble/utils/errors2"
	"github.com/plimble/utils/pool"
	"github.com/tinylib/msgp/msgp"
	"sync"
)

var (
	ErrKeyExist    = errors2.NewInternal("Key already exists")
	ErrKeyNotExist = errors2.NewInternal("Key is not exists")
	ErrNotFound    = errors2.NewNotFound("not found")
)

type ASSingle struct {
	client   *aerospike.Client
	ns       string
	pool     *pool.BufferPool
	iterPool *sync.Pool
}

func New(client *aerospike.Client, ns string, poolSize int) *ASSingle {
	as := &ASSingle{
		client: client,
		ns:     ns,
		pool:   pool.NewBufferPool(poolSize),
	}

	as.iterPool = &sync.Pool{}

	as.iterPool.New = func() interface{} {
		return &Iterator{
			itrPool: as.iterPool,
		}
	}

	return as
}

func (a *ASSingle) Close() {
	a.client.Close()
}

func (a *ASSingle) errGet(record *aerospike.Record, err error) error {
	switch {
	case err == nil && record != nil:
		return nil
	case err != nil:
		return errors2.NewInternal(err.Error())
	case record == nil:
		return ErrNotFound
	}

	return nil
}

func (a *ASSingle) errPut(err error) error {
	if err == nil {
		return nil
	}

	switch err.Error() {
	case "Key already exists":
		return ErrKeyExist
	default:
		return errors2.NewInternal(err.Error())
	}

	return nil
}

func (a *ASSingle) Put(policy *aerospike.WritePolicy, set, key string, val msgp.Marshaler) error {
	k, err := aerospike.NewKey(a.ns, set, key)
	if err != nil {
		return errors2.NewInternal(err.Error())
	}

	data, err := val.MarshalMsg(nil)
	if err != nil {
		return errors2.NewInternal(err.Error())
	}

	bin := aerospike.NewBin("", data)

	return a.errPut(a.client.PutBins(policy, k, bin))
}

func (a *ASSingle) Delete(policy *aerospike.WritePolicy, set, key string) error {
	k, err := aerospike.NewKey(a.ns, set, key)
	if err != nil {
		return errors2.NewInternal(err.Error())
	}

	exist, err := a.client.Delete(policy, k)
	if err != nil {
		return errors2.NewInternal(err.Error())
	}

	if !exist {
		return ErrKeyNotExist
	}

	return nil
}

func (a *ASSingle) Get(policy *aerospike.BasePolicy, set, key string, val msgp.Unmarshaler) error {
	k, err := aerospike.NewKey(a.ns, set, key)
	if err != nil {
		return errors2.NewInternal(err.Error())
	}

	record, err := a.client.Get(policy, k)
	if err := a.errGet(record, err); err != nil {
		return err
	}

	if _, err := val.UnmarshalMsg(record.Bins[""].([]byte)); err != nil {
		return errors2.NewInternal(err.Error())
	}

	return nil
}

func (a *ASSingle) Exist(policy *aerospike.BasePolicy, set string, key string) (bool, error) {
	k, err := aerospike.NewKey(a.ns, set, key)
	if err != nil {
		return false, errors2.NewInternal(err.Error())
	}

	exist, err := a.client.Exists(policy, k)
	if err != nil {
		return false, errors2.NewInternal(err.Error())
	}

	return exist, nil
}

func (a *ASSingle) MGet(policy *aerospike.BasePolicy, set string, keys ...string) (*Iterator, error) {
	kList := make([]*aerospike.Key, len(keys))
	for i, key := range keys {
		kList[i], _ = aerospike.NewKey(a.ns, set, key)
	}

	records, err := a.client.BatchGet(policy, kList, "")
	if err != nil {
		return nil, errors2.NewInternal(err.Error())
	}

	it := a.iterPool.Get().(*Iterator)
	it.index = 0
	it.size = 0
	it.recordLen = len(records)
	it.records = records

	for i := 0; i < it.recordLen; i++ {
		if records[i] != nil {
			it.size++
		}
	}

	return it, nil
}
