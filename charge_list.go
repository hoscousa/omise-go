/*
!!! DO NOT MODIFY !!!

autogenerated
on: Thu Nov 05 02:49:57 +0000 2015
by: chakrit
*/

package omise

type ChargeList struct {
	List
  Data []*Charge `json:"data"`
}

func (list *ChargeList) Find(id string) *Charge {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}
