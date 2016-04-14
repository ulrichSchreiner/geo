package sfo

import(
	"sort"
	"github.com/skypies/geo"
)

// A bunch of constants relating to SFO

var (
	KLatlongSFO = geo.Latlong{37.6188172, -122.3754281}
	KLatlongSJC = geo.Latlong{37.3639472, -121.9289375}	
	KLatlongSERFR1 = geo.Latlong{37.221516, -121.992987} // This is the centerpoint for maps viewport

	KBoxSFO120K = KLatlongSFO.Box(80,80)

	KBoxSnarfingCatchment = KLatlongSFO.Box(125,125)  // The box in which we look for new flights

	// Boxes used in a few reports
	KBoxSFO10K = KLatlongSFO.Box(12,12)
	KBoxPaloAlto20K = geo.Latlong{37.433536,-122.1310187}.Box(6,7)
	
	// http://www.myaviationinfo.com/FixState.php?FixState=CALIFORNIA
	KFixes = map[string]geo.Latlong{
		// SERFR2 & WWAVS1
		"SERFR": geo.Latlong{36.0683056, -121.3646639},
		"NRRLI": geo.Latlong{36.4956000, -121.6994000},
		"WWAVS": geo.Latlong{36.7415306, -121.8942333},
		"EPICK": geo.Latlong{36.9508222, -121.9526722},
		"EDDYY": geo.Latlong{37.3264500, -122.0997083},
		"SWELS": geo.Latlong{37.3681556, -122.1160806},
		"MENLO": geo.Latlong{37.4636861, -122.1536583},
		"WPOUT": geo.Latlong{37.1194861, -122.2927417},
		"THEEZ": geo.Latlong{37.5034694, -122.4247528},
		"WESLA": geo.Latlong{37.6643722, -122.4802917},
		"MVRKK": geo.Latlong{37.7369722, -122.4544500},

		// BRIXX
		"CORKK": geo.Latlong{37.7335889, -122.4975500},
		"BRIXX": geo.Latlong{37.6178444, -122.3745278},
		"LUYTA": geo.Latlong{37.2948889, -122.2045528},
		"JILNA": geo.Latlong{37.2488056, -122.1495000},
		"YADUT": geo.Latlong{37.2039889, -122.0232778},

		// BIGSURTWO
		"CARME": geo.Latlong{36.4551833, -121.8797139},
		"ANJEE": geo.Latlong{36.7462861, -121.9648917},
		"SKUNK": geo.Latlong{37.0075944, -122.0332278},
		"BOLDR": geo.Latlong{37.1708861, -122.0761667},

		// Things for SFO arrivals
		"HEMAN": geo.Latlong{37.5338500, -122.1733333},
		"DUYET": geo.Latlong{37.5674000, -122.2529278},
		"NEPIC": geo.Latlong{37.5858944, -122.2968833},
		
		// Things for Oceanic
		"PPEGS": geo.Latlong{37.3920722, -122.2817222},
		"ALLBE": geo.Latlong{37.5063889, -127.0000000},
		"ALCOA": geo.Latlong{37.8332528, -125.8345250},
		"CINNY": geo.Latlong{36.1816667, -124.7600000},
		"PAINT": geo.Latlong{38.0000000, -125.5000000},
		"OSI"  : geo.Latlong{37.3925000, -122.2813000},
		"PIRAT": geo.Latlong{37.2576500, -122.8633528},

		"PONKE": geo.Latlong{37.4588167, -121.9960528},
		"WETOR": geo.Latlong{37.4847194, -122.0571417},

		// Things for SILCN3
		"VLLEY": geo.Latlong{36.5091667, -121.4402778},
		"GUUYY": geo.Latlong{36.7394444, -121.5411111},
		"SSEBB": geo.Latlong{36.9788889, -121.6425000},
		"GSTEE": geo.Latlong{37.0708333, -121.6716667},
		"KLIDE": geo.Latlong{37.1641667, -121.7130556},
		"BAXBE": geo.Latlong{36.7730556, -121.6263889},
		"APLLE": geo.Latlong{37.0338889, -121.8050000},

		// Personal entries
		"X_RSH": geo.Latlong{36.868582,  -121.691934},
		"X_BLH": geo.Latlong{37.2199471, -122.0425108},
		"X_HBR": geo.Latlong{37.309564,  -122.112378},
	}

	SFOClassBMap = geo.ClassBMap{
		Name: "SFO",
		Center: KLatlongSFO,
		Sectors: []geo.ClassBSector{
			// Magnetic declination at SFO: 13.68
			geo.ClassBSector{
				StartBearing: 0,
				EndBearing: 360,
				Steps: []geo.Cylinder{
					{ 7,  0, 100},   // from origin to  7NM : 100/00 (no floor)
					{10, 15, 100},   // from   7NM  to 10NM : 100/15
					{15, 30, 100},   // from  10NM  to 15NM : 100/30
					{20, 40, 100},   // from  15NM  to 20NM : 100/40
					{25, 60, 100},   // from  20NM  to 25NM : 100/60
					{30, 80, 100},   // from  25NM  to 30NM : 100/80
				},
			},
			// ... more sectors go here !
		},
	}

	// http://flightaware.com/resources/airport/SFO/STAR/SERFR+TWO+(RNAV)/pdf
	Serfr1 = geo.Procedure{
		Name: "SERFR2",
		Departure: false,
		Airport: "SFO",
		Waypoints: []geo.Waypoint{
			{"SERFR", geo.Latlong{},    0,     0,   0, false}, // Many aircraft skip SERFR
			{"NNRLI", geo.Latlong{}, 20000, 20000, 280, true},
			{"WWAVS", geo.Latlong{}, 15000, 19000, 280, true},
			{"EPICK", geo.Latlong{}, 10000, 15000, 280, true},
			{"EDDYY", geo.Latlong{},  6000,  6000, 240, true},  // Delay vectoring inside EPICK-EDDYY
			{"SWELS", geo.Latlong{},  4700,  4700, 240, false},
			{"MENLO", geo.Latlong{},  4000,  4000, 230, false},
		},
	}
)

func ListWaypoints() []string {
	ret := []string{}
	for k,_ := range KFixes { ret = append(ret,k) }
	sort.Strings(ret)
	return ret
}

// {{{ -------------------------={ E N D }=----------------------------------

// Local variables:
// folded-file: t
// end:

// }}}
