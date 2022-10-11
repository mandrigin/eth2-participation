'''
{
  "execution_optimistic": false,
  "data": [
    {
      "index": "0",
      "balance": "31992794164"
    }
'''


import requests
import numpy as np

INITIAL_EPOCH = 'finalized' # 21000

def stats(arr):
    if len(arr) == 0:
        return "empty array"
    aa = np.average(arr)
    print("avg value of arr : ", aa)
    pp = np.percentile(arr, 50)
    print("50th percentile of arr : ", pp)
    if pp > 0 :
        print("99th percentile of arr : ",
               np.percentile(arr, 99))
    else:
        print("99th percentile of arr: ",
               np.percentile(arr, 1))

INITIAL_BALANCE = 32000000000

initial_balances = {}

initial_response = requests.get('http://localhost:4000/eth/v1/beacon/states/'+ str(INITIAL_EPOCH) +'/validator_balances')

for node in initial_response.json()['data']:
    idx = int(node['index'])
    balance = int(node['balance'])
    initial_balances[idx] = balance

response = requests.get('http://localhost:4000/eth/v1/beacon/states/head/validator_balances')

response_json = response.json()



def print_stats(f, t, name):
    print("---")
    print(name)
    print("---")

    validators_negative = []
    validators_negative_balances = []
    validators_positive = []
    validators_positive_balances = []
    all_balances = []

    for node in response_json['data']:
        idx = int(node['index'])
        if idx >= f and idx < t:
            balance = int(node['balance'])

            all_balances.append(balance)

            balance_diff = balance - initial_balances.get(idx, INITIAL_BALANCE)

            if balance_diff < 0:
                validators_negative.append(idx)
                validators_negative_balances.append(balance_diff)
            else:
                validators_positive.append(idx)
                validators_positive_balances.append(balance_diff)


    positive = len(validators_positive)
    negative = len(validators_negative)
    print("## validators_positive", positive)
    stats(validators_positive_balances)
    print("## validators_negative", negative)
    stats(validators_negative_balances)
    print("%% negative validators", negative / (positive + negative) * 100)
    print ("ALL BALANCES")
    stats(all_balances)

print_stats(0, 10000, "TOTAL")
print_stats(0, 2000, "Nethermind")
print_stats(2000, 4000, "Gateway")
print_stats(4000, 6000, "Gnosis")

