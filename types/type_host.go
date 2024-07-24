// Copyright 2024 The tantianran Authors. All rights reserved.
// Use of this source code is governed by an Apache-2.0 license
// that can be found in the LICENSE file.
//
// Package api defines structs for request and response bodies used in interactions with the Zabbix API.
// This file specifically defines structs related to user-related API methods.
// Each struct corresponds to a specific API method's request or response format, ensuring type safety
// and clarity in data exchange with the Zabbix server.
//
// The structs defined here are integral to encoding and decoding JSON-RPC messages sent and received
// by the API client.

package types

type HostCreateRequest struct {
	JSONRPC string           `json:"jsonrpc"`
	Method  string           `json:"method"`
	Params  HostCreateParams `json:"params"`
	ID      int              `json:"id"`
}

type HostCreateParams struct {
	HostId            int             `json:"hostid,omitempty"`
	Host              string          `json:"host,omitempty"`
	Description       string          `json:"description,omitempty"`
	Flags             int             `json:"flags,omitempty"`
	InventoryMode     int             `json:"inventory_mode,omitempty"`
	IpmiAuthtype      int             `json:"ipmi_authtype,omitempty"`
	IpmiPassword      string          `json:"ipmi_password,omitempty"`
	IpmiPrivilege     int             `json:"ipmi_privilege,omitempty"`
	IpmiUsername      string          `json:"ipmi_username,omitempty"`
	MaintenanceFrom   int64           `json:"maintenance_from,omitempty"`
	MaintenanceStatus int             `json:"maintenance_status,omitempty"`
	MaintenanceType   int             `json:"maintenance_type,omitempty"`
	Maintenanceid     int             `json:"maintenanceid,omitempty"`
	Name              string          `json:"name,omitempty"`
	MonitoredBy       int             `json:"monitored_by,omitempty"`
	ProxyId           int             `json:"proxyid,omitempty"`
	ProxyGroupid      int             `json:"proxy_groupid,omitempty"`
	Status            int             `json:"status,omitempty"`
	TlsConnect        int             `json:"tls_connect,omitempty"`
	TlsAccept         int             `json:"tls_accept,omitempty"`
	TlsIssuer         string          `json:"tls_issuer,omitempty"`
	TlsSubject        string          `json:"tls_subject,omitempty"`
	TlsPskIdentity    string          `json:"tls_psk_identity,omitempty"`
	TlsPsk            string          `json:"tls_psk,omitempty"`
	ActiveAvailable   int             `json:"active_available,omitempty"`
	AssignedProxyid   int             `json:"assigned_proxyid,omitempty"`
	Groups            GroupArray      `json:"groups,omitempty"`
	Interfaces        InterfacesArray `json:"interfaces,omitempty"`
	Tags              TagArray        `json:"tags,omitempty"`
	Templates         TemplatesArray  `json:"templates,omitempty"`
	Macros            MacrosArray     `json:"macros,omitempty"`
	Inventory         InventoryParams `json:"inventory,omitempty"`
}

type GroupArray struct {
	GroupID string `json:"groupid,omitempty"`
	Name    string `json:"name,omitempty"`
	Flags   int    `json:"flags,omitempty"`
	Uuid    string `json:"uuid,omitempty"`
}

type InterfacesArray struct {
	Interfaceid  int           `json:"interfaceid,omitempty"`
	Available    int           `json:"available,omitempty"`
	Hostid       int           `json:"hostid,omitempty"`
	Type         int           `json:"type,omitempty"`
	Ip           string        `json:"ip,omitempty"`
	Dns          string        `json:"dns,omitempty"`
	Port         string        `json:"port,omitempty"`
	Useip        int           `json:"useip,omitempty"`
	Main         int           `json:"main,omitempty"`
	Details      []DetailArray `json:"details,omitempty"`
	DisableUntil int64         `json:"disable_until,omitempty"`
	Error        string        `json:"error,omitempty"`
	ErrorsFrom   int           `json:"errors_from,omitempty"`
}

type DetailArray struct {
	Version        int    `json:"version,omitempty"`
	Bulk           int    `json:"bulk,omitempty"`
	Community      string `json:"community,omitempty"`
	MaxRepetitions int    `json:"max_repetitions,omitempty"`
	Securityname   string `json:"securityname,omitempty"`
	Securitylevel  int    `json:"securitylevel,omitempty"`
	Authpassphrase string `json:"authpassphrase,omitempty"`
	Privpassphrase string `json:"privpassphrase,omitempty"`
	Authprotocol   int    `json:"authprotocol,omitempty"`
	Privprotocol   int    `json:"privprotocol,omitempty"`
	Contextname    string `json:"contextname,omitempty"`
}

type TagArray struct {
	Tag       string `json:"tag,omitempty"`
	Value     string `json:"value,omitempty"`
	Automatic int    `json:"automatic,omitempty"`
}

type TemplatesArray struct {
	Templateid    string `json:"templateid,omitempty"`
	Host          string `json:"host,omitempty"`
	Description   string `json:"description,omitempty"`
	Name          string `json:"name,omitempty"`
	Uuid          string `json:"uuid,omitempty"`
	VendorName    string `json:"vendor_name,omitempty"`
	VendorVersion string `json:"vendor_version,omitempty"`
}

type MacrosArray struct {
	Globalmacroid int    `json:"globalmacroid,omitempty"`
	Macro         string `json:"macro,omitempty"`
	Value         string `json:"value,omitempty"`
	Type          int    `json:"type,omitempty"`
	Description   string `json:"description,omitempty"`
}

type InventoryParams struct {
	Alias            string `json:"alias,omitempty"`
	AssetTag         string `json:"asset_tag,omitempty"`
	Chassis          string `json:"chassis,omitempty"`
	Contact          string `json:"contact,omitempty"`
	ContractNumber   string `json:"contract_number,omitempty"`
	DateHwDecomm     string `json:"date_hw_decomm,omitempty"`
	DateHwExpiry     string `json:"date_hw_expiry,omitempty"`
	DateHwInstall    string `json:"date_hw_install,omitempty"`
	DateHwPurchase   string `json:"date_hw_purchase,omitempty"`
	DeploymentStatus string `json:"deployment_status,omitempty"`
	Hardware         string `json:"hardware,omitempty"`
	HardwareFull     string `json:"hardware_full,omitempty"`
	HostNetmask      string `json:"host_netmask,omitempty"`
	HostNetworks     string `json:"host_networks,omitempty"`
	HostRouter       string `json:"host_router,omitempty"`
	HwArch           string `json:"hw_arch,omitempty"`
	InstallerName    string `json:"installer_name,omitempty"`
	Location         string `json:"location,omitempty"`
	LocationLat      string `json:"location_lat,omitempty"`
	LocationLon      string `json:"location_lon,omitempty"`
	MacAddressA      string `json:"macaddress_a,omitempty"`
	MacAddressB      string `json:"macaddress_b,omitempty"`
	Model            string `json:"model,omitempty"`
	Name             string `json:"name,omitempty"`
	Notes            string `json:"notes,omitempty"`
	OobIP            string `json:"oob_ip,omitempty"`
	OobNetmask       string `json:"oob_netmask,omitempty"`
	OobRouter        string `json:"oob_router,omitempty"`
	OS               string `json:"os,omitempty"`
	OSFull           string `json:"os_full,omitempty"`
	OSShort          string `json:"os_short,omitempty"`
	Poc1Cell         string `json:"poc_1_cell,omitempty"`
	Poc1Email        string `json:"poc_1_email,omitempty"`
	Poc1Name         string `json:"poc_1_name,omitempty"`
	Poc1Notes        string `json:"poc_1_notes,omitempty"`
	Poc1PhoneA       string `json:"poc_1_phone_a,omitempty"`
	Poc1PhoneB       string `json:"poc_1_phone_b,omitempty"`
	Poc1Screen       string `json:"poc_1_screen,omitempty"`
	Poc2Cell         string `json:"poc_2_cell,omitempty"`
	Poc2Email        string `json:"poc_2_email,omitempty"`
	Poc2Name         string `json:"poc_2_name,omitempty"`
	Poc2Notes        string `json:"poc_2_notes,omitempty"`
	Poc2PhoneA       string `json:"poc_2_phone_a,omitempty"`
	Poc2PhoneB       string `json:"poc_2_phone_b,omitempty"`
	Poc2Screen       string `json:"poc_2_screen,omitempty"`
	SerialnoA        string `json:"serialno_a,omitempty"`
	SerialnoB        string `json:"serialno_b,omitempty"`
	SiteAddressA     string `json:"site_address_a,omitempty"`
	SiteAddressB     string `json:"site_address_b,omitempty"`
	SiteAddressC     string `json:"site_address_c,omitempty"`
	SiteCity         string `json:"site_city,omitempty"`
	SiteCountry      string `json:"site_country,omitempty"`
	SiteNotes        string `json:"site_notes,omitempty"`
	SiteRack         string `json:"site_rack,omitempty"`
	SiteState        string `json:"site_state,omitempty"`
	SiteZip          string `json:"site_zip,omitempty"`
	Software         string `json:"software,omitempty"`
	SoftwareAppA     string `json:"software_app_a,omitempty"`
	SoftwareAppB     string `json:"software_app_b,omitempty"`
	SoftwareAppC     string `json:"software_app_c,omitempty"`
	SoftwareAppD     string `json:"software_app_d,omitempty"`
	SoftwareAppE     string `json:"software_app_e,omitempty"`
	SoftwareFull     string `json:"software_full,omitempty"`
	Tag              string `json:"tag,omitempty"`
	Type             string `json:"type,omitempty"`
	TypeFull         string `json:"type_full,omitempty"`
	URLA             string `json:"url_a,omitempty"`
	URLB             string `json:"url_b,omitempty"`
	URLC             string `json:"url_c,omitempty"`
	Vendor           string `json:"vendor,omitempty"`
}

type HostCreateResponse struct {
	JSONRPC string            `json:"jsonrpc"`
	Result  *HostCreateResult `json:"result,omitempty"`
	Error   *HostCreateError  `json:"error,omitempty"`
	ID      int               `json:"id"`
}

type HostCreateResult struct {
	HostIDs []string `json:"hostids"`
}

type HostCreateError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
