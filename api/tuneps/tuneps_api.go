package tuneps

/// ************************ Appel d'offres *******************************************
// AOGetAO
// Struct to represent the JSON response structure
type TunepsAOWebList struct {
	Code    string `json:"code"`
	Payload struct {
		Total int `json:"total"`
		Data  []struct {
			BdRecvEndDt   string `json:"bdRecvEndDt"`
			BidInstNm     string `json:"bidInstNm"`
			BidNmEn       string `json:"bidNmEn"`
			PublicDt      string `json:"publicDt"`
			PublicYn      string `json:"publicYn"`
			BidModSeq     string `json:"bidModSeq"`
			BidNo         string `json:"bidNo"`
			BidNmAr       string `json:"bidNmAr"`
			BidNmFr       string `json:"bidNmFr"`
			EpBidMasterID int    `json:"epBidMasterId"`
		} `json:"data"`
	} `json:"payload"`
}

// AOGetInfo
// Struct to represent the JSON response structure
type TunepsAOInfoWebList struct {
	Code    string `json:"code"`
	Payload struct {
		BdRecvEndDt             string `json:"bdRecvEndDt"`
		BiddingDocPriceCurr     string `json:"biddingDocPriceCurr"`
		EncCert                 string `json:"encCert"`
		GeneralYn               string `json:"generalYn"`
		BdRecvStartDt           string `json:"bdRecvStartDt"`
		BidNo                   string `json:"bidNo"`
		ExecutionPlaceStrEn     string `json:"executionPlaceStrEn"`
		ProcedureTypeStrEn      string `json:"procedureTypeStrEn"`
		PbkStrAr                string `json:"pbkStrAr"`
		GuaranteeTypeStrFr      string `json:"guaranteeTypeStrFr"`
		RealYnStrFr             string `json:"realYnStrFr"`
		FinancialMethodStrEn    string `json:"financialMethodStrEn"`
		BiddingDocPrice         int    `json:"biddingDocPrice"`
		RealYn                  string `json:"realYn"`
		BdOpenDt                string `json:"bdOpenDt"`
		BidNmAr                 string `json:"bidNmAr"`
		OnOffTypeStrEn          string `json:"onOffTypeStrEn"`
		PriceTypeStrFr          string `json:"priceTypeStrFr"`
		BizKindStrEn            string `json:"bizKindStrEn"`
		InternationalBidYnStrAr string `json:"internationalBidYnStrAr"`
		OnOffType               string `json:"onOffType"`
		InstRegNo               string `json:"instRegNo"`
		ConsorYn                string `json:"consorYn"`
		RegTypeStrFr            string `json:"regTypeStrFr"`
		ProcedureTypeStrFr      string `json:"procedureTypeStrFr"`
		PriceTypeStrEn          string `json:"priceTypeStrEn"`
		InternationalBidYnStrID int    `json:"internationalBidYnStrId"`
		FinancialMethodStrFr    string `json:"financialMethodStrFr"`
		BdDepartEn              string `json:"bdDepartEn"`
		BidExpiredDays          string `json:"bidExpiredDays"`
		ConsorYnStrFr           string `json:"consorYnStrFr"`
		PublicYnStrAr           string `json:"publicYnStrAr"`
		PublicYn                string `json:"publicYn"`
		PublicYnStrEn           string `json:"publicYnStrEn"`
		BizKindStrFr            string `json:"bizKindStrFr"`
		EvalMethodStrAr         string `json:"evalMethodStrAr"`
		ExecutionPlace          string `json:"executionPlace"`
		BidInstNm               string `json:"bidInstNm"`
		BdRecvStartDtStr        string `json:"bdRecvStartDtStr"`
		RegTypeStrAr            string `json:"regTypeStrAr"`
		StaffNm                 string `json:"staffNm"`
		BdRecvAddrsEn           string `json:"bdRecvAddrsEn"`
		FinancingYn             string `json:"financingYn"`
		RegType                 string `json:"regType"`
		ConsorYnStrEn           string `json:"consorYnStrEn"`
		InternationalBidYn      string `json:"internationalBidYn"`
		RealYnStrEn             string `json:"realYnStrEn"`
		ExamChmanIDStr          string `json:"examChmanIdStr"`
		PriceType               string `json:"priceType"`
		ExecutionPlaceStrFr     string `json:"executionPlaceStrFr"`
		BdOpenDtStr             string `json:"bdOpenDtStr"`
		FrameworkYn             string `json:"frameworkYn"`
		BdDepartFr              string `json:"bdDepartFr"`
		EvalMethodStrID         int    `json:"evalMethodStrId"`
		BizKind                 string `json:"bizKind"`
		BidModSeq               string `json:"bidModSeq"`
		GuaranteeTypeStrID      int    `json:"guaranteeTypeStrId"`
		BidNmFr                 string `json:"bidNmFr"`
		RegTypeStrID            int    `json:"regTypeStrId"`
		BdRecvAddrsFr           string `json:"bdRecvAddrsFr"`
		ProcedureTypeStrID      int    `json:"procedureTypeStrId"`
		FinancialMethodStrID    int    `json:"financialMethodStrId"`
		InternationalBidYnStrFr string `json:"internationalBidYnStrFr"`
		PbkStrEn                string `json:"pbkStrEn"`
		OnOffTypeStrID          int    `json:"onOffTypeStrId"`
		EpBidMasterID           int    `json:"epBidMasterId"`
		BdDepartAr              string `json:"bdDepartAr"`
		BidNmEn                 string `json:"bidNmEn"`
		FinancialMethod         string `json:"financialMethod"`
		OnOffTypeStrAr          string `json:"onOffTypeStrAr"`
		ExecutionPlaceStrAr     string `json:"executionPlaceStrAr"`
		ProcedureTypeStrAr      string `json:"procedureTypeStrAr"`
		GuaranteeTypeStrAr      string `json:"guaranteeTypeStrAr"`
		FinancialMethodStrAr    string `json:"financialMethodStrAr"`
		PbkStrFr                string `json:"pbkStrFr"`
		EvalMethodStrFr         string `json:"evalMethodStrFr"`
		PublicYnStrFr           string `json:"publicYnStrFr"`
		BizKindStrAr            string `json:"bizKindStrAr"`
		BidPoCd                 string `json:"bidPoCd"`
		InternationalBidYnStrEn string `json:"internationalBidYnStrEn"`
		BdRecvEndDtStr          string `json:"bdRecvEndDtStr"`
		PriceTypeStrID          int    `json:"priceTypeStrId"`
		GuaranteeType           string `json:"guaranteeType"`
		BizKindStrID            string `json:"bizKindStrId"`
		PriceTypeStrAr          string `json:"priceTypeStrAr"`
		RegID                   string `json:"regId"`
		PublicDtStr             string `json:"publicDtStr"`
		ExecutionPlaceStrID     int    `json:"executionPlaceStrId"`
		EvalMethod              string `json:"evalMethod"`
		HasAnnonce              string `json:"hasAnnonce"`
		RefNo                   string `json:"refNo"`
		ProcedureType           string `json:"procedureType"`
		PublicDt                string `json:"publicDt"`
		RegTypeStrEn            string `json:"regTypeStrEn"`
		ConsorYnStrID           int    `json:"consorYnStrId"`
		BdRecvAddrsAr           string `json:"bdRecvAddrsAr"`
		ProcChmanIDStr          string `json:"procChmanIdStr"`
		ConsorYnStrAr           string `json:"consorYnStrAr"`
		RealYnStrID             int    `json:"realYnStrId"`
		RealYnStrAr             string `json:"realYnStrAr"`
		ExamChmanID             string `json:"examChmanId"`
		GuaranteeTypeStrEn      string `json:"guaranteeTypeStrEn"`
		ProcChmanID             string `json:"procChmanId"`
		BidInstCd               string `json:"bidInstCd"`
		PbkStrID                int    `json:"pbkStrId"`
		OnOffTypeStrFr          string `json:"onOffTypeStrFr"`
		EvalMethodStrEn         string `json:"evalMethodStrEn"`
	} `json:"payload"`
}

// AOGetLot
// Struct to represent the JSON response structure
type TunepsAOLotWebList struct {
	Code    string `json:"code"`
	Payload []struct {
		EpBidTypCau          string        `json:"epBidTypCau"`
		BidClsDtlDesc        string        `json:"bidClsDtlDesc"`
		BidClsNm             string        `json:"bidClsNm"`
		CdNmAr               string        `json:"cdNmAr"`
		BidNo                string        `json:"bidNo"`
		GuaranteeAmount      int           `json:"guaranteeAmount"`
		BudgetAmtCurrF       string        `json:"budgetAmtCurrF"`
		CdNm                 string        `json:"cdNm"`
		BudgetAmtCurr        string        `json:"budgetAmtCurr"`
		BudgetAmt            int           `json:"budgetAmt"`
		GuaranteeAmountCurr  string        `json:"guaranteeAmountCurr"`
		BidModSeq            string        `json:"bidModSeq"`
		CdNmFr               string        `json:"cdNmFr"`
		GuaranteeAmountCurrF string        `json:"guaranteeAmountCurrF"`
		Affectationcheck     string        `json:"affectationcheck"`
		BidCls               string        `json:"bidCls"`
		ListLocalArticle     []interface{} `json:"listLocalArticle"`
		EpBidClsID           int           `json:"epBidClsId"`
		EpBidTypEval         string        `json:"epBidTypEval"`
	} `json:"payload"`
}

// AOGetCahierCharge
// Struct to represent the JSON response structure
type TunepsAOCCWebList struct {
	Code    string `json:"code"`
	Payload []struct {
		FileNm            string `json:"fileNm"`
		SeqNo             int    `json:"seqNo"`
		BidAttNodeRef     string `json:"bidAttNodeRef"`
		EpBidAttachFileID int    `json:"epBidAttachFileId"`
		BidModSeq         string `json:"bidModSeq"`
		CdNmAr            string `json:"cdNmAr"`
		BidNo             string `json:"bidNo"`
		CdNmFr            string `json:"cdNmFr"`
		DocCd             string `json:"docCd"`
		FileLoc           string `json:"fileLoc"`
		CdNm              string `json:"cdNm"`
	} `json:"payload"`
}

// AOGetAgrément
type TunepsAOAgWebList struct {
	Code    string `json:"code"`
	Payload []struct {
		LicenseCd      string `json:"licenseCd"`
		BidLicenseCd   string `json:"bidLicenseCd"`
		BidModSeq      string `json:"bidModSeq"`
		BidNo          string `json:"bidNo"`
		EpBidLicenseID int    `json:"epBidLicenseId"`
		BidCls         string `json:"bidCls"`
		Seq            string `json:"seq"`
	} `json:"payload"`
}

// AOGetPRoduit

// Struct to represent the JSON response structure
type TunepsAOPRWebList struct {
	Code    string `json:"code"`
	Payload []struct {
		DeliveryEndDtStr string `json:"deliveryEndDtStr"`
		CateNm           string `json:"cateNm"`
		Unit             string `json:"unit"`
		EpBidClsDtlID    int    `json:"epBidClsDtlId"`
		SeqNo            string `json:"seqNo"`
		DeliveryDays     string `json:"deliveryDays"`
		CateID           string `json:"cateId"`
		BidModSeq        string `json:"bidModSeq"`
		Qty              int    `json:"qty"`
		BidNo            string `json:"bidNo"`
		BidCls           string `json:"bidCls"`
	} `json:"payload"`
}

//************* BOUTON List soumissionaire(s) retenu(s)******************
// AOGetDonneDeBase
// Struct to represent the JSON response structure
type TunepsAODNBWebList struct {
	Code    string `json:"code"`
	Payload struct {
		BdRecvEndDt             string `json:"bdRecvEndDt"`
		BiddingDocPriceCurr     string `json:"biddingDocPriceCurr"`
		EncCert                 string `json:"encCert"`
		GeneralYn               string `json:"generalYn"`
		BdRecvStartDt           string `json:"bdRecvStartDt"`
		BidNo                   string `json:"bidNo"`
		ExecutionPlaceStrEn     string `json:"executionPlaceStrEn"`
		ProcedureTypeStrEn      string `json:"procedureTypeStrEn"`
		PbkStrAr                string `json:"pbkStrAr"`
		GuaranteeTypeStrFr      string `json:"guaranteeTypeStrFr"`
		RealYnStrFr             string `json:"realYnStrFr"`
		FinancialMethodStrEn    string `json:"financialMethodStrEn"`
		BiddingDocPrice         int    `json:"biddingDocPrice"`
		RealYn                  string `json:"realYn"`
		BdOpenDt                string `json:"bdOpenDt"`
		BidNmAr                 string `json:"bidNmAr"`
		OnOffTypeStrEn          string `json:"onOffTypeStrEn"`
		PriceTypeStrFr          string `json:"priceTypeStrFr"`
		BizKindStrEn            string `json:"bizKindStrEn"`
		InternationalBidYnStrAr string `json:"internationalBidYnStrAr"`
		OnOffType               string `json:"onOffType"`
		InstRegNo               string `json:"instRegNo"`
		ConsorYn                string `json:"consorYn"`
		RegTypeStrFr            string `json:"regTypeStrFr"`
		ProcedureTypeStrFr      string `json:"procedureTypeStrFr"`
		PriceTypeStrEn          string `json:"priceTypeStrEn"`
		InternationalBidYnStrID int    `json:"internationalBidYnStrId"`
		FinancialMethodStrFr    string `json:"financialMethodStrFr"`
		BdDepartEn              string `json:"bdDepartEn"`
		BidExpiredDays          string `json:"bidExpiredDays"`
		ConsorYnStrFr           string `json:"consorYnStrFr"`
		PublicYnStrAr           string `json:"publicYnStrAr"`
		PlanNmAr                string `json:"planNmAr"`
		PublicYn                string `json:"publicYn"`
		PublicYnStrEn           string `json:"publicYnStrEn"`
		BizKindStrFr            string `json:"bizKindStrFr"`
		EvalMethodStrAr         string `json:"evalMethodStrAr"`
		TpRecvAddrsFr           string `json:"tpRecvAddrsFr"`
		ExecutionPlace          string `json:"executionPlace"`
		BidInstNm               string `json:"bidInstNm"`
		BdRecvStartDtStr        string `json:"bdRecvStartDtStr"`
		RegTypeStrAr            string `json:"regTypeStrAr"`
		StaffNm                 string `json:"staffNm"`
		BdRecvAddrsEn           string `json:"bdRecvAddrsEn"`
		FinancingYn             string `json:"financingYn"`
		RegType                 string `json:"regType"`
		ConsorYnStrEn           string `json:"consorYnStrEn"`
		InternationalBidYn      string `json:"internationalBidYn"`
		RealYnStrEn             string `json:"realYnStrEn"`
		ExamChmanIDStr          string `json:"examChmanIdStr"`
		PriceType               string `json:"priceType"`
		ExecutionPlaceStrFr     string `json:"executionPlaceStrFr"`
		BdOpenDtStr             string `json:"bdOpenDtStr"`
		PlanModSeq              string `json:"planModSeq"`
		FrameworkYn             string `json:"frameworkYn"`
		BdDepartFr              string `json:"bdDepartFr"`
		EvalMethodStrID         int    `json:"evalMethodStrId"`
		BizKind                 string `json:"bizKind"`
		PqRecvAddrsEn           string `json:"pqRecvAddrsEn"`
		TpRecvAddrsEn           string `json:"tpRecvAddrsEn"`
		BidModSeq               string `json:"bidModSeq"`
		GuaranteeTypeStrID      int    `json:"guaranteeTypeStrId"`
		BidNmFr                 string `json:"bidNmFr"`
		RegTypeStrID            int    `json:"regTypeStrId"`
		BdRecvAddrsFr           string `json:"bdRecvAddrsFr"`
		ProcedureTypeStrID      int    `json:"procedureTypeStrId"`
		FinancialMethodStrID    int    `json:"financialMethodStrId"`
		InternationalBidYnStrFr string `json:"internationalBidYnStrFr"`
		PbkStrEn                string `json:"pbkStrEn"`
		OnOffTypeStrID          int    `json:"onOffTypeStrId"`
		EpBidMasterID           int    `json:"epBidMasterId"`
		PqRecvAddrsFr           string `json:"pqRecvAddrsFr"`
		BdDepartAr              string `json:"bdDepartAr"`
		BidNmEn                 string `json:"bidNmEn"`
		FinancialMethod         string `json:"financialMethod"`
		OnOffTypeStrAr          string `json:"onOffTypeStrAr"`
		ExecutionPlaceStrAr     string `json:"executionPlaceStrAr"`
		ProcedureTypeStrAr      string `json:"procedureTypeStrAr"`
		GuaranteeTypeStrAr      string `json:"guaranteeTypeStrAr"`
		FinancialMethodStrAr    string `json:"financialMethodStrAr"`
		PbkStrFr                string `json:"pbkStrFr"`
		EvalMethodStrFr         string `json:"evalMethodStrFr"`
		PublicYnStrFr           string `json:"publicYnStrFr"`
		BizKindStrAr            string `json:"bizKindStrAr"`
		BidPoCd                 string `json:"bidPoCd"`
		InternationalBidYnStrEn string `json:"internationalBidYnStrEn"`
		BdRecvEndDtStr          string `json:"bdRecvEndDtStr"`
		PriceTypeStrID          int    `json:"priceTypeStrId"`
		GuaranteeType           string `json:"guaranteeType"`
		BizKindStrID            string `json:"bizKindStrId"`
		PriceTypeStrAr          string `json:"priceTypeStrAr"`
		RegID                   string `json:"regId"`
		PublicDtStr             string `json:"publicDtStr"`
		ExecutionPlaceStrID     int    `json:"executionPlaceStrId"`
		EvalMethod              string `json:"evalMethod"`
		HasAnnonce              string `json:"hasAnnonce"`
		RefNo                   string `json:"refNo"`
		ProcedureType           string `json:"procedureType"`
		PublicDt                string `json:"publicDt"`
		PlanNmEn                string `json:"planNmEn"`
		RegTypeStrEn            string `json:"regTypeStrEn"`
		ConsorYnStrID           int    `json:"consorYnStrId"`
		BiddingDocPlace         string `json:"biddingDocPlace"`
		BdRecvAddrsAr           string `json:"bdRecvAddrsAr"`
		ProcChmanIDStr          string `json:"procChmanIdStr"`
		ConsorYnStrAr           string `json:"consorYnStrAr"`
		RealYnStrID             int    `json:"realYnStrId"`
		PqRecvAddrsAr           string `json:"pqRecvAddrsAr"`
		PlanNmFr                string `json:"planNmFr"`
		RealYnStrAr             string `json:"realYnStrAr"`
		ExamChmanID             string `json:"examChmanId"`
		GuaranteeTypeStrEn      string `json:"guaranteeTypeStrEn"`
		ProcChmanID             string `json:"procChmanId"`
		PlanNo                  string `json:"planNo"`
		TpRecvAddrsAr           string `json:"tpRecvAddrsAr"`
		BidInstCd               string `json:"bidInstCd"`
		PbkStrID                int    `json:"pbkStrId"`
		OnOffTypeStrFr          string `json:"onOffTypeStrFr"`
		EvalMethodStrEn         string `json:"evalMethodStrEn"`
	} `json:"payload"`
}

// AOGetAttributionDuMarché
// Struct to represent the JSON response structure
type TunepsAOATTMWebList struct {
	Code    string `json:"code"`
	Payload struct {
		Total int `json:"total"`
		Data  []struct {
			WinnerPublicDt      string `json:"winnerPublicDt"`
			NotWinnerYn         string `json:"notWinnerYn"`
			ProcedureType       string `json:"procedureType"`
			RefNo               string `json:"refNo"`
			BizRegNm            string `json:"bizRegNm"`
			BidNo               string `json:"bidNo"`
			TotalBidPriceCorr   string `json:"totalBidPriceCorr"`
			BizKind             string `json:"bizKind"`
			NotWinnerReason     string `json:"notWinnerReason"`
			BidPrice            string `json:"bidPrice"`
			BidNmEn             string `json:"bidNmEn"`
			RegDt               string `json:"regDt"`
			BidPriceSucc        string `json:"bidPriceSucc"`
			BidNmFr             string `json:" bidNmFr"`
			NbrConsorExist      string `json:"nbrConsorExist"`
			BidModSeq           string `json:"bidModSeq"`
			BidPriceSuccCurr    string `json:"bidPriceSuccCurr"`
			BidNmAr             string `json:"bidNmAr"`
			BizReg              string `json:"bizReg"`
			BidCls              string `json:"bidCls"`
			ExecType            string `json:"execType"`
			TotalBidPriceDcExch string `json:"totalBidPriceDcExch"`
			WinnerAfterDt       string `json:"winnerAfterDt"`
		} `json:"data"`
	} `json:"payload"`
}

// ************************bOUTON RESULTAT*************************************

// données de base
type TunepsAORDNBWebList struct {
	Code    string `json:"code"`
	Payload struct {
		BdRecvEndDt             string `json:"bdRecvEndDt"`
		BiddingDocPriceCurr     string `json:"biddingDocPriceCurr"`
		EncCert                 string `json:"encCert"`
		GeneralYn               string `json:"generalYn"`
		BdRecvStartDt           string `json:"bdRecvStartDt"`
		BidNo                   string `json:"bidNo"`
		ExecutionPlaceStrEn     string `json:"executionPlaceStrEn"`
		ProcedureTypeStrEn      string `json:"procedureTypeStrEn"`
		PbkStrAr                string `json:"pbkStrAr"`
		GuaranteeTypeStrFr      string `json:"guaranteeTypeStrFr"`
		RealYnStrFr             string `json:"realYnStrFr"`
		FinancialMethodStrEn    string `json:"financialMethodStrEn"`
		BiddingDocPrice         int    `json:"biddingDocPrice"`
		RealYn                  string `json:"realYn"`
		BdOpenDt                string `json:"bdOpenDt"`
		BidNmAr                 string `json:"bidNmAr"`
		OnOffTypeStrEn          string `json:"onOffTypeStrEn"`
		PriceTypeStrFr          string `json:"priceTypeStrFr"`
		BizKindStrEn            string `json:"bizKindStrEn"`
		InternationalBidYnStrAr string `json:"internationalBidYnStrAr"`
		OnOffType               string `json:"onOffType"`
		InstRegNo               string `json:"instRegNo"`
		ConsorYn                string `json:"consorYn"`
		RegTypeStrFr            string `json:"regTypeStrFr"`
		ProcedureTypeStrFr      string `json:"procedureTypeStrFr"`
		PriceTypeStrEn          string `json:"priceTypeStrEn"`
		InternationalBidYnStrID int    `json:"internationalBidYnStrId"`
		FinancialMethodStrFr    string `json:"financialMethodStrFr"`
		BdDepartEn              string `json:"bdDepartEn"`
		BidExpiredDays          string `json:"bidExpiredDays"`
		ConsorYnStrFr           string `json:"consorYnStrFr"`
		PublicYnStrAr           string `json:"publicYnStrAr"`
		PlanNmAr                string `json:"planNmAr"`
		PublicYn                string `json:"publicYn"`
		PublicYnStrEn           string `json:"publicYnStrEn"`
		BizKindStrFr            string `json:"bizKindStrFr"`
		EvalMethodStrAr         string `json:"evalMethodStrAr"`
		TpRecvAddrsFr           string `json:"tpRecvAddrsFr"`
		ExecutionPlace          string `json:"executionPlace"`
		BidInstNm               string `json:"bidInstNm"`
		BdRecvStartDtStr        string `json:"bdRecvStartDtStr"`
		RegTypeStrAr            string `json:"regTypeStrAr"`
		StaffNm                 string `json:"staffNm"`
		BdRecvAddrsEn           string `json:"bdRecvAddrsEn"`
		FinancingYn             string `json:"financingYn"`
		RegType                 string `json:"regType"`
		ConsorYnStrEn           string `json:"consorYnStrEn"`
		InternationalBidYn      string `json:"internationalBidYn"`
		RealYnStrEn             string `json:"realYnStrEn"`
		ExamChmanIDStr          string `json:"examChmanIdStr"`
		PriceType               string `json:"priceType"`
		ExecutionPlaceStrFr     string `json:"executionPlaceStrFr"`
		BdOpenDtStr             string `json:"bdOpenDtStr"`
		PlanModSeq              string `json:"planModSeq"`
		FrameworkYn             string `json:"frameworkYn"`
		BdDepartFr              string `json:"bdDepartFr"`
		EvalMethodStrID         int    `json:"evalMethodStrId"`
		BizKind                 string `json:"bizKind"`
		PqRecvAddrsEn           string `json:"pqRecvAddrsEn"`
		TpRecvAddrsEn           string `json:"tpRecvAddrsEn"`
		BidModSeq               string `json:"bidModSeq"`
		GuaranteeTypeStrID      int    `json:"guaranteeTypeStrId"`
		BidNmFr                 string `json:"bidNmFr"`
		RegTypeStrID            int    `json:"regTypeStrId"`
		BdRecvAddrsFr           string `json:"bdRecvAddrsFr"`
		ProcedureTypeStrID      int    `json:"procedureTypeStrId"`
		FinancialMethodStrID    int    `json:"financialMethodStrId"`
		InternationalBidYnStrFr string `json:"internationalBidYnStrFr"`
		PbkStrEn                string `json:"pbkStrEn"`
		OnOffTypeStrID          int    `json:"onOffTypeStrId"`
		EpBidMasterID           int    `json:"epBidMasterId"`
		PqRecvAddrsFr           string `json:"pqRecvAddrsFr"`
		BdDepartAr              string `json:"bdDepartAr"`
		BidNmEn                 string `json:"bidNmEn"`
		FinancialMethod         string `json:"financialMethod"`
		OnOffTypeStrAr          string `json:"onOffTypeStrAr"`
		ExecutionPlaceStrAr     string `json:"executionPlaceStrAr"`
		ProcedureTypeStrAr      string `json:"procedureTypeStrAr"`
		GuaranteeTypeStrAr      string `json:"guaranteeTypeStrAr"`
		FinancialMethodStrAr    string `json:"financialMethodStrAr"`
		PbkStrFr                string `json:"pbkStrFr"`
		EvalMethodStrFr         string `json:"evalMethodStrFr"`
		PublicYnStrFr           string `json:"publicYnStrFr"`
		BizKindStrAr            string `json:"bizKindStrAr"`
		BidPoCd                 string `json:"bidPoCd"`
		InternationalBidYnStrEn string `json:"internationalBidYnStrEn"`
		BdRecvEndDtStr          string `json:"bdRecvEndDtStr"`
		PriceTypeStrID          int    `json:"priceTypeStrId"`
		GuaranteeType           string `json:"guaranteeType"`
		BizKindStrID            string `json:"bizKindStrId"`
		PriceTypeStrAr          string `json:"priceTypeStrAr"`
		RegID                   string `json:"regId"`
		PublicDtStr             string `json:"publicDtStr"`
		ExecutionPlaceStrID     int    `json:"executionPlaceStrId"`
		EvalMethod              string `json:"evalMethod"`
		HasAnnonce              string `json:"hasAnnonce"`
		RefNo                   string `json:"refNo"`
		ProcedureType           string `json:"procedureType"`
		PublicDt                string `json:"publicDt"`
		PlanNmEn                string `json:"planNmEn"`
		RegTypeStrEn            string `json:"regTypeStrEn"`
		ConsorYnStrID           int    `json:"consorYnStrId"`
		BiddingDocPlace         string `json:"biddingDocPlace"`
		BdRecvAddrsAr           string `json:"bdRecvAddrsAr"`
		ProcChmanIDStr          string `json:"procChmanIdStr"`
		ConsorYnStrAr           string `json:"consorYnStrAr"`
		RealYnStrID             int    `json:"realYnStrId"`
		PqRecvAddrsAr           string `json:"pqRecvAddrsAr"`
		PlanNmFr                string `json:"planNmFr"`
		RealYnStrAr             string `json:"realYnStrAr"`
		ExamChmanID             string `json:"examChmanId"`
		GuaranteeTypeStrEn      string `json:"guaranteeTypeStrEn"`
		ProcChmanID             string `json:"procChmanId"`
		PlanNo                  string `json:"planNo"`
		TpRecvAddrsAr           string `json:"tpRecvAddrsAr"`
		BidInstCd               string `json:"bidInstCd"`
		PbkStrID                int    `json:"pbkStrId"`
		OnOffTypeStrFr          string `json:"onOffTypeStrFr"`
		EvalMethodStrEn         string `json:"evalMethodStrEn"`
	} `json:"payload"`
}

// Lot-Resultat
type TunepsAORLotWebList struct {
	Code    string `json:"code"`
	Payload struct {
		Total int `json:"total"`
		Data  []struct {
			RefNo         string `json:"refNo"`
			BizKindFrStr  string `json:"bizKindFrStr"`
			BizKindArStr  string `json:"bizKindArStr"`
			BidNo         string `json:"bidNo"`
			Cnt           string `json:"cnt"`
			BizKindEnStr  string `json:"bizKindEnStr"`
			BizKind       string `json:"bizKind"`
			BidNmEn       string `json:"bidNmEn"`
			RegDt         string `json:"regDt"`
			BidModSeq     string `json:"bidModSeq"`
			BidNmAr       string `json:"bidNmAr"`
			Rk            string `json:"rk"`
			BidFailReason string `json:"bidFailReason"`
			RegID         string `json:"regId"`
			BidCls        string `json:"bidCls"`
			ExecType      string `json:"execType"`
			BidNmFr       string `json:"bidNmFr"`
		} `json:"data"`
	} `json:"payload"`
}

// Soumissionaires
type TunepsAORSWebList struct {
	Code    string `json:"code"`
	Payload struct {
		Total int `json:"total"`
		Data  []struct {
			CdNmStrEn           string `json:"cdNmStrEn"`
			IneligibleCd        string `json:"ineligibleCd"`
			TechEvalResult      string `json:"techEvalResult"`
			BidNo               string `json:"bidNo"`
			CdPfTechFr          string `json:"cdPfTechFr"`
			BdRecvDtStr         string `json:"bdRecvDtStr"`
			EpBidOpenResultID   int    `json:"epBidOpenResultId"`
			CdPfFinaFr          string `json:"cdPfFinaFr"`
			Rank                int    `json:"rank,omitempty"`
			OnlineRegYn         string `json:"onlineRegYn"`
			DecAttachFileCnt    string `json:"decAttachFileCnt"`
			BidCls              string `json:"bidCls"`
			TotalBidPriceDcExch int    `json:"totalBidPriceDcExch"`
			BizRegNo            string `json:"bizRegNo"`
			BizRegNm            string `json:"bizRegNm"`
			CdPfTechEn          string `json:"cdPfTechEn"`
			FinaResultReason    string `json:"finaResultReason"`
			CdNmStrFr           string `json:"cdNmStrFr"`
			Cnt                 int    `json:"cnt"`
			CdPfTechAr          string `json:"cdPfTechAr"`
			IneligibleReason    string `json:"ineligibleReason"`
			CdPfFinaAr          string `json:"cdPfFinaAr"`
			BdRecvDt            string `json:"bdRecvDt"`
			BidModSeq           string `json:"bidModSeq"`
			CdNmStrAr           string `json:"cdNmStrAr"`
			Rk                  int    `json:"rk"`
			BizReg              string `json:"bizReg"`
			FinaEvalResult      string `json:"finaEvalResult"`
			TechResultReason    string `json:"techResultReason"`
			CdPfFinaEn          string `json:"cdPfFinaEn"`
		} `json:"data"`
	} `json:"payload"`
}

/// ************************ Shopping Mall *******************************************

//SMGetSM
// Struct to represent the JSON response structure

type TunepsSMWebList struct {
	Code    string `json:"code"`
	Payload struct {
		Total int `json:"total"`
		Data  []struct {
			PublicDt       string `json:"publicDt"`
			SpShopMasterID int    `json:"spShopMasterId"`
			ShopModSeq     string `json:"shopModSeq"`
			InstNm         string `json:"instNm"`
			ShopNmEn       string `json:"shopNmEn"`
			ShopNmFr       string `json:"shopNmFr"`
			SpRecvEndDt    string `json:"spRecvEndDt"`
			ShopNo         string `json:"shopNo"`
			ShopNmAr       string `json:"shopNmAr"`
		} `json:"data"`
	} `json:"payload"`
}

// SM Informations générales
type TunepsSMInfoWebList struct {
	Code    string `json:"code"`
	Payload struct {
		RefNo               string `json:"refNo"`
		PublicDt            string `json:"publicDt"`
		ShopModSeq          string `json:"shopModSeq"`
		PublicYn            string `json:"publicYn"`
		MsAr                string `json:"msAr"`
		PbkStrEn            string `json:"pbkStrEn"`
		RegNm               string `json:"regNm"`
		BizKindStrFr        string `json:"bizKindStrFr"`
		CdNmAr              string `json:"cdNmAr"`
		PbkStrAr            string `json:"pbkStrAr"`
		ShopNmAr            string `json:"shopNmAr"`
		BcNm                string `json:"bcNm"`
		BcID                string `json:"bcId"`
		EpAr                string `json:"epAr"`
		InstNm              string `json:"instNm"`
		InternationalShopYn string `json:"internationalShopYn"`
		SpRecvEndDtStr      string `json:"spRecvEndDtStr"`
		ShopExpiredDays     int    `json:"shopExpiredDays"`
		SpRecvStartDt       string `json:"spRecvStartDt"`
		EpEn                string `json:"epEn"`
		ExamNm              string `json:"examNm"`
		MsEn                string `json:"msEn"`
		CdNmEn              string `json:"cdNmEn"`
		ShopInstCd          string `json:"shopInstCd"`
		PbkStrFr            string `json:"pbkStrFr"`
		MsFr                string `json:"msFr"`
		PbkCd               int    `json:"pbkCd"`
		BizKindStrEn        string `json:"bizKindStrEn"`
		ShopNmFr            string `json:"shopNmFr"`
		MasterCdBc          string `json:"masterCdBc"`
		BizKindStrAr        string `json:"bizKindStrAr"`
		SpRecvEndDt         string `json:"spRecvEndDt"`
		SpShopMasterID      int    `json:"spShopMasterId"`
		BizKindCd           int    `json:"bizKindCd"`
		BizKindStrID        int    `json:"bizKindStrId"`
		ShopNmEn            string `json:"shopNmEn"`
		PbkStrID            int    `json:"pbkStrId"`
		ExamID              string `json:"examId"`
		EpFr                string `json:"epFr"`
		CdNmFr              string `json:"cdNmFr"`
		RegID               string `json:"regId"`
		PublicDtStr         string `json:"publicDtStr"`
		MasterCdExam        string `json:"masterCdExam"`
		ShopNo              string `json:"shopNo"`
		SpOpenDt            string `json:"spOpenDt"`
		HasAnnonce          string `json:"hasAnnonce"`
	} `json:"payload"`
}

//lot
type TunepsSMLotWebList struct {
	Code    string `json:"code"`
	Payload []struct {
		ShopCls             string `json:"shopCls"`
		BudgetAmt           string `json:"budgetAmt"`
		ShopClsDtlDesc      string `json:"shopClsDtlDesc"`
		ShopModSeq          string `json:"shopModSeq"`
		GuaranteeAmountCurr string `json:"guaranteeAmountCurr"`
		GuaranteeAmount     string `json:"guaranteeAmount"`
		ShopNo              string `json:"shopNo"`
		ShopClsNm           string `json:"shopClsNm"`
		BudgetAmtCurr       string `json:"budgetAmtCurr"`
		SpShopClsID         int    `json:"spShopClsId"`
	} `json:"payload"`
}

//Documents de la consultation
type TunepsSMDocWebList struct {
	Code    string `json:"code"`
	Payload []struct {
		FileNm             string `json:"fileNm"`
		ShopModSeq         string `json:"shopModSeq"`
		SeqNo              int    `json:"seqNo"`
		NodeRef            string `json:"nodeRef"`
		CdNmAr             string `json:"cdNmAr"`
		CdNmFr             string `json:"cdNmFr"`
		DocCd              string `json:"docCd"`
		FileLoc            string `json:"fileLoc"`
		ShopNo             string `json:"shopNo"`
		SpShopAttachFileID int    `json:"spShopAttachFileId"`
		CdNmEn             string `json:"cdNmEn"`
	} `json:"payload"`
}
