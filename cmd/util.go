package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	logging "github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p/core/peer"
)

// TODO boost retrieval car file
func ShowEnv(logg *logging.ZapEventLogger) {
	logg.Infof("CAR_SERVER_URLS: %s", checkEnv("CAR_SERVER_URLS", false))
}

func SetEnv(k, v string) {
	os.Setenv(k, v)
}

func checkEnv(env string, must bool) string {
	e := os.Getenv(env)
	if e == "" {
		if must {
			panic(fmt.Sprintf("env: %s can't empty", env))
		}
	}
	return e
}

func GetAddrInfo(ctx context.Context, api api.Gateway, maddr address.Address) (*peer.AddrInfo, error) {
	minfo, err := api.StateMinerInfo(ctx, maddr, types.EmptyTSK)
	if err != nil {
		return nil, err
	}
	if minfo.PeerId == nil {
		return nil, fmt.Errorf("storage provider %s has no peer ID set on-chain", maddr)
	}

	var maddrs []multiaddr.Multiaddr
	for _, mma := range minfo.Multiaddrs {
		ma, err := multiaddr.NewMultiaddrBytes(mma)
		if err != nil {
			return nil, fmt.Errorf("storage provider %s had invalid multiaddrs in their info: %w", maddr, err)
		}
		maddrs = append(maddrs, ma)
	}
	if len(maddrs) == 0 {
		return nil, fmt.Errorf("storage provider %s has no multiaddrs set on-chain", maddr)
	}

	return &peer.AddrInfo{
		ID:    *minfo.PeerId,
		Addrs: maddrs,
	}, nil
}

func PrintJson(obj interface{}) error {
	resJson, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling json: %w", err)
	}

	fmt.Println(string(resJson))
	return nil
}
