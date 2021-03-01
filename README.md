# altcoinmax

> This program calculates the maximum satoshi value of any altcoin using 
> Coingecko.com api data.

## How it works

This program works off the assumption that Bitcoin will always be the highest 
marketcap coin in cryptocurrency. Therefore the highest value of any altcoin 
will have to equal (`coin_supply/bitcoin_supply`).

To run this application, for now simply clone the repo and run:

```
go run main.go
```

You'll get output similar to:
```
--------------------------------------------------------
Coin Statistics for Ethereum (eth):
Current Price: 0.03142684
All Time High (could be inaccurate): 0.14749840
Maximum satoshi value by circulating supply: 0.18280811
Maximum satoshi value by total supply: 0.00000000
--------------------------------------------------------
Coin Statistics for Cardano (ada):
Current Price: 0.00002677
All Time High (could be inaccurate): 0.00007382
Maximum satoshi value by circulating supply: 0.00065729
Maximum satoshi value by total supply: 0.00046667
--------------------------------------------------------
Coin Statistics for Binance Coin (bnb):
Current Price: 0.00511142
All Time High (could be inaccurate): 0.00609080
Maximum satoshi value by circulating supply: 0.13589273
Maximum satoshi value by total supply: 0.12314285
--------------------------------------------------------
```