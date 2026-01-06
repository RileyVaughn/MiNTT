import numpy as np 

#Searhces for hash and subfucntions memory costs
def SearchSU(fileName, functions):

    mem_dict = {}


    with open(fileName, "r") as f:
        lines = f.read().splitlines()

    for line in lines:
        for func in functions:
            if func in line:
                mem_dict[func] = line.split('\t')[1].split('\t')[0]
        
    return mem_dict


def CalcMemory(label,diff):

    functions=["Hash","ntt_sum","ncc","change_base"]
    memory_dict = SearchSU("./" + diff + "/" + label, functions)

    if "norm" in label:
        norm_utils = ["Norm_AddSub","Norm_Mod257"]
        if "int64" in label:
            memory_dict = memory_dict | SearchSU("./util/util_int64.su"+, norm_utils)
        else:
            memory_dict = memory_dict | SearchSU("./util/util_int16.su"+, norm_utils)
    else:
        simd_utils = ["Norm_AddSub","Norm_Mod257"]
        if "int64" in label:
            memory_dict = memory_dict | SearchSU("./util/util_int64.su"+, simd_utils)
        else:
            memory_dict = memory_dict | SearchSU("./util/util_int16.su"+, simd_utils)

    
    print(memory_dict)



labels = [
    "MiNTT64_norm_int64.su",
    "MiNTT64_simd_int64.su",
    "MiNTT128_norm_int64.su",
    "MiNTT128_simd_int64.su",
    "MiNTT64_norm_int16.su",
    "MiNTT64_simd_int16.su",
    "MiNTT128_norm_int16.su",
    "MiNTT128_simd_int16.su",
    "MiNTT8_norm_int16.su",
    "MiNTT8_simd_int16.su",
    "MiNTT128_norm_int64_QF4.su",
    "MiNTT128_simd_int64_QF4.su",
    "MiNTT64_norm_int64_QF4.su",
    "MiNTT64_simd_int64_QF4.su"
    ]

diffs = [
    "N_128",
    "N_256",
    "N_384",
    "N_512",
    "N_640",
    "N_768",
    "N_896",
    "N_1024"
]



CalcMemory(labels[0],diffs[0])