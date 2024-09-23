package usecases

import "github.com/Archisman-Mridha/chat-service/internal/types"

type Usecases struct {
	db      types.Database
	kvStore types.KVStore
}

func NewUsecases(db types.Database, kvStore types.KVStore) *Usecases {
	return &Usecases{
		db,
		kvStore,
	}
}
