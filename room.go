package main

import "net"

type room struct {
	name    string
	members map[net.Addr]*Client
}

func (r *room) broadcast(sender *Client, msg string) {
	for addr, m := range r.members {
		if addr != sender.conn.RemoteAddr() {
			m.msg(msg)
		}
	}
}
