package computervision

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/cognitiveservices/v3.0-preview/computervision"

// OcrDetectionLanguage enumerates the values for ocr detection language.
type OcrDetectionLanguage string

const (
	// En ...
	En OcrDetectionLanguage = "en"
	// Es ...
	Es OcrDetectionLanguage = "es"
)

// PossibleOcrDetectionLanguageValues returns an array of possible values for the OcrDetectionLanguage const type.
func PossibleOcrDetectionLanguageValues() []OcrDetectionLanguage {
	return []OcrDetectionLanguage{En, Es}
}

// OperationStatusCodes enumerates the values for operation status codes.
type OperationStatusCodes string

const (
	// Failed ...
	Failed OperationStatusCodes = "failed"
	// NotStarted ...
	NotStarted OperationStatusCodes = "notStarted"
	// Running ...
	Running OperationStatusCodes = "running"
	// Succeeded ...
	Succeeded OperationStatusCodes = "succeeded"
)

// PossibleOperationStatusCodesValues returns an array of possible values for the OperationStatusCodes const type.
func PossibleOperationStatusCodesValues() []OperationStatusCodes {
	return []OperationStatusCodes{Failed, NotStarted, Running, Succeeded}
}

// TextRecognitionResultDimensionUnit enumerates the values for text recognition result dimension unit.
type TextRecognitionResultDimensionUnit string

const (
	// Inch ...
	Inch TextRecognitionResultDimensionUnit = "inch"
	// Pixel ...
	Pixel TextRecognitionResultDimensionUnit = "pixel"
)

// PossibleTextRecognitionResultDimensionUnitValues returns an array of possible values for the TextRecognitionResultDimensionUnit const type.
func PossibleTextRecognitionResultDimensionUnitValues() []TextRecognitionResultDimensionUnit {
	return []TextRecognitionResultDimensionUnit{Inch, Pixel}
}

// AnalyzeResults analyze batch operation result.
type AnalyzeResults struct {
	// Version - Version of schema used for this result.
	Version *string `json:"version,omitempty"`
	// ReadResults - Text extracted from the input.
	ReadResults *[]ReadResult `json:"readResults,omitempty"`
}

// Error details about the API request error.
type Error struct {
	// Code - The error code.
	Code interface{} `json:"code,omitempty"`
	// Message - A message explaining the error reported by the service.
	Message *string `json:"message,omitempty"`
	// RequestID - A unique request identifier.
	RequestID *string `json:"requestId,omitempty"`
}

// ImageURL ...
type ImageURL struct {
	// URL - Publicly reachable URL of an image.
	URL *string `json:"url,omitempty"`
}

// Line an object representing a recognized text line.
type Line struct {
	// Language - The BCP-47 language code of the recognized text line. Only provided where the language of the line differs from the page's.
	Language *string `json:"language,omitempty"`
	// BoundingBox - Bounding box of a recognized line.
	BoundingBox *[]float64 `json:"boundingBox,omitempty"`
	// Text - The text content of the line.
	Text *string `json:"text,omitempty"`
	// Words - List of words in the text line.
	Words *[]Word `json:"words,omitempty"`
}

// ReadOperationResult OCR result of the read operation.
type ReadOperationResult struct {
	autorest.Response `json:"-"`
	// Status - Status of the read operation. Possible values include: 'NotStarted', 'Running', 'Failed', 'Succeeded'
	Status OperationStatusCodes `json:"status,omitempty"`
	// CreatedDateTime - Get UTC date time the batch operation was submitted.
	CreatedDateTime *date.Time `json:"createdDateTime,omitempty"`
	// LastUpdatedDateTime - Get last updated UTC date time of this batch operation.
	LastUpdatedDateTime *date.Time `json:"lastUpdatedDateTime,omitempty"`
	// AnalyzeResult - Analyze batch operation result.
	AnalyzeResult *AnalyzeResults `json:"analyzeResult,omitempty"`
}

// ReadResult text extracted from a page in the input document.
type ReadResult struct {
	// Page - The 1-based page number of the recognition result.
	Page *int32 `json:"page,omitempty"`
	// Language - The BCP-47 language code of the recognized text page.
	Language *string `json:"language,omitempty"`
	// Angle - The orientation of the image in degrees in the clockwise direction. Range between [-180, 180).
	Angle *float64 `json:"angle,omitempty"`
	// Width - The width of the image in pixels or the PDF in inches.
	Width *float64 `json:"width,omitempty"`
	// Height - The height of the image in pixels or the PDF in inches.
	Height *float64 `json:"height,omitempty"`
	// Unit - The unit used in the Width, Height and BoundingBox. For images, the unit is 'pixel'. For PDF, the unit is 'inch'. Possible values include: 'Pixel', 'Inch'
	Unit TextRecognitionResultDimensionUnit `json:"unit,omitempty"`
	// Lines - A list of recognized text lines.
	Lines *[]Line `json:"lines,omitempty"`
}

// Word an object representing a recognized word.
type Word struct {
	// BoundingBox - Bounding box of a recognized word.
	BoundingBox *[]float64 `json:"boundingBox,omitempty"`
	// Text - The text content of the word.
	Text *string `json:"text,omitempty"`
	// Confidence - Qualitative confidence measure.
	Confidence *float64 `json:"confidence,omitempty"`
}