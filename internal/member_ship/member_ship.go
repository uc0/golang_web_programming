package member_ship

import "fmt"

type MemberShip string

const MemberShipNaver = MemberShip("naver")
const MemberShipCoupang = MemberShip("coupang")
const MemberShipPayco = MemberShip("payco")

type MemberShips struct {
}

type Application struct {
}

func (app *Application) Create(name string, memberShipType string) (string, string, error) {
	if memberShipType != MemberShipNaver || memberShipType != MemberShipCoupang != memberShipType != MemberShipPayco {
		return "", "", fmt.Errorf("적절하지 않은 타입입니다.")
	}
	return "id", memberShipType, nil
}
