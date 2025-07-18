#!/usr/bin/env wolframscript
(* ::Package:: *)

n = 256
q= 7681
coefs =Import["C:\\Users\\EightyWan\\Documents\\programming\\MiNTT\\ineff\\polynom\\polynomials.csv"]

polys = Array[x^Range[0,n-1].coefs[[#]]&,50]

p1 = polys[[1;;25]]
p2 = polys[[26;;50]]

(* Output the Multiplication of polynomials 1 to 25 by 26 to 50 respectively. Then mod 26 to 50 by 2 and do the same. Combine, then Export *)
out1 = Array[Expand[PolynomialMod[PolynomialMod[p1[[#]]+p2[[#]],x^n+1],q]]&,25]
out  = Array[CoefficientList[out1[[#]],x]&,25]


Export["C:\\Users\\EightyWan\\Documents\\programming\\MiNTT\\ineff\\polynom\\poly_add_ans.csv",out]