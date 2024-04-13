package domains

import "github.com/google/uuid"

// ************************ Appel d'offres *******************************************
// Struct to represent the database model
type TunepsAOList struct {
	ID            uuid.UUID `gorm:"column:id; primaryKey; type:uuid; not null;"` // Unique identifier for the role
	BdRecvEndDt   string    `gorm:"column:bd_recv_end_dt;not null;"`
	BidInstNm     string    `gorm:"column:bid_inst_nm;not null;"`
	BidNmEn       string    `gorm:"column:bid_nm_en;not null;"`
	PublicDt      string    `gorm:"column:public_dt;not null;"`
	Closed        bool      `gorm:"column:closed; not null; default:false"`
	PublicYn      string    `gorm:"column:public_yn;not null;"`
	BidModSeq     string    `gorm:"column:bid_mod_seq;not null;"`
	BidNo         string    `gorm:"column:bid_no;not null;"`
	BidNmAr       string    `gorm:"column:bid_nm_ar;not null;"`
	BidNmFr       string    `gorm:"column:bid_nm_fr;not null;"`
	EpBidMasterID int       `gorm:"column:ep_bid_master_id;not null;"`
}

// Struct to represent the database model
type TunepsAOInfoList struct {
	ID                      uuid.UUID `gorm:"column:id; primaryKey; type:uuid; not null;"`
	BdRecvEndDt             string    `gorm:"column:bd_recv_end_dt ;not null;"`
	BiddingDocPriceCurr     string    `gorm:"column:bidding_doc_price_curr;not null;"`
	EncCert                 string    `gorm:"column:enc_cert ;not null;"`
	GeneralYn               string    `gorm:"column:general_yn ;not null;"`
	BdRecvStartDt           string    `gorm:"column:bd_recv_start_dt ;not null;"`
	BidNo                   string    `gorm:"column:bid_no ;not null;"`
	ExecutionPlaceStrEn     string    `gorm:"column:execution_place_str_en ;not null;"`
	ProcedureTypeStrEn      string    `gorm:"column:procedure_type_str_en ;not null;"`
	PbkStrAr                string    `gorm:"column:pbk_str_ar ;not null;"`
	GuaranteeTypeStrFr      string    `gorm:"column:guarantee_type_str_fr ;not null;"`
	RealYnStrFr             string    `gorm:"column:real_yn_str_fr ;not null;"`
	FinancialMethodStrEn    string    `gorm:"column:financial_method_str_en ;not null;"`
	BiddingDocPrice         int       `gorm:"column:bidding_doc_price ;not null;"`
	RealYn                  string    `gorm:"column:real_yn ;not null;"`
	BdOpenDt                string    `gorm:"column:bd_open_dt ;not null;"`
	BidNmAr                 string    `gorm:"column:bid_nm_ar ;not null;"`
	OnOffTypeStrEn          string    `gorm:"column:on_off_type_str_en ;not null;"`
	PriceTypeStrFr          string    `gorm:"column:price_type_str_fr ;not null;"`
	BizKindStrEn            string    `gorm:"column:biz_kind_str_en ;not null;"`
	InternationalBidYnStrAr string    `gorm:"column:international_bid_yn_str_ar;not null;"`
	OnOffType               string    `gorm:"column:on_off_type;not null;"`
	InstRegNo               string    `gorm:"column:inst_reg_no;not null;"`
	ConsorYn                string    `gorm:"column:consor_yn;not null;"`
	RegTypeStrFr            string    `gorm:"column:reg_type_str_fr ;not null;"`
	ProcedureTypeStrFr      string    `gorm:"column:procedure_type_str_fr ;not null;"`
	PriceTypeStrEn          string    `gorm:"column:price_type_str_en;not null;"`
	InternationalBidYnStrID int       `gorm:"column:international_bid_yn_str_id ;not null;"`
	FinancialMethodStrFr    string    `gorm:"column:financial_method_str_fr;not null;"`
	BdDepartEn              string    `gorm:"column:bd_depart_en;not null;"`
	BidExpiredDays          string    `gorm:"column:bid_expired_days;not null;"`
	ConsorYnStrFr           string    `gorm:"column:consor_yn_str_fr;not null;"`
	PublicYnStrAr           string    `gorm:"column:public_yn_str_ar;not null;"`
	PublicYn                string    `gorm:"column:public_yn;not null;"`
	PublicYnStrEn           string    `gorm:"column:public_yn_str_en ;not null;"`
	BizKindStrFr            string    `gorm:"column:biz_kind_str_fr ;not null;"`
	EvalMethodStrAr         string    `gorm:"column:eval_method_str_ar;not null;"`
	ExecutionPlace          string    `gorm:"column:execution_place;not null;"`
	BidInstNm               string    `gorm:"column:bid_inst_nm;not null;"`
	BdRecvStartDtStr        string    `gorm:"column:bd_recv_start_dt_str;not null;"`
	RegTypeStrAr            string    `gorm:"column:reg_type_str_ar;not null;"`
	StaffNm                 string    `gorm:"column:staff_nm;not null;"`
	BdRecvAddrsEn           string    `gorm:"column:bd_recv_addrs_en;not null;"`
	FinancingYn             string    `gorm:"column:financing_yn;not null;"`
	RegType                 string    `gorm:"column:reg_type;not null;"`
	ConsorYnStrEn           string    `gorm:"column:consor_yn_str_en;not null;"`
	InternationalBidYn      string    `gorm:"column:international_bid_yn;not null;"`
	RealYnStrEn             string    `gorm:"column:real_yn_str_en;not null;"`
	ExamChmanIDStr          string    `gorm:"column:exam_chman_id_str;not null;"`
	PriceType               string    `gorm:"column:price_type;not null;"`
	ExecutionPlaceStrFr     string    `gorm:"column:execution_place_str_fr;not null;"`
	BdOpenDtStr             string    `gorm:"column:bd_open_dt_strv;not null;"`
	FrameworkYn             string    `gorm:"column:framework_yn;not null;"`
	BdDepartFr              string    `gorm:"column:bd_depart_fr;not null;"`
	EvalMethodStrID         int       `gorm:"column:eval_method_str_id;not null;"`
	BizKind                 string    `gorm:"column:biz_kind;not null;"`
	BidModSeq               string    `gorm:"column:bid_mod_seq;not null;"`
	GuaranteeTypeStrID      int       `gorm:"column:guarantee_type_str_id;not null;"`
	BidNmFr                 string    `gorm:"column:bid_nm_fr;not null;not null;"`
	RegTypeStrID            int       `gorm:"column:reg_type_str_id;not null;"`
	BdRecvAddrsFr           string    `gorm:"column:bd_recv_addrs_fr;not null;"`
	ProcedureTypeStrID      int       `gorm:"column:procedure_type_str_id;not null;"`
	FinancialMethodStrID    int       `gorm:"column:financial_method_str_id;not null;"`
	InternationalBidYnStrFr string    `gorm:"column:international_bid_yn_str_fr;not null;"`
	PbkStrEn                string    `gorm:"column:pbk_str_en;not null;"`
	OnOffTypeStrID          int       `gorm:"column:on_off_type_str_id;not null;"`
	EpBidMasterID           int       `gorm:"column:ep_bid_master_id;not null;"`
	BdDepartAr              string    `gorm:"column:bd_depart_ar;not null;"`
	BidNmEn                 string    `gorm:"column:bid_nm_en;not null;"`
	FinancialMethod         string    `gorm:"column:financial_method;not null;"`
	OnOffTypeStrAr          string    `gorm:"column:on_off_type_str_ar;not null;"`
	ExecutionPlaceStrAr     string    `gorm:"column:execution_place_str_ar;not null;"`
	ProcedureTypeStrAr      string    `gorm:"column:procedure_type_str_ar;not null;"`
	GuaranteeTypeStrAr      string    `gorm:"column:guarantee_type_str_ar;not null;"`
	FinancialMethodStrAr    string    `gorm:"column:financial_method_str_ar;not null;"`
	PbkStrFr                string    `gorm:"column:pbk_str_fr;not null;"`
	EvalMethodStrFr         string    `gorm:"column:eval_method_str_fr;not null;"`
	PublicYnStrFr           string    `gorm:"column:public_yn_str_fr;not null;"`
	BizKindStrAr            string    `gorm:"column:biz_kind_str_ar;not null;"`
	BidPoCd                 string    `gorm:"column:bid_po_cd;not null;"`
	InternationalBidYnStrEn string    `gorm:"column:international_bid_yn_str_en;not null;"`
	BdRecvEndDtStr          string    `gorm:"column:bd_recv_end_dt_str;not null;"`
	PriceTypeStrID          int       `gorm:"column:price_type_str_id;not null;"`
	GuaranteeType           string    `gorm:"column:guarantee_type;not null;"`
	BizKindStrID            string    `gorm:"column:biz_kind_str_id;not null;"`
	PriceTypeStrAr          string    `gorm:"column:price_type_str_ar;not null;"`
	RegID                   string    `gorm:"column:reg_id;not null;"`
	PublicDtStr             string    `gorm:"column:public_dt_str;not null;"`
	ExecutionPlaceStrID     int       `gorm:"column:execution_place_str_id;not null;"`
	EvalMethod              string    `gorm:"column:eval_method;not null;"`
	HasAnnonce              string    `gorm:"column:has_annonce;not null;"`
	RefNo                   string    `gorm:"column:ref_no;not null;"`
	ProcedureType           string    `gorm:"column:procedure_type;not null;"`
	PublicDt                string    `gorm:"column:public_dt;not null;"`
	RegTypeStrEn            string    `gorm:"column:reg_type_str_en;not null;"`
	ConsorYnStrID           int       `gorm:"column:consor_yn_str_id;not null;"`
	BdRecvAddrsAr           string    `gorm:"column:bd_recv_addrs_ar;not null;"`
	ProcChmanIDStr          string    `gorm:"column:proc_chman_id_str;not null;"`
	ConsorYnStrAr           string    `gorm:"column:consor_yn_str_ar;not null;"`
	RealYnStrID             int       `gorm:"column:real_yn_str_id;not null;"`
	RealYnStrAr             string    `gorm:"column:real_yn_str_ar;not null;"`
	ExamChmanID             string    `gorm:"column:exam_chman_id;not null;"`
	GuaranteeTypeStrEn      string    `gorm:"column:guarantee_type_str_en;not null;"`
	ProcChmanID             string    `gorm:"column:proc_chman_id;not null;"`
	BidInstCd               string    `gorm:"column:bid_inst_cd;not null;"`
	PbkStrID                int       `gorm:"column:pbk_str_id;not null;"`
	OnOffTypeStrFr          string    `gorm:"column:on_off_type_str_fr;not null;"`
	EvalMethodStrEn         string    `gorm:"column:eval_method_str_en;not null;"`
}

// Struct to represent the database model
type TunepsAOLotList struct {
	ID                   uuid.UUID `gorm:"column:id; primaryKey; type:uuid; not null;"`
	EpBidTypCau          string    `gorm:"column:ep_bid_typ_cau;not null"`
	BidClsDtlDesc        string    `gorm:"column:bid_cls_dtl_desc;not null"`
	BidClsNm             string    `gorm:"column:bid_cls_nm;not null"`
	CdNmAr               string    `gorm:"column:cd_nm_ar;not null"`
	BidNo                string    `gorm:"column:bid_no;not null"`
	GuaranteeAmount      int       `gorm:"column:guarantee_amount;not null"`
	BudgetAmtCurrF       string    `gorm:"column:budget_amt_curr_f;not null"`
	CdNm                 string    `gorm:"column:cd_nm;not null"`
	BudgetAmtCurr        string    `gorm:"column:budget_amt_curr;not null"`
	BudgetAmt            int       `gorm:"column:budget_amt;not null"`
	GuaranteeAmountCurr  string    `gorm:"column:guarantee_amount_curr;not null"`
	BidModSeq            string    `gorm:"column:bid_mod_seq;not null"`
	CdNmFr               string    `gorm:"column:cd_nm_fr;not null"`
	GuaranteeAmountCurrF string    `gorm:"column:guarantee_amount_curr_f;not null"`
	Affectationcheck     string    `gorm:"column:affectation_check;not null"`
	BidCls               string    `gorm:"column:bid_cls;not null"`
	EpBidClsID           int       `gorm:"column:ep_bid_cls_id;not null"`
	EpBidTypEval         string    `gorm:"column:ep_bid_typ_eval;not null"`
}

// Struct to represent the database model
type TunepsAOCCList struct {
	ID                uuid.UUID `gorm:"column:id; primaryKey; type:uuid; not null;"`
	FileNm            string    `gorm:"column:file_nm;not null"`
	SeqNo             int       `gorm:"column:seq_no;not null"`
	BidAttNodeRef     string    `gorm:"column:bid_att_node_ref;not null"`
	EpBidAttachFileID int       `gorm:"column:ep_bid_attach_file_id;not null"`
	BidModSeq         string    `gorm:"column:bid_mod_seq;not null"`
	CdNmAr            string    `gorm:"column:cd_nm_ar;not null"`
	BidNo             string    `gorm:"column:bid_no;not null"`
	CdNmFr            string    `gorm:"column:cd_nm_fr;not null"`
	DocCd             string    `gorm:"column:doc_cd;not null"`
	FileLoc           string    `gorm:"column:file_loc;not null"`
	CdNm              string    `gorm:"column:cd_nm;not null"`
}

// AOGetAgrément
// Struct to represent the database model
type TunepsAOAgList struct {
	ID             uuid.UUID `gorm:"column:id; primaryKey; type:uuid; not null;"`
	LicenseCd      string    `gorm:"column:license_cd;not null"`
	BidLicenseCd   string    `gorm:"column:bid_license_cd;not null"`
	BidModSeq      string    `gorm:"column:bid_mod_seq;not null"`
	BidNo          string    `gorm:"column:bid_no;not null"`
	EpBidLicenseID int       `gorm:"column:ep_bid_license_id;not null"`
	BidCls         string    `gorm:"column:bid_cls;not null"`
	Seq            string    `gorm:"column:seq;not null"`
}

// Struct to represent the database model
type TunepsAOPRList struct {
	ID               uuid.UUID `gorm:"column:id; primaryKey; type:uuid; not null;"`
	DeliveryEndDtStr string    `gorm:"column:delivery_end_dt_str;not null"`
	CateNm           string    `gorm:"column:cate_nm;not null"`
	Unit             string    `gorm:"column:unit;not null"`
	EpBidClsDtlID    int       `gorm:"column:ep_bid_cls_dtl_id;not null"`
	SeqNo            string    `gorm:"column:seq_no;not null"`
	DeliveryDays     string    `gorm:"column:delivery_days;not null"`
	CateID           string    `gorm:"column:cate_id;not null"`
	BidModSeq        string    `gorm:"column:bid_mod_seq;not null"`
	Qty              int       `gorm:"column:qty;not null"`
	BidNo            string    `gorm:"column:bid_no;not null"`
	BidCls           string    `gorm:"column:bid_cls;not null"`
}

//************* BOUTON List soumissionaire(s) retenu(s)******************

// Struct to represent the database model
type TunepsAODNBList struct {
	ID                      uuid.UUID `gorm:"column:id; primaryKey; type:uuid; not null;"`
	BdRecvEndDt             string    `gorm:"column:bd_recv_end_dt;not null"`
	BiddingDocPriceCurr     string    `gorm:"column:bidding_doc_price_curr;not null"`
	EncCert                 string    `gorm:"column:enc_cert;not null"`
	GeneralYn               string    `gorm:"column:general_yn;not null"`
	BdRecvStartDt           string    `gorm:"column:bd_recv_start_dt;not null"`
	BidNo                   string    `gorm:"column:bid_no;not null"`
	ExecutionPlaceStrEn     string    `gorm:"column:execution_place_str_en;not null"`
	ProcedureTypeStrEn      string    `gorm:"column:procedure_type_str_en;not null"`
	PbkStrAr                string    `gorm:"column:pbk_str_ar;not null"`
	GuaranteeTypeStrFr      string    `gorm:"column:guarantee_type_str_fr;not null"`
	RealYnStrFr             string    `gorm:"column:real_yn_str_fr;not null"`
	FinancialMethodStrEn    string    `gorm:"column:financial_method_str_en;not null"`
	BiddingDocPrice         int       `gorm:"column:bidding_doc_price;not null"`
	RealYn                  string    `gorm:"column:real_yn;not null"`
	BdOpenDt                string    `gorm:"column:bd_open_dt;not null"`
	BidNmAr                 string    `gorm:"column:bid_nm_ar;not null"`
	OnOffTypeStrEn          string    `gorm:"column:on_off_type_str_en;not null"`
	PriceTypeStrFr          string    `gorm:"column:price_type_str_fr;not null"`
	BizKindStrEn            string    `gorm:"column:biz_kind_str_en;not null"`
	InternationalBidYnStrAr string    `gorm:"column:international_bid_yn_str_ar;not null"`
	OnOffType               string    `gorm:"column:on_off_type;not null"`
	InstRegNo               string    `gorm:"column:inst_reg_no;not null"`
	ConsorYn                string    `gorm:"column:consor_yn;not null"`
	RegTypeStrFr            string    `gorm:"column:reg_type_str_fr;not null"`
	ProcedureTypeStrFr      string    `gorm:"column:procedure_type_str_fr;not null"`
	PriceTypeStrEn          string    `gorm:"column:price_type_str_en;not null"`
	InternationalBidYnStrID int       `gorm:"column:international_bid_yn_str_id;not null"`
	FinancialMethodStrFr    string    `gorm:"column:financial_method_str_fr;not null"`
	BdDepartEn              string    `gorm:"column:bd_depart_en;not null"`
	BidExpiredDays          string    `gorm:"column:bid_expired_days;not null"`
	ConsorYnStrFr           string    `gorm:"column:consor_yn_str_fr;not null"`
	PublicYnStrAr           string    `gorm:"column:public_yn_str_ar;not null"`
	PlanNmAr                string    `gorm:"column:plan_nm_ar;not null"`
	PublicYn                string    `gorm:"column:public_yn;not null"`
	PublicYnStrEn           string    `gorm:"column:public_yn_str_en;not null"`
	BizKindStrFr            string    `gorm:"column:biz_kind_str_fr;not null"`
	EvalMethodStrAr         string    `gorm:"column:eval_method_str_ar;not null"`
	TpRecvAddrsFr           string    `gorm:"column:tp_recv_addrs_fr;not null"`
	ExecutionPlace          string    `gorm:"column:execution_place;not null"`
	BidInstNm               string    `gorm:"column:bid_inst_nm;not null"`
	BdRecvStartDtStr        string    `gorm:"column:bd_recv_start_dt_str;not null"`
	RegTypeStrAr            string    `gorm:"column:reg_type_str_ar;not null"`
	StaffNm                 string    `gorm:"column:staff_nm;not null"`
	BdRecvAddrsEn           string    `gorm:"column:bd_recv_addrs_en;not null"`
	FinancingYn             string    `gorm:"column:financing_yn;not null"`
	RegType                 string    `gorm:"column:reg_type;not null"`
	ConsorYnStrEn           string    `gorm:"column:consor_yn_str_en;not null"`
	InternationalBidYn      string    `gorm:"column:international_bid_yn;not null"`
	RealYnStrEn             string    `gorm:"column:real_yn_str_en;not null"`
	ExamChmanIDStr          string    `gorm:"column:exam_chman_id_str;not null"`
	PriceType               string    `gorm:"column:price_type;not null"`
	ExecutionPlaceStrFr     string    `gorm:"column:execution_place_str_fr;not null"`
	BdOpenDtStr             string    `gorm:"column:bd_open_dt_str;not null"`
	PlanModSeq              string    `gorm:"column:plan_mod_seq;not null"`
	FrameworkYn             string    `gorm:"column:framework_yn;not null"`
	BdDepartFr              string    `gorm:"column:bd_depart_fr;not null"`
	EvalMethodStrID         int       `gorm:"column:eval_method_str_id;not null"`
	BizKind                 string    `gorm:"column:biz_kind;not null"`
	PqRecvAddrsEn           string    `gorm:"column:pq_recv_addrs_en;not null"`
	TpRecvAddrsEn           string    `gorm:"column:tp_recv_addrs_en;not null"`
	BidModSeq               string    `gorm:"column:bid_mod_seq;not null"`
	GuaranteeTypeStrID      int       `gorm:"column:guarantee_type_str_id;not null"`
	BidNmFr                 string    `gorm:"column:bid_nm_fr;not null"`
	RegTypeStrID            int       `gorm:"column:reg_type_str_id;not null"`
	BdRecvAddrsFr           string    `gorm:"column:bd_recv_addrs_fr;not null"`
	ProcedureTypeStrID      int       `gorm:"column:procedure_type_str_id;not null"`
	FinancialMethodStrID    int       `gorm:"column:financial_method_str_id;not null"`
	InternationalBidYnStrFr string    `gorm:"column:international_bid_yn_str_fr;not null"`
	PbkStrEn                string    `gorm:"column:pbk_str_en;not null"`
	OnOffTypeStrID          int       `gorm:"column:on_off_type_str_id;not null"`
	EpBidMasterID           int       `gorm:"column:ep_bid_master_id;not null"`
	PqRecvAddrsFr           string    `gorm:"column:pq_recv_addrs_fr;not null"`
	BdDepartAr              string    `gorm:"column:bd_depart_ar;not null"`
	BidNmEn                 string    `gorm:"column:bid_nm_en;not null"`
	FinancialMethod         string    `gorm:"column:financial_method;not null"`
	OnOffTypeStrAr          string    `gorm:"column:on_off_type_str_ar;not null"`
	ExecutionPlaceStrAr     string    `gorm:"column:execution_place_str_ar;not null"`
	ProcedureTypeStrAr      string    `gorm:"column:procedure_type_str_ar;not null"`
	GuaranteeTypeStrAr      string    `gorm:"column:guarantee_type_str_ar;not null"`
	FinancialMethodStrAr    string    `gorm:"column:financial_method_str_ar;not null"`
	PbkStrFr                string    `gorm:"column:pbk_str_fr;not null"`
	EvalMethodStrFr         string    `gorm:"column:eval_method_str_fr;not null"`
	PublicYnStrFr           string    `gorm:"column:public_yn_str_fr;not null"`
	BizKindStrAr            string    `gorm:"column:biz_kind_str_ar;not null"`
	BidPoCd                 string    `gorm:"column:bid_po_cd;not null"`
	InternationalBidYnStrEn string    `gorm:"column:international_bid_yn_str_en;not null"`
	BdRecvEndDtStr          string    `gorm:"column:bd_recv_end_dt_str;not null"`
	PriceTypeStrID          int       `gorm:"column:price_type_str_id;not null"`
	GuaranteeType           string    `gorm:"column:guarantee_type;not null"`
	BizKindStrID            string    `gorm:"column:biz_kind_str_id;not null"`
	PriceTypeStrAr          string    `gorm:"column:price_type_str_ar;not null"`
	RegID                   string    `gorm:"column:reg_id;not null"`
	PublicDtStr             string    `gorm:"column:public_dt_str;not null"`
	ExecutionPlaceStrID     int       `gorm:"column:execution_place_str_id;not null"`
	EvalMethod              string    `gorm:"column:eval_method;not null"`
	HasAnnonce              string    `gorm:"column:has_annonce;not null"`
	RefNo                   string    `gorm:"column:ref_no;not null"`
	ProcedureType           string    `gorm:"column:procedure_type;not null"`
	PublicDt                string    `gorm:"column:public_dt;not null"`
	PlanNmEn                string    `gorm:"column:plan_nm_en;not null"`
	RegTypeStrEn            string    `gorm:"column:reg_type_str_en;not null"`
	ConsorYnStrID           int       `gorm:"column:consor_yn_str_id;not null"`
	BiddingDocPlace         string    `gorm:"column:bidding_doc_place;not null"`
	BdRecvAddrsAr           string    `gorm:"column:bd_recv_addrs_ar;not null"`
	ProcChmanIDStr          string    `gorm:"column:proc_chman_id_str;not null"`
	ConsorYnStrAr           string    `gorm:"column:consor_yn_str_ar;not null"`
	RealYnStrID             int       `gorm:"column:real_yn_str_id;not null"`
	PqRecvAddrsAr           string    `gorm:"column:pq_recv_addrs_ar;not null"`
	PlanNmFr                string    `gorm:"column:plan_nm_fr;not null"`
	RealYnStrAr             string    `gorm:"column:real_yn_str_ar;not null"`
	ExamChmanID             string    `gorm:"column:exam_chman_id;not null"`
	GuaranteeTypeStrEn      string    `gorm:"column:guarantee_type_str_en;not null"`
	ProcChmanID             string    `gorm:"column:proc_chman_id;not null"`
	PlanNo                  string    `gorm:"column:plan_no;not null"`
	TpRecvAddrsAr           string    `gorm:"column:tp_recv_addrs_ar;not null"`
	BidInstCd               string    `gorm:"column:bid_inst_cd;not null"`
	PbkStrID                int       `gorm:"column:pbk_str_id;not null"`
	OnOffTypeStrFr          string    `gorm:"column:on_off_type_str_fr;not null"`
	EvalMethodStrEn         string    `gorm:"column:eval_method_str_en;not null"`
}

//AOGetAttributionDuMarché
// Struct to represent the database model
type TunepsAOATTMList struct {
	ID                  uuid.UUID `gorm:"column:id; primaryKey; type:uuid; not null;"`
	WinnerPublicDt      string    `gorm:"column:winner_public_dt;not null"`
	NotWinnerYn         string    `gorm:"column:not_winner_yn;not null"`
	ProcedureType       string    `gorm:"column:procedure_type;not null"`
	RefNo               string    `gorm:"column:ref_no;not null"`
	BizRegNm            string    `gorm:"column:biz_reg_nm;not null"`
	BidNo               string    `gorm:"column:bid_no;not null"`
	TotalBidPriceCorr   string    `gorm:"column:total_bid_price_corr;not null"`
	BizKind             string    `gorm:"column:biz_kind;not null"`
	NotWinnerReason     string    `gorm:"column:not_winner_reason;not null"`
	BidPrice            string    `gorm:"column:bid_price;not null"`
	BidNmEn             string    `gorm:"column:bid_nm_en;not null"`
	RegDt               string    `gorm:"column:reg_dt;not null"`
	BidPriceSucc        string    `gorm:"column:bid_price_succ;not null"`
	BidNmFr             string    `gorm:"column:bid_nm_fr;not null"`
	NbrConsorExist      string    `gorm:"column:nbr_consor_exist;not null"`
	BidModSeq           string    `gorm:"column:bid_mod_seq;not null"`
	BidPriceSuccCurr    string    `gorm:"column:bid_price_succ_curr;not null"`
	BidNmAr             string    `gorm:"column:bid_nm_ar;not null"`
	BizReg              string    `gorm:"column:biz_reg;not null"`
	BidCls              string    `gorm:"column:bid_cls;not null"`
	ExecType            string    `gorm:"column:exec_type;not null"`
	TotalBidPriceDcExch string    `gorm:"column:total_bid_price_dc_exch;not null"`
	WinnerAfterDt       string    `gorm:"column:winner_after_dt;not null"`
}

// ************************bOUTON RESULTAT*************************************

// données de base
// Struct to represent the database model
type TunepsAORDNBList struct {
	ID                      uuid.UUID `gorm:"column:id; primaryKey; type:uuid; not null;"`
	BdRecvEndDt             string    `gorm:"column:bd_recv_end_dt;not null"`
	BiddingDocPriceCurr     string    `gorm:"column:bidding_doc_price_curr;not null"`
	EncCert                 string    `gorm:"column:enc_cert;not null"`
	GeneralYn               string    `gorm:"column:general_yn;not null"`
	BdRecvStartDt           string    `gorm:"column:bd_recv_start_dt;not null"`
	BidNo                   string    `gorm:"column:bid_no;not null"`
	ExecutionPlaceStrEn     string    `gorm:"column:execution_place_str_en;not null"`
	ProcedureTypeStrEn      string    `gorm:"column:procedure_type_str_en;not null"`
	PbkStrAr                string    `gorm:"column:pbk_str_ar;not null"`
	GuaranteeTypeStrFr      string    `gorm:"column:guarantee_type_str_fr;not null"`
	RealYnStrFr             string    `gorm:"column:real_yn_str_fr;not null"`
	FinancialMethodStrEn    string    `gorm:"column:financial_method_str_en;not null"`
	BiddingDocPrice         int       `gorm:"column:bidding_doc_price;not null"`
	RealYn                  string    `gorm:"column:real_yn;not null"`
	BdOpenDt                string    `gorm:"column:bd_open_dt;not null"`
	BidNmAr                 string    `gorm:"column:bid_nm_ar;not null"`
	OnOffTypeStrEn          string    `gorm:"column:on_off_type_str_en;not null"`
	PriceTypeStrFr          string    `gorm:"column:price_type_str_fr;not null"`
	BizKindStrEn            string    `gorm:"column:biz_kind_str_en;not null"`
	InternationalBidYnStrAr string    `gorm:"column:international_bid_yn_str_ar;not null"`
	OnOffType               string    `gorm:"column:on_off_type;not null"`
	InstRegNo               string    `gorm:"column:inst_reg_no;not null"`
	ConsorYn                string    `gorm:"column:consor_yn;not null"`
	RegTypeStrFr            string    `gorm:"column:reg_type_str_fr;not null"`
	ProcedureTypeStrFr      string    `gorm:"column:procedure_type_str_fr;not null"`
	PriceTypeStrEn          string    `gorm:"column:price_type_str_en;not null"`
	InternationalBidYnStrID int       `gorm:"column:international_bid_yn_str_id;not null"`
	FinancialMethodStrFr    string    `gorm:"column:financial_method_str_fr;not null"`
	BdDepartEn              string    `gorm:"column:bd_depart_en;not null"`
	BidExpiredDays          string    `gorm:"column:bid_expired_days;not null"`
	ConsorYnStrFr           string    `gorm:"column:consor_yn_str_fr;not null"`
	PublicYnStrAr           string    `gorm:"column:public_yn_str_ar;not null"`
	PlanNmAr                string    `gorm:"column:plan_nm_ar;not null"`
	PublicYn                string    `gorm:"column:public_yn;not null"`
	PublicYnStrEn           string    `gorm:"column:public_yn_str_en;not null"`
	BizKindStrFr            string    `gorm:"column:biz_kind_str_fr;not null"`
	EvalMethodStrAr         string    `gorm:"column:eval_method_str_ar;not null"`
	TpRecvAddrsFr           string    `gorm:"column:tp_recv_addrs_fr;not null"`
	ExecutionPlace          string    `gorm:"column:execution_place;not null"`
	BidInstNm               string    `gorm:"column:bid_inst_nm;not null"`
	BdRecvStartDtStr        string    `gorm:"column:bd_recv_start_dt_str;not null"`
	RegTypeStrAr            string    `gorm:"column:reg_type_str_ar;not null"`
	StaffNm                 string    `gorm:"column:staff_nm;not null"`
	BdRecvAddrsEn           string    `gorm:"column:bd_recv_addrs_en;not null"`
	FinancingYn             string    `gorm:"column:financing_yn;not null"`
	RegType                 string    `gorm:"column:reg_type;not null"`
	ConsorYnStrEn           string    `gorm:"column:consor_yn_str_en;not null"`
	InternationalBidYn      string    `gorm:"column:international_bid_yn;not null"`
	RealYnStrEn             string    `gorm:"column:real_yn_str_en;not null"`
	ExamChmanIDStr          string    `gorm:"column:exam_chman_id_str;not null"`
	PriceType               string    `gorm:"column:price_type;not null"`
	ExecutionPlaceStrFr     string    `gorm:"column:execution_place_str_fr;not null"`
	BdOpenDtStr             string    `gorm:"column:bd_open_dt_str;not null"`
	PlanModSeq              string    `gorm:"column:plan_mod_seq;not null"`
	FrameworkYn             string    `gorm:"column:framework_yn;not null"`
	BdDepartFr              string    `gorm:"column:bd_depart_fr;not null"`
	EvalMethodStrID         int       `gorm:"column:eval_method_str_id;not null"`
	BizKind                 string    `gorm:"column:biz_kind;not null"`
	PqRecvAddrsEn           string    `gorm:"column:pq_recv_addrs_en;not null"`
	TpRecvAddrsEn           string    `gorm:"column:tp_recv_addrs_en;not null"`
	BidModSeq               string    `gorm:"column:bid_mod_seq;not null"`
	GuaranteeTypeStrID      int       `gorm:"column:guarantee_type_str_id;not null"`
	BidNmFr                 string    `gorm:"column:bid_nm_fr;not null"`
	RegTypeStrID            int       `gorm:"column:reg_type_str_id;not null"`
	BdRecvAddrsFr           string    `gorm:"column:bd_recv_addrs_fr;not null"`
	ProcedureTypeStrID      int       `gorm:"column:procedure_type_str_id;not null"`
	FinancialMethodStrID    int       `gorm:"column:financial_method_str_id;not null"`
	InternationalBidYnStrFr string    `gorm:"column:international_bid_yn_str_fr;not null"`
	PbkStrEn                string    `gorm:"column:pbk_str_en;not null"`
	OnOffTypeStrID          int       `gorm:"column:on_off_type_str_id;not null"`
	EpBidMasterID           int       `gorm:"column:ep_bid_master_id;not null"`
	PqRecvAddrsFr           string    `gorm:"column:pq_recv_addrs_fr;not null"`
	BdDepartAr              string    `gorm:"column:bd_depart_ar;not null"`
	BidNmEn                 string    `gorm:"column:bid_nm_en;not null"`
	FinancialMethod         string    `gorm:"column:financial_method;not null"`
	OnOffTypeStrAr          string    `gorm:"column:on_off_type_str_ar;not null"`
	ExecutionPlaceStrAr     string    `gorm:"column:execution_place_str_ar;not null"`
	ProcedureTypeStrAr      string    `gorm:"column:procedure_type_str_ar;not null"`
	GuaranteeTypeStrAr      string    `gorm:"column:guarantee_type_str_ar;not null"`
	FinancialMethodStrAr    string    `gorm:"column:financial_method_str_ar;not null"`
	PbkStrFr                string    `gorm:"column:pbk_str_fr;not null"`
	EvalMethodStrFr         string    `gorm:"column:eval_method_str_fr;not null"`
	PublicYnStrFr           string    `gorm:"column:public_yn_str_fr;not null"`
	BizKindStrAr            string    `gorm:"column:biz_kind_str_ar;not null"`
	BidPoCd                 string    `gorm:"column:bid_po_cd;not null"`
	InternationalBidYnStrEn string    `gorm:"column:international_bid_yn_str_en;not null"`
	BdRecvEndDtStr          string    `gorm:"column:bd_recv_end_dt_str;not null"`
	PriceTypeStrID          int       `gorm:"column:price_type_str_id;not null"`
	GuaranteeType           string    `gorm:"column:guarantee_type;not null"`
	BizKindStrID            string    `gorm:"column:biz_kind_str_id;not null"`
	PriceTypeStrAr          string    `gorm:"column:price_type_str_ar;not null"`
	RegID                   string    `gorm:"column:reg_id;not null"`
	PublicDtStr             string    `gorm:"column:public_dt_str;not null"`
	ExecutionPlaceStrID     int       `gorm:"column:execution_place_str_id;not null"`
	EvalMethod              string    `gorm:"column:eval_method;not null"`
	HasAnnonce              string    `gorm:"column:has_annonce;not null"`
	RefNo                   string    `gorm:"column:ref_no;not null"`
	ProcedureType           string    `gorm:"column:procedure_type;not null"`
	PublicDt                string    `gorm:"column:public_dt;not null"`
	PlanNmEn                string    `gorm:"column:plan_nm_en;not null"`
	RegTypeStrEn            string    `gorm:"column:reg_type_str_en;not null"`
	ConsorYnStrID           int       `gorm:"column:consor_yn_str_id;not null"`
	BiddingDocPlace         string    `gorm:"column:bidding_doc_place;not null"`
	BdRecvAddrsAr           string    `gorm:"column:bd_recv_addrs_ar;not null"`
	ProcChmanIDStr          string    `gorm:"column:proc_chman_id_str;not null"`
	ConsorYnStrAr           string    `gorm:"column:consor_yn_str_ar;not null"`
	RealYnStrID             int       `gorm:"column:real_yn_str_id;not null"`
	PqRecvAddrsAr           string    `gorm:"column:pq_recv_addrs_ar;not null"`
	PlanNmFr                string    `gorm:"column:plan_nm_fr;not null"`
	RealYnStrAr             string    `gorm:"column:real_yn_str_ar;not null"`
	ExamChmanID             string    `gorm:"column:exam_chman_id;not null"`
	GuaranteeTypeStrEn      string    `gorm:"column:guarantee_type_str_en;not null"`
	ProcChmanID             string    `gorm:"column:proc_chman_id;not null"`
	PlanNo                  string    `gorm:"column:plan_no;not null"`
	TpRecvAddrsAr           string    `gorm:"column:tp_recv_addrs_ar;not null"`
	BidInstCd               string    `gorm:"column:bid_inst_cd;not null"`
	PbkStrID                int       `gorm:"column:pbk_str_id;not null"`
	OnOffTypeStrFr          string    `gorm:"column:on_off_type_str_fr;not null"`
	EvalMethodStrEn         string    `gorm:"column:eval_method_str_en;not null"`
}

// Lot-Resultat
// Struct to represent the database model
type TunepsAORLotList struct {
	ID            uuid.UUID `gorm:"column:id; primaryKey; type:uuid; not null;"`
	RefNo         string    `gorm:"column:ref_no;not null"`
	BizKindFrStr  string    `gorm:"column:biz_kind_fr_str;not null"`
	BizKindArStr  string    `gorm:"column:biz_kind_ar_str;not null"`
	BidNo         string    `gorm:"column:bid_no;not null"`
	Cnt           string    `gorm:"column:cnt;not null"`
	BizKindEnStr  string    `gorm:"column:biz_kind_en_str;not null"`
	BizKind       string    `gorm:"column:biz_kind;not null"`
	BidNmEn       string    `gorm:"column:bid_nm_en;not null"`
	RegDt         string    `gorm:"column:reg_dt;not null"`
	BidModSeq     string    `gorm:"column:bid_mod_seq;not null"`
	BidNmAr       string    `gorm:"column:bid_nm_ar;not null"`
	Rk            string    `gorm:"column:rk;not null"`
	BidFailReason string    `gorm:"column:bid_fail_reason;not null"`
	RegID         string    `gorm:"column:reg_id;not null"`
	BidCls        string    `gorm:"column:bid_cls;not null"`
	ExecType      string    `gorm:"column:exec_type;not null"`
	BidNmFr       string    `gorm:"column:bid_nm_fr;not null"`
}

//Soumissionaires
// Struct to represent the database model
type TunepsAORSList struct {
	ID                  uuid.UUID `gorm:"column:id; primaryKey; type:uuid; not null;"`
	CdNmStrEn           string    `gorm:"column:cd_nm_str_en;not null"`
	IneligibleCd        string    `gorm:"column:ineligible_cd;not null"`
	TechEvalResult      string    `gorm:"column:tech_eval_result;not null"`
	BidNo               string    `gorm:"column:bid_no;not null"`
	CdPfTechFr          string    `gorm:"column:cd_pf_tech_fr;not null"`
	BdRecvDtStr         string    `gorm:"column:bd_recv_dt_str;not null"`
	EpBidOpenResultID   int       `gorm:"column:ep_bid_open_result_id;not null"`
	CdPfFinaFr          string    `gorm:"column:cd_pf_fina_fr;not null"`
	Rank                int       `gorm:"column:rank;not null"`
	OnlineRegYn         string    `gorm:"column:online_reg_yn;not null"`
	DecAttachFileCnt    string    `gorm:"column:dec_attach_file_cnt;not null"`
	BidCls              string    `gorm:"column:bid_cls;not null"`
	TotalBidPriceDcExch int       `gorm:"column:total_bid_price_dc_exch;not null"`
	BizRegNo            string    `gorm:"column:biz_reg_no;not null"`
	BizRegNm            string    `gorm:"column:biz_reg_nm;not null"`
	CdPfTechEn          string    `gorm:"column:cd_pf_tech_en;not null"`
	FinaResultReason    string    `gorm:"column:fina_result_reason;not null"`
	CdNmStrFr           string    `gorm:"column:cd_nm_str_fr;not null"`
	Cnt                 int       `gorm:"column:cnt;not null"`
	CdPfTechAr          string    `gorm:"column:cd_pf_tech_ar;not null"`
	IneligibleReason    string    `gorm:"column:ineligible_reason;not null"`
	CdPfFinaAr          string    `gorm:"column:cd_pf_fina_ar;not null"`
	BdRecvDt            string    `gorm:"column:bd_recv_dt;not null"`
	BidModSeq           string    `gorm:"column:bid_mod_seq;not null"`
	CdNmStrAr           string    `gorm:"column:cd_nm_str_ar;not null"`
	Rk                  int       `gorm:"column:rk;not null"`
	BizReg              string    `gorm:"column:biz_reg;not null"`
	FinaEvalResult      string    `gorm:"column:fina_eval_result;not null"`
	TechResultReason    string    `gorm:"column:tech_result_reason;not null"`
	CdPfFinaEn          string    `gorm:"column:cd_pf_fina_en;not null"`
}

/// ************************ Shopping Mall *******************************************

// Struct to represent the database model
type TunepsSMList struct {
	ID             uuid.UUID `gorm:"column:id; primaryKey; type:uuid; not null;"`
	PublicDt       string    `gorm:"column:public_dt;not null"`
	SpShopMasterID int       `gorm:"column:sp_shop_master_id;not null"`
	ShopModSeq     string    `gorm:"column:shop_mod_seq;not null"`
	InstNm         string    `gorm:"column:inst_nm;not null"`
	ShopNmEn       string    `gorm:"column:shop_nm_en;not null"`
	ShopNmFr       string    `gorm:"column:shop_nm_fr;not null"`
	SpRecvEndDt    string    `gorm:"column:sp_recv_end_dt;not null"`
	ShopNo         string    `gorm:"column:shop_no;not null"`
	ShopNmAr       string    `gorm:"column:shop_nm_ar;not null"`
}

// Struct to represent the database model
type TunepsSMInfoList struct {
	ID                  uuid.UUID `gorm:"column:id; primaryKey; type:uuid; not null;"`
	RefNo               string    `gorm:"column:ref_no;not null"`
	PublicDt            string    `gorm:"column:public_dt;not null"`
	ShopModSeq          string    `gorm:"column:shop_mod_seq;not null"`
	PublicYn            string    `gorm:"column:public_yn;not null"`
	MsAr                string    `gorm:"column:ms_ar;not null"`
	PbkStrEn            string    `gorm:"column:pbk_str_en;not null"`
	RegNm               string    `gorm:"column:reg_nm;not null"`
	BizKindStrFr        string    `gorm:"column:biz_kind_str_fr;not null"`
	CdNmAr              string    `gorm:"column:cd_nm_ar;not null"`
	PbkStrAr            string    `gorm:"column:pbk_str_ar;not null"`
	ShopNmAr            string    `gorm:"column:shop_nm_ar;not null"`
	BcNm                string    `gorm:"column:bc_nm;not null"`
	BcID                string    `gorm:"column:bc_id;not null"`
	EpAr                string    `gorm:"column:ep_ar;not null"`
	InstNm              string    `gorm:"column:inst_nm;not null"`
	InternationalShopYn string    `gorm:"column:international_shop_yn;not null"`
	SpRecvEndDtStr      string    `gorm:"column:sp_recv_end_dt_str;not null"`
	ShopExpiredDays     int       `gorm:"column:shop_expired_days;not null"`
	SpRecvStartDt       string    `gorm:"column:sp_recv_start_dt;not null"`
	EpEn                string    `gorm:"column:ep_en;not null"`
	ExamNm              string    `gorm:"column:exam_nm;not null"`
	MsEn                string    `gorm:"column:ms_en;not null"`
	CdNmEn              string    `gorm:"column:cd_nm_en;not null"`
	ShopInstCd          string    `gorm:"column:shop_inst_cd;not null"`
	PbkStrFr            string    `gorm:"column:pbk_str_fr;not null"`
	MsFr                string    `gorm:"column:ms_fr;not null"`
	PbkCd               int       `gorm:"column:pbk_cd;not null"`
	BizKindStrEn        string    `gorm:"column:biz_kind_str_en;not null"`
	ShopNmFr            string    `gorm:"column:shop_nm_fr;not null"`
	MasterCdBc          string    `gorm:"column:master_cd_bc;not null"`
	BizKindStrAr        string    `gorm:"column:biz_kind_str_ar;not null"`
	SpRecvEndDt         string    `gorm:"column:sp_recv_end_dt;not null"`
	SpShopMasterID      int       `gorm:"column:sp_shop_master_id;not null"`
	BizKindCd           int       `gorm:"column:biz_kind_cd;not null"`
	BizKindStrID        int       `gorm:"column:biz_kind_str_id;not null"`
	ShopNmEn            string    `gorm:"column:shop_nm_en;not null"`
	PbkStrID            int       `gorm:"column:pbk_str_id;not null"`
	ExamID              string    `gorm:"column:exam_id;not null"`
	EpFr                string    `gorm:"column:ep_fr;not null"`
	CdNmFr              string    `gorm:"column:cd_nm_fr;not null"`
	RegID               string    `gorm:"column:reg_id;not null"`
	PublicDtStr         string    `gorm:"column:public_dt_str;not null"`
	MasterCdExam        string    `gorm:"column:master_cd_exam;not null"`
	ShopNo              string    `gorm:"column:shop_no;not null"`
	SpOpenDt            string    `gorm:"column:sp_open_dt;not null"`
	HasAnnonce          string    `gorm:"column:has_annonce;not null"`
}

// Struct to represent the database model
type TunepsSMLotList struct {
	ID                  uuid.UUID `gorm:"column:id; primaryKey; type:uuid; not null;"`
	ShopCls             string    `gorm:"column:shop_cls;not null"`
	BudgetAmt           string    `gorm:"column:budget_amt;not null"`
	ShopClsDtlDesc      string    `gorm:"column:shop_cls_dtl_desc;not null"`
	ShopModSeq          string    `gorm:"column:shop_mod_seq;not null"`
	GuaranteeAmountCurr string    `gorm:"column:guarantee_amount_curr;not null"`
	GuaranteeAmount     string    `gorm:"column:guarantee_amount;not null"`
	ShopNo              string    `gorm:"column:shop_no;not null"`
	ShopClsNm           string    `gorm:"column:shop_cls_nm;not null"`
	BudgetAmtCurr       string    `gorm:"column:budget_amt_curr;not null"`
	SpShopClsID         int       `gorm:"column:sp_shop_cls_id;not null"`
}

// Struct to represent the database model
type TunepsSMDocList struct {
	ID                 uuid.UUID `gorm:"column:id; primaryKey; type:uuid; not null;"`
	FileNm             string    `gorm:"column:file_nm;not null"`
	ShopModSeq         string    `gorm:"column:shop_mod_seq;not null"`
	SeqNo              int       `gorm:"column:seq_no;not null"`
	NodeRef            string    `gorm:"column:node_ref;not null"`
	CdNmAr             string    `gorm:"column:cd_nm_ar;not null"`
	CdNmFr             string    `gorm:"column:cd_nm_fr;not null"`
	DocCd              string    `gorm:"column:doc_cd;not null"`
	FileLoc            string    `gorm:"column:file_loc;not null"`
	ShopNo             string    `gorm:"column:shop_no;not null"`
	SpShopAttachFileID int       `gorm:"column:sp_shop_attach_file_id;not null"`
	CdNmEn             string    `gorm:"column:cd_nm_en;not null"`
}
