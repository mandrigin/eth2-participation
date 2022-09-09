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

    for node in response_json['data']:
        idx = int(node['index'])
        if idx >= f and idx < t:
            balance = int(node['balance'])
            balance_diff = balance - INITIAL_BALANCE
            if balance_diff < 0:
                validators_negative.append(idx)
                validators_negative_balances.append(balance_diff)
            else:
                validators_positive.append(idx)
                validators_positive_balances.append(balance_diff)


    positive = len(validators_positive)
    negative = len(validators_negative)
    print("no validators_positive", positive)
    stats(validators_positive_balances)
    print("no validators_negative", negative)
    stats(validators_negative_balances)
    print("%% negative validators", negative / (positive + negative) * 100)

print_stats(0, 10000, "TOTAL")
print_stats(0, 3000, "Nethermind")
print_stats(3000, 6000, "Gateway")
print_stats(6000, 10000, "Gnosis")

