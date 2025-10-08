package ofprepos

import "ego.dev21/greetings/internal/entities"

type OfpRepository interface {
	GetOFPInfoById(id int)
	CreateOFPInfo(ofp entities.OFP)
	DeleteOFPInfoById(id int)
	GetAllOFPInfo()
}
