// Factory for DNS providers
package dns

import (
	"fmt"

	"github.com/jtuchscherer/lego/acme"
	"github.com/jtuchscherer/lego/providers/dns/auroradns"
	"github.com/jtuchscherer/lego/providers/dns/cloudflare"
	"github.com/jtuchscherer/lego/providers/dns/cloudxns"
	"github.com/jtuchscherer/lego/providers/dns/digitalocean"
	"github.com/jtuchscherer/lego/providers/dns/dnsimple"
	"github.com/jtuchscherer/lego/providers/dns/dnsmadeeasy"
	"github.com/jtuchscherer/lego/providers/dns/dnspod"
	"github.com/jtuchscherer/lego/providers/dns/dyn"
	"github.com/jtuchscherer/lego/providers/dns/exoscale"
	"github.com/jtuchscherer/lego/providers/dns/gandi"
	"github.com/jtuchscherer/lego/providers/dns/gandiv5"
	"github.com/jtuchscherer/lego/providers/dns/googlecloud"
	"github.com/jtuchscherer/lego/providers/dns/godaddy"
	"github.com/jtuchscherer/lego/providers/dns/linode"
	"github.com/jtuchscherer/lego/providers/dns/namecheap"
	"github.com/jtuchscherer/lego/providers/dns/ns1"
	"github.com/jtuchscherer/lego/providers/dns/otc"
	"github.com/jtuchscherer/lego/providers/dns/ovh"
	"github.com/jtuchscherer/lego/providers/dns/pdns"
	"github.com/jtuchscherer/lego/providers/dns/rackspace"
	"github.com/jtuchscherer/lego/providers/dns/rfc2136"
	"github.com/jtuchscherer/lego/providers/dns/route53"
	"github.com/jtuchscherer/lego/providers/dns/vultr"
)

func NewDNSChallengeProviderByName(name string) (acme.ChallengeProvider, error) {
	var err error
	var provider acme.ChallengeProvider
	switch name {
	case "auroradns":
		provider, err = auroradns.NewDNSProvider()
	case "cloudflare":
		provider, err = cloudflare.NewDNSProvider()
	case "cloudxns":
		provider, err = cloudxns.NewDNSProvider()
	case "digitalocean":
		provider, err = digitalocean.NewDNSProvider()
	case "dnsimple":
		provider, err = dnsimple.NewDNSProvider()
	case "dnsmadeeasy":
		provider, err = dnsmadeeasy.NewDNSProvider()
	case "dnspod":
		provider, err = dnspod.NewDNSProvider()
	case "dyn":
		provider, err = dyn.NewDNSProvider()
	case "exoscale":
		provider, err = exoscale.NewDNSProvider()
	case "gandi":
		provider, err = gandi.NewDNSProvider()
	case "gandiv5":
		provider, err = gandiv5.NewDNSProvider()
	case "gcloud":
		provider, err = googlecloud.NewDNSProvider()
	case "godaddy":
		provider, err = godaddy.NewDNSProvider()
	case "linode":
		provider, err = linode.NewDNSProvider()
	case "manual":
		provider, err = acme.NewDNSProviderManual()
	case "namecheap":
		provider, err = namecheap.NewDNSProvider()
	case "rackspace":
		provider, err = rackspace.NewDNSProvider()
	case "route53":
		provider, err = route53.NewDNSProvider()
	case "rfc2136":
		provider, err = rfc2136.NewDNSProvider()
	case "vultr":
		provider, err = vultr.NewDNSProvider()
	case "ovh":
		provider, err = ovh.NewDNSProvider()
	case "pdns":
		provider, err = pdns.NewDNSProvider()
	case "ns1":
		provider, err = ns1.NewDNSProvider()
	case "otc":
		provider, err = otc.NewDNSProvider()
	default:
		err = fmt.Errorf("Unrecognised DNS provider: %s", name)
	}
	return provider, err
}
