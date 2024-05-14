package enum

const (
	BillTypeAccommodation  = iota //住宿
	BillTypeRepast                //餐饮
	BillTypeTrip                  //出行
	BillTypeBusinessDinner        //应酬
	BillTypeProcurement           //采购
	BillTypeTeamActivities        //团建
)

func BillTypeToString(billType int) string {
	switch billType {
	case BillTypeAccommodation:
		return "住宿"
	case BillTypeRepast:
		return "餐饮"
	case BillTypeTrip:
		return "出行"
	case BillTypeBusinessDinner:
		return "应酬"
	case BillTypeProcurement:
		return "采购"
	case BillTypeTeamActivities:
		return "团建"
	}
	return ""
}
