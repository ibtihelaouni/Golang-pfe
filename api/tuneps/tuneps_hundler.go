package tuneps

import (
	"errors"
	"labs/domains"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ************************ Appel d'offres *******************************************
// AO
// GetTuneps retrieves TUNEPS data from the database.
func (db Database) GetTunepsAO(ctx *gin.Context) {
	var tunepsData []domains.TunepsAOList

	if err := db.DB.Find(&tunepsData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	ctx.JSON(http.StatusOK, tunepsData)
}

// GetTunepsAOByID retrieves TUNEPS data from the database by ID.
func (db Database) GetTunepsAOByID(ctx *gin.Context) {
	// Get the ID from the request parameters
	id := ctx.Param("ID") // Remplacez "id" par "ID"

	// Check if the ID is empty
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var tunepsData domains.TunepsAOList

	// Fetch TUNEPS data by ID
	if err := db.DB.First(&tunepsData, "id = ?", id).Error; err != nil {
		// Check if the record is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	// Return JSON response with ID included
	ctx.JSON(http.StatusOK, gin.H{"id": id, "data": tunepsData})
}

// AOInfo
// GetTunepsAOInfo retrieves TUNEPS data from the database.
func (db Database) GetTunepsAOInfo(ctx *gin.Context) {
	var tunepsData []domains.TunepsAOInfoList

	if err := db.DB.Find(&tunepsData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	ctx.JSON(http.StatusOK, tunepsData)
}

// GetTunepsAOInfoByID retrieves TUNEPS data from the database by ID.
func (db Database) GetTunepsAOInfoByID(ctx *gin.Context) {
	// Get the ID from the request parameters
	id := ctx.Param("ID") // Remplacez "id" par "ID"

	// Check if the ID is empty
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var tunepsData domains.TunepsAOInfoList

	// Fetch TUNEPS data by ID
	if err := db.DB.First(&tunepsData, "id = ?", id).Error; err != nil {
		// Check if the record is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	// Return JSON response with ID included
	ctx.JSON(http.StatusOK, gin.H{"id": id, "data": tunepsData})
}

// AOLot
// GetTunepsAOLot retrieves TUNEPS data from the database.
func (db Database) GetTunepsAOLot(ctx *gin.Context) {
	var tunepsData []domains.TunepsAOLotList

	if err := db.DB.Find(&tunepsData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	ctx.JSON(http.StatusOK, tunepsData)
}

// GetTunepsAOLotByID retrieves TUNEPS data from the database by ID.
func (db Database) GetTunepsAOLotByID(ctx *gin.Context) {
	// Get the ID from the request parameters
	id := ctx.Param("ID") // Remplacez "id" par "ID"

	// Check if the ID is empty
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var tunepsData domains.TunepsAOLotList

	// Fetch TUNEPS data by ID
	if err := db.DB.First(&tunepsData, "id = ?", id).Error; err != nil {
		// Check if the record is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	// Return JSON response with ID included
	ctx.JSON(http.StatusOK, gin.H{"id": id, "data": tunepsData})
}

// AOCC
// GetTunepsAOCC retrieves TUNEPS data from the database.
func (db Database) GetTunepsAOCC(ctx *gin.Context) {
	var tunepsData []domains.TunepsAOCCList

	if err := db.DB.Find(&tunepsData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	ctx.JSON(http.StatusOK, tunepsData)
}

// GetTunepsAOCCByID retrieves TUNEPS data from the database by ID.
func (db Database) GetTunepsAOCCByID(ctx *gin.Context) {
	// Get the ID from the request parameters
	id := ctx.Param("ID") // Remplacez "id" par "ID"

	// Check if the ID is empty
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var tunepsData domains.TunepsAOCCList

	// Fetch TUNEPS data by ID
	if err := db.DB.First(&tunepsData, "id = ?", id).Error; err != nil {
		// Check if the record is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	// Return JSON response with ID included
	ctx.JSON(http.StatusOK, gin.H{"id": id, "data": tunepsData})
}

// AOAg
// GetTunepsAOAg retrieves TUNEPS data from the database.
func (db Database) GetTunepsAOAg(ctx *gin.Context) {
	var tunepsData []domains.TunepsAOAgList

	if err := db.DB.Find(&tunepsData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	ctx.JSON(http.StatusOK, tunepsData)
}

// GetTunepsAOAgByID retrieves TUNEPS data from the database by ID.
func (db Database) GetTunepsAOAgByID(ctx *gin.Context) {
	// Get the ID from the request parameters
	id := ctx.Param("ID") // Remplacez "id" par "ID"

	// Check if the ID is empty
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var tunepsData domains.TunepsAOAgList

	// Fetch TUNEPS data by ID
	if err := db.DB.First(&tunepsData, "id = ?", id).Error; err != nil {
		// Check if the record is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	// Return JSON response with ID included
	ctx.JSON(http.StatusOK, gin.H{"id": id, "data": tunepsData})
}

// AOPR
// GetTunepsAOPR retrieves TUNEPS data from the database.
func (db Database) GetTunepsAOPR(ctx *gin.Context) {
	var tunepsData []domains.TunepsAOPRList

	if err := db.DB.Find(&tunepsData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	ctx.JSON(http.StatusOK, tunepsData)
}

// GetTunepsAOPRByID retrieves TUNEPS data from the database by ID.
func (db Database) GetTunepsAOPRByID(ctx *gin.Context) {
	// Get the ID from the request parameters
	id := ctx.Param("ID") // Remplacez "id" par "ID"

	// Check if the ID is empty
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var tunepsData domains.TunepsAOPRList

	// Fetch TUNEPS data by ID
	if err := db.DB.First(&tunepsData, "id = ?", id).Error; err != nil {
		// Check if the record is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	// Return JSON response with ID included
	ctx.JSON(http.StatusOK, gin.H{"id": id, "data": tunepsData})
}

//************* BOUTON List soumissionaire(s) retenu(s)******************

// AODNBB
// GetTunepsAODNB retrieves TUNEPS data from the database.
func (db Database) GetTunepsAODNB(ctx *gin.Context) {
	var tunepsData []domains.TunepsAODNBList

	if err := db.DB.Find(&tunepsData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	ctx.JSON(http.StatusOK, tunepsData)
}

// GetTunepsAODNBByID retrieves TUNEPS data from the database by ID.
func (db Database) GetTunepsAODNBByID(ctx *gin.Context) {
	// Get the ID from the request parameters
	id := ctx.Param("ID") // Remplacez "id" par "ID"

	// Check if the ID is empty
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var tunepsData domains.TunepsAODNBList

	// Fetch TUNEPS data by ID
	if err := db.DB.First(&tunepsData, "id = ?", id).Error; err != nil {
		// Check if the record is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	// Return JSON response with ID included
	ctx.JSON(http.StatusOK, gin.H{"id": id, "data": tunepsData})
}

// ATTM
// GetTunepsAOATTM retrieves TUNEPS data from the database.
func (db Database) GetTunepsAOATTM(ctx *gin.Context) {
	var tunepsData []domains.TunepsAOATTMList

	if err := db.DB.Find(&tunepsData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	ctx.JSON(http.StatusOK, tunepsData)
}

// GetTunepsAOATTMByID retrieves TUNEPS data from the database by ID.
func (db Database) GetTunepsAOATTMByID(ctx *gin.Context) {
	// Get the ID from the request parameters
	id := ctx.Param("ID") // Remplacez "id" par "ID"

	// Check if the ID is empty
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var tunepsData domains.TunepsAOATTMList

	// Fetch TUNEPS data by ID
	if err := db.DB.First(&tunepsData, "id = ?", id).Error; err != nil {
		// Check if the record is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	// Return JSON response with ID included
	ctx.JSON(http.StatusOK, gin.H{"id": id, "data": tunepsData})
}

// ************************bOUTON RESULTAT*************************************

// RDNB
// GetTunepsAORNDB retrieves TUNEPS data from the database.
func (db Database) GetTunepsAORDNB(ctx *gin.Context) {
	var tunepsData []domains.TunepsAORDNBList

	if err := db.DB.Find(&tunepsData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	ctx.JSON(http.StatusOK, tunepsData)
}

// GetTunepsAORDNBByID retrieves TUNEPS data from the database by ID.
func (db Database) GetTunepsAORDNBByID(ctx *gin.Context) {
	// Get the ID from the request parameters
	id := ctx.Param("ID") // Remplacez "id" par "ID"

	// Check if the ID is empty
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var tunepsData domains.TunepsAORDNBList

	// Fetch TUNEPS data by ID
	if err := db.DB.First(&tunepsData, "id = ?", id).Error; err != nil {
		// Check if the record is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	// Return JSON response with ID included
	ctx.JSON(http.StatusOK, gin.H{"id": id, "data": tunepsData})
}

// Lot-Resultat
// GetTunepsAORNDB retrieves TUNEPS data from the database.
func (db Database) GetTunepsAORLot(ctx *gin.Context) {
	var tunepsData []domains.TunepsAORLotList

	if err := db.DB.Find(&tunepsData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	ctx.JSON(http.StatusOK, tunepsData)
}

// GetTunepsAORLotByID retrieves TUNEPS data from the database by ID.
func (db Database) GetTunepsAORLotByID(ctx *gin.Context) {
	// Get the ID from the request parameters
	id := ctx.Param("ID") // Remplacez "id" par "ID"

	// Check if the ID is empty
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var tunepsData domains.TunepsAORLotList

	// Fetch TUNEPS data by ID
	if err := db.DB.First(&tunepsData, "id = ?", id).Error; err != nil {
		// Check if the record is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	// Return JSON response with ID included
	ctx.JSON(http.StatusOK, gin.H{"id": id, "data": tunepsData})
}

//Soumissionaires

// GetTunepsAORS retrieves TUNEPS data from the database.
func (db Database) GetTunepsAORS(ctx *gin.Context) {
	var tunepsData []domains.TunepsAORSList

	if err := db.DB.Find(&tunepsData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	ctx.JSON(http.StatusOK, tunepsData)
}

// GetTunepsAORSByID retrieves TUNEPS data from the database by ID.
func (db Database) GetTunepsAORSByID(ctx *gin.Context) {
	// Get the ID from the request parameters
	id := ctx.Param("ID") // Remplacez "id" par "ID"

	// Check if the ID is empty
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var tunepsData domains.TunepsAORSList

	// Fetch TUNEPS data by ID
	if err := db.DB.First(&tunepsData, "id = ?", id).Error; err != nil {
		// Check if the record is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	// Return JSON response with ID included
	ctx.JSON(http.StatusOK, gin.H{"id": id, "data": tunepsData})
}

/// ************************ Shopping Mall *******************************************

//SHOPPING MALL

// GetTunepsSM retrieves TUNEPS data from the database.
func (db Database) GetTunepsSM(ctx *gin.Context) {
	var tunepsData []domains.TunepsSMList

	if err := db.DB.Find(&tunepsData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	ctx.JSON(http.StatusOK, tunepsData)
}

// GetTunepsAORSByID retrieves TUNEPS data from the database by ID.
func (db Database) GetTunepsSMByID(ctx *gin.Context) {
	// Get the ID from the request parameters
	id := ctx.Param("ID") // Remplacez "id" par "ID"

	// Check if the ID is empty
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var tunepsData domains.TunepsSMList

	// Fetch TUNEPS data by ID
	if err := db.DB.First(&tunepsData, "id = ?", id).Error; err != nil {
		// Check if the record is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	// Return JSON response with ID included
	ctx.JSON(http.StatusOK, gin.H{"id": id, "data": tunepsData})
}

///SM INFO

// GetTunepsSMInfo retrieves TUNEPS data from the database.
func (db Database) GetTunepsSMInfo(ctx *gin.Context) {
	var tunepsData []domains.TunepsSMInfoList

	if err := db.DB.Find(&tunepsData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	ctx.JSON(http.StatusOK, tunepsData)
}

// GetTunepsSMInfoByID retrieves TUNEPS data from the database by ID.
func (db Database) GetTunepsSMInfoByID(ctx *gin.Context) {
	// Get the ID from the request parameters
	id := ctx.Param("ID") // Remplacez "id" par "ID"

	// Check if the ID is empty
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var tunepsData domains.TunepsSMInfoList

	// Fetch TUNEPS data by ID
	if err := db.DB.First(&tunepsData, "id = ?", id).Error; err != nil {
		// Check if the record is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	// Return JSON response with ID included
	ctx.JSON(http.StatusOK, gin.H{"id": id, "data": tunepsData})
}

///SM lot

// GetTunepsSMLot retrieves TUNEPS data from the database.
func (db Database) GetTunepsSMLot(ctx *gin.Context) {
	var tunepsData []domains.TunepsSMLotList

	if err := db.DB.Find(&tunepsData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	ctx.JSON(http.StatusOK, tunepsData)
}

// GetTunepsSMLotByID retrieves TUNEPS data from the database by ID.
func (db Database) GetTunepsSMLotByID(ctx *gin.Context) {
	// Get the ID from the request parameters
	id := ctx.Param("ID") // Remplacez "id" par "ID"

	// Check if the ID is empty
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var tunepsData domains.TunepsSMLotList

	// Fetch TUNEPS data by ID
	if err := db.DB.First(&tunepsData, "id = ?", id).Error; err != nil {
		// Check if the record is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	// Return JSON response with ID included
	ctx.JSON(http.StatusOK, gin.H{"id": id, "data": tunepsData})
}

//SM DOC

// GetTunepsSMDoc retrieves TUNEPS data from the database.
func (db Database) GetTunepsSMDoc(ctx *gin.Context) {
	var tunepsData []domains.TunepsSMDocList

	if err := db.DB.Find(&tunepsData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	ctx.JSON(http.StatusOK, tunepsData)
}

// GetTunepsSMDocByID retrieves TUNEPS data from the database by ID.
func (db Database) GetTunepsSMDocByID(ctx *gin.Context) {
	// Get the ID from the request parameters
	id := ctx.Param("ID") // Remplacez "id" par "ID"

	// Check if the ID is empty
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var tunepsData domains.TunepsSMDocList

	// Fetch TUNEPS data by ID
	if err := db.DB.First(&tunepsData, "id = ?", id).Error; err != nil {
		// Check if the record is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch TUNEPS data"})
		return
	}

	// Return JSON response with ID included
	ctx.JSON(http.StatusOK, gin.H{"id": id, "data": tunepsData})
}
