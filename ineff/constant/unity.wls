#!/usr/bin/env wolframscript
(* ::Package:: *)


(* n = 8
q= 257

(* coprimeq = ResourceFunction["CoprimeIntegerList"][q]
f[x_] := If[Mod[x^n,q] == 1,x,0]
allroots = DeleteCases[Map[f,coprimeq],0]
n = n/2
prevroots =  DeleteCases[Map[f,coprimeq],0]
roots = DeleteCases[allroots, Alternatives @@ prevroots] *)

a = 3
g[i_] := Mod[3^i,q]
roots = FirstPosition[Array[g,q],2]
Print[roots] *)


(* q = 17
omega = 13
invOmega = 4
ninv = 13 *)

(* q = 7681
omega = 3383
invOmega = 4298
ninv = 5761 *)

(* q = 257
omega = 2
invOmega = 129
ninv = 225 *)

(* pmod[o_,x_] = Mod[o^x,q]

omegaMatrix = {
    {1,1,1,1},
    {1,omega,pmod[omega,2],pmod[omega,3]},
    {1,pmod[omega,2],1,pmod[omega,2]},
    {1,pmod[omega,3],pmod[omega,2],omega}}
invOmegaMatrix = {
    {1,1,1,1},
    {1,pmod[invOmega,1],pmod[invOmega,2],pmod[invOmega,3]},
    {1,pmod[invOmega,2],1,pmod[invOmega,2]},
    {1,pmod[invOmega,3],pmod[invOmega,2],invOmega}}

p1 = {1,2,3,4}
p2 = {5,6,7,8}

poly1 = x^Range[0,3].p1
poly2 = x^Range[0,3].p2
Print[CoefficientList[PolynomialMod[PolynomialMod[poly1*poly2,x^4-1],q],x]]

p1hat = Mod[omegaMatrix.p1,q]
p2hat = Mod[omegaMatrix.p2,q]
p3hat = Mod[p1hat*p2hat,q]

Print[Mod[ninv * (invOmegaMatrix.p3hat),q]] *)

(* q = 17
omega = 9
invOmega = 2
ninv = 13 *)

(* n=4
q = 7681
omega = 1925
invOmega = 1213
ninv = 5761
g= {1,2,3,4}
h = {5,6,7,8} *)

n=8
q = 257
omega = 2
invOmega = 129
ninv = 225

g = {1,2,3,4,5,6,7,8}
h = {8,7,6,5,4,3,2,1}


pmod[o_,z_] = Mod[o^z,q]


makeNCCOmegaMatrix[omega_,n_] := Array[Mod[omega^(Mod[2*((#1-1)*(#2-1))+(#2-1),2*n]),q]&,{n,n}]


NegayCylicConvolutionNTT[g_,h_] := Module[{p},

    omegaMatrix = makeNCCOmegaMatrix[omega,n];
    invOmegaMatrix = Transpose[makeNCCOmegaMatrix[invOmega,n]];
    
    ghat = Mod[omegaMatrix.g,q];
    hhat = Mod[omegaMatrix.h,q];
    phat = Mod[ghat*hhat,q];

    p = Mod[ninv * (invOmegaMatrix.phat),q]
]

NCC[g_,h_] := Module[{p},

    polyg = x^Range[0,n-1].g;
    polyh = x^Range[0,n-1].h;

    p = CoefficientList[PolynomialMod[PolynomialMod[polyg*polyh,x^n+1],q],x]
]

Print[NCC[g,h]]
Print[NegayCylicConvolutionNTT[g,h]]

(* NegayCylicConvolutionNTT[g,h] *)

