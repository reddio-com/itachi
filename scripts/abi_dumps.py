import json
import shutil

directory = "conf/pre-build/"
cairo0_file_lists = ["NoValidateAccount.json",
    "ArgentAccount.json",
    "BraavosAccount.json", "BraavosAccountBaseImpl.json", "BraavosCallAggregator.json", "BraavosProxy.json",
    "ERC20.json", "ERC721.json",
    "OpenzeppelinAccount.json",
    "UniversalDeployer.json"
]

cairo1_file_lists = ["ArgentAccountCairoOne.json",
    "OpenZeppelinAccountCairoOne.sierra.json"
]

target_directory = "conf/genesis/"

for file in cairo0_file_lists:
    #directly copy the file to the target directory
    origin_file = directory + file
    target_file = target_directory + file
    shutil.copyfile(origin_file, target_file)

for file in cairo1_file_lists:
    origin_file = directory + file
    with open(origin_file, 'r') as f:
        data = json.load(f)
        data_abi = data['abi']
        data["abi"] = json.dumps(data_abi, indent=2)
    
    data_str = json.dumps(data, indent=1)
    target_file = target_directory + file
    with open(target_file , 'w') as f:
        f.write(data_str)
        f.write('\n')
        f.close()

