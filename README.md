# Uniswap Flashbots Trader

How to place 0 slippage trades on UniswapV2 via Flashbots.

Designed with gas efficiency in mind, DO NOT USE VIA THE MEMPOOL!

## Note:

I may still tidy up the code and docs a bit more, but this repository is designed to be self-contained and specific.

## This repository may be useful as:
- A building block for a MEV searcher (contract + go package to submit Flashbot compatible bundles).
- A contract in which you hold your funds and use to trade with front-running protection.

## Some recommendations:

- Whitelist tokens before trading on them! Many tokens are either scams or Salmonellas, meaning that we do not always get what we expect as `amountOutMin`. See https://github.com/BitBaseBit/LessRekt for a starting point, but be aware that some Salmonella implementations are able to bypass simulation. One way to be safer is to add extra heuristics such as minumum transaction count or minimum traded volume for the pool of interest.
- Always leave balance 1 after selling a token, in order to save gas every other time you trade with it in the future.
- If you want to save gas, you can pay miner bribes in gasFee rather than ETH transfer to miner. But beware that in the case the block gets uncled, a miner may still take your reverting transaction and make you lose some ETH in gasFee!
- Get hands on into coding! No one will give you the full bits of a successful MEV searcher bot for free.

## Homework:

Look into `FlashbotsTrader.sol` and find ways to make `executeTrade` even more gas-efficient. Look into the Flashbots Discord and connect all dots!
