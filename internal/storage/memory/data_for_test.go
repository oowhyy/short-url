package memory

type testStruct struct {
	val   string
	short string
	seed  bool
}

var testDataSave = []testStruct{
	{val: "https://KIapqNzQIcSyGPrIBHvEFH.com", short: "ZaluQgETiT", seed: false},
	{val: "https://sYKlU.com", short: "NcbPYXSoVm", seed: true},
	{val: "https://bcJoQKcOlQJapdtuJWvBnKgpnnmtjfDnKsaNH.com", short: "yCAjypdzRx", seed: true},
	{val: "https://dyknEjSVstEfOKEiYVhVDrnadBIhXxN.com", short: "tIQgxkB_VP", seed: true},
	{val: "https://dltPpZXArpnomYcHETOI.com", short: "oqjMpJKHKx", seed: true},
	{val: "https://CDPNSvheAQZUYVCmbBexTbxntffj.com", short: "hLjQoxIOcg", seed: false},
	{val: "https://chTDVgjA.com", short: "kzNutCecyt", seed: false},
	{val: "https://nAriQrPRESfaBmzDABmATTXsNXL.com", short: "qaZDCJmeQO", seed: false},
	{val: "https://JGgczdCWzOUtYLfmwlktKbzsJC.com", short: "qooelOdLwO", seed: true},
	{val: "https://GodUgJbMaadX.com", short: "bCQgByByjD", seed: true},
	{val: "https://APoyUcfmF.com", short: "uehlKejUoe", seed: false},
	{val: "https://yopZoLuxoYyftuwjVkAEkhnNBypKxhDNaipkp.com", short: "HWRuJGGvdt", seed: true},
	{val: "https://kBqltbXUWLlwclXJULccdGHUIKPuKdocYL.com", short: "_jwhEGtpAr", seed: true},
	{val: "https://wVhfmCqHrytGUxGAr.com", short: "wZawvRmLgK", seed: true},
	{val: "https://TOf.com", short: "ttHLJBIbzz", seed: true},
	{val: "https://hATdCxMnxwCVmyMhYdchoDJFHiTqizxv.com", short: "VdOSiuipNF", seed: true},
	{val: "https://myKJJJGfFLmXFfVmzdFiAAxGikedtIVbHyLTPhCDy.com", short: "icEIGRYyQo", seed: true},
	{val: "https://mtCQlmzZqfxrnHLNxVlnN.com", short: "zbmuErIjVm", seed: true},
	{val: "https://vrEuwGNqpElSLaFoMrnt.com", short: "JcUInFtzzn", seed: false},
	{val: "https://TxtTUcWEfjOtnMonxRWKImNcMtpu.com", short: "EiudYrGxCO", seed: true},
	{val: "https://AnCUFOCxlXeHyf.com", short: "eylJaLYDVo", seed: true},
	{val: "https://euRxpFUxnkAbIAir.com", short: "iVrBKhkaoQ", seed: true},
	{val: "https://TCoU.com", short: "ildQtzppvP", seed: true},
	{val: "https://GRTVYPsmMsmAhKlos.com", short: "gVNGOIQHRf", seed: false},
	{val: "https://xldKEdXCGUAvpdmFYGNDChoCWbAiRqmlzCjybR.com", short: "lkVjrFMtIj", seed: false},
	{val: "https://hzXhyJubYVSTWZdITvQCMtlzroCfCiahyASqomCKC.com", short: "XNvUmrlNcL", seed: true},
	{val: "https://YvIMqtBV.com", short: "KNiukkTbtW", seed: true},
	{val: "https://FAtvDjHHcCuqz.com", short: "mfNAkTFhez", seed: true},
	{val: "https://BaRrpi.com", short: "ePfVdyMUio", seed: true},
	{val: "https://ZHgMAyBYXjMKFiO.com", short: "OarylWaBkk", seed: true},
	{val: "https://hjiWUtQBTBahhuNkkyvpwzUnmblcswdbMn.com", short: "jWreeAxnVA", seed: true},
	{val: "https://SZnYOqobQrBFutVGPJHIr.com", short: "EGxZkFoHFg", seed: true},
	{val: "https://JgHuGffNTyBzkFvdcmtGlikHF.com", short: "baZQJTwfqe", seed: true},
	{val: "https://UpHonMxubmVWlRXvQSzqfiUtxwXrnYkVDXJDCMWMr.com", short: "izsvWzwxDt", seed: false},
	{val: "https://FMLIGKBiPdOhOnzOviwDwgGwfolBzTbwtnlH.com", short: "DhesDuLvWG", seed: true},
	{val: "https://itzyzCWaiCcLtDySgacR.com", short: "jiFHAShcQf", seed: true},
	{val: "https://ZVfoaJDklKaNOGsJ.com", short: "XWpmuxIcpV", seed: true},
	{val: "https://aLixydX.com", short: "WESNNoiTSX", seed: true},
	{val: "https://OZclWafUecXqB.com", short: "LaKZYQffeW", seed: false},
	{val: "https://peKSrQiUrceYSDtzCVkIrFwusMWLghY.com", short: "NhBihwZyYE", seed: true},
	{val: "https://crinuofgvelNdqMqzqRcqGra.com", short: "NpMHcmXWAo", seed: true},
	{val: "https://esPrf.com", short: "UuKkSvBjjJ", seed: false},
	{val: "https://dlrURNzlipOTewdjcHvOSnB.com", short: "ygKeiUqtzK", seed: true},
	{val: "https://ZYWfAqPLBAaU.com", short: "fCoBlPgeNK", seed: true},
	{val: "https://PqOUw.com", short: "PemBd_sIuC", seed: true},
	{val: "https://bUWQbemuoFmovSOkZtsdnpn.com", short: "bcR_Ydjrlv", seed: true},
	{val: "https://dVddXQUIshPDcgBrRMbmsyQixeINJYKBYIWEzjfEfQ.com", short: "ythsELLsiZ", seed: true},
	{val: "https://UCpVkjPQhDvKvXpUIwPqJdrl.com", short: "NCBXRkQVzx", seed: true},
	{val: "https://bXpWzMQyhUkZHrKPJb.com", short: "_RDbdbVDui", seed: true},
	{val: "https://NZWNlXrxAtatFBVtvh.com", short: "rQiCBIfqjI", seed: true},
	{val: "https://kNnVfzNYbA.com", short: "aVoPgJFrYq", seed: false},
	{val: "https://mHBEnHckYYXcKRp.com", short: "mGqixGjzhD", seed: true},
	{val: "https://elTnJeUcIlCUoxAPcQZXYnodqvZHXWptgrAmU.com", short: "qNHLKREvqS", seed: true},
	{val: "https://dvvNIpLkIjShdqTZWFvvylZvdAmVzPWT.com", short: "NRQNRnQZDo", seed: false},
	{val: "https://QUQiREFdkBsD.com", short: "JCxKBNxGbo", seed: true},
	{val: "https://GZVgpuPSLIBhWGD.com", short: "zcqwEhdeOK", seed: false},
	{val: "https://TrMdUfKRrYfPypPqnaH.com", short: "rbhuEGrnua", seed: true},
	{val: "https://XpYhZdIjDPcRwUiSGnWiTMYIxkKCvzpm.com", short: "aUFhhjKMcK", seed: true},
	{val: "https://YeBJKKWUQqmnGfwjOCGrnTovcltLr.com", short: "kxfQggXnhv", seed: true},
	{val: "https://pMTSvFznqHvnxKAbIJwSHKJSU.com", short: "MmiEnGYYng", seed: false},
	{val: "https://FufqXzAdNbCXALaWBcRtdbW.com", short: "GDxNCeOzRm", seed: true},
	{val: "https://AFXQQEFhNJWJM.com", short: "RZcqNiTpTe", seed: true},
	{val: "https://HBDgIrZLZ.com", short: "nvAbUzadsb", seed: true},
	{val: "https://OgvMwtTcmJbNyjmtCQzth.com", short: "MxrimrvGDy", seed: true},
	{val: "https://MXll.com", short: "jgHcNtJnUS", seed: true},
	{val: "https://pWJaIKHwkEusjwTpXjKbmyBPYqfLO.com", short: "DjTuttvHfs", seed: true},
	{val: "https://rcjHOANYOZOGojkQSibYpmPjrRnAmOBDLKEUrqSjOd.com", short: "uuGoLTpQzi", seed: true},
	{val: "https://EMjSuanRuRXjWuVCuWJzSvhycqgYFKGYmsyWMKEzCx.com", short: "CYzhDnjJRo", seed: true},
	{val: "https://RcCYLBHKKoBKFXMBMJhd.com", short: "McWZxOhnxf", seed: true},
	{val: "https://gRgR.com", short: "TLtxmFsRKl", seed: true},
	{val: "https://CArNRBfqWvrRYzbkWvBXyVZjSVEQGoznKJllqM.com", short: "pZysYSLtDZ", seed: true},
	{val: "https://uzEMiUSEqoSR.com", short: "UYtEtOOrEm", seed: true},
	{val: "https://BYuwLKEnuz.com", short: "_CqZlEWqnP", seed: true},
	{val: "https://KRDrwzqFrREiMOOHZbtlMfCZJxIniqUtYVeSh.com", short: "rYCZsaDwOq", seed: true},
	{val: "https://mXxyvpiTcMGFVmsCcUOSnJBkZLNNoRQg.com", short: "cWCPbqLEDE", seed: true},
	{val: "https://dxVEcDhmZwatoFZmeeJLzxbWRTAprCbL.com", short: "ImbJGrmamJ", seed: true},
	{val: "https://tJTb.com", short: "RiCbfkYmea", seed: true},
	{val: "https://mVaRzUKzCLfgRTfNgJYsfUeoeB.com", short: "EIMmHPcXuR", seed: true},
	{val: "https://CXKDzpJMtrfFrVwQnof.com", short: "eWaqsSeMCM", seed: true},
	{val: "https://GVoiTUYhiLeWQkekHaPflLgUWuKcGhkmlcbNKM.com", short: "CLJBihsOdj", seed: true},
	{val: "https://PasqKGIzODcCWoMWxciIEphccGYtmRr.com", short: "YY_izFSlCY", seed: true},
	{val: "https://vHobLGQuaAqlYdbrgKWLqLNuTizwdM.com", short: "kRrPrqvDJO", seed: false},
	{val: "https://SLtmkfzJVmbiatrvX.com", short: "aBkCjgeFTs", seed: true},
	{val: "https://fRIsfmlJNszyOHtHRkJJJIgXdpZtuSLgvxvcVgJda.com", short: "QVVDNnebpi", seed: false},
	{val: "https://nEzzvKVKAhjQfITYtFkqJHGcBnVUpbN.com", short: "XojBCmaoPX", seed: false},
	{val: "https://GSRanyPO.com", short: "pYtpfxxSMj", seed: true},
	{val: "https://cBZIDOMWblUMxsfnYeW.com", short: "ZOwQKsuURq", seed: false},
	{val: "https://LieXFFhGIIrWrMsMBFqvTH.com", short: "sockDhne_s", seed: true},
	{val: "https://FEgonfIRWJnJIwFoFDGioULTtTtXg.com", short: "VBNjqmcwjS", seed: true},
	{val: "https://rAmzBWjwmHJiFnvugdOnnoYjSLyjlNdWXjs.com", short: "M_BkpuMWwR", seed: true},
	{val: "https://sPINWxfulrWddmhhPTIFkqwSgzweyNvUezeH.com", short: "uMSrmlwAbK", seed: false},
	{val: "https://VmQAvjzJGKgW.com", short: "sRTpeloftr", seed: true},
	{val: "https://tHCzoNOCySgwndTjcUVOTwZAtIgHICHfYNoUveNBj.com", short: "lohBcxerOk", seed: true},
	{val: "https://xRBwxUQiadcYcKQMNmA.com", short: "wnqWnszzRI", seed: false},
	{val: "https://IluMbUsiDN.com", short: "qgDPoMXrgo", seed: true},
	{val: "https://hMtxajDYqWRpcNFpnl.com", short: "GFAyqROjlD", seed: false},
	{val: "https://lWaeCfqlssAlORrnriVzkcE.com", short: "XyBUYwOis_", seed: true},
	{val: "https://sqCsmMbiDoyMDVktUVoPhyaNaxhMvUipUOmq.com", short: "AAftAqHQGK", seed: true},
	{val: "https://SERsgVDzcjJfkOJJbtiSwpZkpITCznUEYAAWsXtDqs.com", short: "bSezSZ_cdd", seed: true},
	{val: "https://XYhZqqtoXdeluBTTraOSjvWVrX.com", short: "gwwNbwhiGZ", seed: false},
}

var testDataFindByKey = []testStruct{
	{val: "https://wDlXgZZcqH.com", short: "TwQOcNMNeg", seed: true},
	{val: "https://TKfkJUnRNWxLlSttIilnBDRWN.com", short: "BY_rpYBNcw", seed: true},
	{val: "https://aLql.com", short: "EdliTKSCbZ", seed: true},
	{val: "https://VkrxETTLrxLUejZL.com", short: "JamdunBssO", seed: false},
	{val: "https://PLUxbBXvym.com", short: "qruXmvBxoN", seed: false},
	{val: "https://IlwCmUKzXcNRNKXPaZyjqWVCBtBpOQsde.com", short: "JADhrdurzz", seed: false},
	{val: "https://fSVxShysaIj.com", short: "hMZbUpev_K", seed: true},
	{val: "https://ENGWDDhUTTdTLEFaibRDLpvCXYRx.com", short: "bAZdUdcwPY", seed: true},
	{val: "https://pNmEGIscRWzLeeDTDVIvTIuRHvPPI.com", short: "Ovdtel_Arw", seed: true},
	{val: "https://nqz.com", short: "tQqoVXsufM", seed: true},
	{val: "https://vtLxBeMBxxsrKMkxXYfwpTAHRkTJzxaUl.com", short: "rUqJZFpWsi", seed: false},
	{val: "https://ymLeAiyVdmPArcqmAzXdjIeZhfDzaijhy.com", short: "eGxkeUTkLS", seed: false},
	{val: "https://eojtqeUKJaP.com", short: "mmNwVCRHyA", seed: true},
	{val: "https://mVAuOGnEpbNgXHkwAZhAW.com", short: "AtpPpXpKri", seed: true},
	{val: "https://VJRuEiJkzFTR.com", short: "bFDNwhqyNB", seed: true},
	{val: "https://VszdOfpfdZ.com", short: "ASbSwEwpHm", seed: true},
	{val: "https://FrVkkoIXZucpdeIbVbNXbkoiNzvYSeXdXxE.com", short: "msZrusuJFl", seed: false},
	{val: "https://ocz.com", short: "PbzIviYGnx", seed: false},
	{val: "https://OJCjUNhyaLKirvKzJYREsM.com", short: "LBBfRngnvC", seed: true},
	{val: "https://UArKKCYVAtMtbghvJMMcwgHbvsohiKoAGGOy.com", short: "bgsJibpSop", seed: true},
	{val: "https://UiCwdNAQrMzZQPqDddOrzPHyuSrRTDKdnKRTX.com", short: "wFGeIXYSKj", seed: false},
	{val: "https://PLHJxBhGSpBNkwQXBOP.com", short: "TzKiBZMjLl", seed: true},
	{val: "https://SgnlXLRm.com", short: "OKPpZWgslb", seed: true},
	{val: "https://cXbUbnKacqTguMrPQeAxxTEgjcfOm.com", short: "rYNIyLMcOY", seed: true},
	{val: "https://fwIUfMNATXTJmnATaupZ.com", short: "YjSTMqQjnO", seed: true},
	{val: "https://TJRNtEJEUJVZCgpVROpnjcKEITctvYD.com", short: "wwEOgV_WMQ", seed: true},
	{val: "https://xtsJFE.com", short: "zNz_ZyCxRb", seed: false},
	{val: "https://rVBGEeRthTAdH.com", short: "RcTcoUljHP", seed: true},
	{val: "https://ekkUGuG.com", short: "tbwSbxaWwO", seed: true},
	{val: "https://XTXAu.com", short: "mDVwSovYOD", seed: false},
	{val: "https://WNFhCZgbrGqFftPtNsrBDAvaweQLowvwxNYnASwx.com", short: "eAIHBgYBDG", seed: false},
	{val: "https://hxGgwKgCmFGfeDg.com", short: "XzMkmSSDeD", seed: true},
	{val: "https://MnTFQ.com", short: "WzkVIepzZq", seed: false},
	{val: "https://cSlLAqThTcEEHEaRjgdEcODSLRiKNUdmboFUiy.com", short: "PLfZnQmfLV", seed: true},
	{val: "https://EBfljOBL.com", short: "eeNDljaviT", seed: true},
	{val: "https://jznIIcTltWMDmyIGYxpNDmKZxBeD.com", short: "HKGDssdRuE", seed: true},
	{val: "https://oqRiGymezrniFnVzlZCWCaNCrEU.com", short: "rI_bMNiA_g", seed: false},
	{val: "https://iGIcvqORzcLURaNyTXwPReZBwafKrbIn.com", short: "eRUvL_lrlk", seed: false},
	{val: "https://kBpuMAkPMaYVYeGwLrmoQjPjuYVSiRCCKq.com", short: "KvGGviMHNe", seed: true},
	{val: "https://EKxivAblogJCQ.com", short: "XyAkNeuKGH", seed: true},
	{val: "https://wvpKPuZVaaTOpTiMAPhwdMbR.com", short: "mUqjvNvcvs", seed: false},
	{val: "https://jLGmYVizqxoiFkRQYcxKjTiCyWWLyzI.com", short: "vyvWMROWDd", seed: true},
	{val: "https://xqBfsLiEqUTF.com", short: "cCyGsYRyGf", seed: true},
	{val: "https://rjDfH.com", short: "NeciqtioLl", seed: false},
	{val: "https://VfOsjOyCrODIcWJJhcWUteNJ.com", short: "AwKPCkRnNa", seed: true},
	{val: "https://psaobMnBtlNuV.com", short: "UktyMVOdqR", seed: true},
	{val: "https://OCgHCkF.com", short: "RZMhOPrPRe", seed: false},
	{val: "https://oCYDisrrshFxCvEacsUhzRdjtKZUlCoJKy.com", short: "KxwnLnhYbS", seed: false},
	{val: "https://zHahNoom.com", short: "NRfEiNDARu", seed: false},
	{val: "https://HyIaeIPEMV.com", short: "NCiDJJSZnC", seed: false},
	{val: "https://ZLYmOTvRywyAIAiAjwPrOsQOVgZuqsSrOgyTolYX.com", short: "fATdfDSwhr", seed: false},
	{val: "https://EjFldjRZZsRlODvcoTYotDmIPOuL.com", short: "wYTDRFLPxM", seed: true},
	{val: "https://rphpJFIbQGqdxVcGa.com", short: "NpfqBCiCFG", seed: true},
	{val: "https://UbNivRjJKspxUgVgRBpY.com", short: "ElRkLEfzwS", seed: false},
	{val: "https://TXRsmrLPlQefzddyxTsTPqzgIHclubgfq.com", short: "DIaNAXje_a", seed: false},
	{val: "https://HxNHCYJtNZAWYdsSfXokCLmAXfIYLgYXK.com", short: "FwJMnvacMd", seed: true},
	{val: "https://lyKFBTezSrkCNBBijNqwOReneweLxAauENIZjNDMTs.com", short: "RvPWzLOTlS", seed: false},
	{val: "https://hsheNKfsylKQeeysvlkIoIqhJJ.com", short: "GzrMrdzema", seed: true},
	{val: "https://VUgeegkUVjxFPKUeznwtzOqrhanLCmOg.com", short: "TJhqQgrmnW", seed: false},
	{val: "https://qEDjHrJhMsehWVrpqvdBGnUr.com", short: "UMAnByZxKF", seed: true},
	{val: "https://yYfnssOwlLdJ.com", short: "qAOEzzXZqA", seed: true},
	{val: "https://ELoyrLeRxMDnzXBdKkyASMVrNTqk.com", short: "iRmCkfMQHw", seed: true},
	{val: "https://bAlUzdkwCNUehIJDKXfWiWwAIXpuIXuSIzStikZXyW.com", short: "NtDXLrOsml", seed: true},
	{val: "https://vnrY.com", short: "iZNiNNciKf", seed: false},
	{val: "https://QDqJqfX.com", short: "PsZyGnrdLx", seed: false},
	{val: "https://ZLN.com", short: "pNyTCnCnOS", seed: true},
	{val: "https://XTNRfqLtEvDVSBWvPKlJOmrbKCDZ.com", short: "sk_Dz_QFvC", seed: true},
	{val: "https://wIVvPCqsyCMyXp.com", short: "vEyVzYTiuk", seed: true},
	{val: "https://UPMbgooCWcbBuCSNyQfJEJ.com", short: "hEnfAPvef_", seed: true},
	{val: "https://rVhuRXTbcDtpMs.com", short: "yErEDTZzJm", seed: true},
	{val: "https://ZDRixWjrUoVVjoeJlZIGBattxoj.com", short: "SUlLGxdiem", seed: false},
	{val: "https://bKRjKEkckqqfRowHsbNjHvg.com", short: "IAtXAvJoXI", seed: false},
	{val: "https://nomkJvLkqpnyVngEJJn.com", short: "EVhGBRoYSr", seed: true},
	{val: "https://IEkmgjRtuKZFkbeXSrvhYMDPQPfknRxjwG.com", short: "uLfngkZyIa", seed: false},
	{val: "https://gLvbzWWVxyafLHykSjgo.com", short: "DvyHbTUqjw", seed: false},
	{val: "https://tQABoSlmiHMexQpwUFjVMYnEOCvVPok.com", short: "PpCTqUckYi", seed: false},
	{val: "https://DoXApVQyvIjExGXdCwPmOGGbNsIyB.com", short: "LRMCdswxri", seed: true},
	{val: "https://cxrff.com", short: "mzXEQLCEjR", seed: true},
	{val: "https://pnoijnviwVYFmqeOUvryWdKMzztWrBzuiXBjDuX.com", short: "YWVnfUXPXv", seed: true},
	{val: "https://zXlgLHJCzDsIKRinaTyoriYMFJGA.com", short: "gWvLeDGXbK", seed: true},
	{val: "https://WkRkSJcAupaZjBleIT.com", short: "lZtAoATPuu", seed: false},
	{val: "https://HlzkBjfjZtvVPCllshDwslyHxCjJOG.com", short: "dmeLJcOPMY", seed: true},
	{val: "https://HwhIxmOoYAJiZecRWMeTvmfLCIUcdrMOw.com", short: "sRjNeUbLji", seed: false},
	{val: "https://tfLNokrMhAVVsbtcov.com", short: "UcCMMLSovv", seed: true},
	{val: "https://EIdYGGjX.com", short: "sDhCbUjX_r", seed: true},
	{val: "https://BlbabKtETuMnwOkgGtKLscbIOUmyMHkbSnRYjomz.com", short: "sSFKNL_spR", seed: false},
	{val: "https://aklsiBlPXxnGWvSoLBlvnFVkoTRonNbraxZhbwnPyk.com", short: "gMSBYnLPjw", seed: true},
	{val: "https://sHoKoncWwOIISL.com", short: "WFPiWYNVfF", seed: true},
	{val: "https://QvTwvOyxSSlWdhLQFFmvvL.com", short: "ZHIjEoHAPb", seed: true},
	{val: "https://XRMBKnQCynnVecMsaubJcKyKcwprYZzFMFLG.com", short: "KXrUpgluBs", seed: true},
	{val: "https://ufRPVvTAuFfzSvNmvlJbcRaWluRxnaUyHJeUni.com", short: "NadfFDaWgT", seed: true},
	{val: "https://bmmDJKrsrTsvHXUjmUQshFCMSvrkubtjlTgUqR.com", short: "WTTljBtpOp", seed: true},
	{val: "https://bXrSvNDrdLPkXAwtPVOGHUXsCImFSPjDvdiVHdry.com", short: "SJiYyoKknt", seed: true},
	{val: "https://BFqsdtzdyxAehLQlUVvYgXPGccRiW.com", short: "yABGnVnNI_", seed: true},
	{val: "https://aJraYIBEM.com", short: "SEPUSfCfRx", seed: false},
	{val: "https://ssxZhgvtkNvLXCUXCSoOvQoBKagou.com", short: "gwooXZWEMQ", seed: true},
	{val: "https://KhWnNwkCpOMlOvlPZATkKIGGfwGBGuwdXBnV.com", short: "ZOJcCqboMV", seed: true},
	{val: "https://sNoWDzgsddvtdFOcGLYSHQYuoKaARhvIU.com", short: "iuHhboiPsW", seed: true},
	{val: "https://fCaYFDlMtszhA.com", short: "MsiBJYOvRV", seed: true},
	{val: "https://rDRxAdPicjbknMUoWtdsMRO.com", short: "eckncHMqAN", seed: false},
}

var testDataFindByValue = []testStruct{
	{val: "https://MDtFlEPt.com", short: "dZmQSAOcnR", seed: false},
	{val: "https://ZVpCSjGJyymE.com", short: "ZGoVrGslkX", seed: true},
	{val: "https://pApCnPUYFkjTpGXYKbRRWAdSUdNHqjRlLLNUziV.com", short: "czcRceZmHY", seed: true},
	{val: "https://EoyibPrKbaSLwIswvFoGAsusekEIguVnkkLKNpRMe.com", short: "BQwV_TTUef", seed: true},
	{val: "https://gpVTYHyVVkeOfi.com", short: "tnIRFHPJdB", seed: true},
	{val: "https://lxfDmsqCQqf.com", short: "geaUGkwbcb", seed: true},
	{val: "https://DjmTJilWnKQfXHTzHEYZxtVOGFboSwWBwx.com", short: "thIOGcfxrv", seed: true},
	{val: "https://iYcmpjRD.com", short: "PdlfNVfyXL", seed: true},
	{val: "https://VAutGbfYkglzObMohbxpleflBGIyRDfVpvxOMZzyR.com", short: "ZEbACaxwBp", seed: false},
	{val: "https://KnCRAWHuVKXOf.com", short: "whrBORnDxI", seed: true},
	{val: "https://SJEBggQM.com", short: "XUwPoJNGAz", seed: true},
	{val: "https://WbmLsNUsJIGUpgaRmJrvRiHVB.com", short: "tILrrAqFtM", seed: true},
	{val: "https://QOKmeWjNjdLlnknXgzmVLHvXZS.com", short: "CvtzYbIzlo", seed: true},
	{val: "https://LfSy.com", short: "OWiuxYQ_vJ", seed: true},
	{val: "https://ZirkwhjulWYxeKNRbpkYUBmehwpZjVPBHJoxDDJx.com", short: "YYRtZ_IIwx", seed: true},
	{val: "https://SaLfOsnqzjRcDxlvOZvcHmgmNSIh.com", short: "CkhRhiuFjJ", seed: false},
	{val: "https://YzQefwGcUFdJmduZrBbJNJXVucmNXgqOYJX.com", short: "GyrfRwQfek", seed: true},
	{val: "https://ynZrrPrlfJCrvQdcKfUCprNmChHXUtYBKAOxDxqcN.com", short: "BrjNoWjHLc", seed: true},
	{val: "https://GRJVakNUGBAPEIzcLbuEyIGzREm.com", short: "xCvrMVqJsx", seed: true},
	{val: "https://OjZGJBSMqUGLKeIBFePiYDITjZerZF.com", short: "VKNwNxJBuk", seed: true},
	{val: "https://NAJg.com", short: "YkfMYhOpxk", seed: true},
	{val: "https://GbAFEHBsCzrnMvRTEdjQCTJMhcjjnFwrp.com", short: "UEXvBlukkH", seed: false},
	{val: "https://fXSQSQtFbOZhPHxYL.com", short: "lZrtfWWTvi", seed: true},
	{val: "https://ckcUhOkh.com", short: "YnCNBIMcab", seed: false},
	{val: "https://BNjweBjmQcEKDAjDu.com", short: "VBekBnmsJO", seed: true},
	{val: "https://rIyXWbagpOQeaerRgpmJfCo.com", short: "UTpELQAcoX", seed: true},
	{val: "https://WvxkxLAdNqEksktBIuMo.com", short: "D_TjHofsHf", seed: false},
	{val: "https://rfpoBTDGdIIynmQWxHaUOcbNJIjaSosgxhs.com", short: "FoXYTDNaZa", seed: true},
	{val: "https://XKhwgeMzcCCEFHqSsvHenHY.com", short: "CjFqiPmJXR", seed: true},
	{val: "https://eFxkAmDEsjsGTfrFBgxtPEfvIIWPDtUzjvoiqdV.com", short: "ptdjKfOLiQ", seed: true},
	{val: "https://RBzzkBLvmqKJZHPxColJGLGPZIaTHoISIrcEPxoHGR.com", short: "fXjtKGjLSf", seed: false},
	{val: "https://vAUUixyaeNplCdjJHksqcSoxKMCxtSiumVgMYcga.com", short: "boTptFFgXC", seed: true},
	{val: "https://uvkdGyZPpKcoohjobubkxifXAIsDohyJN.com", short: "TKRWSWnTjE", seed: false},
	{val: "https://tKSAhvAXIypXabuZIVbMnPyne.com", short: "Gi_cMmhBxq", seed: false},
	{val: "https://YzaFerglbhtdJeBvkAnyGGBNcmjMrSwRncN.com", short: "RfIslBbsvy", seed: true},
	{val: "https://aoVBQGqjdHpsURhagvIvylLlngPVjRHiYtgaHmFwf.com", short: "PzSrFINPlf", seed: true},
	{val: "https://uXioccyeoHQMDc.com", short: "UaYc_EGLvP", seed: true},
	{val: "https://dBRWAVAMiNOqfHYTIS.com", short: "AlouQXDqBx", seed: true},
	{val: "https://VaOjGfibavArGDbzXjvjKAfbLtGkNBy.com", short: "UFbUTWLlbj", seed: true},
	{val: "https://IYPhvrSinTeLlBPvedUYezayysLKRboRWrYp.com", short: "Amp_UBokcB", seed: true},
	{val: "https://aHeLEAgiQatTDwbrKrPjnwaqN.com", short: "OYNowB_ktR", seed: true},
	{val: "https://OOmJpKtaVWhnxWryJA.com", short: "qnkUAYwZSD", seed: true},
	{val: "https://sshCNZOpWMEDinBCMLHgqyPPsCKYiexPMJD.com", short: "jpInUpkKzh", seed: false},
	{val: "https://wPwKhQVGrTnGk.com", short: "DQQZYRLFGI", seed: true},
	{val: "https://keYPcLnstbhJWdaKSttLbcvU.com", short: "ePPooeLDzB", seed: true},
	{val: "https://SrHcCpyoh.com", short: "ZOXMhnEsR_", seed: true},
	{val: "https://LrCsNvClpMHjordCrcMIYEUiMt.com", short: "XLDsHLzReC", seed: true},
	{val: "https://BzIdyegyTK.com", short: "nUgbM_EZtf", seed: true},
	{val: "https://eolWMrIATqntxdCpJLapUyHDvnmZS.com", short: "eZxVVXYPpo", seed: false},
	{val: "https://GcsxMQsYIMvSbHeWfFmPgrPfwNewloVy.com", short: "hChqvNvBEK", seed: true},
	{val: "https://vCNPGuLQvAtIbyZqlWtXqkixyZ.com", short: "Hpjxhhm_zB", seed: true},
	{val: "https://FwHUpbDtcplfiQGayUslmjYqmeBNZMngTWShIuzQ.com", short: "yZLAGcREfM", seed: false},
	{val: "https://PRqjrVdeaVHpqiRpjclAMIFGvWOnxyOKkIaBTz.com", short: "JGtAFSzzVK", seed: false},
	{val: "https://PSsiqo.com", short: "nQUPmeUDwA", seed: true},
	{val: "https://zKZRCxILjpjlyVRwvmffLGopHggLzsTXWPdHkJp.com", short: "KjGwnFdeDb", seed: false},
	{val: "https://TFtFgxsbhZHveZu.com", short: "dHdvBOaHZE", seed: false},
	{val: "https://iqosWgkguEOVAjvaOmYhVHqIk.com", short: "HtbBARWPwW", seed: true},
	{val: "https://cIghzbzgMiTj.com", short: "hmqVKYyjea", seed: false},
	{val: "https://TMkYKychyoJLYuvx.com", short: "TuCQbgMAwo", seed: true},
	{val: "https://aUWmKEon.com", short: "uonNNsENnb", seed: true},
	{val: "https://isDXkuehXvYXgRWRzpfiaIMrdJtWJPh.com", short: "xQiCQdzLKE", seed: true},
	{val: "https://QcELovlVTBiAmGzAJeVSb.com", short: "BlyMvrOagu", seed: true},
	{val: "https://lchvKnduuIYgTeWy.com", short: "nBDiCFUbgf", seed: false},
	{val: "https://ykzTCsKZxvGe.com", short: "yQjRIszDqC", seed: true},
	{val: "https://HCkxFcyONfmHGeDTHvPKSTR.com", short: "MWpjjfddZY", seed: true},
	{val: "https://RzBxWPOjJEdYPCzXeqEc.com", short: "DdhFixML_X", seed: true},
	{val: "https://rDBvlDkyBLyxPUjSuhayUp.com", short: "pAFqhGzSsE", seed: true},
	{val: "https://GQpiPnHi.com", short: "fvnreT_OxP", seed: true},
	{val: "https://dWsJEsPFOQUhFwULIlRoO.com", short: "YaUFRSD_SE", seed: false},
	{val: "https://LRRJdxQFaTViiGSbuTQCmefieaiD.com", short: "XJdEwveGMl", seed: true},
	{val: "https://UeYkvWUIfgciyBOUrCqSvNiZlSzwWUzZIbNlwxng.com", short: "_HEkaEvQKx", seed: false},
	{val: "https://EnZbNvCCPiDsWxbOIjKgjZkfwjoHeBUZBEHRo.com", short: "hAUiXilibb", seed: false},
	{val: "https://lvOUGIUUutiCTIvFszPSnQc.com", short: "CCOK_EKK_Z", seed: false},
	{val: "https://pAMtZxerRFvBed.com", short: "ZRUy_JCkeP", seed: false},
	{val: "https://IiSgvDzorXFeAxYGfGEswbaCBvHgNLVIQXdxH.com", short: "yrJISRYrPG", seed: true},
	{val: "https://tuqJFfHDkruNfIpP.com", short: "YyVjoD_dsT", seed: false},
	{val: "https://peOzkIUzCEHsbKxvfpMjuayKGOpyHG.com", short: "RZhkXHAcZB", seed: true},
	{val: "https://XrrxmcdspQXBTGxGUmXzlxAjxaSWvnZbpWPboPJ.com", short: "DDXfIjvSAm", seed: false},
	{val: "https://irCpEHLLCWIaxSdhdDbtp.com", short: "QIDIOLmcqb", seed: false},
	{val: "https://vnHnvhLbKaUCCgMciDPPsAnhFtiTSbBRJxnym.com", short: "JXCJwrCsrk", seed: false},
	{val: "https://WsOPkBNQlqnujPDwQD.com", short: "SHHgezHEGH", seed: false},
	{val: "https://RQeXchJGRRNqKsDvHBudoUZQzwgwbRoRMnaeUZWb.com", short: "YnlrBHtsnQ", seed: true},
	{val: "https://zyjWRZmSYGwjARwCxXoDZtBsC.com", short: "yIDghbYivR", seed: true},
	{val: "https://Njr.com", short: "aPoODfPUsa", seed: true},
	{val: "https://dihYsCn.com", short: "FqfUe_Udee", seed: true},
	{val: "https://TnubZdwZDkwGhKAnHQBFavTzyCvKHPZmQl.com", short: "lOiabMiOMG", seed: false},
	{val: "https://tcODRNAIfhJLIssYprltuMMKzKRDArpPsu.com", short: "brQKHslKXZ", seed: true},
	{val: "https://lcZHhdroabmgo.com", short: "yhnfGBRMkQ", seed: true},
	{val: "https://GhPOyhGKwQihLxMT.com", short: "BCxoxewAsP", seed: false},
	{val: "https://LihdcKbiW.com", short: "gHnKZOYxUR", seed: true},
	{val: "https://grRXLdvnNbpvMXGSqboUBNYYt.com", short: "oTVaLQPYcv", seed: true},
	{val: "https://ytEVxzfsmWviZAKQXbZUqvvHVdNAzOM.com", short: "KKBQCziAOV", seed: true},
	{val: "https://hxeRmNpMSanNwu.com", short: "qHdJSzZRSN", seed: false},
	{val: "https://VUjfRVLYBdloVXIBpbkLWXxFbbbnVV.com", short: "BSdwClqyoK", seed: true},
	{val: "https://UPecgcyB.com", short: "FcPPVMzSZY", seed: true},
	{val: "https://ObGeNTzdyfayykKMalUNqfokGoPwkel.com", short: "zNeToQddUy", seed: false},
	{val: "https://QadckEjzhUOSWjkqCiorsnPVWjByrqkGdwj.com", short: "LwnfRmkkjF", seed: false},
	{val: "https://liKMpfTHSntpmrCAHSH.com", short: "aPZhmSJXhu", seed: true},
	{val: "https://tNsgHHaDNfQMXPYQotllwvCKEexZD.com", short: "InFXnGzgmv", seed: false},
	{val: "https://lzLofxJcPDEbDDOyXlHyxZftMNxLA.com", short: "N_qWNkhsOI", seed: true},
}
