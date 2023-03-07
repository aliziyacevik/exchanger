#! ./venv/bin/python3.10

import sys


ERROR_MSG = "Error in: scripts/initialize_env"
USAGE = "Usage: ./path/to/initialize_env.py --CONFIG_NAME=CONFIG_VALUE. " 
def parse(args):
    
    if len(args) == 1:
        raise SystemExit(USAGE)
    
    m = {}
    args.pop(0)
    for arg in args:
        count = 0
        if len(arg) < 2:
            raise SystemExit(ERROR_MSG +'.parse()' +'Error while parsing.')
        if arg[0:2] != '--':
            raise SystemExit(ERROR_MSG +'.parse()' +'Error while parsing.')
        arg = arg[2:]
        if arg == arg.split("="):
            raise SystemExit(ERROR_MSG + '.parse()' + USAGE+ "Did you forget to put '=' sign ?") 
        if len(arg.split("=")) != 2:
            raise SystemExit(ERROR_MSG + '.parse()' + USAGE)
        
        arg.strip(",'/\.-*+`")
        arg = arg.split("=")
        
        m[arg[0]]= arg[1]
    return m    

def write_to_env(m):
    with open(".env", "x") as f:
        for key, value in m.items():
            arg = key + '=' + value
            f.write(arg)
            f.write('\n')

if __name__ == "__main__":
    arguments_to_write = parse(sys.argv)
    write_to_env(arguments_to_write)
