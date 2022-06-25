import requests

# Ledger test
url = 'http://localhost:8080/api/ledgers'
ledger = {'title': "New Ledger", 'members':['6ba7b810-9dad-11d1-80b4-00c04fd430c8']}

x = requests.post(url, json = ledger)
print(x)

