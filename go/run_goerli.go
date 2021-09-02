package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	tradeExec "github.com/happenwah/uniswapTrader/go/tradeExecution"

	trader "github.com/happenwah/uniswapTrader/go/Trader"
)

var (
	g_client *ethclient.Client
	g_GOERLI *big.Int = big.NewInt(5)
)

const (
	g_contractAddress string = "0x7546Fa4BF2625066C71DCff86B17D6b73FC3C5fE"
	g_WETHGoerli      string = "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"
	g_WETHAddress     string = g_WETHGoerli
	g_uniV2Router     string = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
	g_flashbotsRelay  string = "https://relay-goerli.flashbots.net"
	g_nBundles        uint64 = 10
)

var (
	g_zk       *tradeExec.Zk
	g_instance *trader.DEXTrader
)

func TestGoerliBuyAndSell(zk *tradeExec.Zk, chainID *big.Int, client *ethclient.Client, instance *trader.DEXTrader, isMempool bool, gasPrice *big.Int, g_client *ethclient.Client) {
	// Token: Goerli gDAI
	token := "0x9c69cf4e75099bfdcc9e5d97446b1b289881aade"
	token0, poolAddress, _ := tradeExec.GetPairAddress(g_uniV2Router, token, g_WETHGoerli)
	// Dummy values have been chosen
	WETHAmountBuy_wei := big.NewInt(1e17)
	TokenAmountBuy := big.NewInt(1e17)
	WETHAmountSell_wei := big.NewInt(109)
	minerAmount_wei := big.NewInt(11239578123888112)

	tradeExec.ExecuteDEXBuyAndSell(isMempool, zk, chainID, client, instance, g_flashbotsRelay,
		common.HexToAddress(g_WETHAddress), common.HexToAddress(token),
		poolAddress, token0, WETHAmountBuy_wei, TokenAmountBuy, WETHAmountSell_wei,
		gasPrice, minerAmount_wei, g_nBundles)
}

func main() {
	provider := "/home/happenwah/geth/datadir/geth.ipc"

	g_client, err := ethclient.Dial(provider)

	if err != nil {
		log.Fatalf("There was a problem dialing the client: %v ", err)
	} else {
		fmt.Println("Successfully connected to eth node")
	}

	walletAddressIdx := 2
	g_zk = tradeExec.Init(walletAddressIdx)
	g_instance = tradeExec.SetContractTrader(g_client, g_contractAddress)
	gasPrice := big.NewInt(10e9)
	// Set to true to submit via mempool
	// Set to false to submit to Flashbots Goerli relay
	isMempool := false
	TestGoerliBuyAndSell(g_zk, g_GOERLI, g_client, g_instance, isMempool, gasPrice, g_client)

	// CONVERT 1 ETH INTO WETH
	//tradeExec.ConvertETHtoWETH(g_zk, g_client, g_GOERLI, g_instance, g_WETHAddress, big.NewInt(1e18), gasPrice)

	// DESTROY CONTRACT
	//tradeExec.Kill(g_zk, g_client, g_GOERLI, g_instance, gasPrice)
}
