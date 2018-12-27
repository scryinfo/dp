package contractclient

import (
	"../contractinterface"
	"../util/storage/ipfsaccess"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"../core/chainoperations"
	"./contractinterfacewrapper"
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

func NewContractClient(myAddr string, clientIdentityStream string, clientPassword string) (*ContractClient) {
	return &ContractClient{
		address: myAddr,
		identityStream: clientIdentityStream,
		password: clientPassword,
	}
}

func (contractClient *ContractClient) Initialize(ethNodeAddr string,
	                                             protocolContractAddr string,
	                                             tokenContractAddr string,
	                                             ipfsNodeAddr string,
	                                             ) (*Connector, error) {
	err := ipfsaccess.GetInstance().Initialize(ipfsNodeAddr)
	if err != nil {
		fmt.Println("failed to initialize ipfs. error: " + err.Error())
		return nil, err
	}

	connector, err := NewConnector(ethNodeAddr, protocolContractAddr, tokenContractAddr)
	if err != nil {
		fmt.Println("failed to initialize connector. error: " + err.Error())
		return nil, err
	}

	opts := chainoperations.BuildTransactOpts(common.HexToAddress(contractClient.address), contractClient.identityStream, contractClient.password)
	contractinterfacewrapper.Initialize(connector.protocolClient, opts)

	return connector, nil
}


type Connector struct {
	ctx context.Context
	conn *ethclient.Client
	protocolClient *contractinterface.ScryProtocol
}

func NewConnector(ethNodeAddr string,
	              protocolContractAddr string,
	              tokenContractAddr string) (*Connector, error) {
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

	return &Connector{
		ctx: context.Background(),
		conn: cn,
		protocolClient: pc,
	}, nil
}


