package vbb

import "encoding/xml"

type DepartureBoard struct {
	XMLName           xml.Name `xml:"DepartureBoard"`
	Text              string   `xml:",chardata"`
	ServerVersion     string   `xml:"serverVersion,attr"`
	DialectVersion    string   `xml:"dialectVersion,attr"`
	PlanRtTs          string   `xml:"planRtTs,attr"`
	RequestId         string   `xml:"requestId,attr"`
	Xmlns             string   `xml:"xmlns,attr"`
	TechnicalMessages struct {
		Text             string `xml:",chardata"`
		TechnicalMessage []struct {
			Text string `xml:",chardata"`
			Key  string `xml:"key,attr"`
		} `xml:"TechnicalMessage"`
	} `xml:"TechnicalMessages"`
	Departure []struct {
		Text             string `xml:",chardata"`
		Name             string `xml:"name,attr"`
		Type             string `xml:"type,attr"`
		Stop             string `xml:"stop,attr"`
		Stopid           string `xml:"stopid,attr"`
		StopExtId        string `xml:"stopExtId,attr"`
		Lon              string `xml:"lon,attr"`
		Lat              string `xml:"lat,attr"`
		PrognosisType    string `xml:"prognosisType,attr"`
		Time             string `xml:"time,attr"`
		Date             string `xml:"date,attr"`
		Track            string `xml:"track,attr"`
		RtTime           string `xml:"rtTime,attr"`
		RtDate           string `xml:"rtDate,attr"`
		Reachable        string `xml:"reachable,attr"`
		Direction        string `xml:"direction,attr"`
		DirectionFlag    string `xml:"directionFlag,attr"`
		JourneyDetailRef struct {
			Text string `xml:",chardata"`
			Ref  string `xml:"ref,attr"`
		} `xml:"JourneyDetailRef"`
		JourneyStatus string `xml:"JourneyStatus"`
		ProductAtStop struct {
			Text          string `xml:",chardata"`
			Name          string `xml:"name,attr"`
			InternalName  string `xml:"internalName,attr"`
			DisplayNumber string `xml:"displayNumber,attr"`
			Num           string `xml:"num,attr"`
			Line          string `xml:"line,attr"`
			CatOut        string `xml:"catOut,attr"`
			CatIn         string `xml:"catIn,attr"`
			CatCode       string `xml:"catCode,attr"`
			Cls           string `xml:"cls,attr"`
			CatOutS       string `xml:"catOutS,attr"`
			CatOutL       string `xml:"catOutL,attr"`
			OperatorCode  string `xml:"operatorCode,attr"`
			Operator      string `xml:"operator,attr"`
			Admin         string `xml:"admin,attr"`
			MatchId       string `xml:"matchId,attr"`
			Icon          struct {
				Text            string `xml:",chardata"`
				Res             string `xml:"res,attr"`
				ForegroundColor struct {
					Text string `xml:",chardata"`
					R    string `xml:"r,attr"`
					G    string `xml:"g,attr"`
					B    string `xml:"b,attr"`
					Hex  string `xml:"hex,attr"`
				} `xml:"foregroundColor"`
				BackgroundColor struct {
					Text string `xml:",chardata"`
					R    string `xml:"r,attr"`
					G    string `xml:"g,attr"`
					B    string `xml:"b,attr"`
					Hex  string `xml:"hex,attr"`
				} `xml:"backgroundColor"`
			} `xml:"icon"`
			OperatorInfo struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				NameS string `xml:"nameS,attr"`
				NameN string `xml:"nameN,attr"`
				NameL string `xml:"nameL,attr"`
				ID    string `xml:"id,attr"`
			} `xml:"operatorInfo"`
		} `xml:"ProductAtStop"`
		Product struct {
			Text          string `xml:",chardata"`
			Name          string `xml:"name,attr"`
			InternalName  string `xml:"internalName,attr"`
			DisplayNumber string `xml:"displayNumber,attr"`
			Num           string `xml:"num,attr"`
			Line          string `xml:"line,attr"`
			CatOut        string `xml:"catOut,attr"`
			CatIn         string `xml:"catIn,attr"`
			CatCode       string `xml:"catCode,attr"`
			Cls           string `xml:"cls,attr"`
			CatOutS       string `xml:"catOutS,attr"`
			CatOutL       string `xml:"catOutL,attr"`
			OperatorCode  string `xml:"operatorCode,attr"`
			Operator      string `xml:"operator,attr"`
			Admin         string `xml:"admin,attr"`
			RouteIdxFrom  string `xml:"routeIdxFrom,attr"`
			RouteIdxTo    string `xml:"routeIdxTo,attr"`
			MatchId       string `xml:"matchId,attr"`
			Icon          struct {
				Text            string `xml:",chardata"`
				Res             string `xml:"res,attr"`
				ForegroundColor struct {
					Text string `xml:",chardata"`
					R    string `xml:"r,attr"`
					G    string `xml:"g,attr"`
					B    string `xml:"b,attr"`
					Hex  string `xml:"hex,attr"`
				} `xml:"foregroundColor"`
				BackgroundColor struct {
					Text string `xml:",chardata"`
					R    string `xml:"r,attr"`
					G    string `xml:"g,attr"`
					B    string `xml:"b,attr"`
					Hex  string `xml:"hex,attr"`
				} `xml:"backgroundColor"`
			} `xml:"icon"`
			OperatorInfo struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				NameS string `xml:"nameS,attr"`
				NameN string `xml:"nameN,attr"`
				NameL string `xml:"nameL,attr"`
				ID    string `xml:"id,attr"`
			} `xml:"operatorInfo"`
		} `xml:"Product"`
		Notes struct {
			Text string `xml:",chardata"`
			Note []struct {
				Text         string `xml:",chardata"`
				Key          string `xml:"key,attr"`
				Type         string `xml:"type,attr"`
				Priority     string `xml:"priority,attr"`
				RouteIdxFrom string `xml:"routeIdxFrom,attr"`
				RouteIdxTo   string `xml:"routeIdxTo,attr"`
				TxtN         string `xml:"txtN,attr"`
			} `xml:"Note"`
		} `xml:"Notes"`
		AltId    string `xml:"altId"`
		Platform struct {
			Text     string `xml:",chardata"`
			Type     string `xml:"type,attr"`
			AttrText string `xml:"text,attr"`
		} `xml:"platform"`
	} `xml:"Departure"`
}
