#! ./venv/bin/python3.10

"""
    Fetches the initial data and stores in both csv and json files.
"""
import requests
import os
import json
import csv
import sys

from dotenv import load_dotenv

def fetch(URL, payload, headers):
    print("trying to fetch data from: ", URL)
    response = requests.request("GET", URL, headers=headers, data=payload)
    if response.status_code != 200:
        print("Error:   response status code:   ", response.status_code)
        raise SystemExit("Failed trying to fetch from: " + URL)
    print("fetched succesfully from: ", URL)
    return response

def get_symbols():
    API_KEY = os.environ.get("FIXER_IO_KEY")
    url = "https://api.apilayer.com/fixer/symbols"
    
    payload = {}
    headers ={
            "apikey":   API_KEY,
        }
    response = fetch(url, payload, headers)
   
   try:
        data_json = json.loads(response.text)
        data_json.pop('success', None)
        data_json = data_json['symbols']
        data_text = json.dumps(data_json).replace('{', '').replace("}",'').replace('"', '').replace(',','\n').replace(":",",").replace(" ", '').replace("\t", '')
   except:
       raise SystemExit("fetch_initial_data.get_symbols: Error occured while adjusting symbols data")
    
   print("symbols fetched..")
    return data_text

def get_currency(currency_name):
    url = "https://api.apilayer.com/fixer/latest?base=" + currency_name
    API_KEY = os.environ.get("FIXER_IO_KEY")

    payload = {}
    headers ={"apikey": API_KEY}

    response = fetch(url, payload, headers)
    print("currency: ","'",currency_name,"'", "fetched" )
    return response.text 


def get_currencies(filename):
    with open(filename, 'r') as file:
        reader = csv.reader(file)
        count = 0
        return_text = "["
        for row in reader:
            base = row[0]
            if base != 'Value':
                response_text = get_currency(base) 
   
                try:
                    data_json = json.loads(response_text)
                    data_json.pop('success', None)
                    data_json.pop('timestamp', None)
                    data_json.pop('date', None)
        
                    response_text = json.dumps(data_json)
                    return_text += response_text 
                except:
                    raise SystemError("fetch_initial_data.get_currencies: Error while adjusting currencies")
                if count == 5:
                    break
                else:
                    return_text += ","
                count += 1
    
    print(return_text)
    print(count ," currencies have been fetched")
    return return_text + ']'

def save_to(data_text, file_name, formatt):
    fullname = file_name + '.' + formatt
    if formatt == 'csv':
        with open(fullname, "w") as file:
            headers = ['Value', 'Description']
            writer = csv.writer(file)
            writer.writerow(headers)
            
            rows = data_text.split('\n')
            for row in rows:
                row = row.split(',')
                writer.writerow(row)

    elif formatt == 'json':i
        data_json = json.loads(data_text)
        
        with open(fullname, 'w') as f:
                #outfile.write(data_text)
                json.dump(data_json,f)
    else:
        raise SystemExit("only csv and json formats are supported")
    print("File saved to:", fullname)

if __name__ == "__main__":
    load_dotenv()
    
    filename = 'symbols'
    symbols_as_text = get_symbols()
    save_to(symbols_as_text, filename, 'csv')
    
    currencies_as_text = get_currencies(filename + '.csv') 
    save_to(currencies_as_text, 'currencies', 'json')

