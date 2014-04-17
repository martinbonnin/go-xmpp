// Copyright 2013 Flo Lauber <dev@qatfy.at>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO(flo):
//   - support password protected MUC rooms
//   - cleanup signatures of join/leave functions
package xmpp

import (
	"fmt"
)

const (
	nsMUC     = "http://jabber.org/protocol/muc"
	nsMUCUser = "http://jabber.org/protocol/muc#user"
)

// xep-0045 7.2
func (c *Client) JoinMUC(jid string) {
	fmt.Fprintf(c.conn, "<presence to='%s' id='99'>\n"+
		"<x xmlns='%s'><history maxchars=\"0\" maxstanzas=\"0\" /></x>\n"+
		"</presence>",
		xmlEscape(jid), nsMUC)
}

// xep-0045 7.14
func (c *Client) LeaveMUC(jid string) {
	fmt.Fprintf(c.conn, "<presence from='%s' to='%s' type='unavailable' />",
		c.jid, xmlEscape(jid))
}
