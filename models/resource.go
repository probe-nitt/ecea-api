package models

type StudyMaterialRequest struct {
	Name            string `json:"name" form:"name"`
	Subject         string `json:"subject" form:"subject"`
	SubjectCategory string `json:"subject_category" form:"category"`
	SubjectCode     string `json:"subject_code" form:"code"`
}
type StudyMaterials struct {
	Name            string `json:"name"`
	Subject         string `json:"subject"`
	SubjectCategory string `json:"subject_category"`
	SubjectCode     string `hson:"subject_code"`
	DocumentURL     string `json:"document_url"`
}

type SubjectCategory string

const (
	DigitalElectronics           SubjectCategory = "DIGITAL ELECTRONICS"
	AnalogElectronics            SubjectCategory = "ANALOG ELECTRONICS"
	Telecommunication            SubjectCategory = "TELECOMMUNICATION"
	CommunicationChanneling      SubjectCategory = "COMMUNICATION CHANNELING"
	SystemDesignAndArchitechture SubjectCategory = "SYSTEM DESIGN AND ARCHITECTURE"
	BasicEngineering             SubjectCategory = "BASIC ENGINEERING"
	Mathematics                  SubjectCategory = "MATHEMATICS"
	Others                       SubjectCategory = "OTHERS"
)
