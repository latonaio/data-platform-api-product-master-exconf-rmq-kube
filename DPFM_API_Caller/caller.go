package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-product-master-exconf-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-master-exconf-rmq-kube/DPFM_API_Output_Formatter"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
)

type ExistenceConf struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewExistenceConf(ctx context.Context, db *database.Mysql, l *logger.Logger) *ExistenceConf {
	return &ExistenceConf{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (e *ExistenceConf) Conf(
	accepter []string,
	input *dpfm_api_input_reader.SDC,
) dpfm_api_output_formatter.ProductMaster {
	var general *dpfm_api_output_formatter.General
	var businessPartner *dpfm_api_output_formatter.BusinessPartner
	var bPPlant *dpfm_api_output_formatter.BPPlant
	var storageLocation *dpfm_api_output_formatter.StorageLocation
	var mRPArea *dpfm_api_output_formatter.MRPArea
	var workScheduling *dpfm_api_output_formatter.WorkScheduling
	var accounting *dpfm_api_output_formatter.Accounting
	var productDescription *dpfm_api_output_formatter.ProductDescription
	var productDescByBP *dpfm_api_output_formatter.ProductDescByBP
	var tax *dpfm_api_output_formatter.Tax
	var allergen *dpfm_api_output_formatter.Allergen
	var calories *dpfm_api_output_formatter.Calories
	var nutritionalInfo *dpfm_api_output_formatter.NutritionalInfo
	var quality *dpfm_api_output_formatter.Quality
	var storageBin *dpfm_api_output_formatter.StorageBin

	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				general = e.confProductMasterGeneral(input)
			}()
		case "BusinessPartner":
			func() {
				businessPartner = e.confProductMasterBusinessPartner(input)
			}()
		case "BPPlant":
			func() {
				bPPlant = e.confProductMasterBPPlant(input)
			}()
		case "StorageLocation":
			func() {
				storageLocation = e.confProductMasterStorageLocation(input)
			}()
		case "MRPArea":
			func() {
				mRPArea = e.confProductMasterMRPArea(input)
			}()
		case "WorkScheduling":
			func() {
				workScheduling = e.confProductMasterWorkScheduling(input)
			}()
		case "Accounting":
			func() {
				accounting = e.confProductMasterAccounting(input)
			}()
		case "ProductDescription":
			func() {
				productDescription = e.confProductMasterProductDescription(input)
			}()
		case "ProductDescByBP":
			func() {
				productDescByBP = e.confProductMasterProductDescByBP(input)
			}()
		case "Quality":
			func() {
				quality = e.confProductMasterQuality(input)
			}()
		case "Allergen":
			func() {
				allergen = e.confProductMasterAllergen(input)
			}()
		case "NutritionalInfo":
			func() {
				nutritionalInfo = e.confProductMasterNutritionalInfo(input)
			}()
		case "Calories":
			func() {
				calories = e.confProductMasterCalories(input)
			}()
		case "Tax":
			func() {
				tax = e.confProductMasterTax(input)
			}()
		case "StorageBin":
			func() {
				storageBin = e.confProductMasterStorageBin(input)
			}()
		}
	}

	res := dpfm_api_output_formatter.ProductMaster{
		General:            general,
		BusinessPartner:    businessPartner,
		BPPlant:            bPPlant,
		StorageLocation:    storageLocation,
		MRPArea:            mRPArea,
		WorkScheduling:     workScheduling,
		Accounting:         accounting,
		ProductDescription: productDescription,
		ProductDescByBP:    productDescByBP,
		Tax:                tax,
		Allergen:           allergen,
		Calories:           calories,
		NutritionalInfo:    nutritionalInfo,
		Quality:            quality,
		StorageBin:         storageBin,
	}

	return res
}
