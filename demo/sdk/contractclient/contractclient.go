package contractclient

import (
	"../core"
	"../contractinterface"
	"../core/chainoperations"
	"../util/storage/ipfsaccess"
	"./contractinterfacewrapper"
	"../core/chainevents"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)


type ContractOptions struct {
	ethNodeAddr string
	protocolContractAddr string
	tokenContractAddr string
	ipfsNodeAddr string
}

type ContractClient struct {
	contractOptions ContractOptions
	address string
	identityStream string
	password string
}

func NewContractClient(publicKey string, clientIdentityStream string, clientPassword string) (*ContractClient) {
	return &ContractClient{
		address: publicKey,
		identityStream: clientIdentityStream,
		password: clientPassword,
	}
}

func (contractClient *ContractClient) Initialize(ethNodeAddr string,
	                                             protocolContractAddr string,
	                                             protocolContractABI string,
	                                             ipfsNodeAddr string,
	                                             ) (*connector, error) {
    defer func() {
    	if err := recover(); err != nil {
    		fmt.Println("failed to initialize contract client, error:", err)
		}
	}()

	err := ipfsaccess.GetInstance().Initialize(ipfsNodeAddr)
	if err != nil {
		fmt.Println("failed to initialize ipfs. error: " + err.Error())
		return nil, err
	}

	connector, err := newConnector(ethNodeAddr, protocolContractAddr)
	if err != nil {
		fmt.Println("failed to initialize connector. error: " + err.Error())
		return nil, err
	}

	opts := chainoperations.BuildTransactOpts(common.HexToAddress(contractClient.address), contractClient.identityStream, contractClient.password)
	contractinterfacewrapper.Initialize(connector.protocolClient, opts)

	core.StartEngine(connector.conn, protocolContractAddr, protocolContractABI)

	return connector, nil
}

func (contractClient *ContractClient) SubscribeEvent(eventName string, callback chainevents.EventCallback)  {
}

type connector struct {
	ctx context.Context
	conn *ethclient.Client
	protocolClient *contractinterface.ScryProtocol
}

func newConnector(ethNodeAddr string,
	              protocolContractAddr string) (*connector, error) {
	cn, err := ethclient.Dial(ethNodeAddr)
	if err != nil {
		fmt.Println("failed to connect to node:" + ethNodeAddr)
		return nil, err
	}

	pc, err := contractinterface.NewScryProtocol(common.HexToAddress(protocolContractAddr), cn)
	if err != nil {
		fmt.Println("failed to create scry protocol client:" + protocolContractAddr)
		return nil, err
	}

	return &connector{
		ctx: context.Background(),
		conn: cn,
		protocolClient: pc,
	}, nil
}


