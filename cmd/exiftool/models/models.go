package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Exifdata struct {
	SourceFile          string
	DateTimeOriginal    *time.Time
	CreateDate          *time.Time
	FileAccessDate      string
	FileInodeChangeDate string
	FileModifyDate      *time.Time
	MediaCreateDate     string
	MediaModifyDate     string
	ModifyDate          *time.Time
	TrackCreateDate     string
	TrackModifyDate     string
	// AudioBitsPerSample   int      `json:"AudioBitsPerSample,omitempty"`
	// AudioChannels        int      `json:"AudioChannels,omitempty"`
	// AudioFormat          string   `json:"AudioFormat,omitempty"`
	// AudioSampleRate      int      `json:"AudioSampleRate,omitempty"`
	// AudioSetting         string   `json:"AudioSetting,omitempty"`
	// AutoISOMax           int      `json:"AutoISOMax,omitempty"`
	// AutoISOMin           int      `json:"AutoISOMin,omitempty"`
	// AutoRotation         string   `json:"AutoRotation,omitempty"`
	// AvgBitrate           string   `json:"AvgBitrate,omitempty"`
	// BackgroundColor      string   `json:"BackgroundColor,omitempty"`
	// Balance              int      `json:"Balance,omitempty"`
	// BitDepth             int      `json:"BitDepth,omitempty"`
	// CameraSerialNumber   string   `json:"CameraSerialNumber,omitempty"`
	// ColorMode            string   `json:"ColorMode,omitempty"`
	// ColorRepresentation  string   `json:"ColorRepresentation,omitempty"`
	// CompatibleBrands     []string `json:"CompatibleBrands,omitempty"`
	// CompressorID         string   `json:"CompressorID,omitempty"`
	// CompressorName       string   `json:"CompressorName,omitempty"`
	// CurrentTime          string   `json:"CurrentTime,omitempty"`
	// DeviceName           string   `json:"DeviceName,omitempty"`
	// DigitalZoom          string   `json:"DigitalZoom,omitempty"`
	// Directory            string   `json:"Directory,omitempty"`
	// Duration             string   `json:"Duration,omitempty"`
	// ExifToolVersion      float32  `json:"ExifToolVersion,omitempty"`
	// ExposureCompensation float32  `json:"ExposureCompensation,omitempty"`
	// FieldOfView          string   `json:"FieldOfView,omitempty"`
	FileName          string
	FilePermissions   string
	FileSize          string
	FileType          string
	FileTypeExtension string
	FirmwareVersion   string
	GPSCoordinates    string
	GPSLatitude       string
	GPSLongitude      string
	GPSPosition       string
	GPSLatitudeRef    string
	GPSLongitudeRef   string
	GPSAltitudeRef    string
	GPSTimeStamp      string
	GPSDateStamp      string
	GPSAltitude       string
	GPSDateTime       *time.Time
	// FontName           string  `json:"FontName,omitempty"`
	// GenBalance         int     `json:"GenBalance,omitempty"`
	// GenFlags           string  `json:"GenFlags,omitempty"`
	// GenGraphicsMode    string  `json:"GenGraphicsMode,omitempty"`
	// GenMediaVersion    int     `json:"GenMediaVersion,omitempty"`
	// GenOpColor         string  `json:"GenOpColor,omitempty"`
	// GraphicsMode       string  `json:"GraphicsMode,omitempty"`
	// HandlerClass       string  `json:"HandlerClass,omitempty"`
	// HandlerDescription string  `json:"HandlerDescription,omitempty"`
	// HandlerType        string  `json:"HandlerType,omitempty"`
	// ImageHeight        int     `json:"ImageHeight,omitempty"`
	// ImageSize          string  `json:"ImageSize,omitempty"`
	// ImageWidth         int     `json:"ImageWidth,omitempty"`
	// LensSerialNumber   string  `json:"LensSerialNumber,omitempty"`
	// MajorBrand         string  `json:"MajorBrand,omitempty"`
	// MatrixStructure    string  `json:"MatrixStructure,omitempty"`
	// MediaDuration      string  `json:"MediaDuration,omitempty"`
	// MediaHeaderVersion int     `json:"MediaHeaderVersion,omitempty"`
	// MediaTimeScale     int     `json:"MediaTimeScale,omitempty"`
	// Megapixels         float32 `json:"Megapixels,omitempty"`
	// MetaFormat         string  `json:"MetaFormat,omitempty"`
	// MIMEType           string  `json:"MIMEType,omitempty"`
	// MinorVersion       string  `json:"MinorVersion,omitempty"`
	// Model              string  `json:"Model,omitempty"`
	// MovieDataOffset    int     `json:"MovieDataOffset,omitempty"`
	// MovieDataSize      int     `json:"MovieDataSize,omitempty"`
	// MovieHeaderVersion int     `json:"MovieHeaderVersion,omitempty"`
	// NextTrackID        int     `json:"NextTrackID,omitempty"`
	// OpColor            string  `json:"OpColor,omitempty"`
	// OtherFormat        string  `json:"OtherFormat,omitempty"`
	// PosterTime         string  `json:"PosterTime,omitempty"`
	// PreferredRate      int     `json:"PreferredRate,omitempty"`
	// PreferredVolume    string  `json:"PreferredVolume,omitempty"`
	// PreviewDuration    string  `json:"PreviewDuration,omitempty"`
	// PreviewTime        string  `json:"PreviewTime,omitempty"`
	// ProTune            string  `json:"ProTune,omitempty"`
	// Rate               string  `json:"Rate,omitempty"`
	// Rotation           int     `json:"Rotation,omitempty"`
	// SelectionDuration  string  `json:"SelectionDuration,omitempty"`
	// SelectionTime      string  `json:"SelectionTime,omitempty"`
	// SerialNumberHash   string  `json:"SerialNumberHash,omitempty"`
	// Sharpness          string  `json:"Sharpness,omitempty"`
	// SourceImageHeight  int     `json:"SourceImageHeight,omitempty"`
	// SourceImageWidth   int     `json:"SourceImageWidth,omitempty"`
	// TextColor          string  `json:"TextColor,omitempty"`
	// TextFace           string  `json:"TextFace,omitempty"`
	// TextFont           string  `json:"TextFont,omitempty"`
	// TextSize           int     `json:"TextSize,omitempty"`
	// TimeCode           int     `json:"TimeCode,omitempty"`
	// TimeScale          int     `json:"TimeScale,omitempty"`
	// TrackDuration      string  `json:"TrackDuration,omitempty"`
	// TrackHeaderVersion int     `json:"TrackHeaderVersion,omitempty"`
	// TrackID            int     `json:"TrackID,omitempty"`
	// TrackLayer         int     `json:"TrackLayer,omitempty"`
	// TrackVolume        string  `json:"TrackVolume,omitempty"`
	// VideoFrameRate     float32 `json:"VideoFrameRate,omitempty"`
	// Warning            string  `json:"Warning,omitempty"`
	// WhiteBalance       string  `json:"WhiteBalance,omitempty"`
	// XResolution        int     `json:"XResolution,omitempty"`
	// YResolution        int     `json:"YResolution,omitempty"`
}

func (t *Exifdata) Parse(data map[string]interface{}) {
	t.SourceFile = fmt.Sprint(data["SourceFile"])
	t.DateTimeOriginal = parseDate("DateTimeOriginal", data)
	t.CreateDate = parseDate("CreateDate", data)
	t.ModifyDate = parseDate("ModifyDate", data)
	t.FileAccessDate = fmt.Sprint(data["FileAccessDate"])
	t.FileInodeChangeDate = fmt.Sprint(data["FileInodeChangeDate"])
	t.FileModifyDate = parseDateByLayout("FileModifyDate", data, "2006:01:02 15:04:05Z07:00") // "2019:12:30 13:50:06-07:00",
	t.MediaCreateDate = fmt.Sprint(data["MediaCreateDate"])
	t.MediaModifyDate = fmt.Sprint(data["MediaModifyDate"])

	t.TrackCreateDate = fmt.Sprint(data["TrackCreateDate"])
	t.TrackModifyDate = fmt.Sprint(data["TrackModifyDate"])
	t.FileName = fmt.Sprint(data["FileName"])
	t.FilePermissions = fmt.Sprint(data["FilePermissions"])
	t.FileSize = fmt.Sprint(data["FileSize"])
	t.FileType = fmt.Sprint(data["FileType"])
	t.FileTypeExtension = fmt.Sprint(data["FileTypeExtension"])
	t.FirmwareVersion = fmt.Sprint(data["FirmwareVersion"])
	t.GPSCoordinates = fmt.Sprint(data["GPSCoordinates"])
	t.GPSLatitude = fmt.Sprint(data["GPSLatitude"])
	t.GPSLongitude = fmt.Sprint(data["GPSLongitude"])
	t.GPSPosition = fmt.Sprint(data["GPSPosition"])
	t.GPSLatitudeRef = fmt.Sprint(data["GPSLatitudeRef"])
	t.GPSLongitudeRef = fmt.Sprint(data["GPSLongitudeRef"])
	t.GPSAltitudeRef = fmt.Sprint(data["GPSAltitudeRef"])
	t.GPSTimeStamp = fmt.Sprint(data["GPSTimeStamp"])
	t.GPSDateStamp = fmt.Sprint(data["GPSDateStamp"])
	t.GPSAltitude = fmt.Sprint(data["GPSAltitude"])
	t.GPSDateTime = parseDateZ("GPSDateTime", data)

	// _, offset := t.FileModifyDate.Local().Zone()
	// t.GPSDateTime = t.GPSDateTime.Local().Add(time.Duration(offset))
}

// 2020:02:02 13:53:19
func parseDateZ(key string, data map[string]interface{}) *time.Time {
	return parseDateByLayout(key, data, "2006:01:02 15:04:05Z")
}

func parseDate(key string, data map[string]interface{}) *time.Time {
	return parseDateByLayout(key, data, "2006:01:02 15:04:05")
}

func parseDateByLayout(key string, data map[string]interface{}, layout string) *time.Time {
	if value, ok := getStringFromMap(key, data); ok {
		t, err := time.Parse(layout, value)
		checkError(err, value)
		return &t
	}
	return &time.Time{}
}

func getStringFromMap(key string, data map[string]interface{}) (string, bool) {
	if value := data[key]; value != nil {
		if date := fmt.Sprint(value); date != "" {
			return date, true
		}
	}
	return "", false
}

func checkError(err error, message string) {
	if err != nil {
		log.Fatal("could not parse date: ", err, message)
	}
}

// Decode will decode into JSONData
func (t *Exifdata) String() string {
	var info string
	info += fmt.Sprintf("====== %v", t.FileName)
	info += fmt.Sprintf("\nDateTimeOriginal:\t %v", t.DateTimeOriginal.Format("2006-01-02T15:04:05 MST"))
	info += fmt.Sprintf("\nCreateDate:\t\t %v", t.CreateDate.Format("2006-01-02T15:04:05 MST"))
	if !t.GPSDateTime.IsZero() {
		info += fmt.Sprintf("\nGPSDateTime:\t\t %v", t.GPSDateTime.Format("2006-01-02T15:04:05 MST"))
	}
	return info
}

// Set
func (t *Exifdata) SetZoneOffset(datum *time.Time, offset int) {
	hours := time.Duration(int(time.Hour) * offset).Seconds()
	*datum = datum.In(time.FixedZone("", int(hours)))
}

// Decode will decode into JSONData
func (t *Exifdata) Decode(data []byte) error {
	return json.Unmarshal(data, t)
}

// OtherJSONExamples shows ways to use types
// beyond structs and other useful functions
func OtherJSONExamples() error {
	var res []map[string]string
	err := json.Unmarshal([]byte(`[{"key": "value"}]`), &res)
	if err != nil {
		return err
	}

	fmt.Println("We can unmarshal into a map instead of a struct:", res)

	b := bytes.NewReader([]byte(`[{"key2": "value2"}]`))
	decoder := json.NewDecoder(b)

	if err := decoder.Decode(&res); err != nil {
		return err
	}

	fmt.Println("we can also use decoders/encoders to work with streams:", res)

	return nil
}
