package plugin

import (
	"context"

	"github.com/hashicorp/vault/logical"
)

// StorageServer is a net/rpc compatible structure for serving
type StorageServer struct {
	impl logical.Storage
}

func (s *StorageServer) List(prefix string, reply *StorageListReply) error {
	keys, err := s.impl.List(context.Background(), prefix)
	*reply = StorageListReply{
		Keys:  keys,
		Error: wrapError(err),
	}
	return nil
}

func (s *StorageServer) Get(key string, reply *StorageGetReply) error {
	storageEntry, err := s.impl.Get(context.Background(), key)
	*reply = StorageGetReply{
		StorageEntry: storageEntry,
		Error:        wrapError(err),
	}
	return nil
}

func (s *StorageServer) Put(entry *logical.StorageEntry, reply *StoragePutReply) error {
	err := s.impl.Put(context.Background(), entry)
	*reply = StoragePutReply{
		Error: wrapError(err),
	}
	return nil
}

func (s *StorageServer) Delete(key string, reply *StorageDeleteReply) error {
	err := s.impl.Delete(context.Background(), key)
	*reply = StorageDeleteReply{
		Error: wrapError(err),
	}
	return nil
}

type StorageListReply struct {
	Keys  []string
	Error error
}

type StorageGetReply struct {
	StorageEntry *logical.StorageEntry
	Error        error
}

type StoragePutReply struct {
	Error error
}

type StorageDeleteReply struct {
	Error error
}

// NOOPStorage is used to deny access to the storage interface while running a
// backend plugin in metadata mode.
type NOOPStorage struct{}

func (s *NOOPStorage) List(_ context.Context, prefix string) ([]string, error) {
	return []string{}, nil
}

func (s *NOOPStorage) Get(_ context.Context, key string) (*logical.StorageEntry, error) {
	return nil, nil
}

func (s *NOOPStorage) Put(_ context.Context, entry *logical.StorageEntry) error {
	return nil
}

func (s *NOOPStorage) Delete(_ context.Context, key string) error {
	return nil
}
