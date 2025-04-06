package nft

const (
	// This is our file. 'enable' will include it into main
	Default_Rules=[]string{"meta l4proto icmp icmp type echo-request limit rate over 10/second burst 4 packets drop comment ","meta l4proto ipv6-icmp icmpv6 type echo-request limit rate over 10/second burst 4 packets drop comment "}
,"tcp sport ssh ct state new limit rate 15/minute accept comment ",
"meta l4proto { icmp, ipv6-icmp } accept comment",
"ip protocol igmp accept comment" )


