package being

type bid int64

// Being - npc, player, organization
type Being struct {
	id            bid
	name          string
	speices       speices
	attribute     attribute
	skillList     []skill
	equipmentList []equipment
	assetList     []asset
	relationList  []relation
	relationship  relationship
}
