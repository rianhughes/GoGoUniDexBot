## Go Go Uni Dex Bot
This code finds and trades profitable closed-loop atomic arbitrage paths of length 2 and 3 between any set of pools in UniswapV2 (including forks) and UniswapV3.

While the current state of this code is not meaningfully profitable, it provided me with an exellent opportunity to learn Go, Solidity, and how Ethereum works at a deeper level. In that sense I consider it a success. It was also a lot of fun to build! :)

The code was heavily inspired by cbertmillers simple-arbitrage repo, although deviates in meaningful ways (written in Go instead of TS, considers UniV3, etc.).

## Why Release this code
The purpose of writting this code was to learn the fundamentals about on-chain trading, and there are some other strategies (hopefully less crowded!) that I would like to try out. Also, I've learned a lot from other searchers/traders trading code, so I would like to contribute back to the community.

## Structure of the main go file

The main go file is structured into three parts:

**Part One**: 
Use custom on-chain contract to gather pool data for Uniswap V2 (tokens, pool address, etc). I added UniswapV3 data by hand given there are so few pools.
This section also finds all loops (of length 2 and 3) that start and end in the same asset across these set of pools.
I perform some basic filtering to the pools, eg if they have less than 1 WETH.
 
**Part Two** : 
This section calculates the profit along any path.
For UniswapV2, we can calculate the price impact (up to slippage) of each pool by simply knowing the reseves. So we update the reserves by batching queries to an on-chain contract.
For UniswapV3, the mechanics are very different (given multiple ticks), meaning to calculate the price impact (up to slippage) we would need to know the liquidity for the range of ticks we would cross by trading assets. A simple (and slow) method is to query the Uniswapv3 quoter function. A better way would be to run your own node and query the liquidity and caluclate the price impact directly.
We end this section by selecting the most profitable arbitrage path.

**Part Three**: 
This section take our proiftable arbitrage trade, and creates a transaction with some nice features by using a custom execution contract.
The nice features are that it reverts if it's not profitable, etc.
We then create a flashbots bundle and submit it to the flashbots relay.
Finally, we (hopefully) make a nice profit ;).

## To Run this code you need to:
1. Deploy the BundleExecutor.sol contract, and insert its address into PrivateData.go
2. Insert the private key associated with your funds into PrivateData.go
3. Create a flashbots ID, and insert it into global.go
4. Insert the public address associated with your funds into global.go

## Ways to improve this code and open questions:
1. UniswapV3 pricing needs to be improved. Suggest running your own node and querying directly.
2. This code updates all local reserves every block. This could be improved by only updating when a pool is interacted with. 
3. This code only reselves an arbitrage path if it is present after a block is built. It should be possible to resolve within the block.     This is going to be a lot more competitive than this code (and I suspect others are doing this).
4. The code also neglects optimising the position size of the trade. Simply looks for profit for a fixed input amount.  
5. The network lag is significant, and this code should really be run in/on the node. The speed increase will be significant.
