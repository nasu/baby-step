### HOWTO

```
solc --bin --abi [sol's filename]
geth --dev --datadir . console
geth> var data = "0x[data]"
geth> var abi = [abi]
geth> var factory = eth.contract(abi)
geth> var contract = factory.new({from: eth.accounts[0], data: data, gas: 100000})
geth> contract.[function]
geth> contract.[function].sendTransaction([args...], {from: eth.accounts[0], gas: 100000})
```

or

like following [TBW]

```
geth --dev --datadir . --verbosity 0 console --exec "eth.accounts[0]"
