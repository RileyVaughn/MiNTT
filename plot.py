import pandas as pd
import matplotlib.pyplot as plt
import numpy as np
from io import StringIO

# #Takes a df of runtimes and plots graph
# def Plot(df):

#     colors = plt.cm.tab20.colors

#     for i, row_label in enumerate(df.index):
#         plt.plot(
#             df.columns, 
#             df.loc[row_label], 
#             label=row_label, 
#             color=colors[i % len(colors)], 
#             linestyle="--",   # hyphenated line
#             marker='o'
#         )

#     plt.xlabel("Security Parameter (N)")
#     plt.ylabel("Throughput Rate (MB/s)")
#     plt.title("MiNTT Throughput Rate by Parameterization")
#     plt.legend()  
#     # plt.grid(True)
#     plt.show(

def PlotSeparate(df):
    # Separate rows by whether they contain "SIMD"
    simd_df = df[df.index.str.contains("SIMD")]
    normal_df = df[~df.index.str.contains("SIMD")]

    colors = plt.cm.tab20.colors  # enough distinct colors

    # --- Plot Normal Data ---
    plt.figure(figsize=(10, 5))
    for i, row_label in enumerate(normal_df.index):
        plt.plot(
            normal_df.columns,
            normal_df.loc[row_label],
            label=row_label,
            color=colors[i % len(colors)],
            linestyle="--",
            marker='o'
        )
    plt.xlabel("Security Parameter (N)")
    plt.ylabel("Throughput Rate (Mbps)")
    plt.title("MiNTT Throughput Rate by Parameterization")
    plt.legend()  

    plt.show()

    # --- Plot SIMD Data ---
    plt.figure(figsize=(10, 5))
    for i, row_label in enumerate(simd_df.index):
        plt.plot(
            simd_df.columns,
            simd_df.loc[row_label],
            label=row_label,
            color=colors[i % len(colors)],
            linestyle="--",
            marker='o'
        )
    plt.xlabel("Security Parameter (N)")
    plt.ylabel("Throughput Rate (Mbps)")
    plt.title("MiNTT Throughput Rate by Parameterization")
    plt.legend()  
    plt.show()
















raw_data = """24.711	16.957	14.063	11.671	10.325	8.845	8.035	7.072
53.112	34.478	27.840	21.811	18.447	15.669	13.989	12.651
24.076	21.262	18.785	14.836	13.923	12.803	11.951	10.569
60.514	49.382	41.206	32.429	29.141	26.568	23.591	17.854
7.992	5.261	4.354	3.202	2.933	2.528	2.267	2.054
13.926	8.838	6.933	5.520	4.769	4.010	3.569	3.036
9.009	7.353	6.414	5.022	4.570	3.955	3.812	3.331
19.180	14.489	11.308	9.061	8.005	6.080	6.470	5.270
2.478	1.266	0.823	0.628	0.495	0.413	0.354	0.312
3.581	1.752	1.182	0.884	0.715	0.603	0.514	0.449
24.259	20.948	18.473	15.538	12.754	12.796	11.870	10.732
62.360	51.861	41.535	34.658	30.163	23.002	20.480	21.602
24.249	17.874	14.575	11.792	10.358	8.568	7.925	6.838
56.456	36.437	27.425	22.575	19.090	15.693	14.457	12.895"""

# Clean thousand separators (commas inside numbers)
clean_data = raw_data.replace(",", "")

# Load into DataFrame
df = pd.read_csv(StringIO(clean_data), sep="\s+", header=None)

df.columns = ["128", "256", "384", "512", "640", "768", "896", "1024"]
df.index = [
 'n=64 q=257 int64',
 'n=64 q=257 int64 SIMD',
 'n=128 q=257 int64',
 'n=128 q=257 int64 SIMD',
 'n=64 q=257 int16',
 'n=64 q=257 int16 SIMD',
 'n=128 q=257 int16',
 'n=128 q=257 int16 SIMD',
 'n=8 q=257 int16',
 'n=8 q=257 int16 SIMD',
 'n=128 q=65537 int64',
 'n=128 q=65537 int64 SIMD',
 'n=64 q=65537 int64',
 'n=64 q=65537 int64 SIMD'
]


PlotSeparate(df)
