package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"

	agg "github.com/ldkingvivi/go-aggregate"
	"github.com/p4gefau1t/trojan-go/common/geodata"
	// "github.com/v2fly/v2ray-core/v4/app/router"
)


func test1(list []*net.IPNet) []*net.IPNet {
	cidrs := make([]agg.CidrEntry, 0, 35000)
	for _, v := range list {
		
		cidrs = append(cidrs, agg.NewBasicCidrEntry(v))
	}

	// empty merge func will do the basic merge
	result := agg.Aggregate(cidrs, func(_, _ agg.CidrEntry) {})
	cidrList := make([]*net.IPNet, 0, 20000)
	for _, cidr := range result {
		cidrList = append(cidrList, cidr.GetNetwork())

	}
	return cidrList
}

func main() {
	// _, err := geodata.Decode("/goProject/geoip/geoip-2.dat", "cn")

	// if err != nil {
	// 	fmt.Println("++++++++++",err.Error())
	// }

	var filepath string
	var mode string
	var coutry string
	var site string
	var vmode string
	flag.StringVar(&filepath,"f","","filePath")
	flag.StringVar(&mode,"m","ip","Mode ip or site")
	flag.StringVar(&coutry,"c","cn","Country")
	flag.StringVar(&site,"s","gfw","Site-category if mode is site It's must like gfw、cn、 google、category-ads-all、geolocation-!cn、facebook..... ")
	flag.StringVar(&vmode,"v","all","4 for ipv4  6 for ipv6 ")
	flag.Parse()
	if mode=="ip"{
		c := geodata.NewGeodataLoader()
		c2, err := c.LoadIP(filepath, coutry)
		if err != nil {
			fmt.Println("---------", err.Error())
		}
		list :=make([]*net.IPNet,0,30000)
		for _, v := range c2 {  
			ll:=net.IPNet{
				IP: (net.IP)(v.GetIp()),
				Mask: []byte{ uint8( v.GetPrefix())},
			}
		
			list = append(list, &ll)
		}
		resList:=make([]string,0,350000)
		for _, v := range list {
			s := strconv.Itoa(int(v.Mask[0]))
				res:=v.IP.String()+"/"+s
					resList = append(resList, res)
		}
		switch vmode{
		case "all":
			for _, v := range resList {
				fmt.Println(v)
			}
		case "4":
			for _, v := range resList {
				if !strings.Contains(v,":"){
					fmt.Println(v)
				}
			}
		case "6":
			for _, v := range resList {
				if strings.Contains(v,":"){
					fmt.Println(v)
				}
			}
		}
	}else{
		c:=geodata.NewGeodataLoader()
			d, err := c.LoadSite(filepath, site)
	if err != nil {
		fmt.Println("---------", err.Error())
	}
	for _, v := range d {
		if !(strings.Contains(v.Value,"\\") || strings.Contains(v.Value,"?")){
			fmt.Println(v.Value)
		}
		
	}
	}
	




	// d, err := c.LoadSite("/goProject/geoip/geosite.dat", "gfw")
	// if err != nil {
	// 	fmt.Println("---------", err.Error())
	// }
	// for _, v := range d {
	// 	fmt.Println(v)
	// }
	
	// newDomain:=router.Domain{
	// 	Type: router.Domain_Full,
	// 	Value: "okok.ok",
	// }
	// d = append(d, &newDomain)
	// dm, err := router.NewDomainMatcher(d)

	// b := dm.ApplyDomain("okok.ok")
	// fmt.Println(b)

}
