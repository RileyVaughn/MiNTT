import pandas as pd
import matplotlib.pyplot as plt
import numpy as np
from io import StringIO


def Plot(mean_df,stdev_df):
    # Separate rows by whether they contain "SIMD"
    simd_df = mean_df[mean_df.index.str.contains("SIMD")]
    normal_df = mean_df[~mean_df.index.str.contains("SIMD")]

    simd_stdev_df = stdev_df[stdev_df.index.str.contains("SIMD")]
    normal_stdev_df = stdev_df[~stdev_df.index.str.contains("SIMD")]

    colors = plt.cm.tab20.colors  # enough distinct colors

   
   #Normal Plot
    plt.figure(figsize=(10, 5))
    for i, row_label in enumerate(normal_df.index):
        clr = colors[i % len(colors)]
        x = normal_df.columns
        y = normal_df.loc[row_label]
        std = normal_stdev_df.loc[row_label]

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

        plt.fill_between(
            x,
            y-std,
            y+std,
            color=clr,
            alpha=0.2
        )
    plt.xlabel("Security Parameter (N)")
    plt.ylabel("Rate of Throughput Mean \u00B1 StdDev (Mbps)")
    plt.title("MiNTT Throughput Rate by Parameterization")
    plt.legend()  
    plt.savefig("normal_throughput.png")
    plt.savefig("normal_throughput.eps", format="eps")



    #SIMD plot
    plt.figure(figsize=(10, 5))
    for i, row_label in enumerate(simd_df.index):
        clr = colors[i % len(colors)]
        x = simd_df.columns
        y = simd_df.loc[row_label]
        std = simd_stdev_df.loc[row_label]

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

        plt.fill_between(
            x,
            y-std,
            y+std,
            color=clr,
            alpha=0.2
        )
    plt.xlabel("Security Parameter (N)")
    plt.ylabel("Rate of Throughput Mean \u00B1 StdDev (Mbps)")
    plt.title("MiNTT Throughput Rate by Parameterization (SIMD)")
    plt.legend()  
    plt.savefig("SIMD_throughput.png")
    plt.savefig("SIMD_throughput.eps", format="eps")














mean_data = """29.07035429	20.0384418	16.85911226	13.66961536	12.07036014	10.37576107	9.406034139	8.36636243
62.46610997	39.84366894	32.16589261	24.79752455	21.77142274	18.47812139	16.52926442	14.4649567
28.74968805	24.27383925	21.63475082	17.56472418	14.33938027	14.29605576	12.86875592	12.31958742
69.97934637	56.94231625	48.18739543	36.55547182	33.59540863	30.21058245	27.73182936	23.89998107
11.4426477	8.1437312	7.038452844	5.824307637	5.225030071	4.414855939	4.073126357	3.682874752
25.26482005	15.51066021	13.47032319	10.96052024	9.981371572	8.711861971	7.997421466	6.859077076
11.59630368	10.26111509	8.940929329	7.207015579	6.80966168	6.33708681	6.071461646	5.390100953
26.54500207	23.21340413	20.36799114	16.26705075	15.07391715	14.02438035	13.00496233	11.5610707
5.060533113	2.513958251	1.717023072	1.306247594	1.042101064	0.8685841701	0.7453610426	0.6563674336
9.698765765	4.832893884	3.25324171	2.435032766	1.95894247	1.635189505	1.398669564	1.228901589
28.58907151	24.41857203	21.55707606	18.0717013	16.52993011	15.12544878	13.83944715	12.57933443
74.612536	59.80075575	48.37275476	40.0093772	34.75583071	31.53539399	28.23234598	25.10966004
29.13804416	20.83403147	17.11943481	13.97043481	12.18818194	10.44645472	9.497982489	8.499572531
65.69056604	43.00352763	33.27725952	26.19312009	21.72862387	19.02042481	16.87773064	14.45333386"""

stdev_data = """1.020000013	0.6782361002	0.4734421378	0.2462100489	0.2133849048	0.2508401702	0.2007648528	0.1352408234
4.193056195	1.58268385	2.672553363	1.124143236	0.8988325093	0.4784963832	0.3420191053	0.4401423926
1.117825054	1.387179867	0.4320366748	0.2965313058	0.7737352659	0.9496179973	0.9929569884	0.1371361797
3.362796324	2.588592404	1.775706557	1.076021873	1.080407132	0.7364852815	0.6718806999	0.991760607
0.2768275178	0.1271134197	0.0852761313	0.1046881503	0.08279275153	0.1534056033	0.1125474758	0.0709695742
0.8171215756	0.4257091538	0.2879866098	0.1717557309	0.2599182149	0.1839822698	0.1014555176	0.1065828075
0.4517632454	0.1689540405	0.5492801853	0.4933278444	0.3674436784	0.3005373901	0.2248482443	0.1771865461
0.8565195228	0.6792948814	1.900533618	0.5258513715	0.5162576992	0.2176387734	0.293238197	0.1660738195
0.1869418856	0.1454276346	0.04487126954	0.02091668407	0.01360309947	0.01158914035	0.009606904162	0.008598525136
0.2798445255	0.1194638817	0.04082631981	0.03528773858	0.0326260789	0.02176182812	0.02132708701	0.01747203529
1.498787164	0.5513335438	0.6599865405	0.4456672609	0.1949427615	0.295488986	0.2344424508	0.1789195586
3.057126598	2.365291466	1.436752514	1.012262554	1.423314346	1.106792593	0.5233414628	0.7619555707
1.933283329	0.3738657966	0.2558367226	0.1820947024	0.140842989	0.1634920973	0.1691135228	0.1280959525
2.64003289	1.433574151	0.9235894995	0.9688869774	0.6709397479	0.341832231	0.4181771366	0.293235603"""


# Clean thousand separators (commas inside numbers)
clean_mean = mean_data.replace(",", "")
clean_stdev = stdev_data.replace(",", "")

# Load into DataFrame
mean_df = pd.read_csv(StringIO(clean_mean), sep="\s+", header=None)
stdev_df = pd.read_csv(StringIO(clean_stdev), sep="\s+", header=None)

mean_df.columns = ["128", "256", "384", "512", "640", "768", "896", "1024"]
mean_df.index = [
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

stdev_df.columns = ["128", "256", "384", "512", "640", "768", "896", "1024"]
stdev_df.index = [
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



Plot(mean_df,stdev_df)
