package tuneps

import (
	"labs/domains"
	"log"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Database represents the database instance for the tuneps package.
type Database struct {
	DB *gorm.DB
}

// NewTunepsRepository initialise la base de données Tuneps.
func NewTunepsRepository(db *gorm.DB) {

	// Migrer automatiquement la structure Tuneps
	if err := db.AutoMigrate(&domains.TunepsAOList{}, &domains.TunepsAOInfoList{}, &domains.TunepsAOLotList{}, &domains.TunepsAOCCList{}, &domains.TunepsAOAgList{}, &domains.TunepsAOPRList{}, &domains.TunepsAODNBList{}, &domains.TunepsAOATTMList{}, &domains.TunepsAORDNBList{}, &domains.TunepsAORLotList{}, &domains.TunepsAORSList{}, &domains.TunepsSMList{}, &domains.TunepsSMInfoList{}, &domains.TunepsSMLotList{}, &domains.TunepsSMDocList{}); err != nil {
		logrus.Fatal("An error occurred during automatic migration of the TunepsAOList structure. Error: ", err)
	}

	// Initialiser les données TunepsAOList
	if err := FUNCGetAOList(db); err != nil {
		log.Fatalf("Erreur lors de l'initialisation des données TunepsAOList : %v", err)
	}

	// Initialiser les données TunepsAOInfoList

	if err := FUNCGetAOInfoList(db); err != nil {
		log.Fatal("Erreur lors de l'initialisation des données TunepsAOInfoList : ", err)
	}

	//Initialiser les données TunepsAOLotList

	if err := FUNCGetAOLotList(db); err != nil {
		log.Fatal("Erreur lors de l'initialisation des données TunepsAOLotList : ", err)
	}

	// Initialiser les données TunepsAOCCList

	if err := FUNCGetAOCCList(db); err != nil {
		log.Fatal("Erreur lors de l'initialisation des données TunepsAOCCList : ", err)
	}

	// Initialiser les donnéesTunepsAOAgList

	if err := FUNCGetAOAgList(db); err != nil {
		log.Fatal("Erreur lors de l'initialisation des données TunepsAOAgList : ", err)
	}

	// Initialiser les donnéesTunepsAOPRList

	if err := FUNCGetAOProdList(db); err != nil {
		log.Fatal("Erreur lors de l'initialisation des données TunepsAOPRList : ", err)
	}

	// *********************** BOUTON List soumissionaire(s) retenu(s)******************
	// Initialiser les donnéesAOGetDonneDeBase

	if err := FUNCGetAODNBList(db); err != nil {
		log.Fatal("Erreur lors de l'initialisation des données TunepsAODNBList : ", err)
	}

	// Initialiser les données AOGetAttributionDuMarché

	if err := FUNCGetAOATTMList(db); err != nil {
		log.Fatal("Erreur lors de l'initialisation des données TunepsAOATTMList : ", err)
	}

	//************************bOUTON RESULTAT*************************************

	// Initialiser les donnéesAOGetDonneDeBase

	if err := FUNCGetAORDNBList(db); err != nil {
		log.Fatal("Erreur lors de l'initialisation des données TunepsAORDNBList : ", err)
	}
	// Initialiser les donnéesTunepsAORLotList

	if err := FUNCGetAORLotList(db); err != nil {
		log.Fatal("Erreur lors de l'initialisation des données TunepsAORLotList : ", err)
	}

	// Initialiser les donnéesTunepsAORSList

	if err := FUNCGetAORSList(db); err != nil {
		log.Fatal("Erreur lors de l'initialisation des données TunepsAORSList : ", err)
	}

	// ************************ Shopping Mall *******************************************
	// Initialiser les donnéesTunepsAOSMList

	if err := FUNCGetAOSMList(db); err != nil {
		log.Fatal("Erreur lors de l'initialisation des données TunepsAOSMList : ", err)
	}

	// Initialiser les donnéesTunepsAOSMInfoList

	if err := FUNCGetAOSMInfoList(db); err != nil {
		log.Fatal("Erreur lors de l'initialisation des données TunepsAOSMInfoList : ", err)
	}

	// Initialiser les donnéesTunepsAOSMLotList

	if err := FUNCGetAOSMLotList(db); err != nil {
		log.Fatal("Erreur lors de l'initialisation des données TunepsAOSMLotList : ", err)
	}

	// Initialiser les donnéesTunepsAOSMDocList

	if err := FUNCGetAOSMDocList(db); err != nil {
		log.Fatal("Erreur lors de l'initialisation des données TunepsAOSMDocList : ", err)
	}

}
