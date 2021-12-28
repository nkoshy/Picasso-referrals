# This folder contains
- Smart contracts for smart contracts for Picasso referal program.
- Test cases 
- Deployment scripts
- Go wrapper for solidity contracts

# Setup
We are using hardhat for compiling, testing and deploying smart contracts.

## Install Hardhat
- https://hardhat.org/tutorial/setting-up-the-environment.html

## Install Dependencies
- cd solidity && npm install

## Compile smart contracts
- cd solidity && npx hardhat compile

## Run Tests
- cd solidity && npx hardhat test

## Deploy Smart contracts
- cd solidity && npx hardhat run scripts/deploy.js --network matic

## Generate wrappers for solidity contracts
- Install abigen - https://www.metachris.com/2021/05/creating-go-bindings-for-ethereum-smart-contracts/
- make gen