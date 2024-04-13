package tuneps

import (
	"fmt"
	"labs/domains"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ************************ Appel d'offres *******************************************

// FUNCGetAOList récupère des données à partir de REQAOGetList, les traite, et les sauvegarde dans la base de données.
func FUNCGetAOList(db *gorm.DB) error {

	// Obtient la date et l'heure actuelles.
	today := time.Now()

	// Récupère les données de la liste AO depuis la source externe en utilisant la fonction REQAOGetList.
	reqAOList, err := REQAOGetList()
	if err != nil {
		return err
	}
	fmt.Println("Données de REQAOGetList :", reqAOList.Payload.Data)

	for _, reqAO := range reqAOList.Payload.Data {
		closed := false

		fmt.Println("Data from REQAOGetList:", reqAO)

		// Analyse la date de fin du AO et vérifie si elle est antérieure à la date actuelle.
		endDate, err := time.Parse("2006-01-02 15:04:05.9", reqAO.BdRecvEndDt)
		if err != nil || endDate.Before(today) {
			closed = true
		}

		// Crée un objet TunepsAOList avec les détails du AO et l'état fermé.
		offre := domains.TunepsAOList{
			ID:            uuid.New(),
			BdRecvEndDt:   reqAO.BdRecvEndDt,
			BidInstNm:     reqAO.BidInstNm,
			BidNmEn:       reqAO.BidNmEn,
			PublicDt:      reqAO.PublicDt,
			PublicYn:      reqAO.PublicYn,
			BidModSeq:     reqAO.BidModSeq,
			BidNo:         reqAO.BidNo,
			BidNmAr:       reqAO.BidNmAr,
			BidNmFr:       reqAO.BidNmFr,
			EpBidMasterID: reqAO.EpBidMasterID,
			Closed:        closed,
		}
		fmt.Println("Data to be inserted:", offre)

		// Insère les données dans la base de données en utilisant la fonction Create de domaines.
		if err = domains.Create(db, &offre); err != nil {
			fmt.Println("Error inserting data into the database:", err)
			return err
		}
	}
	fmt.Println("Data AO saved to the  database successfully!")
	return nil

}

// FUNCGetAOInfoList récupère des données à partir de REQAOGetInfoList, les traite, et les sauvegarde dans la base de données.
func FUNCGetAOInfoList(db *gorm.DB) error {

	// Récupère les données de la liste AO depuis la source externe en utilisant la fonction REQAOGetInfoList.
	reqAOInfoList, err := REQAOGetInfoList()
	if err != nil {
		return err
	}
	fmt.Println("Data from REQAOGetInfoList:", reqAOInfoList.Payload)
	// Crée un objet TunepsAOList avec les détails du AOInfo
	infoOffre := domains.TunepsAOInfoList{
		ID:                      uuid.New(),
		BdRecvEndDt:             reqAOInfoList.Payload.BdRecvEndDt,
		BiddingDocPriceCurr:     reqAOInfoList.Payload.BiddingDocPriceCurr,
		EncCert:                 reqAOInfoList.Payload.EncCert,
		GeneralYn:               reqAOInfoList.Payload.GeneralYn,
		BdRecvStartDt:           reqAOInfoList.Payload.BdRecvStartDt,
		BidNo:                   reqAOInfoList.Payload.BidNo,
		ExecutionPlaceStrEn:     reqAOInfoList.Payload.ProcedureTypeStrEn,
		ProcedureTypeStrEn:      reqAOInfoList.Payload.ProcedureTypeStrEn,
		PbkStrAr:                reqAOInfoList.Payload.PbkStrAr,
		GuaranteeTypeStrFr:      reqAOInfoList.Payload.GuaranteeTypeStrFr,
		RealYnStrFr:             reqAOInfoList.Payload.RealYnStrFr,
		FinancialMethodStrEn:    reqAOInfoList.Payload.FinancialMethodStrEn,
		BiddingDocPrice:         reqAOInfoList.Payload.BiddingDocPrice,
		RealYn:                  reqAOInfoList.Payload.RealYn,
		BdOpenDt:                reqAOInfoList.Payload.BdOpenDt,
		BidNmAr:                 reqAOInfoList.Payload.BidNmAr,
		OnOffTypeStrEn:          reqAOInfoList.Payload.OnOffTypeStrEn,
		PriceTypeStrFr:          reqAOInfoList.Payload.PriceTypeStrFr,
		BizKindStrEn:            reqAOInfoList.Payload.BizKindStrEn,
		InternationalBidYnStrAr: reqAOInfoList.Payload.InternationalBidYnStrAr,
		OnOffType:               reqAOInfoList.Payload.OnOffType,
		InstRegNo:               reqAOInfoList.Payload.InstRegNo,
		ConsorYn:                reqAOInfoList.Payload.ConsorYn,
		RegTypeStrFr:            reqAOInfoList.Payload.RegTypeStrFr,
		ProcedureTypeStrFr:      reqAOInfoList.Payload.ProcedureTypeStrFr,
		PriceTypeStrEn:          reqAOInfoList.Payload.PriceTypeStrEn,
		InternationalBidYnStrID: reqAOInfoList.Payload.InternationalBidYnStrID,
		FinancialMethodStrFr:    reqAOInfoList.Payload.FinancialMethodStrFr,
		BdDepartEn:              reqAOInfoList.Payload.BdDepartEn,
		BidExpiredDays:          reqAOInfoList.Payload.BidExpiredDays,
		ConsorYnStrFr:           reqAOInfoList.Payload.ConsorYnStrFr,
		PublicYnStrAr:           reqAOInfoList.Payload.PublicYnStrAr,
		PublicYn:                reqAOInfoList.Payload.PublicYn,
		PublicYnStrEn:           reqAOInfoList.Payload.PublicYnStrEn,
		BizKindStrFr:            reqAOInfoList.Payload.BizKindStrFr,
		EvalMethodStrAr:         reqAOInfoList.Payload.EvalMethodStrAr,
		ExecutionPlace:          reqAOInfoList.Payload.ExecutionPlace,
		BidInstNm:               reqAOInfoList.Payload.BidInstNm,
		BdRecvStartDtStr:        reqAOInfoList.Payload.BdRecvStartDtStr,
		RegTypeStrAr:            reqAOInfoList.Payload.RegTypeStrAr,
		StaffNm:                 reqAOInfoList.Payload.StaffNm,
		BdRecvAddrsEn:           reqAOInfoList.Payload.BdRecvAddrsEn,
		FinancingYn:             reqAOInfoList.Payload.FinancingYn,
		RegType:                 reqAOInfoList.Payload.RegType,
		ConsorYnStrEn:           reqAOInfoList.Payload.ConsorYnStrEn,
		InternationalBidYn:      reqAOInfoList.Payload.InternationalBidYn,
		RealYnStrEn:             reqAOInfoList.Payload.RealYnStrEn,
		ExamChmanIDStr:          reqAOInfoList.Payload.ExamChmanIDStr,
		PriceType:               reqAOInfoList.Payload.PriceType,
		ExecutionPlaceStrFr:     reqAOInfoList.Payload.ExecutionPlaceStrFr,
		BdOpenDtStr:             reqAOInfoList.Payload.BdOpenDtStr,
		FrameworkYn:             reqAOInfoList.Payload.FrameworkYn,
		BdDepartFr:              reqAOInfoList.Payload.BdDepartFr,
		EvalMethodStrID:         reqAOInfoList.Payload.EvalMethodStrID,
		BizKind:                 reqAOInfoList.Payload.BizKind,
		BidModSeq:               reqAOInfoList.Payload.BidModSeq,
		GuaranteeTypeStrID:      reqAOInfoList.Payload.GuaranteeTypeStrID,
		BidNmFr:                 reqAOInfoList.Payload.BidNmFr,
		RegTypeStrID:            reqAOInfoList.Payload.RegTypeStrID,
		BdRecvAddrsFr:           reqAOInfoList.Payload.BdRecvAddrsFr,
		ProcedureTypeStrID:      reqAOInfoList.Payload.ProcedureTypeStrID,
		FinancialMethodStrID:    reqAOInfoList.Payload.FinancialMethodStrID,
		InternationalBidYnStrFr: reqAOInfoList.Payload.InternationalBidYnStrFr,
		PbkStrEn:                reqAOInfoList.Payload.PbkStrEn,
		OnOffTypeStrID:          reqAOInfoList.Payload.OnOffTypeStrID,
		EpBidMasterID:           reqAOInfoList.Payload.EpBidMasterID,
		BdDepartAr:              reqAOInfoList.Payload.BdDepartAr,
		BidNmEn:                 reqAOInfoList.Payload.BidNmEn,
		FinancialMethod:         reqAOInfoList.Payload.FinancialMethod,
		OnOffTypeStrAr:          reqAOInfoList.Payload.OnOffTypeStrAr,
		ExecutionPlaceStrAr:     reqAOInfoList.Payload.ExecutionPlaceStrAr,
		ProcedureTypeStrAr:      reqAOInfoList.Payload.ProcedureTypeStrAr,
		GuaranteeTypeStrAr:      reqAOInfoList.Payload.GuaranteeTypeStrAr,
		FinancialMethodStrAr:    reqAOInfoList.Payload.FinancialMethodStrAr,
		PbkStrFr:                reqAOInfoList.Payload.PbkStrFr,
		EvalMethodStrFr:         reqAOInfoList.Payload.EvalMethodStrFr,
		PublicYnStrFr:           reqAOInfoList.Payload.PublicYnStrFr,
		BizKindStrAr:            reqAOInfoList.Payload.BizKindStrAr,
		BidPoCd:                 reqAOInfoList.Payload.BidPoCd,
		InternationalBidYnStrEn: reqAOInfoList.Payload.InternationalBidYnStrEn,
		BdRecvEndDtStr:          reqAOInfoList.Payload.BdRecvEndDtStr,
		PriceTypeStrID:          reqAOInfoList.Payload.PriceTypeStrID,
		GuaranteeType:           reqAOInfoList.Payload.GuaranteeType,
		BizKindStrID:            reqAOInfoList.Payload.BizKindStrID,
		PriceTypeStrAr:          reqAOInfoList.Payload.PriceTypeStrAr,
		RegID:                   reqAOInfoList.Payload.RegID,
		PublicDtStr:             reqAOInfoList.Payload.PublicDtStr,
		ExecutionPlaceStrID:     reqAOInfoList.Payload.ExecutionPlaceStrID,
		EvalMethod:              reqAOInfoList.Payload.EvalMethod,
		HasAnnonce:              reqAOInfoList.Payload.HasAnnonce,
		RefNo:                   reqAOInfoList.Payload.RefNo,
		ProcedureType:           reqAOInfoList.Payload.ProcedureType,
		PublicDt:                reqAOInfoList.Payload.PublicDt,
		RegTypeStrEn:            reqAOInfoList.Payload.RegTypeStrEn,
		ConsorYnStrID:           reqAOInfoList.Payload.ConsorYnStrID,
		BdRecvAddrsAr:           reqAOInfoList.Payload.BdRecvAddrsAr,
		ProcChmanIDStr:          reqAOInfoList.Payload.ProcChmanIDStr,
		ConsorYnStrAr:           reqAOInfoList.Payload.ConsorYnStrAr,
		RealYnStrID:             reqAOInfoList.Payload.RealYnStrID,
		RealYnStrAr:             reqAOInfoList.Payload.RealYnStrAr,
		ExamChmanID:             reqAOInfoList.Payload.ExamChmanID,
		GuaranteeTypeStrEn:      reqAOInfoList.Payload.GuaranteeTypeStrEn,
		ProcChmanID:             reqAOInfoList.Payload.ProcChmanID,
		BidInstCd:               reqAOInfoList.Payload.BidInstCd,
		PbkStrID:                reqAOInfoList.Payload.PbkStrID,
		OnOffTypeStrFr:          reqAOInfoList.Payload.OnOffTypeStrFr,
		EvalMethodStrEn:         reqAOInfoList.Payload.EvalMethodStrEn,
	}
	fmt.Println("Data to be inserted:", infoOffre)

	// Insère les données dans la base de données en utilisant la fonction Create de domaines.
	if err = domains.Create(db, &infoOffre); err != nil {
		fmt.Println("Error inserting data into the database:", err)
		return err
	}

	fmt.Println("Data AO Info saved to the  database successfully!")
	return nil

}

// FUNCGetAOLotList récupère des données à partir de REQAOGetLotList, les traite, et les sauvegarde dans la base de données.
func FUNCGetAOLotList(db *gorm.DB) error {

	// Récupère les données de la liste AO depuis la source externe en utilisant la fonction REQAOGetLotList.
	reqAOLotList, err := REQAOGetLotList()
	if err != nil {
		return err
	}
	fmt.Println("Data from REQAOGetLotList:", reqAOLotList.Payload)

	for _, reqAOLot := range reqAOLotList.Payload {
		offreLot := domains.TunepsAOLotList{
			ID:                   uuid.New(),
			EpBidTypCau:          reqAOLot.EpBidTypCau,
			BidClsDtlDesc:        reqAOLot.BidClsDtlDesc,
			BidClsNm:             reqAOLot.BidClsNm,
			CdNmAr:               reqAOLot.CdNmAr,
			BidNo:                reqAOLot.BidNo,
			GuaranteeAmount:      reqAOLot.GuaranteeAmount,
			BudgetAmtCurrF:       reqAOLot.BudgetAmtCurrF,
			CdNm:                 reqAOLot.CdNm,
			BudgetAmtCurr:        reqAOLot.BudgetAmtCurr,
			BudgetAmt:            reqAOLot.BudgetAmt,
			GuaranteeAmountCurr:  reqAOLot.GuaranteeAmountCurr,
			BidModSeq:            reqAOLot.BidModSeq,
			CdNmFr:               reqAOLot.CdNmFr,
			GuaranteeAmountCurrF: reqAOLot.GuaranteeAmountCurrF,
			Affectationcheck:     reqAOLot.Affectationcheck,
			BidCls:               reqAOLot.BidCls,
			EpBidClsID:           reqAOLot.EpBidClsID,
			EpBidTypEval:         reqAOLot.EpBidTypEval,
		}

		// Insère les données dans la base de données en utilisant la fonction Create de domaines.
		if err = domains.Create(db, &offreLot); err != nil {
			fmt.Println("Error inserting data into the database:", err)
			return err
		}
	}
	fmt.Println("Data lot  saved to the  database successfully!")
	return nil

}

// FUNCGetAOCCList récupère des données à partir de REQAOGetInfoList, les traite, et les sauvegarde dans la base de données.
func FUNCGetAOCCList(db *gorm.DB) error {
	// Récupère les données de la liste AO depuis la source externe en utilisant la fonction REQAOGetLotList.
	reqAOCCList, err := REQAOGetCCList()
	if err != nil {
		return err
	}
	fmt.Println("Data from REQAOGetCCList:", reqAOCCList.Payload)

	for _, reqAOCC := range reqAOCCList.Payload {

		offreCC := domains.TunepsAOCCList{
			ID:                uuid.New(),
			FileNm:            reqAOCC.FileNm,
			SeqNo:             reqAOCC.SeqNo,
			BidAttNodeRef:     reqAOCC.BidAttNodeRef,
			EpBidAttachFileID: reqAOCC.EpBidAttachFileID,
			BidModSeq:         reqAOCC.BidModSeq,
			CdNmAr:            reqAOCC.CdNmAr,
			BidNo:             reqAOCC.BidNo,
			CdNmFr:            reqAOCC.CdNmFr,
			DocCd:             reqAOCC.DocCd,
			FileLoc:           reqAOCC.FileLoc,
			CdNm:              reqAOCC.CdNm,
		}
		// Insère les données dans la base de données en utilisant la fonction Create de domaines.
		if err = domains.Create(db, &offreCC); err != nil {
			fmt.Println("Error inserting data into the database:", err)
			return err
		}
	}
	fmt.Println("Data cahier de charge  saved to the  database successfully!")
	return nil

}

// FUNCGetAOAgList récupère des données à partir de REQAOGetAgList, les traite, et les sauvegarde dans la base de données.
func FUNCGetAOAgList(db *gorm.DB) error {
	// Récupère les données de la liste AO depuis la source externe en utilisant la fonction REQAOGetAgList.
	reqAOAgList, err := REQAOGetAgList()
	if err != nil {
		return err
	}
	fmt.Println("Data from REQAOGetAgList:", reqAOAgList.Payload)

	for _, reqAOAg := range reqAOAgList.Payload {
		offreAg := domains.TunepsAOAgList{
			ID:             uuid.New(),
			LicenseCd:      reqAOAg.LicenseCd,
			BidLicenseCd:   reqAOAg.BidLicenseCd,
			BidModSeq:      reqAOAg.BidModSeq,
			BidNo:          reqAOAg.BidNo,
			EpBidLicenseID: reqAOAg.EpBidLicenseID,
			BidCls:         reqAOAg.BidCls,
			Seq:            reqAOAg.Seq,
		}
		// Insère les données dans la base de données en utilisant la fonction Create de domaines.
		if err = domains.Create(db, &offreAg); err != nil {
			fmt.Println("Error inserting data into the database:", err)
			return err
		}
	}
	fmt.Println("Data agrément  saved to the  database successfully!")
	return nil

}

// FUNCGetAOProdList récupère des données à partir de REQAOGetProdList, les traite, et les sauvegarde dans la base de données.
func FUNCGetAOProdList(db *gorm.DB) error {

	// Récupère les données de la liste AO depuis la source externe en utilisant la fonction REQAOGetAgList.
	reqAOPRList, err := REQAOGetProdList()
	if err != nil {
		return err
	}
	fmt.Println("Data from REQAOGetProdList:", reqAOPRList.Payload)

	for _, reqAOPR := range reqAOPRList.Payload {
		offrePR := domains.TunepsAOPRList{
			ID:               uuid.New(),
			DeliveryEndDtStr: reqAOPR.DeliveryEndDtStr,
			CateNm:           reqAOPR.CateNm,
			Unit:             reqAOPR.Unit,
			EpBidClsDtlID:    reqAOPR.EpBidClsDtlID,
			SeqNo:            reqAOPR.SeqNo,
			DeliveryDays:     reqAOPR.DeliveryDays,
			CateID:           reqAOPR.CateID,
			BidModSeq:        reqAOPR.BidModSeq,
			Qty:              reqAOPR.Qty,
			BidNo:            reqAOPR.BidNo,
			BidCls:           reqAOPR.BidCls,
		}
		// Insère les données dans la base de données en utilisant la fonction Create de domaines.
		if err = domains.Create(db, &offrePR); err != nil {
			fmt.Println("Error inserting data into the database:", err)
			return err
		}
	}
	fmt.Println("Data Produit  saved to the  database successfully!")
	return nil

}

// ************* BOUTON List soumissionaire(s) retenu(s)******************
// FUNCGetAODNBList récupère des données à partir de REQAOGetList, les traite, et les sauvegarde dans la base de données.
func FUNCGetAODNBList(db *gorm.DB) error {
	// Récupère les données de la liste AO depuis la source externe en utilisant la fonction REQAOGetDNBList.
	reqAODNBList, err := REQAOGetDNBList()
	if err != nil {
		return err
	}
	fmt.Println("Data from REQAOGetDNBList:", reqAODNBList.Payload)

	offreDNB := domains.TunepsAODNBList{
		ID:                      uuid.New(),
		BdRecvEndDt:             reqAODNBList.Payload.BdRecvEndDt,
		BiddingDocPriceCurr:     reqAODNBList.Payload.BiddingDocPriceCurr,
		EncCert:                 reqAODNBList.Payload.EncCert,
		GeneralYn:               reqAODNBList.Payload.GeneralYn,
		BdRecvStartDt:           reqAODNBList.Payload.BdRecvStartDt,
		BidNo:                   reqAODNBList.Payload.BidNo,
		ExecutionPlaceStrEn:     reqAODNBList.Payload.ExecutionPlaceStrEn,
		ProcedureTypeStrEn:      reqAODNBList.Payload.ProcedureTypeStrEn,
		PbkStrAr:                reqAODNBList.Payload.PbkStrAr,
		GuaranteeTypeStrFr:      reqAODNBList.Payload.GuaranteeTypeStrFr,
		RealYnStrFr:             reqAODNBList.Payload.RealYnStrFr,
		FinancialMethodStrEn:    reqAODNBList.Payload.FinancialMethodStrEn,
		BiddingDocPrice:         reqAODNBList.Payload.BiddingDocPrice,
		RealYn:                  reqAODNBList.Payload.RealYn,
		BdOpenDt:                reqAODNBList.Payload.BdOpenDt,
		BidNmAr:                 reqAODNBList.Payload.BidNmAr,
		OnOffTypeStrEn:          reqAODNBList.Payload.OnOffTypeStrEn,
		PriceTypeStrFr:          reqAODNBList.Payload.PriceTypeStrFr,
		BizKindStrEn:            reqAODNBList.Payload.BizKindStrEn,
		InternationalBidYnStrAr: reqAODNBList.Payload.InternationalBidYnStrAr,
		OnOffType:               reqAODNBList.Payload.OnOffType,
		InstRegNo:               reqAODNBList.Payload.InstRegNo,
		ConsorYn:                reqAODNBList.Payload.ConsorYn,
		RegTypeStrFr:            reqAODNBList.Payload.RegTypeStrFr,
		ProcedureTypeStrFr:      reqAODNBList.Payload.ProcedureTypeStrFr,
		PriceTypeStrEn:          reqAODNBList.Payload.PriceTypeStrEn,
		InternationalBidYnStrID: reqAODNBList.Payload.InternationalBidYnStrID,
		FinancialMethodStrFr:    reqAODNBList.Payload.FinancialMethodStrFr,
		BdDepartEn:              reqAODNBList.Payload.BdDepartEn,
		BidExpiredDays:          reqAODNBList.Payload.BidExpiredDays,
		ConsorYnStrFr:           reqAODNBList.Payload.ConsorYnStrFr,
		PublicYnStrAr:           reqAODNBList.Payload.PublicYnStrAr,
		PlanNmAr:                reqAODNBList.Payload.PlanNmAr,
		PublicYn:                reqAODNBList.Payload.PublicYn,
		PublicYnStrEn:           reqAODNBList.Payload.PublicYnStrEn,
		BizKindStrFr:            reqAODNBList.Payload.BizKindStrFr,
		EvalMethodStrAr:         reqAODNBList.Payload.EvalMethodStrAr,
		TpRecvAddrsFr:           reqAODNBList.Payload.TpRecvAddrsFr,
		ExecutionPlace:          reqAODNBList.Payload.ExecutionPlace,
		BidInstNm:               reqAODNBList.Payload.BidInstNm,
		BdRecvStartDtStr:        reqAODNBList.Payload.BdRecvStartDtStr,
		RegTypeStrAr:            reqAODNBList.Payload.RegTypeStrAr,
		StaffNm:                 reqAODNBList.Payload.StaffNm,
		BdRecvAddrsEn:           reqAODNBList.Payload.BdRecvAddrsEn,
		FinancingYn:             reqAODNBList.Payload.FinancingYn,
		RegType:                 reqAODNBList.Payload.RegType,
		ConsorYnStrEn:           reqAODNBList.Payload.ConsorYnStrEn,
		InternationalBidYn:      reqAODNBList.Payload.InternationalBidYn,
		RealYnStrEn:             reqAODNBList.Payload.RealYnStrEn,
		ExamChmanIDStr:          reqAODNBList.Payload.ExamChmanIDStr,
		PriceType:               reqAODNBList.Payload.PriceType,
		ExecutionPlaceStrFr:     reqAODNBList.Payload.ExecutionPlaceStrFr,
		BdOpenDtStr:             reqAODNBList.Payload.BdOpenDtStr,
		PlanModSeq:              reqAODNBList.Payload.PlanModSeq,
		FrameworkYn:             reqAODNBList.Payload.FrameworkYn,
		BdDepartFr:              reqAODNBList.Payload.BdDepartFr,
		EvalMethodStrID:         reqAODNBList.Payload.EvalMethodStrID,
		BizKind:                 reqAODNBList.Payload.BizKind,
		PqRecvAddrsEn:           reqAODNBList.Payload.PqRecvAddrsEn,
		TpRecvAddrsEn:           reqAODNBList.Payload.TpRecvAddrsEn,
		BidModSeq:               reqAODNBList.Payload.BidModSeq,
		GuaranteeTypeStrID:      reqAODNBList.Payload.GuaranteeTypeStrID,
		BidNmFr:                 reqAODNBList.Payload.BidNmFr,
		RegTypeStrID:            reqAODNBList.Payload.RegTypeStrID,
		BdRecvAddrsFr:           reqAODNBList.Payload.BdRecvAddrsFr,
		ProcedureTypeStrID:      reqAODNBList.Payload.ProcedureTypeStrID,
		FinancialMethodStrID:    reqAODNBList.Payload.FinancialMethodStrID,
		InternationalBidYnStrFr: reqAODNBList.Payload.InternationalBidYnStrFr,
		PbkStrEn:                reqAODNBList.Payload.PbkStrEn,
		OnOffTypeStrID:          reqAODNBList.Payload.OnOffTypeStrID,
		EpBidMasterID:           reqAODNBList.Payload.EpBidMasterID,
		PqRecvAddrsFr:           reqAODNBList.Payload.PqRecvAddrsFr,
		BdDepartAr:              reqAODNBList.Payload.BdDepartAr,
		BidNmEn:                 reqAODNBList.Payload.BidNmEn,
		FinancialMethod:         reqAODNBList.Payload.FinancialMethod,
		OnOffTypeStrAr:          reqAODNBList.Payload.OnOffTypeStrAr,
		ExecutionPlaceStrAr:     reqAODNBList.Payload.ExecutionPlaceStrAr,
		ProcedureTypeStrAr:      reqAODNBList.Payload.ProcedureTypeStrAr,
		GuaranteeTypeStrAr:      reqAODNBList.Payload.GuaranteeTypeStrAr,
		FinancialMethodStrAr:    reqAODNBList.Payload.FinancialMethodStrAr,
		PbkStrFr:                reqAODNBList.Payload.PbkStrFr,
		EvalMethodStrFr:         reqAODNBList.Payload.EvalMethodStrFr,
		PublicYnStrFr:           reqAODNBList.Payload.PublicYnStrFr,
		BizKindStrAr:            reqAODNBList.Payload.BizKindStrAr,
		BidPoCd:                 reqAODNBList.Payload.BidPoCd,
		InternationalBidYnStrEn: reqAODNBList.Payload.InternationalBidYnStrEn,
		BdRecvEndDtStr:          reqAODNBList.Payload.BdRecvEndDtStr,
		PriceTypeStrID:          reqAODNBList.Payload.PriceTypeStrID,
		GuaranteeType:           reqAODNBList.Payload.GuaranteeType,
		BizKindStrID:            reqAODNBList.Payload.BizKindStrID,
		PriceTypeStrAr:          reqAODNBList.Payload.PriceTypeStrAr,
		RegID:                   reqAODNBList.Payload.RegID,
		PublicDtStr:             reqAODNBList.Payload.PublicDtStr,
		ExecutionPlaceStrID:     reqAODNBList.Payload.ExecutionPlaceStrID,
		EvalMethod:              reqAODNBList.Payload.EvalMethod,
		HasAnnonce:              reqAODNBList.Payload.HasAnnonce,
		RefNo:                   reqAODNBList.Payload.RefNo,
		ProcedureType:           reqAODNBList.Payload.ProcedureType,
		PublicDt:                reqAODNBList.Payload.PublicDt,
		PlanNmEn:                reqAODNBList.Payload.PlanNmEn,
		RegTypeStrEn:            reqAODNBList.Payload.RegTypeStrEn,
		ConsorYnStrID:           reqAODNBList.Payload.ConsorYnStrID,
		BiddingDocPlace:         reqAODNBList.Payload.BiddingDocPlace,
		BdRecvAddrsAr:           reqAODNBList.Payload.BdRecvAddrsAr,
		ProcChmanIDStr:          reqAODNBList.Payload.ProcChmanIDStr,
		ConsorYnStrAr:           reqAODNBList.Payload.ConsorYnStrAr,
		RealYnStrID:             reqAODNBList.Payload.RealYnStrID,
		PqRecvAddrsAr:           reqAODNBList.Payload.PqRecvAddrsAr,
		PlanNmFr:                reqAODNBList.Payload.PlanNmFr,
		RealYnStrAr:             reqAODNBList.Payload.RealYnStrAr,
		ExamChmanID:             reqAODNBList.Payload.ExamChmanID,
		GuaranteeTypeStrEn:      reqAODNBList.Payload.GuaranteeTypeStrEn,
		ProcChmanID:             reqAODNBList.Payload.ProcChmanID,
		PlanNo:                  reqAODNBList.Payload.PlanNo,
		TpRecvAddrsAr:           reqAODNBList.Payload.TpRecvAddrsAr,
		BidInstCd:               reqAODNBList.Payload.BidInstCd,
		PbkStrID:                reqAODNBList.Payload.PbkStrID,
		OnOffTypeStrFr:          reqAODNBList.Payload.OnOffTypeStrFr,
		EvalMethodStrEn:         reqAODNBList.Payload.EvalMethodStrEn,
	}

	// Insère les données dans la base de données en utilisant la fonction Create de domaines.
	if err = domains.Create(db, &offreDNB); err != nil {
		fmt.Println("Error inserting data into the database:", err)
		return err
	}

	fmt.Println("Data Données de base   saved to the  database successfully!")
	return nil

}

// FUNCGetAOATTMList récupère des données à partir de REQAOGetATTMList, les traite, et les sauvegarde dans la base de données.
func FUNCGetAOATTMList(db *gorm.DB) error {
	// Récupère les données de la liste AO depuis la source externe en utilisant la fonction REQAOGetATTMList.
	reqAOATTMList, err := REQAOGetATTMList()
	if err != nil {
		return err
	}
	fmt.Println("Data from REQAOGetATTMList:", reqAOATTMList.Payload.Data)

	for _, reqAOATTM := range reqAOATTMList.Payload.Data {

		fmt.Println("Data from REQAOGetATTMList:", reqAOATTM)

		offreATTM := domains.TunepsAOATTMList{
			ID:                  uuid.New(),
			WinnerPublicDt:      reqAOATTM.WinnerPublicDt,
			NotWinnerYn:         reqAOATTM.NotWinnerYn,
			ProcedureType:       reqAOATTM.ProcedureType,
			RefNo:               reqAOATTM.RefNo,
			BizRegNm:            reqAOATTM.BizRegNm,
			BidNo:               reqAOATTM.BidNo,
			TotalBidPriceCorr:   reqAOATTM.TotalBidPriceCorr,
			BizKind:             reqAOATTM.BizKind,
			NotWinnerReason:     reqAOATTM.NotWinnerReason,
			BidPrice:            reqAOATTM.BidPrice,
			BidNmEn:             reqAOATTM.BidNmEn,
			RegDt:               reqAOATTM.RegDt,
			BidPriceSucc:        reqAOATTM.BidPriceSucc,
			BidNmFr:             reqAOATTM.BidNmFr,
			NbrConsorExist:      reqAOATTM.NbrConsorExist,
			BidModSeq:           reqAOATTM.BidModSeq,
			BidPriceSuccCurr:    reqAOATTM.BidPriceSuccCurr,
			BidNmAr:             reqAOATTM.BidNmAr,
			BizReg:              reqAOATTM.BizReg,
			BidCls:              reqAOATTM.BidCls,
			ExecType:            reqAOATTM.ExecType,
			TotalBidPriceDcExch: reqAOATTM.TotalBidPriceDcExch,
			WinnerAfterDt:       reqAOATTM.WinnerAfterDt,
		}

		// Insère les données dans la base de données en utilisant la fonction Create de domaines.
		if err = domains.Create(db, &offreATTM); err != nil {
			fmt.Println("Error inserting data into the database:", err)
			return err
		}
	}
	fmt.Println("Data  Attribution du marché   saved to the  database successfully!")
	return nil

}

// ************************bOUTON RESULTAT*************************************

// FUNCGetAORDNBList récupère des données à partir de REQAOGetRDNBList, les traite, et les sauvegarde dans la base de données.
func FUNCGetAORDNBList(db *gorm.DB) error {
	// Récupère les données de la liste AO depuis la source externe en utilisant la fonction REQAOGetRDNBList.
	reqAORDNBList, err := REQAOGetRDNBList()
	if err != nil {
		return err
	}
	fmt.Println("Data from REQAOGetRDNBList:", reqAORDNBList.Payload)

	offreRDNB := domains.TunepsAORDNBList{
		ID:                      uuid.New(),
		BdRecvEndDt:             reqAORDNBList.Payload.BdRecvEndDt,
		BiddingDocPriceCurr:     reqAORDNBList.Payload.BiddingDocPriceCurr,
		EncCert:                 reqAORDNBList.Payload.EncCert,
		GeneralYn:               reqAORDNBList.Payload.GeneralYn,
		BdRecvStartDt:           reqAORDNBList.Payload.BdRecvStartDt,
		BidNo:                   reqAORDNBList.Payload.BidNo,
		ExecutionPlaceStrEn:     reqAORDNBList.Payload.ExecutionPlaceStrEn,
		ProcedureTypeStrEn:      reqAORDNBList.Payload.ProcedureTypeStrEn,
		PbkStrAr:                reqAORDNBList.Payload.PbkStrAr,
		GuaranteeTypeStrFr:      reqAORDNBList.Payload.GuaranteeTypeStrFr,
		RealYnStrFr:             reqAORDNBList.Payload.RealYnStrFr,
		FinancialMethodStrEn:    reqAORDNBList.Payload.FinancialMethodStrEn,
		BiddingDocPrice:         reqAORDNBList.Payload.BiddingDocPrice,
		RealYn:                  reqAORDNBList.Payload.RealYn,
		BdOpenDt:                reqAORDNBList.Payload.BdOpenDt,
		BidNmAr:                 reqAORDNBList.Payload.BidNmAr,
		OnOffTypeStrEn:          reqAORDNBList.Payload.OnOffTypeStrEn,
		PriceTypeStrFr:          reqAORDNBList.Payload.PriceTypeStrFr,
		BizKindStrEn:            reqAORDNBList.Payload.BizKindStrEn,
		InternationalBidYnStrAr: reqAORDNBList.Payload.InternationalBidYnStrAr,
		OnOffType:               reqAORDNBList.Payload.OnOffType,
		InstRegNo:               reqAORDNBList.Payload.InstRegNo,
		ConsorYn:                reqAORDNBList.Payload.ConsorYn,
		RegTypeStrFr:            reqAORDNBList.Payload.RegTypeStrFr,
		ProcedureTypeStrFr:      reqAORDNBList.Payload.ProcedureTypeStrFr,
		PriceTypeStrEn:          reqAORDNBList.Payload.PriceTypeStrEn,
		InternationalBidYnStrID: reqAORDNBList.Payload.InternationalBidYnStrID,
		FinancialMethodStrFr:    reqAORDNBList.Payload.FinancialMethodStrFr,
		BdDepartEn:              reqAORDNBList.Payload.BdDepartEn,
		BidExpiredDays:          reqAORDNBList.Payload.BidExpiredDays,
		ConsorYnStrFr:           reqAORDNBList.Payload.ConsorYnStrFr,
		PublicYnStrAr:           reqAORDNBList.Payload.PublicYnStrAr,
		PlanNmAr:                reqAORDNBList.Payload.PlanNmAr,
		PublicYn:                reqAORDNBList.Payload.PublicYn,
		PublicYnStrEn:           reqAORDNBList.Payload.PublicYnStrEn,
		BizKindStrFr:            reqAORDNBList.Payload.BizKindStrFr,
		EvalMethodStrAr:         reqAORDNBList.Payload.EvalMethodStrAr,
		TpRecvAddrsFr:           reqAORDNBList.Payload.TpRecvAddrsFr,
		ExecutionPlace:          reqAORDNBList.Payload.ExecutionPlace,
		BidInstNm:               reqAORDNBList.Payload.BidInstNm,
		BdRecvStartDtStr:        reqAORDNBList.Payload.BdRecvStartDtStr,
		RegTypeStrAr:            reqAORDNBList.Payload.RegTypeStrAr,
		StaffNm:                 reqAORDNBList.Payload.StaffNm,
		BdRecvAddrsEn:           reqAORDNBList.Payload.BdRecvAddrsEn,
		FinancingYn:             reqAORDNBList.Payload.FinancingYn,
		RegType:                 reqAORDNBList.Payload.RegType,
		ConsorYnStrEn:           reqAORDNBList.Payload.ConsorYnStrEn,
		InternationalBidYn:      reqAORDNBList.Payload.InternationalBidYn,
		ExamChmanIDStr:          reqAORDNBList.Payload.ExamChmanIDStr,
		PriceType:               reqAORDNBList.Payload.PriceType,
		ExecutionPlaceStrFr:     reqAORDNBList.Payload.ExecutionPlaceStrFr,
		BdOpenDtStr:             reqAORDNBList.Payload.BdOpenDtStr,
		PlanModSeq:              reqAORDNBList.Payload.PlanModSeq,
		FrameworkYn:             reqAORDNBList.Payload.FrameworkYn,
		BdDepartFr:              reqAORDNBList.Payload.BdDepartFr,
		EvalMethodStrID:         reqAORDNBList.Payload.EvalMethodStrID,
		BizKind:                 reqAORDNBList.Payload.BizKind,
		PqRecvAddrsEn:           reqAORDNBList.Payload.PqRecvAddrsEn,
		TpRecvAddrsEn:           reqAORDNBList.Payload.TpRecvAddrsEn,
		BidModSeq:               reqAORDNBList.Payload.BidModSeq,
		GuaranteeTypeStrID:      reqAORDNBList.Payload.GuaranteeTypeStrID,
		BidNmFr:                 reqAORDNBList.Payload.BidNmFr,
		RegTypeStrID:            reqAORDNBList.Payload.RegTypeStrID,
		BdRecvAddrsFr:           reqAORDNBList.Payload.BdRecvAddrsFr,
		ProcedureTypeStrID:      reqAORDNBList.Payload.ProcedureTypeStrID,
		FinancialMethodStrID:    reqAORDNBList.Payload.FinancialMethodStrID,
		InternationalBidYnStrFr: reqAORDNBList.Payload.InternationalBidYnStrFr,
		PbkStrEn:                reqAORDNBList.Payload.PbkStrEn,
		OnOffTypeStrID:          reqAORDNBList.Payload.OnOffTypeStrID,
		EpBidMasterID:           reqAORDNBList.Payload.EpBidMasterID,
		PqRecvAddrsFr:           reqAORDNBList.Payload.PqRecvAddrsFr,
		BdDepartAr:              reqAORDNBList.Payload.BdDepartAr,
		BidNmEn:                 reqAORDNBList.Payload.BidNmEn,
		FinancialMethod:         reqAORDNBList.Payload.FinancialMethod,
		OnOffTypeStrAr:          reqAORDNBList.Payload.OnOffTypeStrAr,
		ExecutionPlaceStrAr:     reqAORDNBList.Payload.ExecutionPlaceStrAr,
		ProcedureTypeStrAr:      reqAORDNBList.Payload.ProcedureTypeStrAr,
		GuaranteeTypeStrAr:      reqAORDNBList.Payload.GuaranteeTypeStrAr,
		FinancialMethodStrAr:    reqAORDNBList.Payload.FinancialMethodStrAr,
		PbkStrFr:                reqAORDNBList.Payload.PbkStrFr,
		EvalMethodStrFr:         reqAORDNBList.Payload.EvalMethodStrFr,
		PublicYnStrFr:           reqAORDNBList.Payload.PublicYnStrFr,
		BizKindStrAr:            reqAORDNBList.Payload.BizKindStrAr,
		BidPoCd:                 reqAORDNBList.Payload.BidPoCd,
		InternationalBidYnStrEn: reqAORDNBList.Payload.InternationalBidYnStrEn,
		BdRecvEndDtStr:          reqAORDNBList.Payload.BdRecvEndDtStr,
		PriceTypeStrID:          reqAORDNBList.Payload.PriceTypeStrID,
		GuaranteeType:           reqAORDNBList.Payload.GuaranteeType,
		BizKindStrID:            reqAORDNBList.Payload.BizKindStrID,
		PriceTypeStrAr:          reqAORDNBList.Payload.PriceTypeStrAr,
		RegID:                   reqAORDNBList.Payload.RegID,
		PublicDtStr:             reqAORDNBList.Payload.PublicDtStr,
		ExecutionPlaceStrID:     reqAORDNBList.Payload.ExecutionPlaceStrID,
		EvalMethod:              reqAORDNBList.Payload.EvalMethod,
		HasAnnonce:              reqAORDNBList.Payload.HasAnnonce,
		RefNo:                   reqAORDNBList.Payload.RefNo,
		ProcedureType:           reqAORDNBList.Payload.ProcedureType,
		PublicDt:                reqAORDNBList.Payload.PublicDt,
		PlanNmEn:                reqAORDNBList.Payload.PlanNmEn,
		RegTypeStrEn:            reqAORDNBList.Payload.RegTypeStrEn,
		ConsorYnStrID:           reqAORDNBList.Payload.ConsorYnStrID,
		BiddingDocPlace:         reqAORDNBList.Payload.BiddingDocPlace,
		BdRecvAddrsAr:           reqAORDNBList.Payload.BdRecvAddrsAr,
		ProcChmanIDStr:          reqAORDNBList.Payload.ProcChmanIDStr,
		ConsorYnStrAr:           reqAORDNBList.Payload.ConsorYnStrAr,
		RealYnStrID:             reqAORDNBList.Payload.RealYnStrID,
		PqRecvAddrsAr:           reqAORDNBList.Payload.PqRecvAddrsAr,
		PlanNmFr:                reqAORDNBList.Payload.PlanNmFr,
		RealYnStrAr:             reqAORDNBList.Payload.RealYnStrAr,
		ExamChmanID:             reqAORDNBList.Payload.ExamChmanID,
		GuaranteeTypeStrEn:      reqAORDNBList.Payload.GuaranteeTypeStrEn,
		ProcChmanID:             reqAORDNBList.Payload.ProcChmanID,
		PlanNo:                  reqAORDNBList.Payload.PlanNo,
		TpRecvAddrsAr:           reqAORDNBList.Payload.TpRecvAddrsAr,
		BidInstCd:               reqAORDNBList.Payload.BidInstCd,
		PbkStrID:                reqAORDNBList.Payload.PbkStrID,
		OnOffTypeStrFr:          reqAORDNBList.Payload.OnOffTypeStrFr,
		EvalMethodStrEn:         reqAORDNBList.Payload.EvalMethodStrEn,
	}

	// Insère les données dans la base de données en utilisant la fonction Create de domaines.
	if err = domains.Create(db, &offreRDNB); err != nil {
		fmt.Println("Error inserting data into the database:", err)
		return err
	}

	fmt.Println("Data  Données de base-RESULTATS   saved to the  database successfully!")
	return nil
}

// FUNCGetAORLotList récupère des données à partir de REQAOGetRLotList, les traite, et les sauvegarde dans la base de données.
func FUNCGetAORLotList(db *gorm.DB) error {
	// Récupère les données de la liste AO depuis la source externe en utilisant la fonction REQAOGetRLotList.
	reqAORLotList, err := REQAOGetRLotList()
	if err != nil {
		return err
	}
	fmt.Println("Data from REQAOGetRLotList:", reqAORLotList.Payload.Data)

	for _, reqAORLot := range reqAORLotList.Payload.Data {
		offreRLot := domains.TunepsAORLotList{
			ID: uuid.New(),

			RefNo:         reqAORLot.RefNo,
			BizKindFrStr:  reqAORLot.BizKindFrStr,
			BizKindArStr:  reqAORLot.BizKindArStr,
			BidNo:         reqAORLot.BidNo,
			Cnt:           reqAORLot.Cnt,
			BizKindEnStr:  reqAORLot.BizKindEnStr,
			BizKind:       reqAORLot.BizKind,
			BidNmEn:       reqAORLot.BidNmEn,
			RegDt:         reqAORLot.RegDt,
			BidModSeq:     reqAORLot.BidModSeq,
			BidNmAr:       reqAORLot.BidNmAr,
			Rk:            reqAORLot.Rk,
			BidFailReason: reqAORLot.BidFailReason,
			RegID:         reqAORLot.RegID,
			BidCls:        reqAORLot.BidCls,
			ExecType:      reqAORLot.ExecType,
			BidNmFr:       reqAORLot.BidNmFr,
		}

		// Insère les données dans la base de données en utilisant la fonction Create de domaines.
		if err = domains.Create(db, &offreRLot); err != nil {
			fmt.Println("Error inserting data into the database:", err)
			return err
		}
	}

	fmt.Println("Data  Lot-RESULTATS   saved to the  database successfully!")
	return nil

}

// FUNCGetAORSList récupère des données à partir de REQAOGetRSList, les traite, et les sauvegarde dans la base de données.
func FUNCGetAORSList(db *gorm.DB) error {
	// Récupère les données de la liste AO depuis la source externe en utilisant la fonction REQAOGetRSList.
	reqAORSList, err := REQAOGetRSList()
	if err != nil {
		return err
	}
	fmt.Println("Data from REQAOGetRSList:", reqAORSList.Payload.Data)

	for _, reqAORS := range reqAORSList.Payload.Data {
		offreRS := domains.TunepsAORSList{
			ID:                  uuid.New(),
			CdNmStrEn:           reqAORS.CdNmStrEn,
			IneligibleCd:        reqAORS.IneligibleCd,
			TechEvalResult:      reqAORS.TechEvalResult,
			BidNo:               reqAORS.BidNo,
			CdPfTechFr:          reqAORS.CdPfTechFr,
			BdRecvDtStr:         reqAORS.BdRecvDtStr,
			EpBidOpenResultID:   reqAORS.EpBidOpenResultID,
			CdPfFinaFr:          reqAORS.CdPfFinaFr,
			Rank:                reqAORS.Rank,
			OnlineRegYn:         reqAORS.OnlineRegYn,
			DecAttachFileCnt:    reqAORS.DecAttachFileCnt,
			BidCls:              reqAORS.BidCls,
			TotalBidPriceDcExch: reqAORS.TotalBidPriceDcExch,
			BizRegNo:            reqAORS.BizRegNo,
			BizRegNm:            reqAORS.BizRegNm,
			CdPfTechEn:          reqAORS.CdPfTechEn,
			FinaResultReason:    reqAORS.FinaResultReason,
			CdNmStrFr:           reqAORS.CdNmStrFr,
			Cnt:                 reqAORS.Cnt,
			CdPfTechAr:          reqAORS.CdPfTechAr,
			IneligibleReason:    reqAORS.IneligibleReason,
			CdPfFinaAr:          reqAORS.CdPfFinaAr,
			BdRecvDt:            reqAORS.BdRecvDt,
			BidModSeq:           reqAORS.BidModSeq,
			CdNmStrAr:           reqAORS.CdNmStrAr,
			Rk:                  reqAORS.Rk,
			BizReg:              reqAORS.BizReg,
			FinaEvalResult:      reqAORS.FinaEvalResult,
			TechResultReason:    reqAORS.TechResultReason,
			CdPfFinaEn:          reqAORS.CdPfFinaEn,
		}
		// Insère les données dans la base de données en utilisant la fonction Create de domaines.
		if err = domains.Create(db, &offreRS); err != nil {
			fmt.Println("Error inserting data into the database:", err)
			return err
		}
	}

	fmt.Println("Data  Soumissionaires-RESULTATS   saved to the  database successfully!")
	return nil
}

/// ************************ Shopping Mall *******************************************

// FUNCGetAOSMList récupère des données à partir de REQAOGetSMList, les traite, et les sauvegarde dans la base de données.
func FUNCGetAOSMList(db *gorm.DB) error {
	// Récupère les données de la liste AO depuis la source externe en utilisant la fonction REQAOGetSMList.
	reqAOSMList, err := REQAOGetSMList()
	if err != nil {
		return err
	}
	fmt.Println("Data from REQAOGetSMList:", reqAOSMList.Payload.Data)

	for _, reqAOSM := range reqAOSMList.Payload.Data {
		offreSM := domains.TunepsSMList{
			ID:             uuid.New(),
			PublicDt:       reqAOSM.PublicDt,
			SpShopMasterID: reqAOSM.SpShopMasterID,
			ShopModSeq:     reqAOSM.ShopModSeq,
			InstNm:         reqAOSM.InstNm,
			ShopNmEn:       reqAOSM.ShopNmEn,
			ShopNmFr:       reqAOSM.ShopNmFr,
			SpRecvEndDt:    reqAOSM.SpRecvEndDt,
			ShopNo:         reqAOSM.ShopNo,
			ShopNmAr:       reqAOSM.ShopNmAr,
		}
		// Insère les données dans la base de données en utilisant la fonction Create de domaines.
		if err = domains.Create(db, &offreSM); err != nil {
			fmt.Println("Error inserting data into the database:", err)
			return err
		}
	}

	fmt.Println("Data  Shopping Mall   saved to the  database successfully!")
	return nil
}

// FUNCGetAOSMInfoList récupère des données à partir de REQAOGetSMInfoList, les traite, et les sauvegarde dans la base de données.
func FUNCGetAOSMInfoList(db *gorm.DB) error {
	// Récupère les données de la liste AO depuis la source externe en utilisant la fonction REQAOGetSMInfoList.
	reqAOSMInfoList, err := REQAOGetSMInfoList()
	if err != nil {
		return err
	}

	offreSMInfo := domains.TunepsSMInfoList{
		ID:                  uuid.New(),
		RefNo:               reqAOSMInfoList.Payload.RefNo,
		PublicDt:            reqAOSMInfoList.Payload.PublicDt,
		ShopModSeq:          reqAOSMInfoList.Payload.ShopModSeq,
		PublicYn:            reqAOSMInfoList.Payload.PublicYn,
		MsAr:                reqAOSMInfoList.Payload.MsAr,
		PbkStrEn:            reqAOSMInfoList.Payload.PbkStrEn,
		RegNm:               reqAOSMInfoList.Payload.RegNm,
		BizKindStrFr:        reqAOSMInfoList.Payload.BizKindStrFr,
		CdNmAr:              reqAOSMInfoList.Payload.CdNmAr,
		PbkStrAr:            reqAOSMInfoList.Payload.PbkStrAr,
		ShopNmAr:            reqAOSMInfoList.Payload.ShopNmAr,
		BcNm:                reqAOSMInfoList.Payload.BcNm,
		BcID:                reqAOSMInfoList.Payload.BcID,
		EpAr:                reqAOSMInfoList.Payload.EpAr,
		InstNm:              reqAOSMInfoList.Payload.InstNm,
		InternationalShopYn: reqAOSMInfoList.Payload.InternationalShopYn,
		SpRecvEndDtStr:      reqAOSMInfoList.Payload.SpRecvEndDtStr,
		ShopExpiredDays:     reqAOSMInfoList.Payload.ShopExpiredDays,
		SpRecvStartDt:       reqAOSMInfoList.Payload.SpRecvStartDt,
		EpEn:                reqAOSMInfoList.Payload.EpEn,
		ExamNm:              reqAOSMInfoList.Payload.ExamNm,
		MsEn:                reqAOSMInfoList.Payload.MsEn,
		CdNmEn:              reqAOSMInfoList.Payload.CdNmEn,
		ShopInstCd:          reqAOSMInfoList.Payload.ShopInstCd,
		PbkStrFr:            reqAOSMInfoList.Payload.PbkStrFr,
		MsFr:                reqAOSMInfoList.Payload.MsFr,
		PbkCd:               reqAOSMInfoList.Payload.PbkCd,
		BizKindStrEn:        reqAOSMInfoList.Payload.BizKindStrEn,
		ShopNmFr:            reqAOSMInfoList.Payload.ShopNmFr,
		MasterCdBc:          reqAOSMInfoList.Payload.MasterCdBc,
		BizKindStrAr:        reqAOSMInfoList.Payload.BizKindStrAr,
		SpRecvEndDt:         reqAOSMInfoList.Payload.SpRecvEndDt,
		SpShopMasterID:      reqAOSMInfoList.Payload.SpShopMasterID,
		BizKindCd:           reqAOSMInfoList.Payload.BizKindCd,
		BizKindStrID:        reqAOSMInfoList.Payload.BizKindStrID,
		ShopNmEn:            reqAOSMInfoList.Payload.ShopNmEn,
		PbkStrID:            reqAOSMInfoList.Payload.PbkStrID,
		ExamID:              reqAOSMInfoList.Payload.ExamID,
		EpFr:                reqAOSMInfoList.Payload.EpFr,
		CdNmFr:              reqAOSMInfoList.Payload.CdNmFr,
		RegID:               reqAOSMInfoList.Payload.RegID,
		PublicDtStr:         reqAOSMInfoList.Payload.PublicDtStr,
		MasterCdExam:        reqAOSMInfoList.Payload.MasterCdExam,
		ShopNo:              reqAOSMInfoList.Payload.ShopNo,
		SpOpenDt:            reqAOSMInfoList.Payload.SpOpenDt,
		HasAnnonce:          reqAOSMInfoList.Payload.HasAnnonce,
	}

	// Insère les données dans la base de données en utilisant la fonction Create de domaines.
	if err = domains.Create(db, &offreSMInfo); err != nil {
		fmt.Println("Error inserting data into the database:", err)
		return err
	}

	fmt.Println("Data  Shopping Mall Information générales   saved to the  database successfully!")
	return nil
}

// FUNCGetAOSMLotList récupère des données à partir de REQAOGetSMLotList, les traite, et les sauvegarde dans la base de données.
func FUNCGetAOSMLotList(db *gorm.DB) error {
	// Récupère les données de la liste AO depuis la source externe en utilisant la fonction REQAOGetSMLotList.
	reqAOSMLotList, err := REQAOGetSMLotList()
	if err != nil {
		return err
	}

	for _, reqAOSMLot := range reqAOSMLotList.Payload {
		offreSMLot := domains.TunepsSMLotList{
			ID:                  uuid.New(),
			ShopCls:             reqAOSMLot.ShopCls,
			BudgetAmt:           reqAOSMLot.BudgetAmt,
			ShopClsDtlDesc:      reqAOSMLot.ShopClsDtlDesc,
			ShopModSeq:          reqAOSMLot.ShopModSeq,
			GuaranteeAmountCurr: reqAOSMLot.GuaranteeAmountCurr,
			GuaranteeAmount:     reqAOSMLot.GuaranteeAmount,
			ShopNo:              reqAOSMLot.ShopNo,
			ShopClsNm:           reqAOSMLot.ShopClsNm,
			BudgetAmtCurr:       reqAOSMLot.BudgetAmtCurr,
			SpShopClsID:         reqAOSMLot.SpShopClsID,
		}
		// Insère les données dans la base de données en utilisant la fonction Create de domaines.
		if err = domains.Create(db, &offreSMLot); err != nil {
			fmt.Println("Error inserting data into the database:", err)
			return err
		}
	}
	fmt.Println("Data  Lot Shopping Mall  saved to the  database successfully!")
	return nil

}

////Documents de la consultation

// FUNCGetAOSMDocList récupère des données à partir de REQAOGetSMDocList, les traite, et les sauvegarde dans la base de données.
func FUNCGetAOSMDocList(db *gorm.DB) error {
	// Récupère les données de la liste AO depuis la source externe en utilisant la fonction REQAOGetSMDocList.
	reqAOSMDocList, err := REQAOGetSMDocList()
	if err != nil {
		return err
	}
	for _, reqAOSMDoc := range reqAOSMDocList.Payload {
		offreSMDoc := domains.TunepsSMDocList{
			ID:                 uuid.New(),
			FileNm:             reqAOSMDoc.FileNm,
			ShopModSeq:         reqAOSMDoc.ShopModSeq,
			SeqNo:              reqAOSMDoc.SeqNo,
			NodeRef:            reqAOSMDoc.NodeRef,
			CdNmAr:             reqAOSMDoc.CdNmAr,
			CdNmFr:             reqAOSMDoc.CdNmFr,
			DocCd:              reqAOSMDoc.DocCd,
			FileLoc:            reqAOSMDoc.FileLoc,
			ShopNo:             reqAOSMDoc.ShopNo,
			SpShopAttachFileID: reqAOSMDoc.SpShopAttachFileID,
			CdNmEn:             reqAOSMDoc.CdNmEn,
		}
		// Insère les données dans la base de données en utilisant la fonction Create de domaines.
		if err = domains.Create(db, &offreSMDoc); err != nil {
			fmt.Println("Error inserting data into the database:", err)
			return err
		}
	}
	fmt.Println("Data  Documents  Shopping Mall  saved to the  database successfully!")
	return nil

}
