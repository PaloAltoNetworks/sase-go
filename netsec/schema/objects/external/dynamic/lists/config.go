package lists

/*
Config ""

Parent chains:
* "Config"

Args:

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param. String length must not exceed 63 characters.

Param Type (required, TypeObject): The Type param. TypeObject instance.
*/
type Config struct {
    ObjectId string `json:"id,omitempty"`
    Name string `json:"name"`
    Type TypeObject `json:"type"`
}

/*
TypeObject ""

Parent chains:
* "Config type"

Args:

Param Domain (DomainObject): The Domain param. DomainObject instance.

Param Imei (ImeiObject): The Imei param. ImeiObject instance.

Param Imsi (ImsiObject): The Imsi param. ImsiObject instance.

Param Ip (IpObject): The Ip param. IpObject instance.

Param PredefinedIp (PredefinedIpObject): The PredefinedIp param. PredefinedIpObject instance.

Param PredefinedUrl (PredefinedUrlObject): The PredefinedUrl param. PredefinedUrlObject instance.

Param Url (UrlObject): The Url param. UrlObject instance.
*/
type TypeObject struct {
    Domain *DomainObject `json:"domain,omitempty"`
    Imei *ImeiObject `json:"imei,omitempty"`
    Imsi *ImsiObject `json:"imsi,omitempty"`
    Ip *IpObject `json:"ip,omitempty"`
    PredefinedIp *PredefinedIpObject `json:"predefined_ip,omitempty"`
    PredefinedUrl *PredefinedUrlObject `json:"predefined_url,omitempty"`
    Url *UrlObject `json:"url,omitempty"`
}

/*
DomainObject ""

Parent chains:
* "Config type domain"

Args:

Param Auth (AuthObject): The Auth param. AuthObject instance.

Param CertificateProfile (string): The CertificateProfile param. Default: "None".

Param Description (string): The Description param. String lengh must be between 0 and 255 characters.

Param ExceptionList (list of strings): The ExceptionList param.

Param ExpandDomain (bool): The ExpandDomain param. Default: false

Param Recurring (required, RecurringObject): The Recurring param. RecurringObject instance.

Param Url (required, string): The Url param. String lengh must be between 0 and 255 characters. Default: "http://".
*/
type DomainObject struct {
    Auth *AuthObject `json:"auth,omitempty"`
    CertificateProfile string `json:"certificate_profile,omitempty"`
    Description string `json:"description,omitempty"`
    ExceptionList []string `json:"exception_list,omitempty"`
    ExpandDomain bool `json:"expand_domain,omitempty"`
    Recurring RecurringObject `json:"recurring"`
    Url string `json:"url"`
}

/*
AuthObject ""

Parent chains:
* "Config type domain auth"
* "Config type imei auth"
* "Config type imsi auth"
* "Config type ip auth"
* "Config type url auth"

Args:

Param Password (required, string): The Password param. String length must not exceed 255 characters.

Param Username (required, string): The Username param. String lengh must be between 1 and 255 characters.
*/
type AuthObject struct {
    Password string `json:"password"`
    Username string `json:"username"`
}

/*
RecurringObject ""

Parent chains:
* "Config type domain recurring"
* "Config type imei recurring"
* "Config type imsi recurring"
* "Config type ip recurring"
* "Config type url recurring"

Args:

Param Daily (DailyObject): The Daily param. DailyObject instance. One of these params should be specified:  daily, five_minute, hourly, monthly, or weekly.

Param FiveMinute (interface{}): The FiveMinute param. interface{} instance. One of these params should be specified:  daily, five_minute, hourly, monthly, or weekly.

Param Hourly (interface{}): The Hourly param. interface{} instance. One of these params should be specified:  daily, five_minute, hourly, monthly, or weekly.

Param Monthly (MonthlyObject): The Monthly param. MonthlyObject instance. One of these params should be specified:  daily, five_minute, hourly, monthly, or weekly.

Param Weekly (WeeklyObject): The Weekly param. WeeklyObject instance. One of these params should be specified:  daily, five_minute, hourly, monthly, or weekly.
*/
type RecurringObject struct {
    Daily *DailyObject `json:"daily,omitempty"`
    FiveMinute interface{} `json:"five_minute,omitempty"`
    Hourly interface{} `json:"hourly,omitempty"`
    Monthly *MonthlyObject `json:"monthly,omitempty"`
    Weekly *WeeklyObject `json:"weekly,omitempty"`
}

/*
DailyObject ""

Parent chains:
* "Config type domain recurring daily"

Args:

Param At (required, string): The At param. String lengh must be between 2 and 2 characters. Default: "00".
*/
type DailyObject struct {
    At string `json:"at"`
}

/*
MonthlyObject ""

Parent chains:
* "Config type domain recurring monthly"

Args:

Param At (required, string): The At param. String lengh must be between 2 and 2 characters. Default: "00".

Param DayOfMonth (required, int64): The DayOfMonth param. Value must be between 1 and 31.
*/
type MonthlyObject struct {
    At string `json:"at"`
    DayOfMonth int64 `json:"day_of_month"`
}

/*
WeeklyObject ""

Parent chains:
* "Config type domain recurring weekly"

Args:

Param At (required, string): The At param. String lengh must be between 2 and 2 characters. Default: "00".

Param DayOfWeek (required, string): The DayOfWeek param. String values: []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
*/
type WeeklyObject struct {
    At string `json:"at"`
    DayOfWeek string `json:"day_of_week"`
}

/*
ImeiObject ""

Parent chains:
* "Config type imei"

Args:

Param Auth (AuthObject): The Auth param. AuthObject instance.

Param CertificateProfile (string): The CertificateProfile param. Default: "None".

Param Description (string): The Description param. String lengh must be between 0 and 255 characters.

Param ExceptionList (list of strings): The ExceptionList param.

Param Recurring (required, RecurringObject): The Recurring param. RecurringObject instance.

Param Url (required, string): The Url param. String lengh must be between 0 and 255 characters. Default: "http://".
*/
type ImeiObject struct {
    Auth *AuthObject `json:"auth,omitempty"`
    CertificateProfile string `json:"certificate_profile,omitempty"`
    Description string `json:"description,omitempty"`
    ExceptionList []string `json:"exception_list,omitempty"`
    Recurring RecurringObject `json:"recurring"`
    Url string `json:"url"`
}

/*
ImsiObject ""

Parent chains:
* "Config type imsi"

Args:

Param Auth (AuthObject): The Auth param. AuthObject instance.

Param CertificateProfile (string): The CertificateProfile param. Default: "None".

Param Description (string): The Description param. String lengh must be between 0 and 255 characters.

Param ExceptionList (list of strings): The ExceptionList param.

Param Recurring (required, RecurringObject): The Recurring param. RecurringObject instance.

Param Url (required, string): The Url param. String lengh must be between 0 and 255 characters. Default: "http://".
*/
type ImsiObject struct {
    Auth *AuthObject `json:"auth,omitempty"`
    CertificateProfile string `json:"certificate_profile,omitempty"`
    Description string `json:"description,omitempty"`
    ExceptionList []string `json:"exception_list,omitempty"`
    Recurring RecurringObject `json:"recurring"`
    Url string `json:"url"`
}

/*
IpObject ""

Parent chains:
* "Config type ip"

Args:

Param Auth (AuthObject): The Auth param. AuthObject instance.

Param CertificateProfile (string): The CertificateProfile param. Default: "None".

Param Description (string): The Description param. String lengh must be between 0 and 255 characters.

Param ExceptionList (list of strings): The ExceptionList param.

Param Recurring (required, RecurringObject): The Recurring param. RecurringObject instance.

Param Url (required, string): The Url param. String lengh must be between 0 and 255 characters. Default: "http://".
*/
type IpObject struct {
    Auth *AuthObject `json:"auth,omitempty"`
    CertificateProfile string `json:"certificate_profile,omitempty"`
    Description string `json:"description,omitempty"`
    ExceptionList []string `json:"exception_list,omitempty"`
    Recurring RecurringObject `json:"recurring"`
    Url string `json:"url"`
}

/*
PredefinedIpObject ""

Parent chains:
* "Config type predefined_ip"

Args:

Param Description (string): The Description param. String lengh must be between 0 and 255 characters.

Param ExceptionList (list of strings): The ExceptionList param.

Param Url (required, string): The Url param.
*/
type PredefinedIpObject struct {
    Description string `json:"description,omitempty"`
    ExceptionList []string `json:"exception_list,omitempty"`
    Url string `json:"url"`
}

/*
PredefinedUrlObject ""

Parent chains:
* "Config type predefined_url"

Args:

Param Description (string): The Description param. String lengh must be between 0 and 255 characters.

Param ExceptionList (list of strings): The ExceptionList param.

Param Url (required, string): The Url param.
*/
type PredefinedUrlObject struct {
    Description string `json:"description,omitempty"`
    ExceptionList []string `json:"exception_list,omitempty"`
    Url string `json:"url"`
}

/*
UrlObject ""

Parent chains:
* "Config type url"

Args:

Param Auth (AuthObject): The Auth param. AuthObject instance.

Param CertificateProfile (string): The CertificateProfile param. Default: "None".

Param Description (string): The Description param. String lengh must be between 0 and 255 characters.

Param ExceptionList (list of strings): The ExceptionList param.

Param Recurring (required, RecurringObject): The Recurring param. RecurringObject instance.

Param Url (required, string): The Url param. String lengh must be between 0 and 255 characters. Default: "http://".
*/
type UrlObject struct {
    Auth *AuthObject `json:"auth,omitempty"`
    CertificateProfile string `json:"certificate_profile,omitempty"`
    Description string `json:"description,omitempty"`
    ExceptionList []string `json:"exception_list,omitempty"`
    Recurring RecurringObject `json:"recurring"`
    Url string `json:"url"`
}