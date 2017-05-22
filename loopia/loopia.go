package main

import(
	"fmt"
	"github.com/daedaluz/go-loopia"
	"strings"
	"os"
	"strconv"
	"text/tabwriter"
)

var help = `
loopia
    list   [subdomain.]domain.td
    add    subdomain.domain.td [ip]
    del    subdomain.domain.td
    zadd   subdomain.domain.td type data [ttl]
    zdel   subdomain.domain.td id
`

func splitname(name string) (string, string) {
	parts := strings.Split(name, ".")
	domains := parts[len(parts)-2:]
	hosts := parts[:len(parts)-2]
	domain := strings.Join(domains, ".")
	host := strings.Join(hosts, ".")
	return host, domain
}
var(
	USERNAME = "dummy"
	PASSWORD = "dummy"
)

func main() {

	w := tabwriter.NewWriter(os.Stdout, 3, 4, 2, ' ', 0)

	api := loopia.NewClient("tux@loopiaapi", "tux linux")
	args := os.Args
	if len(args) < 3 {
		fmt.Println(help)
		return
	}
	switch args[1] {
	case "list":
		subd, domain := splitname(args[2])
		if subd == "" {
			fmt.Printf("%s:\n", domain)
			subdomains := api.GetSubdomains(domain)
			for _, subdomain := range subdomains {
				fmt.Printf("  %s.%s\n", subdomain, domain)
			}
			return
		} else {
			fmt.Println(subd, domain)
			zones := api.GetZoneRecords(subd, domain)
			fmt.Fprintf(w, "  %s\t%s\t%s\t%s\t%s\n", "Type", "TTL", "Prio", "Rdata", "RecordId")
			fmt.Fprintf(w, "  %s\t%s\t%s\t%s\t%s\n", "----", "----", "----", "-----", "--------")
			for _, zone := range(zones) {
				fmt.Fprintf(w, "  %s\t%d\t%d\t%s\t%d\n", zone.Type, zone.TTL, zone.Priority, zone.Rdata, zone.RecordId)
			}
			w.Flush()
		}
	case "add":
		subd, domain := splitname(args[2])
		fmt.Println("Adding subdomain:")
		fmt.Println(api.AddSubdomain(subd, domain))
		if len(args) >= 3 {
			fmt.Println("Adding record:")
			fmt.Println(api.AddZoneRecord(subd, domain, &loopia.Record{
					Type: "A",
					TTL: 3600,
					Priority: 0,
					Rdata: args[3],
			}))
		}

	case "del":
		subd, domain := splitname(args[2])
		fmt.Println(api.RemoveSubdomain(subd, domain))

	case "zadd":
		subd, domain := splitname(args[2])
		record := loopia.Record {
			Type: args[3],
			TTL: 3600,
			Priority: 0,
			Rdata: args[4],
		}
		if len(args) >= 6 {
			ttl, err := strconv.ParseInt(args[6], 10, 64)
			if err != nil{
				fmt.Println(help)
				return
			}
			record.TTL = int(ttl)
		}
		fmt.Println(api.AddZoneRecord(subd, domain, &record))

	case "zdel":
		if len(args) < 4 {
			fmt.Println(help)
			return
		}
		subd, domain := splitname(args[2])
		id, _ := strconv.ParseInt(args[3], 10, 64)
		if id == 0 {
			fmt.Println(help)
			return
		}
		fmt.Println(api.RemoveZoneRecord(subd, domain, int(id)))
	default:
		fmt.Println(help)
		return
	}
}
