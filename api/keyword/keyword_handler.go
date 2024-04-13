package keyword

import (
	"labs/constants"
	"labs/domains"
	"labs/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// CreateKeyword		Handles the creation of a new service.
// @Summary        	Create keyword
// @Description    	Create a new keyword.
// @Tags			keywords
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			serviceID		path			string				true		"Service ID"
// @Param			request			body			keywords.keywordIn		true		"Keyword query params"
// @Success			201				{object}		utils.ApiResponses
// @Failure			400				{object}		utils.ApiResponses	"Invalid request"
// @Failure			401				{object}		utils.ApiResponses	"Unauthorized"
// @Failure			403				{object}		utils.ApiResponses	"Forbidden"
// @Failure			500				{object}		utils.ApiResponses	"Internal Server Error"
// @Router			/keywords/{serviceID}	[post]
// CreateService handles the creation of a new service.
func (db Database) CreateKeyword(ctx *gin.Context) {
	// Parse and validate the service ID from the request parameter
	serviceID, err := uuid.Parse(ctx.Param("serviceID"))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid UUID format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Parse the incoming JSON request into a KeywordIn struct
	keywordInput := new(KeywordIn)
	if err := ctx.ShouldBindJSON(keywordInput); err != nil {
		logrus.Error("Error mapping request from frontend. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Create a new keyword in the database
	dbKeyword := &domains.Keyword{
		ID:        uuid.New(),
		Name:      keywordInput.Name,
		ServiceID: serviceID,
	}
	if err := domains.Create(db.DB, dbKeyword); err != nil {
		logrus.Error("Error saving data to the database. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.UNKNOWN_ERROR, utils.Null())
		return
	}

	// Respond with success
	utils.BuildResponse(ctx, http.StatusCreated, constants.CREATED, utils.Null())
}

// ReadKeywords handles the retrieval of all keywords.
// @Summary        	Get keywords
// @Description    	Get all keywords.
// @Tags			Keywords
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			page			query		int			false		"Page"
// @Param			limit			query		int			false		"Limit"
// @Param			serviceID		path		string		true		"Service ID"
// @Success			200				{object}	keywords.KeywordsPagination
// @Failure			400				{object}	utils.ApiResponses		"Invalid request"
// @Failure			401				{object}	utils.ApiResponses		"Unauthorized"
// @Failure			403				{object}	utils.ApiResponses		"Forbidden"
// @Failure			500				{object}	utils.ApiResponses		"Internal Server Error"
// @Router			/keywords/{serviceID}	[get]
func (db Database) ReadKeywords(ctx *gin.Context) {
	// Parse and validate the page from the request parameter
	page, err := strconv.Atoi(ctx.DefaultQuery("page", strconv.Itoa(constants.DEFAULT_PAGE_PAGINATION)))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid INT format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Parse and validate the limit from the request parameter
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", strconv.Itoa(constants.DEFAULT_LIMIT_PAGINATION)))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid INT format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Parse and validate the company ID from the request parameter
	serviceID, err := uuid.Parse(ctx.Param("serviceID"))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid UUID format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Check if the user's value is among the allowed choices
	validChoices := utils.ResponseLimitPagination()
	isValidChoice := false
	for _, choice := range validChoices {
		if uint(limit) == choice {
			isValidChoice = true
			break
		}
	}

	// If the value is invalid, set it to default DEFAULT_LIMIT_PAGINATION
	if !isValidChoice {
		limit = constants.DEFAULT_LIMIT_PAGINATION
	}

	// Generate offset
	offset := (page - 1) * limit

	// Retrieve all service data from the database
	keywords, err := ReadAllPagination(db.DB, []domains.Keyword{}, serviceID, limit, offset)
	if err != nil {
		logrus.Error("Error occurred while finding all keyword data. Error: ", err)
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.UNKNOWN_ERROR, utils.Null())
		return
	}

	// Retrieve total count
	count, err := domains.ReadTotalCount(db.DB, &domains.Keyword{}, "service_id", serviceID)
	if err != nil {
		logrus.Error("Error occurred while finding total count. Error: ", err)
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.UNKNOWN_ERROR, utils.Null())
		return
	}

	// Generate a keyword structure as a response
	response := KeywordsPagination{}
	dataTableKeyword := []KeywordsTable{}
	for _, keyword := range keywords {
		dataTableKeyword = append(dataTableKeyword, KeywordsTable{
			ID:        keyword.ID,
			Name:      keyword.Name,
			CreatedAt: keyword.CreatedAt,
		})
	}
	response.Items = dataTableKeyword
	response.Page = uint(page)
	response.Limit = uint(limit)
	response.TotalCount = count

	// Respond with success
	utils.BuildResponse(ctx, http.StatusOK, constants.SUCCESS, response)
}

// ReadKeywordsList 	Handles the retrieval the list of all keywords.
// @Summary        	Get list of  keywords
// @Description    	Get list of all keywords.
// @Tags			keywords
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			serviceID			path			string			true	"Service ID"
// @Success			200					{array}			keywords.KeywordsList
// @Failure			400					{object}		utils.ApiResponses		"Invalid request"
// @Failure			401					{object}		utils.ApiResponses		"Unauthorized"
// @Failure			403					{object}		utils.ApiResponses		"Forbidden"
// @Failure			500					{object}		utils.ApiResponses		"Internal Server Error"
// @Router			/keywords/{serviceID}/list	[get]
func (db Database) ReadKeywordsList(ctx *gin.Context) {

	// Parse and validate the servic ID from the request parameter
	serviceID, err := uuid.Parse(ctx.Param("serviceID"))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid UUID format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Retrieve all service data from the database
	keywords, err := ReadAllList(db.DB, []domains.Keyword{}, serviceID)
	if err != nil {
		logrus.Error("Error occurred while finding all service data. Error: ", err)
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.UNKNOWN_ERROR, utils.Null())
		return
	}

	// Generate a keyword structure as a response
	keywordsList := []KeywordsList{}
	for _, service := range keywords {
		keywordsList = append(keywordsList, KeywordsList{
			ID:   service.ID,
			Name: service.Name,
		})
	}

	// Respond with success
	utils.BuildResponse(ctx, http.StatusOK, constants.SUCCESS, keywordsList)
}

// ReadKeywordsCount handles the retrieval of the number of all keywords.
// @Summary        	Get number of keywords
// @Description    	Get number of all keywords.
// @Tags			Keywords
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			serviceID				path			string		true	"Service ID"
// @Success			200						{object}		services.ServicesCount
// @Failure			400						{object}		utils.ApiResponses	"Invalid request"
// @Failure			401						{object}		utils.ApiResponses	"Unauthorized"
// @Failure			403						{object}		utils.ApiResponses	"Forbidden"
// @Failure			500						{object}		utils.ApiResponses	"Internal Server Error"
// @Router			/keywords/{serviceID}/count	[get]
func (db Database) ReadKeywordsCount(ctx *gin.Context) {
	// Parse and validate the service ID from the request parameter
	serviceID, err := uuid.Parse(ctx.Param("serviceID"))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid UUID format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Retrieve the number of keywords from the database
	keywordsCount, err := domains.ReadTotalCount(db.DB, &[]domains.Keyword{}, "service_id", serviceID)
	if err != nil {
		logrus.Error("Error occurred while finding the number of keywords. Error: ", err)
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.UNKNOWN_ERROR, utils.Null())
		return
	}

	// Generate a response structure
	keywordsCountResponse := KeywordsCount{
		Count: keywordsCount,
	}

	// Respond with success
	utils.BuildResponse(ctx, http.StatusOK, constants.SUCCESS, keywordsCountResponse)
}

// ReadKeyword handles the retrieval of one keyword.
// @Summary        	Get keyword
// @Description    	Get one keyword.
// @Tags			Keywords
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			serviceID			path			string			true	"Service ID"
// @Param			ID					path			string			true	"Keyword ID"
// @Success			200					{object}		services.ServicesDetails
// @Failure			400					{object}		utils.ApiResponses		"Invalid request"
// @Failure			401					{object}		utils.ApiResponses		"Unauthorized"
// @Failure			403					{object}		utils.ApiResponses		"Forbidden"
// @Failure			500					{object}		utils.ApiResponses		"Internal Server Error"
// @Router			/keywords/{serviceID}/{ID}	[get]
func (db Database) ReadKeyword(ctx *gin.Context) {

	// Parse and validate the keyword ID from the request parameter
	keywordID, err := uuid.Parse(ctx.Param("ID"))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid UUID format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Retrieve keyword data from the database
	keyword, err := ReadByID(db.DB, domains.Keyword{}, keywordID)
	if err != nil {
		logrus.Error("Error retrieving keyword data from the database. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.DATA_NOT_FOUND, utils.Null())
		return
	}

	// Generate a keyword structure as a response
	details := KeywordsDetails{
		ID:        keyword.ID,
		Name:      keyword.Name,
		CreatedAt: keyword.CreatedAt,
	}

	// Respond with success
	utils.BuildResponse(ctx, http.StatusOK, constants.SUCCESS, details)
}

// UpdateKeyword handles the update of a keyword.
// @Summary        	Update keyword
// @Description    	Update keyword.
// @Tags            Keywords
// @Accept          json
// @Produce         json
// @Security        ApiKeyAuth
// @Param           serviceId       path            string              true        "Service ID"
// @Param           ID              path            string              true        "Keyword ID"
// @Param           request         body            keywords.KeywordIn true        "Keyword query params"
// @Success         200             {object}        utils.ApiResponses
// @Failure         400             {object}        utils.ApiResponses        "Invalid request"
// @Failure         401             {object}        utils.ApiResponses        "Unauthorized"
// @Failure         403             {object}        utils.ApiResponses        "Forbidden"
// @Failure         500             {object}        utils.ApiResponses        "Internal Server Error"
// @Router          /keywords/{serviceId}/{ID} [put]
func (db Database) UpdateKeyword(ctx *gin.Context) {
	// Parse and validate the service ID from the request parameter
	serviceID, err := uuid.Parse(ctx.Param("serviceId"))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid UUID format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Parse and validate the keyword ID from the request parameter
	keywordID, err := uuid.Parse(ctx.Param("ID"))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid UUID format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Parse the incoming JSON request into a KeywordIn struct
	keywordInput := new(KeywordIn)
	if err := ctx.ShouldBindJSON(keywordInput); err != nil {
		logrus.Error("Error mapping request from frontend. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Check if the keyword with the specified ID exists
	if err = domains.CheckByID(db.DB, &domains.Keyword{}, keywordID); err != nil {
		logrus.Error("Error checking if the keyword with the specified ID exists. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Update the keyword data in the database
	dbKeyword := &domains.Keyword{
		Name:      keywordInput.Name,
		ServiceID: serviceID,
	}
	if err = domains.Update(db.DB, dbKeyword, keywordID); err != nil {
		logrus.Error("Error updating keyword data in the database. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.UNKNOWN_ERROR, utils.Null())
		return
	}

	// Respond with success
	utils.BuildResponse(ctx, http.StatusOK, constants.SUCCESS, utils.Null())
}

// DeleteKeyword handles the deletion of a keyword.
// @Summary        	Delete keyword
// @Description    	Delete one keyword.
// @Tags            Keywords
// @Produce         json
// @Security        ApiKeyAuth
// @Param           serviceId       path            string              true        "Service ID"
// @Param           ID              path            string              true        "Keyword ID"
// @Success         200             {object}        utils.ApiResponses
// @Failure         400             {object}        utils.ApiResponses        "Invalid request"
// @Failure         401             {object}        utils.ApiResponses        "Unauthorized"
// @Failure         403             {object}        utils.ApiResponses        "Forbidden"
// @Failure         500             {object}        utils.ApiResponses        "Internal Server Error"
// @Router          /keywords/{serviceId}/{ID} [delete]

func (db Database) DeleteKeyword(ctx *gin.Context) {

	// Parse and validate the keyword ID from the request parameter
	keywordID, err := uuid.Parse(ctx.Param("ID"))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid UUID format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Check if the keyword with the specified ID exists
	if err := domains.CheckByID(db.DB, &domains.Keyword{}, keywordID); err != nil {
		logrus.Error("Error checking if the keyword with the specified ID exists. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusNotFound, constants.DATA_NOT_FOUND, utils.Null())
		return
	}

	// Delete the keyword data from the database
	if err := domains.Delete(db.DB, &domains.Keyword{}, keywordID); err != nil {
		logrus.Error("Error deleting keyword data from the database. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.UNKNOWN_ERROR, utils.Null())
		return
	}

	// Respond with success
	utils.BuildResponse(ctx, http.StatusOK, constants.SUCCESS, utils.Null())
}
