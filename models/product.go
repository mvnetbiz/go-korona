// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Product product
//
// swagger:model Product
type Product struct {

	// indicates whether the object is active for use or not
	// Read Only: true
	Active *bool `json:"active,omitempty"`

	// alternative sector
	AlternativeSector *ModelReference `json:"alternativeSector,omitempty"`

	// assortment
	Assortment *ModelReference `json:"assortment,omitempty"`

	// codes
	Codes []*ProductCode `json:"codes"`

	// commodity group
	CommodityGroup *ModelReference `json:"commodityGroup,omitempty"`

	// conversion
	Conversion bool `json:"conversion,omitempty"`

	// costs
	Costs float64 `json:"costs,omitempty"`

	// deactivated
	Deactivated bool `json:"deactivated,omitempty"`

	// deposit
	Deposit bool `json:"deposit,omitempty"`

	// descriptions
	Descriptions []*ProductDescription `json:"descriptions"`

	// discountable
	Discountable bool `json:"discountable,omitempty"`

	// global object uuid (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)
	ID string `json:"id,omitempty"`

	// image
	Image *ModelReference `json:"image,omitempty"`

	// info texts
	InfoTexts []*ModelReference `json:"infoTexts"`

	// item sequence
	ItemSequence *ModelReference `json:"itemSequence,omitempty"`

	// last purchase price
	LastPurchasePrice float64 `json:"lastPurchasePrice,omitempty"`

	// listed
	Listed bool `json:"listed,omitempty"`

	// listed since
	// Format: date-time
	ListedSince strfmt.DateTime `json:"listedSince,omitempty"`

	// max price
	MaxPrice float64 `json:"maxPrice,omitempty"`

	// media urls
	// Unique: true
	MediaUrls []*MediaURL `json:"mediaUrls"`

	// min price
	MinPrice float64 `json:"minPrice,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// number of the object, like it is set in backoffice; will be removed when active=false
	Number string `json:"number,omitempty"`

	// packaging quantity
	PackagingQuantity float64 `json:"packagingQuantity,omitempty"`

	// packaging required
	PackagingRequired bool `json:"packagingRequired,omitempty"`

	// packaging unit
	// Enum: [CUBIC_INCH CUBIC_METER FLUID_OUNCE GALLON_FL GRAM INCH KILOGRAM LITER METER MILLILITER OUNCE POUND SQUARE_INCH SQUARE_METER]
	PackagingUnit string `json:"packagingUnit,omitempty"`

	// personalization required
	PersonalizationRequired bool `json:"personalizationRequired,omitempty"`

	// price changable
	PriceChangable bool `json:"priceChangable,omitempty"`

	// prices
	Prices []*ProductPrice `json:"prices"`

	// print tickets separately
	PrintTicketsSeparately bool `json:"printTicketsSeparately,omitempty"`

	// production type
	ProductionType *ModelReference `json:"productionType,omitempty"`

	// quantity denomination
	QuantityDenomination float64 `json:"quantityDenomination,omitempty"`

	// recommended retail price
	RecommendedRetailPrice float64 `json:"recommendedRetailPrice,omitempty"`

	// related products
	RelatedProducts *ModelReference `json:"relatedProducts,omitempty"`

	// the revision number of the object. revision numbers are unique per object-type. there is is no object of the same type with identical revision numbers.
	// Read Only: true
	Revision int64 `json:"revision,omitempty"`

	// sector
	Sector *ModelReference `json:"sector,omitempty"`

	// serial number required
	SerialNumberRequired bool `json:"serialNumberRequired,omitempty"`

	// stock return unsellable
	StockReturnUnsellable bool `json:"stockReturnUnsellable,omitempty"`

	// subproduct presentation
	// Enum: [DEFAULT HIDE_ALL HIDE_PRICES HIDE_QUANTITIES]
	SubproductPresentation string `json:"subproductPresentation,omitempty"`

	// subproducts
	Subproducts []*ProductSubproduct `json:"subproducts"`

	// supplier prices
	SupplierPrices []*ProductSupplierPrice `json:"supplierPrices"`

	// tags
	Tags []*ModelReference `json:"tags"`

	// ticket definition
	TicketDefinition *ModelReference `json:"ticketDefinition,omitempty"`

	// track inventory
	TrackInventory bool `json:"trackInventory,omitempty"`
}

// Validate validates this product
func (m *Product) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlternativeSector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssortment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommodityGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescriptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfoTexts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemSequence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateListedSince(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaUrls(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackagingUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedProducts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubproductPresentation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubproducts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupplierPrices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTicketDefinition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Product) validateAlternativeSector(formats strfmt.Registry) error {

	if swag.IsZero(m.AlternativeSector) { // not required
		return nil
	}

	if m.AlternativeSector != nil {
		if err := m.AlternativeSector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alternativeSector")
			}
			return err
		}
	}

	return nil
}

func (m *Product) validateAssortment(formats strfmt.Registry) error {

	if swag.IsZero(m.Assortment) { // not required
		return nil
	}

	if m.Assortment != nil {
		if err := m.Assortment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assortment")
			}
			return err
		}
	}

	return nil
}

func (m *Product) validateCodes(formats strfmt.Registry) error {

	if swag.IsZero(m.Codes) { // not required
		return nil
	}

	for i := 0; i < len(m.Codes); i++ {
		if swag.IsZero(m.Codes[i]) { // not required
			continue
		}

		if m.Codes[i] != nil {
			if err := m.Codes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("codes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Product) validateCommodityGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.CommodityGroup) { // not required
		return nil
	}

	if m.CommodityGroup != nil {
		if err := m.CommodityGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commodityGroup")
			}
			return err
		}
	}

	return nil
}

func (m *Product) validateDescriptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Descriptions) { // not required
		return nil
	}

	for i := 0; i < len(m.Descriptions); i++ {
		if swag.IsZero(m.Descriptions[i]) { // not required
			continue
		}

		if m.Descriptions[i] != nil {
			if err := m.Descriptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("descriptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Product) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *Product) validateInfoTexts(formats strfmt.Registry) error {

	if swag.IsZero(m.InfoTexts) { // not required
		return nil
	}

	for i := 0; i < len(m.InfoTexts); i++ {
		if swag.IsZero(m.InfoTexts[i]) { // not required
			continue
		}

		if m.InfoTexts[i] != nil {
			if err := m.InfoTexts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("infoTexts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Product) validateItemSequence(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemSequence) { // not required
		return nil
	}

	if m.ItemSequence != nil {
		if err := m.ItemSequence.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("itemSequence")
			}
			return err
		}
	}

	return nil
}

func (m *Product) validateListedSince(formats strfmt.Registry) error {

	if swag.IsZero(m.ListedSince) { // not required
		return nil
	}

	if err := validate.FormatOf("listedSince", "body", "date-time", m.ListedSince.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateMediaUrls(formats strfmt.Registry) error {

	if swag.IsZero(m.MediaUrls) { // not required
		return nil
	}

	if err := validate.UniqueItems("mediaUrls", "body", m.MediaUrls); err != nil {
		return err
	}

	for i := 0; i < len(m.MediaUrls); i++ {
		if swag.IsZero(m.MediaUrls[i]) { // not required
			continue
		}

		if m.MediaUrls[i] != nil {
			if err := m.MediaUrls[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mediaUrls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var productTypePackagingUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CUBIC_INCH","CUBIC_METER","FLUID_OUNCE","GALLON_FL","GRAM","INCH","KILOGRAM","LITER","METER","MILLILITER","OUNCE","POUND","SQUARE_INCH","SQUARE_METER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		productTypePackagingUnitPropEnum = append(productTypePackagingUnitPropEnum, v)
	}
}

const (

	// ProductPackagingUnitCUBICINCH captures enum value "CUBIC_INCH"
	ProductPackagingUnitCUBICINCH string = "CUBIC_INCH"

	// ProductPackagingUnitCUBICMETER captures enum value "CUBIC_METER"
	ProductPackagingUnitCUBICMETER string = "CUBIC_METER"

	// ProductPackagingUnitFLUIDOUNCE captures enum value "FLUID_OUNCE"
	ProductPackagingUnitFLUIDOUNCE string = "FLUID_OUNCE"

	// ProductPackagingUnitGALLONFL captures enum value "GALLON_FL"
	ProductPackagingUnitGALLONFL string = "GALLON_FL"

	// ProductPackagingUnitGRAM captures enum value "GRAM"
	ProductPackagingUnitGRAM string = "GRAM"

	// ProductPackagingUnitINCH captures enum value "INCH"
	ProductPackagingUnitINCH string = "INCH"

	// ProductPackagingUnitKILOGRAM captures enum value "KILOGRAM"
	ProductPackagingUnitKILOGRAM string = "KILOGRAM"

	// ProductPackagingUnitLITER captures enum value "LITER"
	ProductPackagingUnitLITER string = "LITER"

	// ProductPackagingUnitMETER captures enum value "METER"
	ProductPackagingUnitMETER string = "METER"

	// ProductPackagingUnitMILLILITER captures enum value "MILLILITER"
	ProductPackagingUnitMILLILITER string = "MILLILITER"

	// ProductPackagingUnitOUNCE captures enum value "OUNCE"
	ProductPackagingUnitOUNCE string = "OUNCE"

	// ProductPackagingUnitPOUND captures enum value "POUND"
	ProductPackagingUnitPOUND string = "POUND"

	// ProductPackagingUnitSQUAREINCH captures enum value "SQUARE_INCH"
	ProductPackagingUnitSQUAREINCH string = "SQUARE_INCH"

	// ProductPackagingUnitSQUAREMETER captures enum value "SQUARE_METER"
	ProductPackagingUnitSQUAREMETER string = "SQUARE_METER"
)

// prop value enum
func (m *Product) validatePackagingUnitEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, productTypePackagingUnitPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Product) validatePackagingUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.PackagingUnit) { // not required
		return nil
	}

	// value enum
	if err := m.validatePackagingUnitEnum("packagingUnit", "body", m.PackagingUnit); err != nil {
		return err
	}

	return nil
}

func (m *Product) validatePrices(formats strfmt.Registry) error {

	if swag.IsZero(m.Prices) { // not required
		return nil
	}

	for i := 0; i < len(m.Prices); i++ {
		if swag.IsZero(m.Prices[i]) { // not required
			continue
		}

		if m.Prices[i] != nil {
			if err := m.Prices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Product) validateProductionType(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductionType) { // not required
		return nil
	}

	if m.ProductionType != nil {
		if err := m.ProductionType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("productionType")
			}
			return err
		}
	}

	return nil
}

func (m *Product) validateRelatedProducts(formats strfmt.Registry) error {

	if swag.IsZero(m.RelatedProducts) { // not required
		return nil
	}

	if m.RelatedProducts != nil {
		if err := m.RelatedProducts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relatedProducts")
			}
			return err
		}
	}

	return nil
}

func (m *Product) validateSector(formats strfmt.Registry) error {

	if swag.IsZero(m.Sector) { // not required
		return nil
	}

	if m.Sector != nil {
		if err := m.Sector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sector")
			}
			return err
		}
	}

	return nil
}

var productTypeSubproductPresentationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","HIDE_ALL","HIDE_PRICES","HIDE_QUANTITIES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		productTypeSubproductPresentationPropEnum = append(productTypeSubproductPresentationPropEnum, v)
	}
}

const (

	// ProductSubproductPresentationDEFAULT captures enum value "DEFAULT"
	ProductSubproductPresentationDEFAULT string = "DEFAULT"

	// ProductSubproductPresentationHIDEALL captures enum value "HIDE_ALL"
	ProductSubproductPresentationHIDEALL string = "HIDE_ALL"

	// ProductSubproductPresentationHIDEPRICES captures enum value "HIDE_PRICES"
	ProductSubproductPresentationHIDEPRICES string = "HIDE_PRICES"

	// ProductSubproductPresentationHIDEQUANTITIES captures enum value "HIDE_QUANTITIES"
	ProductSubproductPresentationHIDEQUANTITIES string = "HIDE_QUANTITIES"
)

// prop value enum
func (m *Product) validateSubproductPresentationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, productTypeSubproductPresentationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Product) validateSubproductPresentation(formats strfmt.Registry) error {

	if swag.IsZero(m.SubproductPresentation) { // not required
		return nil
	}

	// value enum
	if err := m.validateSubproductPresentationEnum("subproductPresentation", "body", m.SubproductPresentation); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateSubproducts(formats strfmt.Registry) error {

	if swag.IsZero(m.Subproducts) { // not required
		return nil
	}

	for i := 0; i < len(m.Subproducts); i++ {
		if swag.IsZero(m.Subproducts[i]) { // not required
			continue
		}

		if m.Subproducts[i] != nil {
			if err := m.Subproducts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subproducts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Product) validateSupplierPrices(formats strfmt.Registry) error {

	if swag.IsZero(m.SupplierPrices) { // not required
		return nil
	}

	for i := 0; i < len(m.SupplierPrices); i++ {
		if swag.IsZero(m.SupplierPrices[i]) { // not required
			continue
		}

		if m.SupplierPrices[i] != nil {
			if err := m.SupplierPrices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supplierPrices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Product) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Product) validateTicketDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.TicketDefinition) { // not required
		return nil
	}

	if m.TicketDefinition != nil {
		if err := m.TicketDefinition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ticketDefinition")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Product) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Product) UnmarshalBinary(b []byte) error {
	var res Product
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}