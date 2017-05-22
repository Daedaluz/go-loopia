package loopia

import(
	"github.com/kolo/xmlrpc"
)

const(
	SERVER = "https://api.loopia.se/RPCSERV"
)

type Client struct {
	username string
	password string
	client *xmlrpc.Client
}

func NewClient(username, password string) *Client {
	cli, _ := xmlrpc.NewClient(SERVER, nil)
	return &Client {
		username: username,
		password: password,
		client: cli,
	}
}

const (
	TYPE_DOMAIN = "LOOPIADOMAIN"
	TYPE_DNS = "LOOPIADNS"
	TYPE_PRIVATE = "HOSTING_PRIVATE"
	TYPE_BUSSINESS = "HOSTING_BUSSINESS"
	TYPE_PLUS = "HOSTING_BUSSINESS_PLUS"

	STATUS_OK = "OK"
	STATUS_ERROR = "AUTH_ERROR"
	STATUS_OCCUPIED = "DOMAIN_OCCUPIED"
	STATUS_LIMITED = "RATE_LIMITED"
	STATUS_BAD = "BAD_INDATA"
	STATUS_UNKNOWN = "UNKNOWN_ERROR"
	STATUS_FUNDS = "INSUFFICIENT_FUNDS"

	CONFIG_NOCONFIG = "NO_CONFIG"
	CONFIG_PARKING = "PARKING"
	CONFIG_UNIX = "HOSTING_UNIX"
	CONFIG_AUTOBAHN = "HOSTING_AUTOBAHN"
	CONFIG_WINDOWS = "HOSTING_WINDOWS"

	RENEWAL_NORMAL = "NORMAL"
	RENEWAL_DEACTIVATED ="DEACTIVATED"
	RENEWAL_NOT_LOOPIA = "NOT_HANDLED_BY_LOOPIA"

	ORDER_DELETED = "DELETED"
	ORDER_PENDING = "PENDING"
	ORDER_PROCESSED = "PROCESSED"
)

type Contact struct {
	FirstName string
	LastName string
	Company string
	Street string
	Street2 string
	Zip string
	City string
	Country string
	Orgno string
	Phone string
	Cell string
	Fax string
	Email string
}

type Domain struct {
	Domain string `xmlrpc:"domain"`
	Paid int `xmlrpc:"paid"`
	Registered int `xmlrpc:"registered"`
	Renewal string `xmlrpc:"renewal_status"`
	Expiration string `xmlrpc:"expiration_date"`
	Reference int `xmlrpc:"reference_no"`
}

type Record struct {
	Type string `xmlrpc:"type"`
	TTL int `xmlrpc:"ttl"`
	Priority int `xmlrpc:"priority"`
	Rdata string `xmlrpc:"rdata"`
	RecordId int `xmlrpc:"record_id"`
}

func (c *Client) GetDomain(domain string) Domain {
	result := Domain{}
	c.client.Call("getDomain", []interface{}{c.username, c.password, "", domain}, &result)
	return result
}

func (c *Client) GetDomains() []Domain {
	result := make([]Domain, 0, 10)
	c.client.Call("getDomains", []interface{}{c.username, c.password}, &result)
	return result
}

func (c *Client) GetSubdomains(domain string) []string {
	result := make([]string, 0, 10)
	c.client.Call("getSubdomains", []interface{}{c.username, c.password, "", domain}, &result)
	return result
}

func (c *Client) AddSubdomain(subdomain, domain string) string {
	var result string
	c.client.Call("addSubdomain", []interface{}{c.username, c.password, "", domain, subdomain}, &result)
	return result
}

func (c *Client) RemoveSubdomain(subdomain, domain string) string {
	var result string
	c.client.Call("removeSubdomain", []interface{}{c.username, c.password, "", domain, subdomain}, &result)
	return result
}

func (c *Client) AddZoneRecord(subdomain, domain string, record *Record) string {
	var result string
	c.client.Call("addZoneRecord", []interface{}{c.username, c.password, "", domain, subdomain, record}, &result)
	return result
}

func (c *Client) UpdateZoneRecord(subdomain, domain string, record *Record) string {
	var result string
	c.client.Call("updateZoneRecord", []interface{}{c.username, c.password, "", domain, subdomain, record}, &result)
	return result
}

func (c *Client) RemoveZoneRecord(subdomain, domain string, record_id int) string {
	var result string
	c.client.Call("removeZoneRecord", []interface{}{c.username, c.password, "", domain, subdomain, record_id}, &result)
	return result
}

func (c *Client) GetZoneRecords(subdomain, domain string) []Record {
	result := make([]Record, 0, 10)
	c.client.Call("getZoneRecords", []interface{}{c.username, c.password, "", domain, subdomain}, &result)
	return result
}
