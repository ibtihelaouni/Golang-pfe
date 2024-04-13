package tuneps

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CompanyRouterInit initializes the routes related to companies.
func TunepsRouterInit(router *gin.RouterGroup, db *gorm.DB) {

	// Initialize database instance
	baseInstance := Database{DB: db}

	// Automigrate / Update table
	NewTunepsRepository(db)

	// Private
	tuneps := router.Group("/tuneps")
	{
		// ************************ Appel d'offres *******************************************
		// GET endpoint to retrieve all AO for tuneps
		tuneps.GET("/AO", baseInstance.GetTunepsAO)
		// GET endpoint to retrieve  AO by ID  for tuneps
		tuneps.GET("/AO/:ID", baseInstance.GetTunepsAOByID)

		// GET endpoint to retrieve all AO Info for tuneps
		tuneps.GET("/AOInfo", baseInstance.GetTunepsAOInfo)
		// GET endpoint to retrieve  AO Info  by ID  for tuneps
		tuneps.GET("/AOInfo/:ID", baseInstance.GetTunepsAOInfoByID)

		// GET endpoint to retrieve all AO Lot for tuneps
		tuneps.GET("/AOLot", baseInstance.GetTunepsAOLot)
		// GET endpoint to retrieve  AO Lot  by ID  for tuneps
		tuneps.GET("/AOLot/:ID", baseInstance.GetTunepsAOLotByID)

		// GET endpoint to retrieve all AO CC for tuneps
		tuneps.GET("/AOCC", baseInstance.GetTunepsAOCC)
		// GET endpoint to retrieve  AO CC  by ID  for tuneps
		tuneps.GET("/AOCC/:ID", baseInstance.GetTunepsAOCCByID)

		// GET endpoint to retrieve all AO Ag for tuneps
		tuneps.GET("/AOAg", baseInstance.GetTunepsAOAg)
		// GET endpoint to retrieve  AO Ag  by ID  for tuneps
		tuneps.GET("/AOAg/:ID", baseInstance.GetTunepsAOAgByID)

		// GET endpoint to retrieve all AO PR for tuneps
		tuneps.GET("/AOPR", baseInstance.GetTunepsAOPR)
		// GET endpoint to retrieve  AO PR  by ID  for tuneps
		tuneps.GET("/AOPR/:ID", baseInstance.GetTunepsAOPRByID)

		//************* BOUTON List soumissionaire(s) retenu(s)******************

		// GET endpoint to retrieve all AO DNB for tuneps
		tuneps.GET("/AODNB", baseInstance.GetTunepsAODNB)
		// GET endpoint to retrieve  AO DNB  by ID  for tuneps
		tuneps.GET("/AODNB/:ID", baseInstance.GetTunepsAODNBByID)

		// GET endpoint to retrieve all AO ATTM for tuneps
		tuneps.GET("/AOATTM", baseInstance.GetTunepsAOATTM)
		// GET endpoint to retrieve  AO ATTM  by ID  for tuneps
		tuneps.GET("/AOATTM/:ID", baseInstance.GetTunepsAOATTMByID)

		// ************************bOUTON RESULTAT*************************************

		// GET endpoint to retrieve all AO RDNB for tuneps
		tuneps.GET("/AORDNB", baseInstance.GetTunepsAORDNB)
		// GET endpoint to retrieve  AO RDNB  by ID  for tuneps
		tuneps.GET("/AORDNB/:ID", baseInstance.GetTunepsAORDNBByID)

		// GET endpoint to retrieve all AO RESULTAT Lot for tuneps
		tuneps.GET("/AORLot", baseInstance.GetTunepsAORLot)
		// GET endpoint to retrieve  AO RESULTAT Lot  by ID  for tuneps
		tuneps.GET("/AORLot/:ID", baseInstance.GetTunepsAORLotByID)

		// GET endpoint to retrieve all AO RS for tuneps
		tuneps.GET("/AORS", baseInstance.GetTunepsAORS)
		// GET endpoint to retrieve  AO RS  by ID  for tuneps
		tuneps.GET("/AORS/:ID", baseInstance.GetTunepsAORSByID)

		/// ************************ Shopping Mall *******************************************

		// GET endpoint to retrieve all SM for tuneps
		tuneps.GET("/SM", baseInstance.GetTunepsSM)
		// GET endpoint to retrieve  SM by ID  for tuneps
		tuneps.GET("/SM/:ID", baseInstance.GetTunepsSMByID)

		// GET endpoint to retrieve all SM info for tuneps
		tuneps.GET("/SMInfo", baseInstance.GetTunepsSMInfo)
		// GET endpoint to retrieve  SM info by ID  for tuneps
		tuneps.GET("/SMInfo/:ID", baseInstance.GetTunepsSMInfoByID)

		// GET endpoint to retrieve all SM Lot for tuneps
		tuneps.GET("/SMLot", baseInstance.GetTunepsSMLot)
		// GET endpoint to retrieve  SM Lot  by ID  for tuneps
		tuneps.GET("/SMLot/:ID", baseInstance.GetTunepsSMLotByID)

		// GET endpoint to retrieve all SM doc for tuneps
		tuneps.GET("/SMDoc", baseInstance.GetTunepsSMDoc)
		// GET endpoint to retrieve  SM doc  by ID  for tuneps
		tuneps.GET("/SMDoc/:ID", baseInstance.GetTunepsSMDocByID)

	}
}
