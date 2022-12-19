package main

type comandID int

const (
	CMD_NICK comandID = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

type command struct {
	id     comandID
	client *Client
	args   []string
}
