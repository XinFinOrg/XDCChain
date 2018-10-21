package blocksigner

 import (
	"context"
	"math/big"
	"testing"
	"time"

 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

 var (
	key, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	addr   = crypto.PubkeyToAddress(key.PublicKey)
)

 func TestBlockSigner(t *testing.T) {
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{addr: {Balance: big.NewInt(1000000000)}})
	transactOpts := bind.NewKeyedTransactor(key)

 	_, blockSigner, err := DeployBlockSigner(transactOpts, contractBackend)
	if err != nil {
		t.Fatalf("can't deploy root registry: %v", err)
	}
	contractBackend.Commit()

	d := time.Now().Add(1000 * time.Millisecond)
	ctx, _ := context.WithDeadline(context.Background(), d)
	code, _ := contractBackend.CodeAt(ctx, blockSignerAddress, nil)
	t.Log("contract code", common.ToHex(code))
	f := func(key, val common.Hash) bool {
		t.Log(key.Hex(), val.Hex())
		return true
	}
	contractBackend.ForEachStorageAt(ctx, blockSignerAddress, nil, f) 
	
	signers, err := blockSigner.GetSigners(big.NewInt(0))
	if err != nil {
		t.Fatalf("can't get candidates: %v", err)
	}
	for _, it := range signers {
		t.Log("signer", it.String())
	}
	contractBackend.Commit()
}