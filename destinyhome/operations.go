package destinyhome

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

var gearHashMap = map[string]int{
	"kinetic":    1498876634,
	"special":    2465295065,
	"power":      953998645,
	"head":       3448274439,
	"arms":       3551918588,
	"chest":      14239492,
	"legs":       20886954,
	"class_item": 1585787867,
}

var operations = map[string]func(map[string]string) (string, error){
	"get_equiped_item": func(params map[string]string) (string, error) {

		const operation = "get_equiped_item"

		// Get the user given the username.
		user, err := repo.GetUser(params["username"])
		if err != nil {
			return "", errors.Wrap(err, operation)
		}

		// Convert the guardianIndex into an int.
		gi, err := strconv.Atoi(params["guardianIndex"])
		if err != nil {
			return "", errors.Wrap(err, operation)
		}

		// Get character equipment.
		equipment, err := bungieSrv.Destiny2().
			GetCharacter(user.MembershipType, user.MembershipID, user.Characters[gi].ID).
			Component(204).
			Do()

		if err != nil {
			return "", errors.Wrap(err, operation)
		}

		if equipment.ErrorCode.IsError() {
			return "", equipment
		}

		// Range over the equipment items.
		for _, v := range equipment.Response.Get().CharacterEquipment.Equipment.Data.Items {

			// Chack for the right bucket.
			if v.BucketHash == gearHashMap[params["bucket"]] {

				// Get the item definition and return.
				item, err := bungosrv.Destiny2.
					GetDestinyEntityDefinition("DestinyInventoryItemDefinition", strconv.Itoa(v.ItemHash)).Do()

				s := fmt.Sprintf("You have the %s %s equiped.", item.Response.DisplayProperties.Name, params["bucket"])
				return s, errors.Wrap(err, operation)
			}
		}

		return "", nil
	},
}
