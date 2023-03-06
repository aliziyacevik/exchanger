#! ./venv/bin/python3.10

"""
    Fetches the initial data and stores in both csv and json files.
"""
import requests
import os
import json
import sys

import pandas as pd # for converting json to csv


from dotenv import load_dotenv

#data = None 

def fetch(URL, payload, headers):
    print("trying to fetch data from: ", URL)
    response = requests.request("GET", URL, headers=headers, data=payload)
    if response.status_code != 200:
        print("Error:   response status code:   ", response.status_code)
        print("Failed trying to fetch from: ", URL)
        raise SystemExit()
    print("fetched succesfully from: ", URL)
    return response


def get_symbols():
    API_KEY = os.environ.get("FIXER_IO_KEY")
    API_KEY = "B8Pl9MAMErjNKChI8h0BEzMyb6Y986nS"
    url = "https://api.apilayer.com/fixer/symbols"
    
    payload = {}
    headers ={
            "apikey":   API_KEY,
        }
    response = fetch(url, payload, headers)
    
    data_frame = pd.read_json(response.text)
    data_frame = data_frame.drop(columns = 'success')
   
    print("symbols fetched..")
    return data_frame.to_json()

def get_currency(currency_name):
    url = "https://api.apilayer.com/fixer/latest?base=" + currency_name
    API_KEY = os.environ.get("FIXER_IO_KEY")

    payload = {}
    headers ={"apikey": API_KEY}

    response = fetch(url, payload, headers)
    print("currency: ","'",currency_name,"'", "fetched" )
    return response.text 


def get_currencies(data_as_text):
    data = json.loads(data_as_text)
    data = data['symbols']
    
    response_text = ""
    count = 0
    for value, desc in data.items():
        count += 1
        response_text += get_currency(value) 
        if count == 1:
            break
    data_frame = pd.read_json(response_text)
    data_frame = data_frame.drop(columns = ['success', 'timestamp', 'date'])

    print(count ," currencies have been fetched")
    return data_frame.to_json()

def save_to(data, file_name, formatt):
    fullname = file_name + '.' + formatt
    if formatt == 'csv':
        data_frame = pd.read_json(data)
        data_frame.to_csv(fullname)
    elif formatt =='json':
        data_frame = pd.read_json(data)
        data_frame.to_json(fullname)
    else:
        print("only json and csv formats are supported.")
        raise SystemExit()

    print("File saved to:", fullname)

if __name__ == "__main__":
    load_dotenv()
    
    symbols_as_text = get_symbols()
    currencies_as_text = get_currencies(symbols_as_text)
    
    save_to(symbols_as_text, 'symbols', 'csv')
    save_to(currencies_as_text, 'currencies', 'json')
