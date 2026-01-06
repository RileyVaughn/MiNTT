import numpy as np 

#Searhces for hash and subfucntions memory costs
def SearchSU(fileName, functions):

    mem_dict = {}


    with open(fileName, "r") as f:
        lines = f.read().splitlines()

    for line in lines:
        for func in functions:
            if func in line:
                mem_dict[func] = int(line.split('\t')[1].split('\t')[0])
        
    return mem_dict


def CalcMemory(label,diff):

    functions=["Hash","ntt_sum","ncc","change_base"]
    memory_dict = SearchSU("./" + diff + "/" + label, functions)

    if "norm" in label:
        norm_utils = ["Norm_AddSub","Norm_Mod257"]
        if "int64" in label:
            int_type=64
            memory_dict = memory_dict | SearchSU("./util/util_int64.su", norm_utils)
        else:
            int_type=16
            memory_dict = memory_dict | SearchSU("./util/util_int16.su", norm_utils)
    else:
        simd_utils = ["SIMD_Mod257"]
        if "int64" in label:
            int_type=64
            simd_utils.append("SIMD_Q_reduce")
            memory_dict = memory_dict | SearchSU("./util/util_int64.su", simd_utils)
        else:
            int_type=16
            simd_utils.append("SIMD_Center257")
            memory_dict = memory_dict | SearchSU("./util/util_int16.su", simd_utils)

    N= int(diff.split('_')[1])
    if "MiNTT64" in label:
        n = 64
    elif "MiNTT128" in label:
        n=128
    else:
        n=8
    d=N//n
    if "QF4" in label:
        q = 65537
        m=2*d*17
    else:
        q = 257
        m=2*d*9
    in_size = n*m//8
    out_size=in_size//2

    A = m*N*int_type//8
    NTT8 = 8*256*int_type//8
    if n==8:
        mult = 0
    else:
        mult = n * int_type//8

    memory_dict["A_table"] = A
    memory_dict["NTT8_table"] = NTT8
    memory_dict["Mult_table"] = mult
    memory_dict["Input_size"] = in_size
    memory_dict["Output_size"] = out_size

    #print(memory_dict)
    return sum(memory_dict.values())/1000



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

memory_calc = {}
for label in labels:
    mem_per_label = []
    for diff in diffs:
         mem_per_label.append(CalcMemory(label,diff))
    memory_calc[label] = mem_per_label
print(memory_calc)
