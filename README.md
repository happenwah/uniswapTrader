# Uniswap Flashbots Trader

How to place 0 slippage trades on UniswapV2 via Flashbots.

Designed with gas efficiency in mind, DO NOT USE ON THE MEMPOOL!

This is meant to be useful as either:
- A building block for a MEV searcher (contract + go package to submit Flashbot compatible bundles).
- A contract in which you hold your funds and use to trade with front-running protection.

## Some recommendations:

- Whitelist tokens before trading on them! Many of them are either scams or Salmonellas, meaning that we do not always get what we expect as amountOut.
- Always leave balance 1 after selling a token, in order to save gas every other time you trade in the future.
- If you want to save gas, you can pay miner bribes in gasFee rather than ETH transfer to miner. But beware that in the case the block gets uncled, a miner may still take your reverting transaction and make you lose some ETH in gasFee.
