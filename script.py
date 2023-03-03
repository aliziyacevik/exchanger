#! ./venv/bin/python3.10
import requests
import os
import json

import pandas as pd # for converting json to csv


from dotenv import load_dotenv

#data = None 

def get_symbols():
    API_KEY = "B8Pl9MAMErjNKChI8h0BEzMyb6Y986nS"#os.environ.get("FIXER_IO_KEY")
    url = "https://api.apilayer.com/fixer/symbols"
    
    payload = {}
    headers ={
            "apikey":   API_KEY
        }
    response = requests.request("GET", url, headers=headers, data=payload)
    if response.status_code != 200:
        pass
    data_text = response.text    #string or text data
    data_json = json.loads(data_text) #json data

    return data_json

def convert_to_csv(data):
    data_frame = pd.read_json(data)
    data_frame.to_csv('symbols.csv')

if __name__ == "__main__":
    load_dotenv()
    data_json = get_symbols()
    convert_to_csv(data_json)

