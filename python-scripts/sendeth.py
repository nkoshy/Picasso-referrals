
import time
import contract_abi
from web3 import Web3, HTTPProvider

contract_address     = "0x3A45F15e1d9a9447d5dEB82248291cC0737Fe7eD"
wallet_private_key   = "48c2341e77218d5a094d50b2b7ee49bd3e4a189fda134cec6501d5f222e037ca"
wallet_address       = "0x3cB34536F22003E9c2733d67f28A1148dB45F5Ff"

w3 = Web3(HTTPProvider("https://rpc-mumbai.maticvigil.com/v1/3b64186f6a398fe17e87f323298a22404746cb5c"))


contract = w3.eth.contract(address = contract_address, abi = contract_abi.abi)

def getFeeRecepient(userAddress):
    fee_receipient = contract.functions.getFeeRecepient(userAddress).call()
    print(fee_receipient)
    return fee_receipient

def refer_user(referalCode, userAddress):
    nonce = w3.eth.getTransactionCount(wallet_address)
    txn_dict = contract.functions.referUser(referalCode, userAddress).buildTransaction({
        'chainId': 80001,
        'gas': 140000,
        'gasPrice': w3.toWei('40', 'gwei'),
        'nonce': nonce,
    })

    signed_txn = w3.eth.account.signTransaction(txn_dict, private_key=wallet_private_key)
    result = w3.eth.sendRawTransaction(signed_txn.rawTransaction)
    tx_receipt = w3.eth.getTransactionReceipt(result)

    count = 0
    while tx_receipt is None and (count < 30):
        time.sleep(10)
        tx_receipt = w3.eth.getTransactionReceipt(result)
        print(tx_receipt)

    if tx_receipt is None:
        return {'status': 'failed', 'error': 'timeout'}

    processed_receipt = contract.events.OpinionBroadcast().processReceipt(tx_receipt)
    print(processed_receipt)

    # output = "Address {} Referred: {}"\
    #     .format(processed_receipt[0].args._soapboxer, processed_receipt[0].args._opinion)
    # print(output)

    return {'status': 'added', 'processed_receipt': processed_receipt} 

# refer_user("ABC","0xF955C57f9EA9Dc8781965FEaE0b6A2acE2BAD6f3")
getFeeRecepient("0xF955C57f9EA9Dc8781965FEaE0b6A2acE2BAD6f3")