
import time
import contract_abi
from web3 import Web3, HTTPProvider

contract_address     = "0x93fb5919A628C90f0Bd40f9e1dee61E256e37E12"
wallet_private_key   = "403a6d93a86b9c8e876408903de9588ef4b31763366bc56fbfdde006b846ed7d"
wallet_address       = "0xdD698f9193fD7398eF00D06365c49BD93F771520"

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