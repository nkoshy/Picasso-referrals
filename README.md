# picasso-referal

Picasso Referal enables user referals on picasso exchange.

### Workflow
## Influencer Onboarding
1. Picasso admin sends a `createInfluencer` transaction to `PicassoReferal.sol` on Polygon chain to Add/Create Influencer with influencer details like referalCode, influencerName, influencerAddress. 
2. By default, referals are disabled for new influencers. 
3. Picasso admin sends a `enableRefereals` transaction to `PicassoReferal.sol` to enable referals for a referal code.
4. Onboarding of Influencer is completed.

## Refer User
1. Influencer shares referal link with his users.
2. when user clicks referal link, he is navigated to picasso exchange website and asked to connect wallet. 
3. Once user clicks connect wallet, a POST call is send to AWS lambda function.
4. AWS lamda function sends a `referUser` transaction to `PicassoReferal.sol` passing `referalCode` and `userAddress`. Upon which the smart contract will add influencer address as the fee_reciepeint for this user.

## Set fee_recipient to influencer 
1. When user clicks on `ConnectWallet` on picasso exchange, we make a `getFeeRecepient` call with userAddres to `PicassoReferal.sol` and set it on exchange. 


### Steps to setup
1. Clone `git clone https://github.com/PicassoExchange/picasso-referal`
2. Install - `cd picasso-referal && make install`
3. Configure - `Create .env file with required config. Refer to .env.example`
4. Run - `picasso-referal refer-user`


### Useful commands 
```
 picasso-referals refer-user --help
picasso-referal config envname local appLogLevel debug

Usage: picasso-referal refer-user [OPTIONS]

Start user referals

Options:
      --eth-chain-id               Specify Chain ID of the Ethereum network. (env $ETH_CHAIN_ID) (default 80001)
      --eth-node-http              Specify HTTP endpoint for an Ethereum node. (env $ETH_RPC) (default "https://rpc-mumbai.matic.today")
      --eth_gas_price_adjustment   gas price adjustment for Ethereum transactions (env $ETH_GAS_PRICE_ADJUSTMENT) (default 1.3)
      --eth-keystore-dir           Specify Ethereum keystore dir (Geth-format) prefix. (env $ETH_KEYSTORE_DIR)
      --eth-from                   Specify the from address. If specified, must exist in keystore, ledger or match the privkey. (env $ETH_FROM)
      --eth-passphrase             Passphrase to unlock the private key from armor, if empty then stdin is used. (env $ETH_PASSPHRASE)
      --eth-pk                     Provide a raw Ethereum private key of the validator in hex. USE FOR TESTING ONLY! (env $ETH_PK)
      --eth-use-ledger             Use the Ethereum app on hardware ledger to sign transactions. (env $ETH_USE_LEDGER)
      --picasso-referal-contract   Specify the Picasso Referal contract address. (env $ETH_PICASSO_REFERAL_CONTRACT)
```