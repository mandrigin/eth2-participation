# usage python3 balance_history.py <validator_idx> <from_slot> <to_slot>
import requests
import sys


INITIAL_BALANCE = 32000000000

validator_idx = sys.argv[1]

current_state = int(sys.argv[2])

to_state = int(sys.argv[3])

balance_history = {}

print("Stats for validator", validator_idx)

prev_balance = INITIAL_BALANCE

while current_state <= to_state:

    response = requests.get('http://localhost:4000/eth/v1/beacon/states/'+str(current_state)+'/validator_balances?id='+str(validator_idx))
    response_json = response.json()

    balance = int(response.json()['data'][0]['balance'])

    diff = balance-prev_balance

    if diff != 0:
        print("slot", current_state, "balance", balance, "diff", diff)

    prev_balance = balance
    current_state+=1

