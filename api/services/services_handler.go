package services

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

// CreateService 	Handles the creation of a new service.
// @Summary        	Create service
// @Description    	Create a new service.
// @Tags			Services
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			companyID		path			string				true		"Company ID"
// @Param			request			body			services.ServicesIn	true	"Service query params"
// @Success			201				{object}		utils.ApiResponses
// @Failure			400				{object}		utils.ApiResponses	"Invalid request"
// @Failure			401				{object}		utils.ApiResponses	"Unauthorized"
// @Failure			403				{object}		utils.ApiResponses	"Forbidden"
// @Failure			500				{object}		utils.ApiResponses	"Internal Server Error"
// @Router			/services/{companyID}	[post]
// CreateService handles the creation of a new service.
func (db Database) CreateService(ctx *gin.Context) {
	// Extract JWT values from the context
	session := utils.ExtractJWTValues(ctx)

	// Parse and validate the company ID from the request parameter
	companyID, err := uuid.Parse(ctx.Param("companyID"))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid UUID format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Check if the employee belongs to the specified company
	if err := domains.CheckEmployeeBelonging(db.DB, companyID, session.UserID, session.CompanyID); err != nil {
		logrus.Error("Error verifying employee belonging. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Parse the incoming JSON request into a ServiceIn struct
	serviceInput := new(ServicesIn)
	if err := ctx.ShouldBindJSON(serviceInput); err != nil {
		logrus.Error("Error mapping request from frontend. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Create a new service in the database
	dbService := &domains.Services{
		ID:          uuid.New(),
		Name:        serviceInput.Name,
		Description: serviceInput.Description,
		CompanyID:   companyID,
	}
	if err := domains.Create(db.DB, dbService); err != nil {
		logrus.Error("Error saving data to the database. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.UNKNOWN_ERROR, utils.Null())
		return
	}

	// Create a new keyword in the database
	dbKeyword := &domains.Keyword{
		ID:   uuid.New(),
		Name: serviceInput.KeywordName,
	}
	if err := domains.Create(db.DB, dbKeyword); err != nil {
		logrus.Error("Error creating a new keyword in the database. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.UNKNOWN_ERROR, utils.Null())
		return
	}

	// Répondre avec succès
	utils.BuildResponse(ctx, http.StatusCreated, constants.CREATED, utils.Null())
}

// ReadServices 	Handles the retrieval of all services.
// @Summary        	Get services
// @Description    	Get all services.
// @Tags			Services
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			page			query		int			false		"Page"
// @Param			limit			query		int			false		"Limit"
// @Param			companyID		path		string		true		"Company ID"
// @Success			200				{object}	services.ServicesPagination
// @Failure			400				{object}	utils.ApiResponses		"Invalid request"
// @Failure			401				{object}	utils.ApiResponses		"Unauthorized"
// @Failure			403				{object}	utils.ApiResponses		"Forbidden"
// @Failure			500				{object}	utils.ApiResponses		"Internal Server Error"
// @Router			/services/{companyID}	[get]
func (db Database) ReadServices(ctx *gin.Context) {
	// Extract JWT values from the context
	session := utils.ExtractJWTValues(ctx)

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
	companyID, err := uuid.Parse(ctx.Param("companyID"))
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

	// Check if the employee belongs to the specified company
	if err := domains.CheckEmployeeBelonging(db.DB, companyID, session.UserID, session.CompanyID); err != nil {
		logrus.Error("Error verifying employee belonging. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Retrieve all service data from the database
	services, err := ReadAllPagination(db.DB, []domains.Services{}, session.CompanyID, limit, offset)
	if err != nil {
		logrus.Error("Error occurred while finding all service data. Error: ", err)
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.UNKNOWN_ERROR, utils.Null())
		return
	}

	// Retrieve total count
	count, err := domains.ReadTotalCount(db.DB, &domains.Services{}, "company_id", session.CompanyID)
	if err != nil {
		logrus.Error("Error occurred while finding total count. Error: ", err)
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.UNKNOWN_ERROR, utils.Null())
		return
	}

	// Generate a service structure as a response
	response := ServicesPagination{}
	dataTableService := []ServicesTable{}
	for _, service := range services {
		dataTableService = append(dataTableService, ServicesTable{
			ID:          service.ID,
			Name:        service.Name,
			Description: service.Description,
			CreatedAt:   service.CreatedAt,
		})
	}
	response.Items = dataTableService
	response.Page = uint(page)
	response.Limit = uint(limit)
	response.TotalCount = count

	// Respond with success
	utils.BuildResponse(ctx, http.StatusOK, constants.SUCCESS, response)
}

// ReadServicesList Handles the retrieval the list of all services.
// @Summary        	Get list of  services
// @Description    	Get list of all services.
// @Tags			Services
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			companyID			path			string			true	"Company ID"
// @Success			200					{array}			services.ServicesList
// @Failure			400					{object}		utils.ApiResponses		"Invalid request"
// @Failure			401					{object}		utils.ApiResponses		"Unauthorized"
// @Failure			403					{object}		utils.ApiResponses		"Forbidden"
// @Failure			500					{object}		utils.ApiResponses		"Internal Server Error"
// @Router			/services/{companyID}/list	[get]
func (db Database) ReadServicesList(ctx *gin.Context) {
	// Extract JWT values from the context
	session := utils.ExtractJWTValues(ctx)

	// Parse and validate the company ID from the request parameter
	companyID, err := uuid.Parse(ctx.Param("companyID"))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid UUID format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Check if the employee belongs to the specified company
	if err := domains.CheckEmployeeBelonging(db.DB, companyID, session.UserID, session.CompanyID); err != nil {
		logrus.Error("Error verifying employee belonging. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Retrieve all service data from the database
	services, err := ReadAllList(db.DB, []domains.Services{}, session.CompanyID)
	if err != nil {
		logrus.Error("Error occurred while finding all service data. Error: ", err)
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.UNKNOWN_ERROR, utils.Null())
		return
	}

	// Generate a service structure as a response
	servicesList := []ServicesList{}
	for _, service := range services {
		servicesList = append(servicesList, ServicesList{
			ID:   service.ID,
			Name: service.Name,
			// Ajoutez d'autres champs de service ici...
		})
	}

	// Respond with success
	utils.BuildResponse(ctx, http.StatusOK, constants.SUCCESS, servicesList)
}

// ReadServicesCount Handles the retrieval the number of all services.
// @Summary        	Get number of  services
// @Description    	Get number of all services.
// @Tags			Services
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			companyID				path			string		true	"Company ID"
// @Success			200						{object}		services.ServicesCount
// @Failure			400						{object}		utils.ApiResponses	"Invalid request"
// @Failure			401						{object}		utils.ApiResponses	"Unauthorized"
// @Failure			403						{object}		utils.ApiResponses	"Forbidden"
// @Failure			500						{object}		utils.ApiResponses	"Internal Server Error"
// @Router			/services/{companyID}/count	[get]
func (db Database) ReadServicesCount(ctx *gin.Context) {

	// Extract JWT values from the context
	session := utils.ExtractJWTValues(ctx)

	// Parse and validate the company ID from the request parameter
	companyID, err := uuid.Parse(ctx.Param("companyID"))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid UUID format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Check if the employee belongs to the specified company
	if err := domains.CheckEmployeeBelonging(db.DB, companyID, session.UserID, session.CompanyID); err != nil {
		logrus.Error("Error verifying employee belonging. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Retrieve all user data from the database
	services, err := domains.ReadTotalCount(db.DB, &[]domains.Services{}, "company_id", session.CompanyID)
	if err != nil {
		logrus.Error("Error occurred while finding all services data. Error: ", err)
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.UNKNOWN_ERROR, utils.Null())
		return
	}

	// Generate a user structure as a response
	servicesCount := ServicesCount{
		Count: services,
	}

	// Respond with success
	utils.BuildResponse(ctx, http.StatusOK, constants.SUCCESS, servicesCount)
}

// ReadService 		Handles the retrieval of one service.
// @Summary        	Get service
// @Description    	Get one service.
// @Tags			Services
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			companyID			path			string			true	"Company ID"
// @Param			ID					path			string			true	"Service ID"
// @Success			200					{object}		services.ServicesDetails
// @Failure			400					{object}		utils.ApiResponses		"Invalid request"
// @Failure			401					{object}		utils.ApiResponses		"Unauthorized"
// @Failure			403					{object}		utils.ApiResponses		"Forbidden"
// @Failure			500					{object}		utils.ApiResponses		"Internal Server Error"
// @Router			/services/{companyID}/{ID}	[get]
func (db Database) ReadService(ctx *gin.Context) {

	// Extract JWT values from the context
	session := utils.ExtractJWTValues(ctx)

	// Parse and validate the company ID from the request parameter
	companyID, err := uuid.Parse(ctx.Param("companyID"))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid UUID format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Parse and validate the service ID from the request parameter
	objectID, err := uuid.Parse(ctx.Param("ID"))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid UUID format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Check if the employee belongs to the specified company
	if err := domains.CheckEmployeeBelonging(db.DB, companyID, session.UserID, session.CompanyID); err != nil {
		logrus.Error("Error verifying employee belonging. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Retrieve user data from the database
	service, err := ReadByID(db.DB, domains.Services{}, objectID)
	if err != nil {
		logrus.Error("Error retrieving service data from the database. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.DATA_NOT_FOUND, utils.Null())
		return
	}

	// Generate a user structure as a response
	details := ServicesDetails{
		ID:          service.ID,
		Name:        service.Name,
		Description: service.Description,
		CreatedAt:   service.CreatedAt,
	}

	// Respond with success
	utils.BuildResponse(ctx, http.StatusOK, constants.SUCCESS, details)
}

// UpdateService 	Handles the update of a service.
// @Summary        	Update service
// @Description    	Update service.
// @Tags			Services
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			companyID			path			string				true	"Company ID"
// @Param			ID					path			string				true	"Service ID"
// @Param			request				body			services.ServicesIn		true	"Service query params"
// @Success			200					{object}		utils.ApiResponses
// @Failure			400					{object}		utils.ApiResponses			"Invalid request"
// @Failure			401					{object}		utils.ApiResponses			"Unauthorized"
// @Failure			403					{object}		utils.ApiResponses			"Forbidden"
// @Failure			500					{object}		utils.ApiResponses			"Internal Server Error"
// @Router			/services/{companyID}/{ID}	[put]
func (db Database) UpdateService(ctx *gin.Context) {

	// Extract JWT values from the context
	session := utils.ExtractJWTValues(ctx)

	// Parse and validate the company ID from the request parameter
	companyID, err := uuid.Parse(ctx.Param("companyID"))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid UUID format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Parse and validate the user ID from the request parameter
	objectID, err := uuid.Parse(ctx.Param("ID"))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid UUID format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Check if the employee belongs to the specified company
	if err := domains.CheckEmployeeBelonging(db.DB, companyID, session.UserID, session.CompanyID); err != nil {
		logrus.Error("Error verifying employee belonging. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Parse the incoming JSON request into a UserIn struct
	service := new(ServicesIn)
	if err := ctx.ShouldBindJSON(service); err != nil {
		logrus.Error("Error mapping request from frontend. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Check if the service with the specified ID exists
	if err = domains.CheckByID(db.DB, &domains.Services{}, objectID); err != nil {
		logrus.Error("Error checking if the user with the specified ID exists. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Update the user data in the database
	dbService := &domains.Services{
		Name:        service.Name,
		Description: service.Description,
	}
	if err = domains.Update(db.DB, dbService, objectID); err != nil {
		logrus.Error("Error updating service data in the database. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.UNKNOWN_ERROR, utils.Null())
		return
	}

	// Respond with success
	utils.BuildResponse(ctx, http.StatusOK, constants.SUCCESS, utils.Null())
}

// DeleteService	Handles the deletion of a service.
// @Summary        	Delete service
// @Description    	Delete one service.
// @Tags			Services
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			companyID			path			string			true	"Company ID"
// @Param			ID					path			string			true	"service ID"
// @Success			200					{object}		utils.ApiResponses
// @Failure			400					{object}		utils.ApiResponses		"Invalid request"
// @Failure			401					{object}		utils.ApiResponses		"Unauthorized"
// @Failure			403					{object}		utils.ApiResponses		"Forbidden"
// @Failure			500					{object}		utils.ApiResponses		"Internal Server Error"
// @Router			/services/{companyID}/{ID}	[delete]
func (db Database) DeleteService(ctx *gin.Context) {

	// Extract JWT values from the context
	session := utils.ExtractJWTValues(ctx)

	// Parse and validate the company ID from the request parameter
	companyID, err := uuid.Parse(ctx.Param("companyID"))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid UUID format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Parse and validate the user ID from the request parameter
	objectID, err := uuid.Parse(ctx.Param("ID"))
	if err != nil {
		logrus.Error("Error mapping request from frontend. Invalid UUID format. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Check if the employee belongs to the specified company
	if err := domains.CheckEmployeeBelonging(db.DB, companyID, session.UserID, session.CompanyID); err != nil {
		logrus.Error("Error verifying employee belonging. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.INVALID_REQUEST, utils.Null())
		return
	}

	// Check if the service with the specified ID exists
	if err := domains.CheckByID(db.DB, &domains.Services{}, objectID); err != nil {
		logrus.Error("Error checking if the service with the specified ID exists. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusNotFound, constants.DATA_NOT_FOUND, utils.Null())
		return
	}

	// Delete the service data from the database
	if err := domains.Delete(db.DB, &domains.Services{}, objectID); err != nil {
		logrus.Error("Error deleting service data from the database. Error: ", err.Error())
		utils.BuildErrorResponse(ctx, http.StatusBadRequest, constants.UNKNOWN_ERROR, utils.Null())
		return
	}

	// Respond with success
	utils.BuildResponse(ctx, http.StatusOK, constants.SUCCESS, utils.Null())
}
