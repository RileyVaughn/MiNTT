import pandas as pd
import matplotlib.pyplot as plt
import numpy as np
from io import StringIO

#assume all df are the same shape
def Plot(df, name):

    colors = plt.cm.tab20.colors
    n_cols = df.shape[1]
    x = np.arange(n_cols)
    bottom = np.zeros(n_cols)

    plt.figure(figsize=(10, 5))
    for i, row_label in enumerate(df.index):
        
        plt.bar(
            x,
            df.iloc[i].values,
            bottom = bottom,
            color=colors[i],
            label=row_label,
            edgecolor="black"
        )

        bottom+=df.iloc[i].values

    plt.ylim(0,1750)
    plt.xticks(x, df.columns)
    plt.xlabel("Security Parameter (N)")
    plt.ylabel("Runtime (µs)")
    plt.title("Micro-benchmarks of Sub-functions ("+name +")")
    plt.legend()  
    plt.savefig(name+".png")


def PlotAll(dfs, names):
    colors = plt.cm.tab20.colors
    n_cols = dfs[0].shape[1]
    x = np.arange(n_cols)
    fig, axes = plt.subplots(2,2,figsize=(10, 7))
    axes = axes.flatten()

    for ax,df,name in zip(axes,dfs,names):
        bottom = np.zeros(n_cols)

        for i, row_label in enumerate(df.index):
            ax.bar(
                x,
                df.iloc[i].values,
                bottom = bottom,
                color=colors[i],
                label=row_label,
                edgecolor="black"
            )
            bottom+=df.iloc[i].values

        ax.set_ylim(0,1750)
        ax.set_xticks(x, df.columns)
        ax.set_xlabel("Security Parameter (N)")
        ax.set_ylabel("Runtime (µs)")
        ax.set_title(name)

    handles,labels = axes[0].get_legend_handles_labels()
    fig.legend(
        handles,
        labels,
        # loc="upper right"
        bbox_to_anchor=(1.02, 0.9)

    )
    fig.suptitle("Sub-function Micro-benchmarks of MiNTT Variants")
    fig.set_constrained_layout(True)
    fig.set_constrained_layout_pads(w_pad=0.1, h_pad=0.1)
    plt.savefig("micro_bench.png")
    plt.savefig("micro_bench.pdf", format="pdf")




norm128_int64_QF4_data = """16.70779297	31.24890215	45.0614225	59.98210871	75.4237962	85.56695989	101.1916269	113.4228571
1.467382527	2.809449788	4.023592206	4.990240957	6.193122276	7.102520854	8.594338962	9.187315914
38.83916998	72.63684633	106.5025952	140.4968844	173.8154814	198.1190295	232.9290363	258.9805344
17.95726398	69.40277548	144.1275687	272.3401967	397.957593	566.909659	751.2207411	994.9959507
1.141390548	2.12707973	3.108821466	3.827362135	4.810250658	5.483460884	6.685684314	7.271047611"""

simd128_int64_QF4_data = """6.401877132	12.75993185	26.4739387	34.78760794	43.56403772	49.57740324	57.24285078	67.7524928
0.562252756	1.147188713	2.000447477	2.558310975	2.989363117	3.385608073	4.224713271	4.579884267
14.88189341	29.65996068	15.49788768	21.63431259	23.44648916	29.05934929	30.4888224	35.21753328
6.880633356	28.33938553	89.21717206	156.4116396	240.3823057	328.9242873	443.9130501	581.686039
0.437343344	0.86855507	1.76244613	2.156991809	2.658691648	3.064207194	3.653827799	4.042794093"""

norm128_int16_data = """8.086458668	21.48026474	30.63193629	43.25826575	50.3125048	57.10006021	63.09736349	73.61958558
59.80746031	131.3348348	219.2998343	354.5249558	461.163774	582.4619088	708.1236847	900.9625389
21.97931816	21.18138354	30.30616291	42.09773256	48.70926744	56.16007599	60.82165817	71.69554651
8.710038337	47.9593069	102.6889656	195.0093456	280.0344936	388.8729332	489.328444	655.9805419
0.7586490251	2.581165084	3.610008196	4.486553165	5.637070181	6.12675999	6.809690241	7.541487663"""

simd128_int16_data = """15.11639136	6.731130029	10.22148756	14.66861732	18.21188728	21.09270895	24.50043137	28.73204551
1.156274613	57.78976149	97.82225153	159.4736378	215.0004228	271.8455696	342.016902	428.7139639
8.751553884	18.56467813	27.38348403	40.082988	50.39970383	57.79722312	67.51581076	78.6012934
17.363453	15.52277219	33.39161233	67.85610797	96.97174168	140.2495676	183.9453023	258.7566896
1.010279401	0.6446353321	0.859179824	1.190974633	1.533229178	1.87056102	2.092820234	2.353880029"""

norm128_int64_QF4_df = pd.read_csv(StringIO(norm128_int64_QF4_data), sep="\s+", header=None)
simd128_int64_QF4_df = pd.read_csv(StringIO(simd128_int64_QF4_data), sep="\s+", header=None)
norm128_int16_df = pd.read_csv(StringIO(norm128_int16_data), sep="\s+", header=None)
simd128_int16_df = pd.read_csv(StringIO(simd128_int16_data), sep="\s+", header=None)


given_order = ["NTT Lookup Tables", "Modulo", "NTT (other)", "Key Calculation", "Base Change"]
new_order = ["Key Calculation","NTT (other)","NTT Lookup Tables","Modulo","Base Change"]
labels = ["128", "256", "384", "512", "640", "768", "896", "1024"]

norm128_int64_QF4_df.index = given_order
norm128_int64_QF4_df.columns = labels
norm128_int64_QF4_df = norm128_int64_QF4_df.loc[new_order]

simd128_int64_QF4_df.index = given_order
simd128_int64_QF4_df.columns = labels
simd128_int64_QF4_df = simd128_int64_QF4_df.loc[new_order]

norm128_int16_df.index = given_order
norm128_int16_df.columns = labels
norm128_int16_df = norm128_int16_df.loc[new_order]

simd128_int16_df.index = given_order
simd128_int16_df.columns = labels
simd128_int16_df = simd128_int16_df.loc[new_order]


names = ["n=128 q=65537 int64","n=128 q=65537 int64 SIMD","n=128 q=257 int16 ","n=128 q=257 int16 SIMD"]
dfs=[norm128_int64_QF4_df,simd128_int64_QF4_df,norm128_int16_df,simd128_int16_df]

# Plot(dfs[0], names[0])
# Plot(dfs[1], names[1])
# Plot(dfs[2], names[2])
# Plot(dfs[3], names[3])
PlotAll(dfs,names)