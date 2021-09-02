package tradeExecution

import (
	"bufio"
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync"

	trader "github.com/happenwah/uniswapTrader/go/Trader"

	jsoniter "github.com/json-iterator/go"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

var g_abi abi.ABI

// EOA struct
type Zk struct {
	Mu          sync.Mutex
	D           *ecdsa.PrivateKey
	FromAddress common.Address
}

var json = jsoniter.ConfigFastest

// Get the address of a UniswapV2 / SushiswapV2 pair given two contract addresses
// The first string returned is token0
func GetPairAddress(contractAddress, token, otherToken string) (common.Address, common.Address, error) {
	otherTokenInt := new(big.Int)
	otherTokenInt.SetString(otherToken[2:], 16)
	factoryffString := ""
	initCodeString := ""

	if strings.ToLower(contractAddress) == strings.ToLower("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D") {
		//0xff + UniswapV2 Factory
		factoryffString = "0xff5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
		initCodeString = "96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f"
	} else if contractAddress == strings.ToLower("0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F") {
		//0xff + SushiswapV2 Factory
		factoryffString = "0xffC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac"
		initCodeString = "e18a34eb0e04b04f7a0ac29a6e80748dca96319b42c54d679cb821dca90c6303"
	}

	if factoryffString == "" || initCodeString == "" {
		fmt.Printf("GetPairAdddress: failed set factoryffString or initCodeString")
		return common.Address{}, common.Address{}, errors.New("GetPairAdddress: failed set factoryffString or initCodeString")
	}

	// remove 0x
	tokenNorm := token[2:]
	// Get big.Int to compare
	tokenInt := new(big.Int)
	tokenInt.SetString(tokenNorm, 16)

	// Cmp returns -1 if tokenInt < otherTokenInt
	i := tokenInt.Cmp(otherTokenInt)

	if i == -1 {
		token0String := token
		token1String := otherToken
		token01String := token0String + token1String[2:]
		token01, err := hexutil.Decode(token01String)

		if err != nil {
			return common.Address{}, common.Address{}, err
		}

		salt := crypto.Keccak256(token01)
		finalString := factoryffString + hexutil.Encode(salt)[2:] + initCodeString
		pairAddress, err := hexutil.Decode(finalString)

		if err != nil {
			return common.Address{}, common.Address{}, err
		}

		pairAddressFinal := hexutil.Encode(crypto.Keccak256(pairAddress))

		return common.HexToAddress(token), common.HexToAddress(pairAddressFinal), nil

	} else if i == 1 {
		token0String := otherToken
		token1String := token
		token01String := token0String + token1String[2:]
		token01, err := hexutil.Decode(token01String)

		if err != nil {
			return common.Address{}, common.Address{}, err
		}

		salt := crypto.Keccak256(token01)
		finalString := factoryffString + hexutil.Encode(salt)[2:] + initCodeString
		pairAddress, err := hexutil.Decode(finalString)

		if err != nil {
			return common.Address{}, common.Address{}, err
		}

		pairAddressFinal := hexutil.Encode(crypto.Keccak256(pairAddress))

		return common.HexToAddress(otherToken), common.HexToAddress(pairAddressFinal), nil

	} else {
		return common.Address{}, common.Address{}, errors.New("Something went wrong, WETHint & token were equal")
	}
}

// Checks token balance of a given address
func CheckTokenBalance(tokenAddress, address common.Address, client *ethclient.Client) (error, *big.Int) {
	token, err := NewIERC20(tokenAddress, client)
	if err != nil {
		fmt.Printf("Failed to instantiate IERC20 of: %v\n", tokenAddress)
		return err, nil
	}

	tokenBalance, err := token.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		fmt.Printf("Failed to retrieve balance of token %v at address %v\n", tokenAddress, address)
		return err, nil
	}

	return nil, tokenBalance
}

// Setup wallet struct from mnemonic input,
// picking one of its EOA addresses via walletAddressIdx
func (zk *Zk) SetWalletFromMnemonic(walletAddressIdx int) {
	fmt.Print("Enter wallet mnemonic: ")
	inputReader := bufio.NewReader(os.Stdin)
	mnemonic, _ := inputReader.ReadString('\n')

	wallet, err := hdwallet.NewFromMnemonic(strings.TrimSpace(mnemonic))

	if err != nil {
		log.Fatal(err)
	}

	derivationPath := fmt.Sprintf("m/44'/60'/0'/0/%v", walletAddressIdx)
	path := hdwallet.MustParseDerivationPath(derivationPath)
	account, err := wallet.Derive(path, true)

	if err != nil {
		log.Fatal(err)
	}

	d, err := wallet.PrivateKey(account)

	if err != nil {
		log.Fatal(err)
	}

	zk.D = d
	zk.FromAddress = account.Address

	fmt.Printf("Parsed wallet address with index %v: %v\n", walletAddressIdx, zk.FromAddress)

	fmt.Println("Wallet is loaded")
}

// Loads instance of DEXTrader contract at its deployed address
func SetContractTrader(client *ethclient.Client, _address string) *trader.DEXTrader {
	address := common.HexToAddress(_address)
	instance, err := trader.NewDEXTrader(address, client)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("DEXTrader has been loaded from address: %v\n", address)
	return instance
}

// Simple example of submitting a WETH -> Token Buy,
// followed by a Token -> WETH sell.
// Works for both Flashbots and mempool, as dictated by isMempool bool.
func ExecuteDEXBuyAndSell(isMempool bool,
	zk *Zk,
	chainID *big.Int,
	client *ethclient.Client,
	instance *trader.DEXTrader,
	flashbotsRelay string,
	WETHAddress,
	tokenAddress,
	poolAddress,
	token0 common.Address,
	WETHAmountBuy_wei,
	TokenAmountBuy,
	WETHAmountSell_wei,
	gasPrice,
	minerAmount *big.Int,
	nBundles uint64) (DEXTradeReturn, error) {

	var txs = make(map[*big.Int][2]common.Hash)
	// Get latest known block
	blockNumber, err := client.BlockNumber(context.Background())

	if err != nil {
		fmt.Printf("Failed to retrieve block number with error: %v\n", err)
		out := DEXTradeReturn{
			Executed:      false,
			BlockDeadline: nil,
		}
		return out, err
	}

	blockStart := blockNumber + 1
	blockDeadline := blockNumber + nBundles

	// Get current pending nonce
	nonce, err := client.PendingNonceAt(context.Background(), zk.FromAddress)

	if err != nil {
		fmt.Printf("Failed to retrieve pending nonce with error: %v\n", err)
		out := DEXTradeReturn{
			Executed:      false,
			BlockDeadline: nil,
		}
		return out, err
	}

	// Submit nBundles copies of the same bundle to Flashbots,
	// one for each target blockNumberToExecute
	if !isMempool {
		for blockToExecute := blockStart; blockToExecute <= blockDeadline; blockToExecute++ {
			blockNumberToExecute := big.NewInt(int64(blockToExecute))

			// Buy: WETH -> Token
			txBuy, ourRawTxBuy, err := GetOurRawTx(false,
				zk,
				chainID,
				client,
				nonce,
				token0,
				instance,
				gasPrice,
				WETHAddress,
				poolAddress,
				big.NewInt(0), // We do not pay any miner bribe on the first tx
				WETHAmountBuy_wei,
				TokenAmountBuy,
				blockNumberToExecute)

			if err != nil {
				fmt.Print(err)
				out := DEXTradeReturn{
					Executed:      false,
					BlockDeadline: nil,
				}
				return out, err
			}

			// Sell: Token -> WETH
			txSell, ourRawTxSell, err := GetOurRawTx(false,
				zk,
				chainID,
				client,
				nonce+1,
				token0,
				instance,
				gasPrice,
				tokenAddress,
				poolAddress,
				minerAmount,
				new(big.Int).Sub(TokenAmountBuy, big.NewInt(1)),
				WETHAmountSell_wei,
				blockNumberToExecute)

			bundle := []string{ourRawTxBuy, ourRawTxSell}

			flashBotsResponse, err := SendFlashbotsBundle(zk, client, flashbotsRelay, bundle, blockNumberToExecute)
			fmt.Printf("Flashbots Response: %v\n", flashBotsResponse)

			if err != nil {
				fmt.Printf("Failed to send bundle with error: %v\n", err)
				out := DEXTradeReturn{
					Executed:      false,
					BlockDeadline: nil,
				}
				return out, err
			}

			fmt.Printf("BlockToExecute: %v || ourHashes: [%v, %v]\n", blockNumberToExecute, txBuy.Hash(), txSell.Hash())

			txs[big.NewInt(int64(blockToExecute))] = [2]common.Hash{txBuy.Hash(), txSell.Hash()}
		}

		out := DEXTradeReturn{
			Executed:      true,
			Txs:           txs,
			BlockDeadline: big.NewInt(int64(blockDeadline)),
		}

		return out, nil
	} else { // Otherwise we submit to mempool
		// Buy: WETH -> Token
		txBuy, _, err := GetOurRawTx(true,
			zk,
			chainID,
			client,
			nonce,
			token0,
			instance,
			gasPrice,
			WETHAddress,
			poolAddress,
			big.NewInt(0), // We do not pay any miner bribe on the first tx
			WETHAmountBuy_wei,
			TokenAmountBuy,
			big.NewInt(0))

		if err != nil {
			fmt.Println(err)
			out := DEXTradeReturn{
				Executed:      false,
				BlockDeadline: nil,
			}
			return out, err
		}

		// Sell: Token -> WETH
		txSell, _, err := GetOurRawTx(true,
			zk,
			chainID,
			client,
			nonce+1,
			token0,
			instance,
			gasPrice,
			tokenAddress,
			poolAddress,
			minerAmount,
			TokenAmountBuy,
			WETHAmountSell_wei,
			big.NewInt(0))

		txs[big.NewInt(int64(blockStart))] = [2]common.Hash{txBuy.Hash(), txSell.Hash()}

		out := DEXTradeReturn{
			Executed:      true,
			Txs:           txs,
			BlockDeadline: big.NewInt(int64(blockDeadline)),
		}

		return out, nil

	}
}

// Given a user's tx Hash, we use an RPC method to get its raw signed data
func GetTargetRawTx(hash string) string {
	req := fmt.Sprintf(`{"jsonrpc": "2.0", "id": 42, "method": "eth_getRawTransactionByHash", "params": ["%v"]}`, hash)
	jsonValue := []byte(req)
	resp, err := http.Post("http://localhost:8545", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	rawTxTarget := json.Get(bodyBytes, "result").ToString()
	return rawTxTarget
}

// Submit transactions using our EOA, either in mempool or Flashbots format
func GetOurRawTx(isMempool bool,
	zk *Zk,
	chainID *big.Int,
	client *ethclient.Client,
	nonce uint64,
	token0 common.Address,
	instance *trader.DEXTrader,
	gasPrice *big.Int,
	inputToken,
	poolAddress common.Address,
	minerAmount,
	inputTokenAmount,
	outputTokenAmount,
	blockNumberToExecute *big.Int) (*types.Transaction, string, error) {

	var tx *types.Transaction
	d := zk.D
	auth, err := bind.NewKeyedTransactorWithChainID(d, chainID)

	if err != nil {
		log.Fatal(err)
		return &types.Transaction{}, "", err
	}

	auth.Context = context.Background()
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	//If Flashbots, we don't send to mempool
	if !isMempool {
		auth.NoSend = true
	} else {
		auth.NoSend = false
	}

	inputIsToken0 := (strings.ToLower(inputToken.String()) == strings.ToLower(token0.String()))

	// Build pairAddressWithData by concatenating (blockNumber, inputIsToken0, poolAddress)
	poolAddressWithDataHex := hex.EncodeToString(poolAddress.Bytes())

	if inputIsToken0 {
		poolAddressWithDataHex = "01" + poolAddressWithDataHex
	} else {
		poolAddressWithDataHex = "00" + poolAddressWithDataHex
	}

	blockNumberToExecuteHex := fmt.Sprintf("%v", blockNumberToExecute.Text(16))
	poolAddressWithDataHex = blockNumberToExecuteHex + poolAddressWithDataHex
	numLeftZeros := 64 - len(poolAddressWithDataHex)

	if numLeftZeros <= 0 {
		log.Fatal("poolAddressWithData contains more than 32 bytes")
		return &types.Transaction{}, "", err
	}

	poolAddressWithDataHex = strings.Repeat("0", numLeftZeros) + poolAddressWithDataHex
	poolAddressWithDataBytes, err := hex.DecodeString(poolAddressWithDataHex)

	if err != nil {
		log.Fatal(err)
		return &types.Transaction{}, "", err
	}

	poolAddressWithData := new(big.Int).SetBytes(poolAddressWithDataBytes)

	// Build amounts by concatenating (outputAmount, minerAmount)
	minerAmountHex := fmt.Sprintf("%v", minerAmount.Text(16))

	numLeftZeros = 32 - len(minerAmountHex)

	if numLeftZeros <= 0 {
		log.Fatal("minerAmount contains more than 16 bytes")
		return &types.Transaction{}, "", err
	}

	minerAmountHex = strings.Repeat("0", numLeftZeros) + minerAmountHex

	outputAmountHex := fmt.Sprintf("%v", outputTokenAmount.Text(16))
	numLeftZeros = 32 - len(outputAmountHex)

	if numLeftZeros <= 0 {
		log.Fatal("outputTokenAmount contains more than 16 bytes")
	}

	outputAmountHex = strings.Repeat("0", numLeftZeros) + outputAmountHex
	amountsHex := outputAmountHex + minerAmountHex
	amountsBytes, err := hex.DecodeString(amountsHex)

	if err != nil {
		log.Fatal(err)
		return &types.Transaction{}, "", err
	}

	amounts := new(big.Int).SetBytes(amountsBytes)

	tx, err = instance.ExecuteTrade(auth,
		poolAddressWithData,
		amounts,
		inputToken,
		inputTokenAmount)

	if err != nil {
		fmt.Println(err)
		return &types.Transaction{}, "", errors.New("Failed to create a tx object in GetOurRawTx")
	}

	// If Flashbots, we return our raw signed tx
	if !isMempool {
		// We sign using our private key
		// Note: we can use any other private key, this is just for relay stats purposes
		ourSignedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), d)

		if err != nil {
			return &types.Transaction{}, "", err
		}

		ourSignedTxRlp, err := rlp.EncodeToBytes(ourSignedTx)

		if err != nil {
			fmt.Printf("Failed to RLP encode transaction with error: %v\n", err)
			return &types.Transaction{}, "", errors.New(fmt.Sprintf("%v\n", err))
		}

		ourRawTxHex := "0x" + hex.EncodeToString(ourSignedTxRlp)

		return tx, ourRawTxHex, nil
	}

	return tx, "", nil
}

// Sends bundle to Flashbots relay, targetting a specific blockNumber to execute
func SendFlashbotsBundle(zk *Zk, client *ethclient.Client, flashbotsRelay string, bundle []string, blockNumberToExecute *big.Int) (string, error) {
	flashbotXHeader := "X-Flashbots-Signature"

	blockNumberHex := fmt.Sprintf("0x%v", blockNumberToExecute.Text(16))

	params := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "eth_sendBundle",
		"params": []interface{}{
			map[string]interface{}{
				"txs":         bundle,
				"blockNumber": blockNumberHex,
			},
		},
	}

	payload, _ := json.Marshal(params)

	jsonPayload := []byte(payload)
	req, err := http.NewRequest("POST", flashbotsRelay, bytes.NewBuffer(jsonPayload))

	if err != nil {
		fmt.Printf("Failed to create new Flashbots request with err: %v\n", err)
	}

	headerValue, err := crypto.Sign(
		accounts.TextHash([]byte(hexutil.Encode(crypto.Keccak256(jsonPayload)))),
		zk.D,
	)

	if err != nil {
		fmt.Printf("Failed to create Flashbots signed payload with err: %v\n", err)
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add(flashbotXHeader, flashbotsHeader(headerValue, zk.D))

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		fmt.Printf("Failed to do Flashbots request in execute_trades: 348 with error: %v\n", err)
		return "", err
	}

	res, _ := ioutil.ReadAll(resp.Body)
	return string(res), nil
}

func flashbotsHeader(signature []byte, privateKey *ecdsa.PrivateKey) string {
	return crypto.PubkeyToAddress(privateKey.PublicKey).Hex() +
		":" + hexutil.Encode(signature)
}

type DEXTradeReturn struct {
	Executed      bool
	Txs           map[*big.Int][2]common.Hash
	BlockDeadline *big.Int
}

// Converts ETHAmount_wei into WETH
// Note: Assumes contract already has ETH
func ConvertETHtoWETH(zk *Zk,
	client *ethclient.Client,
	chainID *big.Int,
	instance *trader.DEXTrader,
	WETHAddress string,
	ETHAmount_wei,
	gasPrice *big.Int) (*types.Transaction, error) {
	d := zk.D
	auth, err := bind.NewKeyedTransactorWithChainID(d, chainID)

	if err != nil {
		log.Fatal(err)
		return &types.Transaction{}, err
	}

	// Get current pending nonce
	nonce, err := client.PendingNonceAt(context.Background(), zk.FromAddress)

	if err != nil {
		fmt.Printf("Failed to retrieve pending nonce with error: %v\n", err)
		return &types.Transaction{}, err
	}

	auth.Context = context.Background()
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(200000)
	auth.GasPrice = gasPrice

	tx, err := instance.ConvertETHtoWETH(auth, ETHAmount_wei)

	if err != nil {
		fmt.Printf("Failed to execute ConvertETHtoWETH with error: %v\n", err)
		return &types.Transaction{}, err
	}

	fmt.Printf("ConvertETHtoWETH tx hash: %v\n", tx.Hash().Hex())
	return tx, nil

}

// Selfdestructs instance and sends any ETH to our EOA
func Kill(zk *Zk,
	client *ethclient.Client,
	chainID *big.Int,
	instance *trader.DEXTrader,
	gasPrice *big.Int) (*types.Transaction, error) {
	d := zk.D
	auth, err := bind.NewKeyedTransactorWithChainID(d, chainID)

	if err != nil {
		log.Fatal(err)
		return &types.Transaction{}, err
	}

	// Get current pending nonce
	nonce, err := client.PendingNonceAt(context.Background(), zk.FromAddress)

	if err != nil {
		fmt.Printf("Failed to retrieve pending nonce with error: %v\n", err)
		return &types.Transaction{}, err
	}

	auth.Context = context.Background()
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(200000)
	auth.GasPrice = gasPrice

	tx, err := instance.Kill(auth)

	if err != nil {
		fmt.Printf("Failed to execute Kill with error: %v\n", err)
		return &types.Transaction{}, err
	}

	fmt.Printf("Kill tx hash: %v\n", tx.Hash().Hex())
	return tx, nil

}

func Init(walletAddressIdx int) *Zk {
	zk := Zk{}
	zk.SetWalletFromMnemonic(walletAddressIdx)
	return &zk
}
