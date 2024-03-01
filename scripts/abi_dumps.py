import json

directory = "conf/pre-build/"
file_lists = ["NoValidateAccount.json",
    "ArgentAccount.json", "ArgentAccountCairoOne.json",
    "BraavosAccount.json", "BraavosAccountBaseImpl.json", "BraavosCallAggregator.json", "BraavosProxy.json",
    "ERC20.json", "ERC721.json",
    "OpenzeppelinAccount.json", "OpenZeppelinAccountCairoOne.sierra.json",
    "UniversalDeployer.json"
]

target_directory = "conf/genesis/"

for file in file_lists:
    origin_file = directory + file
    with open(origin_file, 'r') as f:
        data = json.load(f)
        data_abi = data['abi']
        # data_abi = json.dumps(data_abi, indent=4)
        data["abi"] = json.dumps(data_abi, indent=4)
    
    data_str = json.dumps(data, indent=1)
    target_file = target_directory + file
    with open(target_file , 'w') as f:
        f.write(data_str)
        f.write('\n')
        f.close()