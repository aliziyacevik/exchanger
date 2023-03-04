#! ./venv/bin/python3.10

"""
    Fetches the initial data and stores in both csv and json files.
"""
import requests
import os
import json

import pandas as pd # for converting json to csv


from dotenv import load_dotenv

#data = None 

def get_symbols():
    API_KEY = os.environ.get("FIXER_IO_KEY")
    url = "https://api.apilayer.com/fixer/symbols"
    
    payload = {}
    headers ={
            "apikey":   API_KEY,
        }
    response = requests.request("GET", url, headers=headers, data=payload)
    if response.status_code != 200:
        pass
    data_text = response.text    #string or text data in JSON format.
    return data_text

def get_currencies(data_text):
    API_KEY = "B8Pl9MAMErjNKChI8h0BEzMyb6Y986nS"
    #API_KEY =  "sK6UdTomBVyeYF1jr03XE4KTWMjNFpV2"

    data = json.loads(data_text)
    data = data['symbols']
    
    url = "https://api.apilayer.com/fixer/latest?base="
    return_text = ""
    for value, desc in data.items():
        payload = {}
        headers= {
            "apikey":   API_KEY,     
        }
        response = requests.request('GET', url+value, headers=headers, data = payload)
        return_text += response.text
        print(return_text)

    #status_code = response.status_code
    #result = response.text
    return return_text

def get_currency(currency_name):
    pass

def save_to_csv(data, csv_name):
    data_frame = pd.read_json(data)
    data_frame = data_frame.drop(columns = 'success')
    data_frame.to_csv(csv_name + '.csv')

    
if __name__ == "__main__":
    load_dotenv()
    
    symbols_data_text = get_symbols()
    currencies_data_text = get_currencies(symbols_data_text)
    
    save_to_csv(symbols_data_text, 'symbols')
    save_to_csv(currencies_data_text, 'currencies')
