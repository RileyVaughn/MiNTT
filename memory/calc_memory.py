import numpy as np 
import pandas as pd
import matplotlib.pyplot as plt

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



def Plot(memory_df):

    # Separate rows by whether they contain "SIMD"
    simd_df = memory_df[memory_df.index.str.contains("simd")]
    normal_df = memory_df[~memory_df.index.str.contains("simd")]

    # print(simd_df)
    # print(normal_df)


    colors = plt.cm.tab20.colors  # enough distinct colors

   
   #Normal Plot
    plt.figure(figsize=(10, 5))
    for i, row_label in enumerate(normal_df.index):
        clr = colors[i % len(colors)]
        x = normal_df.columns
        y = normal_df.loc[row_label]

        plt.plot(
            x,
            y,
            label=row_label,
            color=clr,
            linestyle="-",
            marker='o',
            lw=1.25,
            ms=4
        )
    # plt.yscale("log")
    plt.xlabel("Security Parameter (N)")
    plt.ylabel("Peak Memory Utilization (KB)")
    plt.title("MiNTT Peak Memory by Parameterization")
    plt.legend()  
    plt.savefig("normal_memory.png")
    plt.savefig("normal_memory.pdf", format="pdf")

    #SIMD Plot
    plt.figure(figsize=(10, 5))
    for i, row_label in enumerate(simd_df.index):
        clr = colors[i % len(colors)]
        x = simd_df.columns
        y = simd_df.loc[row_label]

        plt.plot(
            x,
            y,
            label=row_label,
            color=clr,
            linestyle="-",
            marker='o',
            lw=1.25,
            ms=4
        )
    # plt.yscale("log")
    plt.xlabel("Security Parameter (N)")
    plt.ylabel("Peak Memory Utilization (KB)")
    plt.title("MiNTT Peak Memory by Parameterization")
    plt.legend()  
    plt.savefig("SIMD_memory.png")
    plt.savefig("SIMD_memory.pdf", format="pdf")
    









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

memory_df = pd.DataFrame.from_dict(memory_calc, orient="index")
memory_df.columns = ["128", "256", "384", "512", "640", "768", "896", "1024"]
Plot(memory_df)
print(memory_df.loc["MiNTT128_norm_int16.su"])
print(memory_df.loc["MiNTT128_simd_int16.su"])
