/* Go module for Collection+JSON

Version: MPL 1.1/GPL 2.0/LGPL 2.1

The contents of this file are subject to the Mozilla Public License Version
1.1 (the "License"); you may not use this file except in compliance with
the License. You may obtain a copy of the License at
http://www.mozilla.org/MPL/

Software distributed under the License is distributed on an "AS IS" basis,
WITHOUT WARRANTY OF ANY KIND, either express or implied. See the License
for the specific language governing rights and limitations under the
License.

The Initial Developer of the Original Code is
Mozilla Corporation
Portions created by the Initial Developer are Copyright (C) 2014
the Initial Developer. All Rights Reserved.

Contributor(s):
Julien Vehent jvehent@mozilla.com [:ulfr]

Alternatively, the contents of this file may be used under the terms of
either the GNU General Public License Version 2 or later (the "GPL"), or
the GNU Lesser General Public License Version 2.1 or later (the "LGPL"),
in which case the provisions of the GPL or the LGPL are applicable instead
of those above. If you wish to allow use of your version of this file only
under the terms of either the GPL or the LGPL, and not to allow others to
use your version of this file under the terms of the MPL, indicate your
decision by deleting the provisions above and replace them with the notice
and other provisions required by the GPL or the LGPL. If you do not delete
the provisions above, a recipient may use your version of this file under
the terms of any one of the MPL, the GPL or the LGPL.
*/

package cljs

import (
	"fmt"
)

type Query struct {
	Rel    string `json:"rel"`              //required
	Href   string `json:"href"`             //required
	Name   string `json:"name,omitempty"`   //optional
	Prompt string `json:"prompt,omitempty"` //optional
	Data   []Data `json:"data,omitempty"`   //optional
}

func (r *Resource) AddQuery(query Query) (err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// syntax checking
	err = query.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate query: %v", err)
	}
	// allocate in the resource if not exist
	r.Collection.Queries = make([]Query, 0)
	var tmpqueries []Query
	tmpqueries = r.Collection.Queries
	tmpqueries = append(tmpqueries, query)
	r.Collection.Queries = tmpqueries
	return
}

func (query Query) Validate() (err error) {
	if query.Rel == "" {
		return fmt.Errorf("'rel' attr is empty")
	}
	if query.Href == "" {
		return fmt.Errorf("'href' attr is empty")
	}
	return
}
