# picasso-bridge

Picasso bridge enables Pixo token distribution on Ethereum based chains (like polygon) for trade volumes by users on Injective chain in frequent intervals. 

### Workflow
1. User places order through picasso frontend. 
2. The order is broadcasted to injective chain where it will be matched with counter order and executed. 
3. Once order is filled, an event is emitted from Injective chain.
4. `picasso-exchange` will subscribe to these events, and persists picasso exchange specific orders into `picasso-orders` collection in mongo db.
5. `picasso-bridge` will poll mongo db in frequent intervals, retrieves new orders placed through picasso ui, and submits them to `DistributionManager` smart contract deployed on Ethereum based chains. (like Polygon)
6. `DistributionManager` smart contract will issue `PIXO` tokens to user address proportional to the trade volumes submitted.

### Steps to setup
1. Clone `git clone https://github.com/PicassoExchange/picasso-bridge`
2. Install - `cd picasso-bridge && make install`
3. Configure - `Create .env file with required config. Refer to .env.example`
4. Run - `picasso-bridge  distribute-tokens`


### Useful commands 
```
$ picasso-bridge --help

picasso-bridge config envname local appLogLevel debug

Usage: picasso-bridge [OPTIONS] COMMAND [arg...]

picasso-bridge is a executable for orchestrating picasso volumes to polygon chain.

Options:
  -e, --env           The environment name this app runs in. Used for metrics and error reporting. (env $PICASSO_BRIDGE_ENV) (default "local")
  -l, --log-level     Available levels: error, warn, info, debug. (env $PICASSO_BRIDGE_LOG_LEVEL) (default "info")

Commands:
  version             Prints picasso-bridge
  distribute-tokens   Start token distribution

Run 'picasso-bridge COMMAND --help' for more information on a command.
```