import pandas as pd
import matplotlib.pyplot as plt
import numpy as np
from io import StringIO


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
















raw_data = """28.89607946	20.0384418	16.85911226	13.66961536	12.07036014	10.37576107	9.406034139	8.36636243
62.93706294	39.84366894	32.16589261	24.79752455	21.77142274	18.47812139	16.52926442	14.4649567
28.59980139	24.27383925	21.63475082	17.56472418	14.33938027	14.29605576	12.86875592	12.31958742
69.88595001	56.94231625	48.18739543	36.55547182	33.59540863	30.21058245	27.73182936	23.89998107
11.38418665	8.1437312	7.038452844	5.824307637	5.225030071	4.414855939	4.073126357	3.682874752
21.82934455	15.51066021	13.47032319	10.96052024	9.981371572	8.711861971	7.997421466	6.859077076
11.55466399	10.26111509	8.940929329	7.207015579	6.80966168	6.33708681	6.071461646	5.390100953
26.50774293	23.21340413	20.36799114	16.26705075	15.07391715	14.02438035	13.00496233	11.5610707
4.562069088	2.513958251	1.717023072	1.306247594	1.042101064	0.8685841701	0.7453610426	0.6563674336
9.469556854	4.832893884	3.25324171	2.435032766	1.95894247	1.635189505	1.398669564	1.228901589
28.3152676	24.41857203	21.55707606	18.0717013	16.52993011	15.12544878	13.83944715	12.57933443
74.35249095	59.80075575	48.37275476	40.0093772	34.75583071	31.53539399	28.23234598	25.10966004
29.06176962	20.83403147	17.11943481	13.97043481	12.18818194	10.44645472	9.497982489	8.499572531
65.77197437	43.00352763	33.27725952	26.19312009	21.72862387	19.02042481	16.87773064	14.45333386"""

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
