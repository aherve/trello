// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"fmt"
	"time"
)

// List represents Trello lists.
// https://developers.trello.com/reference/#list-object
type List struct {
	client  *Client
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	IDBoard string  `json:"idBoard,omitempty"`
	Closed  bool    `json:"closed"`
	Pos     float32 `json:"pos,omitempty"`
	Board   *Board  `json:"board,omitempty"`
	Cards   []*Card `json:"cards,omitempty"`
}

// CreatedAt returns the time.Time from the list's id.
func (l *List) CreatedAt() time.Time {
	t, _ := IDToTime(l.ID)
	return t
}

// GetList takes a list's id and Arguments and returns the matching list.
func (c *Client) GetList(listID string, args Arguments) (list *List, err error) {
	path := fmt.Sprintf("lists/%s", listID)
	err = c.Get(path, args, &list)
	if list != nil {
		list.client = c
		for i := range list.Cards {
			list.Cards[i].client = c
		}
	}
	return
}

// GetLists takes Arguments and returns the lists of the receiver Board.
func (b *Board) GetLists(args Arguments) (lists []*List, err error) {
	path := fmt.Sprintf("boards/%s/lists", b.ID)
	err = b.client.Get(path, args, &lists)
	for i := range lists {
		lists[i].client = b.client
		for j := range lists[i].Cards {
			lists[i].Cards[j].client = b.client
		}
	}
	return
}
