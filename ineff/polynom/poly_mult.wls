#!/usr/bin/env wolframscript
(* ::Package:: *)

n = 256
q= 7681
coefs =Import["C:\\Users\\EightyWan\\Documents\\programming\\MiNTT\\ineff\\polynom\\polynomials.csv"]

(* f[y_] = x^Range[0,n-1].coefs[[y]] *)

polys = Array[x^Range[0,n-1].coefs[[#]]&,50]

p1 = polys[[1;;25]]
p2 = polys[[26;;50]]

(* Output the Multiplication of polynomials 1 to 25 by 26 to 50 respectively. Then mod 26 to 50 by 2 and do the same. Combine, then Export *)
out1 = Array[Expand[PolynomialMod[PolynomialMod[p1[[#]]*p2[[#]],x^n+1],q]]&,25]
out2 = Array[Expand[PolynomialMod[PolynomialMod[p1[[#]]*PolynomialMod[p2[[#]],2],x^n+1],q]]&,25]
out3 = Join[out1,out2]
out  = Array[CoefficientList[out3[[#]],x]&,50]


Export["C:\\Users\\EightyWan\\Documents\\programming\\MiNTT\\ineff\\polynom\\poly_mult_ans.csv",out]