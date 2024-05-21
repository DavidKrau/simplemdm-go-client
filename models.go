package simplemdm

type Attribute struct {
	Data struct {
		Type       string     `json:"type"`
		ID         string     `json:"id"`
		Attributes Attributes `json:"attributes"`
	} `json:"data"`
}

type AttributeArray struct {
	Data []struct {
		Type       string     `json:"type"`
		ID         string     `json:"id"`
		Attributes Attributes `json:"attributes"`
	} `json:"data"`
}

type SimplemdmDefaultStruct struct {
	Data struct {
		Type          string     `json:"type"`
		ID            int        `json:"id"`
		Attributes    Attributes `json:"attributes"`
		Relationships Relations  `json:"relationships,omitempty"`
	} `json:"data"`
}

type SimpleMDMArayStruct struct {
	Data []struct {
		Type          string     `json:"type"`
		ID            int        `json:"id"`
		Attributes    Attributes `json:"attributes"`
		Relationships Relations  `json:"relationships,omitempty"`
	} `json:"data"`
	HasMore bool `json:"has_more"`
}

type Attributes struct {
	Name                   string `json:"name"`
	AutoDeploy             bool   `json:"auto_deploy"`
	Type                   string `json:"type"`
	InstallType            string `json:"install_type"`
	DefaultValue           string `json:"default_value"`
	ReinstallAfterOsUpdate bool   `json:"reinstall_after_os_update"`
	ProfileIdentifier      string `json:"profile_identifier"`
	UserScope              bool   `json:"user_scope"`
	AttributeSupport       bool   `json:"attribute_support"`
	EscapeAttributes       bool   `json:"escape_attributes"`
	GroupCount             int    `json:"group_count"`
	DeviceCount            int    `json:"device_count"`
	ProfileSHA             string `json:"profile_sha"`
	Source                 string `json:"source"`
	Secret                 bool   `json:"secret"`
	Value                  string `json:"value"`
	DeviceName             string `json:"device_name"`
	EnrollmentURL          string `json:"enrollment_url"`
	Content                string `json:"content"`
	VariableSupport        bool   `json:"variable_support"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
}

type Relations struct {
	Apps             Apps             `json:"apps,omitempty"`
	DeviceGroups     DeviceGroups     `json:"device_groups,omitempty"`
	DeviceGroup      DeviceGroup      `json:"device_group,omitempty"`
	Media            Media            `json:"media,omitempty"`
	Devices          Devices          `json:"devices,omitempty"`
	CustomAttributes CustomAttributes `json:"custom_attribute_values,omitempty"`
}

type Apps struct {
	Data []Data `json:"data,omitempty"`
}

type CustomAttributes struct {
	Data []DataCustomAttributes `json:"data,omitempty"`
}

type DeviceGroups struct {
	Data []Data `json:"data,omitempty"`
}

type DeviceGroup struct {
	Data Data `json:"data,omitempty"`
}

type Media struct {
	Data []Data `json:"data,omitempty"`
}

type Devices struct {
	Data []Data `json:"data,omitempty"`
}

type Data struct {
	Type string `json:"type"`
	ID   int    `json:"id"`
}
type DataCustomAttributes struct {
	Type       string     `json:"type"`
	ID         string     `json:"id"`
	Attributes Attributes `json:"attributes"`
}
