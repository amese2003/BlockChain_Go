package blockchain

import (
	utils "blockchain/Utils"
	"reflect"
	"testing"
)

func TestCreateBlock(t *testing.T) {
	dbStorage = fakeDB{}
	Mempool().Txs["test"] = &Tx{}
	b := createBlock("x", 1, 1)
	if reflect.TypeOf(b) != reflect.TypeOf(&Block{}) {
		t.Error("createBlock() should return an instance of a block")
	}
}

func TestFindBlock(t *testing.T) {
	t.Run("Block not Found", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeFindBlock: func() []byte {
				return nil
			},
		}

		_, err := FindBlock("xx")
		if err == nil {
			t.Error("블록이 없어야 합니다.")
		}
	})

	t.Run("block is found", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeFindBlock: func() []byte {
				b := &Block{
					Height: 1,
				}

				return utils.ToBytes(b)
			},
		}

		block, _ := FindBlock("xx")
		if block.Height != 1 {
			t.Error("블록이 있어야 합니다.")
		}
	})
}
