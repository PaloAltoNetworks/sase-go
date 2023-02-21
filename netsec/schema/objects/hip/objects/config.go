package objects

/*
Config ""

Parent chains:
* "Config"

Args:

Param AntiMalware (AntiMalwareObject): The AntiMalware param. AntiMalwareObject instance.

Param Certificate (CertificateObject): The Certificate param. CertificateObject instance.

Param CustomChecks (CustomChecksObject): The CustomChecks param. CustomChecksObject instance.

Param DataLossPrevention (DataLossPreventionObject): The DataLossPrevention param. DataLossPreventionObject instance.

Param Description (string): The Description param. String lengh must be between 0 and 255 characters.

Param DiskBackup (DiskBackupObject): The DiskBackup param. DiskBackupObject instance.

Param DiskEncryption (DiskEncryptionObject): The DiskEncryption param. DiskEncryptionObject instance.

Param Firewall (FirewallObject): The Firewall param. FirewallObject instance.

Param HostInfo (HostInfoObject): The HostInfo param. HostInfoObject instance.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param MobileDevice (MobileDeviceObject): The MobileDevice param. MobileDeviceObject instance.

Param Name (required, string): The Name param. String length must not exceed 31 characters.

Param NetworkInfo (NetworkInfoObject): The NetworkInfo param. NetworkInfoObject instance.

Param PatchManagement (PatchManagementObject): The PatchManagement param. PatchManagementObject instance.
*/
type Config struct {
    AntiMalware *AntiMalwareObject `json:"anti_malware,omitempty"`
    Certificate *CertificateObject `json:"certificate,omitempty"`
    CustomChecks *CustomChecksObject `json:"custom_checks,omitempty"`
    DataLossPrevention *DataLossPreventionObject `json:"data_loss_prevention,omitempty"`
    Description string `json:"description,omitempty"`
    DiskBackup *DiskBackupObject `json:"disk_backup,omitempty"`
    DiskEncryption *DiskEncryptionObject `json:"disk_encryption,omitempty"`
    Firewall *FirewallObject `json:"firewall,omitempty"`
    HostInfo *HostInfoObject `json:"host_info,omitempty"`
    ObjectId string `json:"id,omitempty"`
    MobileDevice *MobileDeviceObject `json:"mobile_device,omitempty"`
    Name string `json:"name"`
    NetworkInfo *NetworkInfoObject `json:"network_info,omitempty"`
    PatchManagement *PatchManagementObject `json:"patch_management,omitempty"`
}

/*
AntiMalwareObject ""

Parent chains:
* "Config anti_malware"

Args:

Param Criteria (CriteriaObject): The Criteria param. CriteriaObject instance.

Param ExcludeVendor (bool): The ExcludeVendor param. Default: false

Param Vendor (list of VendorObject objects): The Vendor param.
*/
type AntiMalwareObject struct {
    Criteria *CriteriaObject `json:"criteria,omitempty"`
    ExcludeVendor bool `json:"exclude_vendor,omitempty"`
    Vendor []VendorObject `json:"vendor,omitempty"`
}

/*
CriteriaObject ""

Parent chains:
* "Config anti_malware criteria"

Args:

Param IsInstalled (bool): The IsInstalled param. Default: true

Param LastScanTime (LastScanTimeObject): The LastScanTime param. LastScanTimeObject instance.

Param ProductVersion (ProductVersionObject): The ProductVersion param. ProductVersionObject instance.

Param RealTimeProtection (string): The RealTimeProtection param. String values: []string{"no", "yes", "not-available"}

Param VirdefVersion (VirdefVersionObject): The VirdefVersion param. VirdefVersionObject instance.
*/
type CriteriaObject struct {
    IsInstalled bool `json:"is_installed,omitempty"`
    LastScanTime *LastScanTimeObject `json:"last_scan_time,omitempty"`
    ProductVersion *ProductVersionObject `json:"product_version,omitempty"`
    RealTimeProtection string `json:"real_time_protection,omitempty"`
    VirdefVersion *VirdefVersionObject `json:"virdef_version,omitempty"`
}

/*
LastScanTimeObject ""

Parent chains:
* "Config anti_malware criteria last_scan_time"

Args:

Param NotAvailable (interface{}): The NotAvailable param. interface{} instance. One of these params should be specified:  not_available, not_within, or within.

Param NotWithin (NotWithinObject): The NotWithin param. NotWithinObject instance. One of these params should be specified:  not_available, not_within, or within.

Param Within (WithinObject): The Within param. WithinObject instance. One of these params should be specified:  not_available, not_within, or within.
*/
type LastScanTimeObject struct {
    NotAvailable interface{} `json:"not_available,omitempty"`
    NotWithin *NotWithinObject `json:"not_within,omitempty"`
    Within *WithinObject `json:"within,omitempty"`
}

/*
NotWithinObject ""

Parent chains:
* "Config anti_malware criteria last_scan_time not_within"
* "Config disk_backup criteria last_backup_time not_within"

Args:

Param Days (int64): The Days param. Value must be between 1 and 65535. Default: 1

Param Hours (int64): The Hours param. Value must be between 1 and 65535. Default: 24
*/
type NotWithinObject struct {
    Days int64 `json:"days,omitempty"`
    Hours int64 `json:"hours,omitempty"`
}

/*
WithinObject ""

Parent chains:
* "Config anti_malware criteria last_scan_time within"
* "Config disk_backup criteria last_backup_time within"

Args:

Param Days (int64): The Days param. Value must be between 1 and 65535. Default: 1

Param Hours (int64): The Hours param. Value must be between 1 and 65535. Default: 24
*/
type WithinObject struct {
    Days int64 `json:"days,omitempty"`
    Hours int64 `json:"hours,omitempty"`
}

/*
ProductVersionObject ""

Parent chains:
* "Config anti_malware criteria product_version"

Args:

Param Contains (string): The Contains param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, greater_equal, greater_than, is, is_not, less_equal, less_than, not_within, or within.

Param GreaterEqual (string): The GreaterEqual param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, greater_equal, greater_than, is, is_not, less_equal, less_than, not_within, or within.

Param GreaterThan (string): The GreaterThan param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, greater_equal, greater_than, is, is_not, less_equal, less_than, not_within, or within.

Param Is (string): The Is param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, greater_equal, greater_than, is, is_not, less_equal, less_than, not_within, or within.

Param IsNot (string): The IsNot param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, greater_equal, greater_than, is, is_not, less_equal, less_than, not_within, or within.

Param LessEqual (string): The LessEqual param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, greater_equal, greater_than, is, is_not, less_equal, less_than, not_within, or within.

Param LessThan (string): The LessThan param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, greater_equal, greater_than, is, is_not, less_equal, less_than, not_within, or within.

Param NotWithin (NotWithinObject1): The NotWithin param. NotWithinObject1 instance. One of these params should be specified:  contains, greater_equal, greater_than, is, is_not, less_equal, less_than, not_within, or within.

Param Within (WithinObject1): The Within param. WithinObject1 instance. One of these params should be specified:  contains, greater_equal, greater_than, is, is_not, less_equal, less_than, not_within, or within.
*/
type ProductVersionObject struct {
    Contains string `json:"contains,omitempty"`
    GreaterEqual string `json:"greater_equal,omitempty"`
    GreaterThan string `json:"greater_than,omitempty"`
    Is string `json:"is,omitempty"`
    IsNot string `json:"is_not,omitempty"`
    LessEqual string `json:"less_equal,omitempty"`
    LessThan string `json:"less_than,omitempty"`
    NotWithin *NotWithinObject1 `json:"not_within,omitempty"`
    Within *WithinObject1 `json:"within,omitempty"`
}

/*
NotWithinObject1 ""

Parent chains:
* "Config anti_malware criteria product_version not_within"

Args:

Param Versions (required, int64): The Versions param. Value must be between 1 and 65535. Default: 1
*/
type NotWithinObject1 struct {
    Versions int64 `json:"versions"`
}

/*
WithinObject1 ""

Parent chains:
* "Config anti_malware criteria product_version within"

Args:

Param Versions (required, int64): The Versions param. Value must be between 1 and 65535. Default: 1
*/
type WithinObject1 struct {
    Versions int64 `json:"versions"`
}

/*
VirdefVersionObject ""

Parent chains:
* "Config anti_malware criteria virdef_version"

Args:

Param NotWithin (NotWithinObject2): The NotWithin param. NotWithinObject2 instance. One of these params should be specified:  not_within or within.

Param Within (WithinObject2): The Within param. WithinObject2 instance. One of these params should be specified:  not_within or within.
*/
type VirdefVersionObject struct {
    NotWithin *NotWithinObject2 `json:"not_within,omitempty"`
    Within *WithinObject2 `json:"within,omitempty"`
}

/*
NotWithinObject2 ""

Parent chains:
* "Config anti_malware criteria virdef_version not_within"

Args:

Param Days (int64): The Days param. Value must be between 1 and 65535. Default: 1

Param Versions (int64): The Versions param. Value must be between 1 and 65535. Default: 1
*/
type NotWithinObject2 struct {
    Days int64 `json:"days,omitempty"`
    Versions int64 `json:"versions,omitempty"`
}

/*
WithinObject2 ""

Parent chains:
* "Config anti_malware criteria virdef_version within"

Args:

Param Days (int64): The Days param. Value must be between 1 and 65535. Default: 1

Param Versions (int64): The Versions param. Value must be between 1 and 65535. Default: 1
*/
type WithinObject2 struct {
    Days int64 `json:"days,omitempty"`
    Versions int64 `json:"versions,omitempty"`
}

/*
VendorObject ""

Parent chains:
* "Config anti_malware vendor"
* "Config disk_backup vendor"
* "Config disk_encryption vendor"
* "Config firewall vendor"

Args:

Param Name (required, string): The Name param. String length must not exceed 103 characters.

Param Product (list of strings): The Product param.
*/
type VendorObject struct {
    Name string `json:"name"`
    Product []string `json:"product,omitempty"`
}

/*
CertificateObject ""

Parent chains:
* "Config certificate"

Args:

Param Criteria (CriteriaObject1): The Criteria param. CriteriaObject1 instance.
*/
type CertificateObject struct {
    Criteria *CriteriaObject1 `json:"criteria,omitempty"`
}

/*
CriteriaObject1 ""

Parent chains:
* "Config certificate criteria"

Args:

Param CertificateAttributes (list of CertificateAttributesObject objects): The CertificateAttributes param.

Param CertificateProfile (string): The CertificateProfile param.
*/
type CriteriaObject1 struct {
    CertificateAttributes []CertificateAttributesObject `json:"certificate_attributes,omitempty"`
    CertificateProfile string `json:"certificate_profile,omitempty"`
}

/*
CertificateAttributesObject ""

Parent chains:
* "Config certificate criteria certificate_attributes"

Args:

Param Name (required, string): The Name param.

Param Value (string): The Value param. String length must not exceed 1024 characters.
*/
type CertificateAttributesObject struct {
    Name string `json:"name"`
    Value string `json:"value,omitempty"`
}

/*
CustomChecksObject ""

Parent chains:
* "Config custom_checks"

Args:

Param Criteria (required, CriteriaObject2): The Criteria param. CriteriaObject2 instance.
*/
type CustomChecksObject struct {
    Criteria CriteriaObject2 `json:"criteria"`
}

/*
CriteriaObject2 ""

Parent chains:
* "Config custom_checks criteria"

Args:

Param Plist (list of PlistObject objects): The Plist param.

Param ProcessList (list of ProcessListObject objects): The ProcessList param.

Param RegistryKey (list of RegistryKeyObject objects): The RegistryKey param.
*/
type CriteriaObject2 struct {
    Plist []PlistObject `json:"plist,omitempty"`
    ProcessList []ProcessListObject `json:"process_list,omitempty"`
    RegistryKey []RegistryKeyObject `json:"registry_key,omitempty"`
}

/*
PlistObject ""

Parent chains:
* "Config custom_checks criteria plist"

Args:

Param Key (list of KeyObject objects): The Key param.

Param Name (required, string): The Name param. String length must not exceed 1023 characters.

Param Negate (bool): The Negate param. Default: false
*/
type PlistObject struct {
    Key []KeyObject `json:"key,omitempty"`
    Name string `json:"name"`
    Negate bool `json:"negate,omitempty"`
}

/*
KeyObject ""

Parent chains:
* "Config custom_checks criteria plist key"

Args:

Param Name (required, string): The Name param. String length must not exceed 1023 characters.

Param Negate (bool): The Negate param. Default: false

Param Value (string): The Value param. String length must not exceed 1024 characters.
*/
type KeyObject struct {
    Name string `json:"name"`
    Negate bool `json:"negate,omitempty"`
    Value string `json:"value,omitempty"`
}

/*
ProcessListObject ""

Parent chains:
* "Config custom_checks criteria process_list"

Args:

Param Name (required, string): The Name param. String length must not exceed 1023 characters.

Param Running (bool): The Running param. Default: true
*/
type ProcessListObject struct {
    Name string `json:"name"`
    Running bool `json:"running,omitempty"`
}

/*
RegistryKeyObject ""

Parent chains:
* "Config custom_checks criteria registry_key"

Args:

Param DefaultValueData (string): The DefaultValueData param. String length must not exceed 1024 characters.

Param Name (required, string): The Name param. String length must not exceed 1023 characters.

Param Negate (bool): The Negate param. Default: false

Param RegistryValue (list of RegistryValueObject objects): The RegistryValue param.
*/
type RegistryKeyObject struct {
    DefaultValueData string `json:"default_value_data,omitempty"`
    Name string `json:"name"`
    Negate bool `json:"negate,omitempty"`
    RegistryValue []RegistryValueObject `json:"registry_value,omitempty"`
}

/*
RegistryValueObject ""

Parent chains:
* "Config custom_checks criteria registry_key registry_value"

Args:

Param Name (required, string): The Name param. String length must not exceed 1023 characters.

Param Negate (bool): The Negate param. Default: false

Param ValueData (string): The ValueData param. String length must not exceed 1024 characters.
*/
type RegistryValueObject struct {
    Name string `json:"name"`
    Negate bool `json:"negate,omitempty"`
    ValueData string `json:"value_data,omitempty"`
}

/*
DataLossPreventionObject ""

Parent chains:
* "Config data_loss_prevention"

Args:

Param Criteria (CriteriaObject3): The Criteria param. CriteriaObject3 instance.

Param ExcludeVendor (bool): The ExcludeVendor param. Default: false

Param Vendor (list of VendorObject1 objects): The Vendor param.
*/
type DataLossPreventionObject struct {
    Criteria *CriteriaObject3 `json:"criteria,omitempty"`
    ExcludeVendor bool `json:"exclude_vendor,omitempty"`
    Vendor []VendorObject1 `json:"vendor,omitempty"`
}

/*
CriteriaObject3 ""

Parent chains:
* "Config data_loss_prevention criteria"
* "Config firewall criteria"

Args:

Param IsEnabled (string): The IsEnabled param. String values: []string{"no", "yes", "not-available"}

Param IsInstalled (bool): The IsInstalled param. Default: true
*/
type CriteriaObject3 struct {
    IsEnabled string `json:"is_enabled,omitempty"`
    IsInstalled bool `json:"is_installed,omitempty"`
}

/*
VendorObject1 ""

Parent chains:
* "Config data_loss_prevention vendor"
* "Config patch_management vendor"

Args:

Param Name (required, string): The Name param. String length must not exceed 103 characters.

Param Product (list of strings): The Product param.
*/
type VendorObject1 struct {
    Name string `json:"name"`
    Product []string `json:"product,omitempty"`
}

/*
DiskBackupObject ""

Parent chains:
* "Config disk_backup"

Args:

Param Criteria (CriteriaObject4): The Criteria param. CriteriaObject4 instance.

Param ExcludeVendor (bool): The ExcludeVendor param. Default: false

Param Vendor (list of VendorObject objects): The Vendor param.
*/
type DiskBackupObject struct {
    Criteria *CriteriaObject4 `json:"criteria,omitempty"`
    ExcludeVendor bool `json:"exclude_vendor,omitempty"`
    Vendor []VendorObject `json:"vendor,omitempty"`
}

/*
CriteriaObject4 ""

Parent chains:
* "Config disk_backup criteria"

Args:

Param IsInstalled (bool): The IsInstalled param. Default: true

Param LastBackupTime (LastBackupTimeObject): The LastBackupTime param. LastBackupTimeObject instance.
*/
type CriteriaObject4 struct {
    IsInstalled bool `json:"is_installed,omitempty"`
    LastBackupTime *LastBackupTimeObject `json:"last_backup_time,omitempty"`
}

/*
LastBackupTimeObject ""

Parent chains:
* "Config disk_backup criteria last_backup_time"

Args:

Param NotAvailable (interface{}): The NotAvailable param. interface{} instance. One of these params should be specified:  not_available, not_within, or within.

Param NotWithin (NotWithinObject): The NotWithin param. NotWithinObject instance. One of these params should be specified:  not_available, not_within, or within.

Param Within (WithinObject): The Within param. WithinObject instance. One of these params should be specified:  not_available, not_within, or within.
*/
type LastBackupTimeObject struct {
    NotAvailable interface{} `json:"not_available,omitempty"`
    NotWithin *NotWithinObject `json:"not_within,omitempty"`
    Within *WithinObject `json:"within,omitempty"`
}

/*
DiskEncryptionObject ""

Parent chains:
* "Config disk_encryption"

Args:

Param Criteria (CriteriaObject5): The Criteria param. CriteriaObject5 instance.

Param ExcludeVendor (bool): The ExcludeVendor param. Default: false

Param Vendor (list of VendorObject objects): The Vendor param.
*/
type DiskEncryptionObject struct {
    Criteria *CriteriaObject5 `json:"criteria,omitempty"`
    ExcludeVendor bool `json:"exclude_vendor,omitempty"`
    Vendor []VendorObject `json:"vendor,omitempty"`
}

/*
CriteriaObject5 ""

Parent chains:
* "Config disk_encryption criteria"

Args:

Param EncryptedLocations (list of EncryptedLocationsObject objects): The EncryptedLocations param.

Param IsInstalled (bool): The IsInstalled param. Default: true
*/
type CriteriaObject5 struct {
    EncryptedLocations []EncryptedLocationsObject `json:"encrypted_locations,omitempty"`
    IsInstalled bool `json:"is_installed,omitempty"`
}

/*
EncryptedLocationsObject ""

Parent chains:
* "Config disk_encryption criteria encrypted_locations"

Args:

Param EncryptionState (EncryptionStateObject): The EncryptionState param. EncryptionStateObject instance.

Param Name (required, string): The Name param. String length must not exceed 1023 characters.
*/
type EncryptedLocationsObject struct {
    EncryptionState *EncryptionStateObject `json:"encryption_state,omitempty"`
    Name string `json:"name"`
}

/*
EncryptionStateObject ""

Parent chains:
* "Config disk_encryption criteria encrypted_locations encryption_state"

Args:

Param Is (string): The Is param. String values: []string{"encrypted", "unencrypted", "partial", "unknown"} Default: "encrypted".

Param IsNot (string): The IsNot param. String values: []string{"encrypted", "unencrypted", "partial", "unknown"} Default: "encrypted".
*/
type EncryptionStateObject struct {
    Is string `json:"is,omitempty"`
    IsNot string `json:"is_not,omitempty"`
}

/*
FirewallObject ""

Parent chains:
* "Config firewall"

Args:

Param Criteria (CriteriaObject3): The Criteria param. CriteriaObject3 instance.

Param ExcludeVendor (bool): The ExcludeVendor param. Default: false

Param Vendor (list of VendorObject objects): The Vendor param.
*/
type FirewallObject struct {
    Criteria *CriteriaObject3 `json:"criteria,omitempty"`
    ExcludeVendor bool `json:"exclude_vendor,omitempty"`
    Vendor []VendorObject `json:"vendor,omitempty"`
}

/*
HostInfoObject ""

Parent chains:
* "Config host_info"

Args:

Param Criteria (required, CriteriaObject6): The Criteria param. CriteriaObject6 instance.
*/
type HostInfoObject struct {
    Criteria CriteriaObject6 `json:"criteria"`
}

/*
CriteriaObject6 ""

Parent chains:
* "Config host_info criteria"

Args:

Param ClientVersion (ClientVersionObject): The ClientVersion param. ClientVersionObject instance.

Param Domain (DomainObject): The Domain param. DomainObject instance.

Param HostId (HostIdObject): The HostId param. HostIdObject instance.

Param HostName (HostNameObject): The HostName param. HostNameObject instance.

Param Managed (bool): The Managed param.

Param Os (OsObject): The Os param. OsObject instance.

Param SerialNumber (SerialNumberObject): The SerialNumber param. SerialNumberObject instance.
*/
type CriteriaObject6 struct {
    ClientVersion *ClientVersionObject `json:"client_version,omitempty"`
    Domain *DomainObject `json:"domain,omitempty"`
    HostId *HostIdObject `json:"host_id,omitempty"`
    HostName *HostNameObject `json:"host_name,omitempty"`
    Managed bool `json:"managed,omitempty"`
    Os *OsObject `json:"os,omitempty"`
    SerialNumber *SerialNumberObject `json:"serial_number,omitempty"`
}

/*
ClientVersionObject ""

Parent chains:
* "Config host_info criteria client_version"

Args:

Param Contains (string): The Contains param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.

Param Is (string): The Is param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.

Param IsNot (string): The IsNot param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.
*/
type ClientVersionObject struct {
    Contains string `json:"contains,omitempty"`
    Is string `json:"is,omitempty"`
    IsNot string `json:"is_not,omitempty"`
}

/*
DomainObject ""

Parent chains:
* "Config host_info criteria domain"

Args:

Param Contains (string): The Contains param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.

Param Is (string): The Is param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.

Param IsNot (string): The IsNot param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.
*/
type DomainObject struct {
    Contains string `json:"contains,omitempty"`
    Is string `json:"is,omitempty"`
    IsNot string `json:"is_not,omitempty"`
}

/*
HostIdObject ""

Parent chains:
* "Config host_info criteria host_id"

Args:

Param Contains (string): The Contains param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.

Param Is (string): The Is param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.

Param IsNot (string): The IsNot param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.
*/
type HostIdObject struct {
    Contains string `json:"contains,omitempty"`
    Is string `json:"is,omitempty"`
    IsNot string `json:"is_not,omitempty"`
}

/*
HostNameObject ""

Parent chains:
* "Config host_info criteria host_name"

Args:

Param Contains (string): The Contains param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.

Param Is (string): The Is param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.

Param IsNot (string): The IsNot param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.
*/
type HostNameObject struct {
    Contains string `json:"contains,omitempty"`
    Is string `json:"is,omitempty"`
    IsNot string `json:"is_not,omitempty"`
}

/*
OsObject ""

Parent chains:
* "Config host_info criteria os"

Args:

Param Contains (ContainsObject): The Contains param. ContainsObject instance. One of these params should be specified:  contains.
*/
type OsObject struct {
    Contains *ContainsObject `json:"contains,omitempty"`
}

/*
ContainsObject ""

Parent chains:
* "Config host_info criteria os contains"

Args:

Param Apple (string): The Apple param. String lengh must be between 0 and 255 characters. Default: "All". One of these params should be specified:  Apple, Google, Linux, Microsoft, or Other.

Param Google (string): The Google param. String lengh must be between 0 and 255 characters. Default: "All". One of these params should be specified:  Apple, Google, Linux, Microsoft, or Other.

Param Linux (string): The Linux param. String lengh must be between 0 and 255 characters. Default: "All". One of these params should be specified:  Apple, Google, Linux, Microsoft, or Other.

Param Microsoft (string): The Microsoft param. String lengh must be between 0 and 255 characters. Default: "All". One of these params should be specified:  Apple, Google, Linux, Microsoft, or Other.

Param Other (string): The Other param. String lengh must be between 0 and 255 characters. One of these params should be specified:  Apple, Google, Linux, Microsoft, or Other.
*/
type ContainsObject struct {
    Apple string `json:"Apple,omitempty"`
    Google string `json:"Google,omitempty"`
    Linux string `json:"Linux,omitempty"`
    Microsoft string `json:"Microsoft,omitempty"`
    Other string `json:"Other,omitempty"`
}

/*
SerialNumberObject ""

Parent chains:
* "Config host_info criteria serial_number"

Args:

Param Contains (string): The Contains param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.

Param Is (string): The Is param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.

Param IsNot (string): The IsNot param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.
*/
type SerialNumberObject struct {
    Contains string `json:"contains,omitempty"`
    Is string `json:"is,omitempty"`
    IsNot string `json:"is_not,omitempty"`
}

/*
MobileDeviceObject ""

Parent chains:
* "Config mobile_device"

Args:

Param Criteria (CriteriaObject7): The Criteria param. CriteriaObject7 instance.
*/
type MobileDeviceObject struct {
    Criteria *CriteriaObject7 `json:"criteria,omitempty"`
}

/*
CriteriaObject7 ""

Parent chains:
* "Config mobile_device criteria"

Args:

Param Applications (ApplicationsObject): The Applications param. ApplicationsObject instance.

Param DiskEncrypted (bool): The DiskEncrypted param.

Param Imei (ImeiObject): The Imei param. ImeiObject instance.

Param Jailbroken (bool): The Jailbroken param.

Param LastCheckinTime (LastCheckinTimeObject): The LastCheckinTime param. LastCheckinTimeObject instance.

Param Model (ModelObject): The Model param. ModelObject instance.

Param PasscodeSet (bool): The PasscodeSet param.

Param PhoneNumber (PhoneNumberObject): The PhoneNumber param. PhoneNumberObject instance.

Param Tag (TagObject): The Tag param. TagObject instance.
*/
type CriteriaObject7 struct {
    Applications *ApplicationsObject `json:"applications,omitempty"`
    DiskEncrypted bool `json:"disk_encrypted,omitempty"`
    Imei *ImeiObject `json:"imei,omitempty"`
    Jailbroken bool `json:"jailbroken,omitempty"`
    LastCheckinTime *LastCheckinTimeObject `json:"last_checkin_time,omitempty"`
    Model *ModelObject `json:"model,omitempty"`
    PasscodeSet bool `json:"passcode_set,omitempty"`
    PhoneNumber *PhoneNumberObject `json:"phone_number,omitempty"`
    Tag *TagObject `json:"tag,omitempty"`
}

/*
ApplicationsObject ""

Parent chains:
* "Config mobile_device criteria applications"

Args:

Param HasMalware (HasMalwareObject): The HasMalware param. HasMalwareObject instance.

Param HasUnmanagedApp (bool): The HasUnmanagedApp param.

Param Includes (list of IncludesObject objects): The Includes param.
*/
type ApplicationsObject struct {
    HasMalware *HasMalwareObject `json:"has_malware,omitempty"`
    HasUnmanagedApp bool `json:"has_unmanaged_app,omitempty"`
    Includes []IncludesObject `json:"includes,omitempty"`
}

/*
HasMalwareObject ""

Parent chains:
* "Config mobile_device criteria applications has_malware"

Args:

Param No (interface{}): The No param. interface{} instance.

Param Yes (YesObject): The Yes param. YesObject instance.
*/
type HasMalwareObject struct {
    No interface{} `json:"no,omitempty"`
    Yes *YesObject `json:"yes,omitempty"`
}

/*
YesObject ""

Parent chains:
* "Config mobile_device criteria applications has_malware yes"

Args:

Param Excludes (list of ExcludesObject objects): The Excludes param.
*/
type YesObject struct {
    Excludes []ExcludesObject `json:"excludes,omitempty"`
}

/*
ExcludesObject ""

Parent chains:
* "Config mobile_device criteria applications has_malware yes excludes"

Args:

Param Hash (string): The Hash param. String length must not exceed 1024 characters.

Param Name (required, string): The Name param. String length must not exceed 31 characters.

Param Package (string): The Package param. String length must not exceed 1024 characters.
*/
type ExcludesObject struct {
    Hash string `json:"hash,omitempty"`
    Name string `json:"name"`
    Package string `json:"package,omitempty"`
}

/*
IncludesObject ""

Parent chains:
* "Config mobile_device criteria applications includes"

Args:

Param Hash (string): The Hash param. String length must not exceed 1024 characters.

Param Name (required, string): The Name param. String length must not exceed 31 characters.

Param Package (string): The Package param. String length must not exceed 1024 characters.
*/
type IncludesObject struct {
    Hash string `json:"hash,omitempty"`
    Name string `json:"name"`
    Package string `json:"package,omitempty"`
}

/*
ImeiObject ""

Parent chains:
* "Config mobile_device criteria imei"

Args:

Param Contains (string): The Contains param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.

Param Is (string): The Is param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.

Param IsNot (string): The IsNot param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.
*/
type ImeiObject struct {
    Contains string `json:"contains,omitempty"`
    Is string `json:"is,omitempty"`
    IsNot string `json:"is_not,omitempty"`
}

/*
LastCheckinTimeObject ""

Parent chains:
* "Config mobile_device criteria last_checkin_time"

Args:

Param NotWithin (NotWithinObject3): The NotWithin param. NotWithinObject3 instance. One of these params should be specified:  not_within or within.

Param Within (WithinObject3): The Within param. WithinObject3 instance. One of these params should be specified:  not_within or within.
*/
type LastCheckinTimeObject struct {
    NotWithin *NotWithinObject3 `json:"not_within,omitempty"`
    Within *WithinObject3 `json:"within,omitempty"`
}

/*
NotWithinObject3 ""

Parent chains:
* "Config mobile_device criteria last_checkin_time not_within"

Args:

Param Days (required, int64): The Days param. Value must be between 1 and 365. Default: 30
*/
type NotWithinObject3 struct {
    Days int64 `json:"days"`
}

/*
WithinObject3 ""

Parent chains:
* "Config mobile_device criteria last_checkin_time within"

Args:

Param Days (required, int64): The Days param. Value must be between 1 and 365. Default: 30
*/
type WithinObject3 struct {
    Days int64 `json:"days"`
}

/*
ModelObject ""

Parent chains:
* "Config mobile_device criteria model"

Args:

Param Contains (string): The Contains param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.

Param Is (string): The Is param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.

Param IsNot (string): The IsNot param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.
*/
type ModelObject struct {
    Contains string `json:"contains,omitempty"`
    Is string `json:"is,omitempty"`
    IsNot string `json:"is_not,omitempty"`
}

/*
PhoneNumberObject ""

Parent chains:
* "Config mobile_device criteria phone_number"

Args:

Param Contains (string): The Contains param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.

Param Is (string): The Is param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.

Param IsNot (string): The IsNot param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.
*/
type PhoneNumberObject struct {
    Contains string `json:"contains,omitempty"`
    Is string `json:"is,omitempty"`
    IsNot string `json:"is_not,omitempty"`
}

/*
TagObject ""

Parent chains:
* "Config mobile_device criteria tag"

Args:

Param Contains (string): The Contains param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.

Param Is (string): The Is param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.

Param IsNot (string): The IsNot param. String lengh must be between 0 and 255 characters. One of these params should be specified:  contains, is, or is_not.
*/
type TagObject struct {
    Contains string `json:"contains,omitempty"`
    Is string `json:"is,omitempty"`
    IsNot string `json:"is_not,omitempty"`
}

/*
NetworkInfoObject ""

Parent chains:
* "Config network_info"

Args:

Param Criteria (CriteriaObject8): The Criteria param. CriteriaObject8 instance.
*/
type NetworkInfoObject struct {
    Criteria *CriteriaObject8 `json:"criteria,omitempty"`
}

/*
CriteriaObject8 ""

Parent chains:
* "Config network_info criteria"

Args:

Param Network (NetworkObject): The Network param. NetworkObject instance.
*/
type CriteriaObject8 struct {
    Network *NetworkObject `json:"network,omitempty"`
}

/*
NetworkObject ""

Parent chains:
* "Config network_info criteria network"

Args:

Param Is (IsObject): The Is param. IsObject instance.

Param IsNot (IsNotObject): The IsNot param. IsNotObject instance.
*/
type NetworkObject struct {
    Is *IsObject `json:"is,omitempty"`
    IsNot *IsNotObject `json:"is_not,omitempty"`
}

/*
IsObject ""

Parent chains:
* "Config network_info criteria network is"

Args:

Param Mobile (MobileObject): The Mobile param. MobileObject instance.

Param Unknown (interface{}): The Unknown param. interface{} instance.

Param Wifi (WifiObject): The Wifi param. WifiObject instance.
*/
type IsObject struct {
    Mobile *MobileObject `json:"mobile,omitempty"`
    Unknown interface{} `json:"unknown,omitempty"`
    Wifi *WifiObject `json:"wifi,omitempty"`
}

/*
MobileObject ""

Parent chains:
* "Config network_info criteria network is mobile"
* "Config network_info criteria network is_not mobile"

Args:

Param Carrier (string): The Carrier param. String length must not exceed 1023 characters.
*/
type MobileObject struct {
    Carrier string `json:"carrier,omitempty"`
}

/*
WifiObject ""

Parent chains:
* "Config network_info criteria network is wifi"
* "Config network_info criteria network is_not wifi"

Args:

Param Ssid (string): The Ssid param. String length must not exceed 1023 characters.
*/
type WifiObject struct {
    Ssid string `json:"ssid,omitempty"`
}

/*
IsNotObject ""

Parent chains:
* "Config network_info criteria network is_not"

Args:

Param Ethernet (interface{}): The Ethernet param. interface{} instance.

Param Mobile (MobileObject): The Mobile param. MobileObject instance.

Param Unknown (interface{}): The Unknown param. interface{} instance.

Param Wifi (WifiObject): The Wifi param. WifiObject instance.
*/
type IsNotObject struct {
    Ethernet interface{} `json:"ethernet,omitempty"`
    Mobile *MobileObject `json:"mobile,omitempty"`
    Unknown interface{} `json:"unknown,omitempty"`
    Wifi *WifiObject `json:"wifi,omitempty"`
}

/*
PatchManagementObject ""

Parent chains:
* "Config patch_management"

Args:

Param Criteria (CriteriaObject9): The Criteria param. CriteriaObject9 instance.

Param ExcludeVendor (bool): The ExcludeVendor param. Default: false

Param Vendor (list of VendorObject1 objects): The Vendor param.
*/
type PatchManagementObject struct {
    Criteria *CriteriaObject9 `json:"criteria,omitempty"`
    ExcludeVendor bool `json:"exclude_vendor,omitempty"`
    Vendor []VendorObject1 `json:"vendor,omitempty"`
}

/*
CriteriaObject9 ""

Parent chains:
* "Config patch_management criteria"

Args:

Param IsEnabled (string): The IsEnabled param. String values: []string{"no", "yes", "not-available"}

Param IsInstalled (bool): The IsInstalled param. Default: true

Param MissingPatches (MissingPatchesObject): The MissingPatches param. MissingPatchesObject instance.
*/
type CriteriaObject9 struct {
    IsEnabled string `json:"is_enabled,omitempty"`
    IsInstalled bool `json:"is_installed,omitempty"`
    MissingPatches *MissingPatchesObject `json:"missing_patches,omitempty"`
}

/*
MissingPatchesObject ""

Parent chains:
* "Config patch_management criteria missing_patches"

Args:

Param Check (required, string): The Check param. String values: []string{"has-any", "has-none", "has-all"} Default: "has-any".

Param Patches (list of strings): The Patches param.

Param Severity (SeverityObject): The Severity param. SeverityObject instance.
*/
type MissingPatchesObject struct {
    Check string `json:"check"`
    Patches []string `json:"patches,omitempty"`
    Severity *SeverityObject `json:"severity,omitempty"`
}

/*
SeverityObject ""

Parent chains:
* "Config patch_management criteria missing_patches severity"

Args:

Param GreaterEqual (int64): The GreaterEqual param. Value must be between 0 and 100000. One of these params should be specified:  greater_equal, greater_than, is, is_not, less_equal, or less_than.

Param GreaterThan (int64): The GreaterThan param. Value must be between 0 and 100000. One of these params should be specified:  greater_equal, greater_than, is, is_not, less_equal, or less_than.

Param Is (int64): The Is param. Value must be between 0 and 100000. One of these params should be specified:  greater_equal, greater_than, is, is_not, less_equal, or less_than.

Param IsNot (int64): The IsNot param. Value must be between 0 and 100000. One of these params should be specified:  greater_equal, greater_than, is, is_not, less_equal, or less_than.

Param LessEqual (int64): The LessEqual param. Value must be between 0 and 100000. One of these params should be specified:  greater_equal, greater_than, is, is_not, less_equal, or less_than.

Param LessThan (int64): The LessThan param. Value must be between 0 and 100000. One of these params should be specified:  greater_equal, greater_than, is, is_not, less_equal, or less_than.
*/
type SeverityObject struct {
    GreaterEqual int64 `json:"greater_equal,omitempty"`
    GreaterThan int64 `json:"greater_than,omitempty"`
    Is int64 `json:"is,omitempty"`
    IsNot int64 `json:"is_not,omitempty"`
    LessEqual int64 `json:"less_equal,omitempty"`
    LessThan int64 `json:"less_than,omitempty"`
}