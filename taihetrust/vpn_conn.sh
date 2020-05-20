#!/usr/bin/expect

# nohup /home/wangch/mytest/work/vpn_conn.sh > /dev/null 2>&1

set username "wangch@taihetrust.com"
set password "qwe123\`"
set timeout 40

spawn sudo openvpn /etc/openvpn/taihetrust.ovpn
expect {
	"* 的密码*" {
		send "$password\r\n"
		exp_continue
	}
	"*yes/no*" {
		send "yes\r\n"
		exp_continue
	}
	"Enter Auth Username:" {
		send "$username\r"
		exp_continue
	}
	"Enter Auth Password:" {
		send "$password\r"
	}
}
interact
