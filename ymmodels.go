package main

type YMCategory struct {
	Entity      string `json:"entity"`
	FullName    string `json:"fullName"`
	ID          int    `json:"id"`
	IsLeaf      bool   `json:"isLeaf"`
	ModelsCount int    `json:"modelsCount"`
	Name        string `json:"name"`
	Nid         int    `json:"nid"`
	OffersCount int    `json:"offersCount"`
	Type        string `json:"type"`
	ViewType    string `json:"viewType"`
}

type YMLink struct {
	Params struct {
		CatID []string `json:"catId"`
		Hid   []string `json:"hid"`
		Nid   []string `json:"nid"`
	} `json:"params"`
	Target string `json:"target"`
}

type YMPicture struct {
	Entity     string `json:"entity"`
	Height     int    `json:"height"`
	Thumbnails []struct {
		Densities []struct {
			Entity string `json:"entity"`
			ID     string `json:"id"`
			URL    string `json:"url"`
		} `json:"densities"`
		Entity string `json:"entity"`
		Height int    `json:"height"`
		ID     string `json:"id"`
		Width  int    `json:"width"`
	} `json:"thumbnails"`
	URL   string `json:"url"`
	Width int    `json:"width"`
}

type YMRootNavNode struct {
	Entity string `json:"entity"`
	ID     int    `json:"id"`
}

type YMNavigationResp struct {
	Categories struct {
		Category     YMCategory `json:"category"`
		ChildrenType string     `json:"childrenType"`
		Entity       string     `json:"entity"`
		FullName     string     `json:"fullName"`
		HasPromo     bool       `json:"hasPromo"`
		Icons        []struct {
			Entity string `json:"entity"`
			URL    string `json:"url"`
		} `json:"icons"`
		ID       int    `json:"id"`
		IsLeaf   bool   `json:"isLeaf"`
		Link     YMLink `json:"link"`
		Name     string `json:"name"`
		Navnodes []struct {
			URL          string     `json:"url"`
			IsLeaf       bool       `json:"isLeaf"`
			Name         string     `json:"name"`
			Category     YMCategory `json:"category,omitempty"`
			ChildrenType string     `json:"childrenType,omitempty"`
			Entity       string     `json:"entity,omitempty"`
			FullName     string     `json:"fullName,omitempty"`
			HasPromo     bool       `json:"hasPromo,omitempty"`
			ID           int        `json:"id,omitempty"`
			Link         YMLink     `json:"link,omitempty"`
			Navnodes     []struct {
				Category     YMCategory    `json:"category"`
				ChildrenType string        `json:"childrenType"`
				Entity       string        `json:"entity"`
				FullName     string        `json:"fullName"`
				HasPromo     bool          `json:"hasPromo"`
				ID           int           `json:"id"`
				IsLeaf       bool          `json:"isLeaf"`
				Link         YMLink        `json:"link"`
				Name         string        `json:"name"`
				RootNavnode  YMRootNavNode `json:"rootNavnode"`
				Type         string        `json:"type"`
			} `json:"navnodes,omitempty"`
			Pictures    []YMPicture   `json:"pictures,omitempty"`
			RootNavnode YMRootNavNode `json:"rootNavnode,omitempty"`
			Type        string        `json:"type,omitempty"`
		} `json:"navnodes"`
		Pictures    []YMPicture   `json:"pictures"`
		RootNavnode YMRootNavNode `json:"rootNavnode"`
		Type        string        `json:"type"`
	} `json:"categories"`
}

type YMProductFull struct {
	ShowUID string `json:"showUid"`
	Entity  string `json:"entity"`
	Vendor  struct {
		Entity      string `json:"entity"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Website     string `json:"website"`
		Logo        struct {
			Entity     string        `json:"entity"`
			URL        string        `json:"url"`
			Thumbnails []interface{} `json:"thumbnails"`
			Signatures []interface{} `json:"signatures"`
		} `json:"logo"`
		Filter string `json:"filter"`
	} `json:"vendor"`
	Titles struct {
		Raw         string `json:"raw"`
		Highlighted []struct {
			Value string `json:"value"`
		} `json:"highlighted"`
	} `json:"titles"`
	Description                    string `json:"description"`
	EligibleForBookingInUserRegion bool   `json:"eligibleForBookingInUserRegion"`
	Categories                     []struct {
		Entity   string `json:"entity"`
		ID       int    `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"fullName"`
		Type     string `json:"type"`
		CpaType  string `json:"cpaType"`
		IsLeaf   bool   `json:"isLeaf"`
	} `json:"categories"`
	Urls struct {
	} `json:"urls"`
	Navnodes []struct {
		Entity      string `json:"entity"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
		FullName    string `json:"fullName"`
		IsLeaf      bool   `json:"isLeaf"`
		RootNavnode struct {
		} `json:"rootNavnode"`
	} `json:"navnodes"`
	Pictures []struct {
		Entity   string `json:"entity"`
		Original struct {
			ContainerWidth  int    `json:"containerWidth"`
			ContainerHeight int    `json:"containerHeight"`
			URL             string `json:"url"`
			Width           int    `json:"width"`
			Height          int    `json:"height"`
		} `json:"original"`
		Thumbnails []struct {
			ContainerWidth  int    `json:"containerWidth"`
			ContainerHeight int    `json:"containerHeight"`
			URL             string `json:"url"`
			Width           int    `json:"width"`
			Height          int    `json:"height"`
		} `json:"thumbnails"`
		FiltersMatching map[string]interface{} `json:"filtersMatching"`
		Signatures      []interface{}          `json:"signatures"`
	} `json:"pictures"`
	Filters []struct {
		ID          string `json:"id"`
		Type        string `json:"type"`
		Name        string `json:"name"`
		Xslname     string `json:"xslname"`
		SubType     string `json:"subType"`
		Kind        int    `json:"kind"`
		Position    int    `json:"position"`
		Noffers     int    `json:"noffers"`
		ValuesCount int    `json:"valuesCount,omitempty"`
		Values      []struct {
			InitialFound int    `json:"initialFound"`
			Found        int    `json:"found"`
			Value        string `json:"value"`
			Vendor       struct {
				Name   string `json:"name"`
				Entity string `json:"entity"`
				ID     int    `json:"id"`
			} `json:"vendor"`
			ID string `json:"id"`
		} `json:"values"`
		ValuesGroups []struct {
			Type      string   `json:"type"`
			ValuesIds []string `json:"valuesIds"`
		} `json:"valuesGroups,omitempty"`
	} `json:"filters"`
	Meta struct {
	} `json:"meta"`
	Type   string `json:"type"`
	ID     int    `json:"id"`
	Offers struct {
		Count int `json:"count"`
		Items []struct {
			Entity string `json:"entity"`
			Vendor struct {
				Entity      string `json:"entity"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Website     string `json:"website"`
				Logo        struct {
					Entity     string        `json:"entity"`
					URL        string        `json:"url"`
					Thumbnails []interface{} `json:"thumbnails"`
					Signatures []interface{} `json:"signatures"`
				} `json:"logo"`
			} `json:"vendor"`
			Titles struct {
				Raw         string `json:"raw"`
				Highlighted []struct {
					Value string `json:"value"`
				} `json:"highlighted"`
			} `json:"titles"`
			Description                    string `json:"description"`
			EligibleForBookingInUserRegion bool   `json:"eligibleForBookingInUserRegion"`
			Categories                     []struct {
				Entity   string `json:"entity"`
				ID       int    `json:"id"`
				Name     string `json:"name"`
				FullName string `json:"fullName"`
				Type     string `json:"type"`
				CpaType  string `json:"cpaType"`
				IsLeaf   bool   `json:"isLeaf"`
			} `json:"categories"`
			Cpc  string `json:"cpc"`
			Urls struct {
				Direct string `json:"direct"`
			} `json:"urls"`
			Navnodes []struct {
				Entity      string `json:"entity"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
				FullName    string `json:"fullName"`
				IsLeaf      bool   `json:"isLeaf"`
				RootNavnode struct {
				} `json:"rootNavnode"`
			} `json:"navnodes"`
			Pictures []struct {
				Entity   string `json:"entity"`
				Original struct {
					ContainerWidth  int    `json:"containerWidth"`
					ContainerHeight int    `json:"containerHeight"`
					URL             string `json:"url"`
					Width           int    `json:"width"`
					Height          int    `json:"height"`
				} `json:"original"`
				Thumbnails []struct {
					ContainerWidth  int    `json:"containerWidth"`
					ContainerHeight int    `json:"containerHeight"`
					URL             string `json:"url"`
					Width           int    `json:"width"`
					Height          int    `json:"height"`
				} `json:"thumbnails"`
				Signatures []interface{} `json:"signatures"`
			} `json:"pictures"`
			Meta struct {
			} `json:"meta"`
			Model struct {
				ID int `json:"id"`
			} `json:"model"`
			Delivery struct {
				ShopPriorityRegion struct {
					Entity string `json:"entity"`
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Lingua struct {
						Name struct {
							Genitive      string `json:"genitive"`
							Preposition   string `json:"preposition"`
							Prepositional string `json:"prepositional"`
							Accusative    string `json:"accusative"`
						} `json:"name"`
					} `json:"lingua"`
				} `json:"shopPriorityRegion"`
				ShopPriorityCountry struct {
					Entity string `json:"entity"`
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Lingua struct {
						Name struct {
							Genitive      string `json:"genitive"`
							Preposition   string `json:"preposition"`
							Prepositional string `json:"prepositional"`
							Accusative    string `json:"accusative"`
						} `json:"name"`
					} `json:"lingua"`
				} `json:"shopPriorityCountry"`
				IsPriorityRegion bool `json:"isPriorityRegion"`
				IsCountrywide    bool `json:"isCountrywide"`
				IsAvailable      bool `json:"isAvailable"`
				HasPickup        bool `json:"hasPickup"`
				HasLocalStore    bool `json:"hasLocalStore"`
				HasPost          bool `json:"hasPost"`
				Region           struct {
					Entity string `json:"entity"`
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Lingua struct {
						Name struct {
							Genitive      string `json:"genitive"`
							Preposition   string `json:"preposition"`
							Prepositional string `json:"prepositional"`
							Accusative    string `json:"accusative"`
						} `json:"name"`
					} `json:"lingua"`
				} `json:"region"`
				AvailableServices []struct {
					ServiceID   int    `json:"serviceId"`
					ServiceName string `json:"serviceName"`
				} `json:"availableServices"`
				Price struct {
					Currency           string `json:"currency"`
					Value              string `json:"value"`
					IsDeliveryIncluded bool   `json:"isDeliveryIncluded"`
				} `json:"price"`
				IsFree         bool `json:"isFree"`
				IsDownloadable bool `json:"isDownloadable"`
				InStock        bool `json:"inStock"`
				PostAvailable  bool `json:"postAvailable"`
				Options        []struct {
					Price struct {
						Currency           string `json:"currency"`
						Value              string `json:"value"`
						IsDeliveryIncluded bool   `json:"isDeliveryIncluded"`
					} `json:"price"`
					DayFrom     int    `json:"dayFrom"`
					DayTo       int    `json:"dayTo"`
					IsDefault   bool   `json:"isDefault"`
					ServiceID   string `json:"serviceId"`
					PartnerType string `json:"partnerType"`
					Region      struct {
						Entity string `json:"entity"`
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Lingua struct {
							Name struct {
								Genitive      string `json:"genitive"`
								Preposition   string `json:"preposition"`
								Prepositional string `json:"prepositional"`
								Accusative    string `json:"accusative"`
							} `json:"name"`
						} `json:"lingua"`
					} `json:"region"`
				} `json:"options"`
			} `json:"delivery"`
			Shop struct {
				Entity              string  `json:"entity"`
				ID                  int     `json:"id"`
				Name                string  `json:"name"`
				GradesCount         int     `json:"gradesCount"`
				OverallGradesCount  int     `json:"overallGradesCount"`
				QualityRating       int     `json:"qualityRating"`
				IsGlobal            bool    `json:"isGlobal"`
				IsCpaPrior          bool    `json:"isCpaPrior"`
				IsCpaPartner        bool    `json:"isCpaPartner"`
				IsNewRating         bool    `json:"isNewRating"`
				NewGradesCount      int     `json:"newGradesCount"`
				NewQualityRating    float64 `json:"newQualityRating"`
				NewQualityRating3M  float64 `json:"newQualityRating3M"`
				RatingToShow        float64 `json:"ratingToShow"`
				NewGradesCount3M    int     `json:"newGradesCount3M"`
				Status              string  `json:"status"`
				Cutoff              string  `json:"cutoff"`
				OutletsCount        int     `json:"outletsCount"`
				StoresCount         int     `json:"storesCount"`
				PickupStoresCount   int     `json:"pickupStoresCount"`
				DepotStoresCount    int     `json:"depotStoresCount"`
				PostomatStoresCount int     `json:"postomatStoresCount"`
				BookNowStoresCount  int     `json:"bookNowStoresCount"`
				Subsidies           bool    `json:"subsidies"`
				Feed                struct {
					ID         string `json:"id"`
					OfferID    string `json:"offerId"`
					CategoryID string `json:"categoryId"`
				} `json:"feed"`
				HomeRegion struct {
					Entity string `json:"entity"`
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Lingua struct {
						Name struct {
						} `json:"name"`
					} `json:"lingua"`
				} `json:"homeRegion"`
			} `json:"shop"`
			ReturnPolicy      string `json:"returnPolicy"`
			WareID            string `json:"wareId"`
			ClassifierMagicID string `json:"classifierMagicId"`
			Prices            struct {
				Currency           string `json:"currency"`
				Value              string `json:"value"`
				IsDeliveryIncluded bool   `json:"isDeliveryIncluded"`
				RawValue           string `json:"rawValue"`
			} `json:"prices"`
			Manufacturer struct {
				Entity   string `json:"entity"`
				Warranty bool   `json:"warranty"`
			} `json:"manufacturer"`
			Seller struct {
				Price                    string `json:"price"`
				Currency                 string `json:"currency"`
				SellerToUserExchangeRate int    `json:"sellerToUserExchangeRate"`
			} `json:"seller"`
			IsRecommendedByVendor bool `json:"isRecommendedByVendor"`
			Benefit               struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				IsPrimary   bool   `json:"isPrimary"`
			} `json:"benefit"`
			Outlet struct {
				Entity          string   `json:"entity"`
				ID              string   `json:"id"`
				Name            string   `json:"name"`
				Type            string   `json:"type"`
				Purpose         []string `json:"purpose"`
				ServiceID       int      `json:"serviceId"`
				Email           string   `json:"email"`
				IsMarketBranded bool     `json:"isMarketBranded"`
				Shop            struct {
					ID int `json:"id"`
				} `json:"shop"`
				Address struct {
					FullAddress  string `json:"fullAddress"`
					Country      string `json:"country"`
					Region       string `json:"region"`
					Locality     string `json:"locality"`
					Street       string `json:"street"`
					Km           string `json:"km"`
					Building     string `json:"building"`
					Block        string `json:"block"`
					Wing         string `json:"wing"`
					Estate       string `json:"estate"`
					Entrance     string `json:"entrance"`
					Floor        string `json:"floor"`
					Room         string `json:"room"`
					OfficeNumber string `json:"office_number"`
					Note         string `json:"note"`
				} `json:"address"`
				Telephones []struct {
					Entity          string `json:"entity"`
					CountryCode     string `json:"countryCode"`
					CityCode        string `json:"cityCode"`
					TelephoneNumber string `json:"telephoneNumber"`
					ExtensionNumber string `json:"extensionNumber"`
				} `json:"telephones"`
				WorkingTime []struct {
					DaysFrom  string `json:"daysFrom"`
					DaysTo    string `json:"daysTo"`
					HoursFrom string `json:"hoursFrom"`
					HoursTo   string `json:"hoursTo"`
				} `json:"workingTime"`
				SelfDeliveryRule struct {
					WorkInHoliday          bool   `json:"workInHoliday"`
					Currency               string `json:"currency"`
					Cost                   string `json:"cost"`
					ShipperHumanReadableID string `json:"shipperHumanReadableId"`
				} `json:"selfDeliveryRule"`
				Region struct {
					Entity string `json:"entity"`
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Lingua struct {
						Name struct {
							Genitive      string `json:"genitive"`
							Preposition   string `json:"preposition"`
							Prepositional string `json:"prepositional"`
							Accusative    string `json:"accusative"`
						} `json:"name"`
					} `json:"lingua"`
				} `json:"region"`
				GpsCoord struct {
					Longitude string `json:"longitude"`
					Latitude  string `json:"latitude"`
				} `json:"gpsCoord"`
			} `json:"outlet"`
			BundleSettings struct {
				QuantityLimit struct {
					Minimum int `json:"minimum"`
					Step    int `json:"step"`
				} `json:"quantityLimit"`
			} `json:"bundleSettings"`
			PrepayEnabled    bool   `json:"prepayEnabled"`
			PromoCodeEnabled bool   `json:"promoCodeEnabled"`
			FeedGroupID      string `json:"feedGroupId"`
			IsFulfillment    bool   `json:"isFulfillment"`
		} `json:"items"`
	} `json:"offers"`
	IsNew  bool `json:"isNew"`
	Prices struct {
		Min      string `json:"min"`
		Max      string `json:"max"`
		Currency string `json:"currency"`
		Avg      string `json:"avg"`
	} `json:"prices"`
	Opinions     int     `json:"opinions"`
	Rating       float64 `json:"rating"`
	Reviews      int     `json:"reviews"`
	ReasonsToBuy []struct {
		FactorName     string  `json:"factor_name,omitempty"`
		Type           string  `json:"type"`
		FactorID       string  `json:"factor_id,omitempty"`
		Value          float64 `json:"value"`
		FactorPriority string  `json:"factor_priority,omitempty"`
		ID             string  `json:"id"`
		AuthorPuid     string  `json:"author_puid,omitempty"`
		FeedbackID     string  `json:"feedback_id,omitempty"`
		Text           string  `json:"text,omitempty"`
	} `json:"reasonsToBuy"`
	Specs struct {
		Friendly    []string `json:"friendly"`
		Friendlyext []struct {
			Type       string `json:"type"`
			Value      string `json:"value"`
			UsedParams []int  `json:"usedParams"`
		} `json:"friendlyext"`
	} `json:"specs"`
	Lingua struct {
		Type struct {
			Nominative string `json:"nominative"`
			Genitive   string `json:"genitive"`
			Dative     string `json:"dative"`
			Accusative string `json:"accusative"`
		} `json:"type"`
	} `json:"lingua"`
	RetailersCount int `json:"retailersCount"`
	SkuStats       struct {
		TotalCount         int `json:"totalCount"`
		BeforeFiltersCount int `json:"beforeFiltersCount"`
		AfterFiltersCount  int `json:"afterFiltersCount"`
	} `json:"skuStats"`
	Promo struct {
		WhitePromoCount int `json:"whitePromoCount"`
	} `json:"promo"`
}

type YMProductStripped struct {
	ShowUID string `json:"showUid"`
	Vendor  struct {
		// ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"vendor"`
	Titles struct {
		Raw string `json:"raw"`
	} `json:"titles"`
	Description string `json:"description"`
	Type        string `json:"type"`
	// ID          int    `json:"id"`
	Offers struct {
		Count int `json:"count"`
	} `json:"offers"`
	IsNew  bool `json:"isNew"`
	Prices struct {
		Min      string `json:"min"`
		Max      string `json:"max"`
		Currency string `json:"currency"`
		Avg      string `json:"avg"`
	} `json:"prices"`
	Opinions       int     `json:"opinions"`
	Rating         float64 `json:"rating"`
	Reviews        int     `json:"reviews"`
	RetailersCount int     `json:"retailersCount"`
	SkuStats       struct {
		TotalCount         int `json:"totalCount"`
		BeforeFiltersCount int `json:"beforeFiltersCount"`
		AfterFiltersCount  int `json:"afterFiltersCount"`
	} `json:"skuStats"`
	Promo struct {
		WhitePromoCount int `json:"whitePromoCount"`
	} `json:"promo"`
}

type YMOfferFull struct {
	Entity string `json:"entity"`
	Titles struct {
		Raw         string `json:"raw"`
		Highlighted []struct {
			Value string `json:"value"`
		} `json:"highlighted"`
	} `json:"titles"`
	Description                    string `json:"description"`
	EligibleForBookingInUserRegion bool   `json:"eligibleForBookingInUserRegion"`
	Categories                     []struct {
		Entity   string `json:"entity"`
		ID       int    `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"fullName"`
		Type     string `json:"type"`
		CpaType  string `json:"cpaType"`
		IsLeaf   bool   `json:"isLeaf"`
	} `json:"categories"`
	Cpc  string `json:"cpc"`
	Urls struct {
		Encrypted string `json:"encrypted"`
		CallPhone string `json:"callPhone"`
		Offercard string `json:"offercard"`
		Direct    string `json:"direct"`
	} `json:"urls"`
	Navnodes []struct {
		Entity      string `json:"entity"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
		FullName    string `json:"fullName"`
		IsLeaf      bool   `json:"isLeaf"`
		RootNavnode struct {
		} `json:"rootNavnode"`
	} `json:"navnodes"`
	Pictures []struct {
		Entity   string `json:"entity"`
		Original struct {
			ContainerWidth  int    `json:"containerWidth"`
			ContainerHeight int    `json:"containerHeight"`
			URL             string `json:"url"`
			Width           int    `json:"width"`
			Height          int    `json:"height"`
		} `json:"original"`
		Thumbnails []struct {
			ContainerWidth  int    `json:"containerWidth"`
			ContainerHeight int    `json:"containerHeight"`
			URL             string `json:"url"`
			Width           int    `json:"width"`
			Height          int    `json:"height"`
		} `json:"thumbnails"`
		Signatures []interface{} `json:"signatures"`
	} `json:"pictures"`
	Meta struct {
	} `json:"meta"`
	Delivery struct {
		ShopPriorityRegion struct {
			Entity string `json:"entity"`
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Lingua struct {
				Name struct {
					Genitive      string `json:"genitive"`
					Preposition   string `json:"preposition"`
					Prepositional string `json:"prepositional"`
					Accusative    string `json:"accusative"`
				} `json:"name"`
			} `json:"lingua"`
		} `json:"shopPriorityRegion"`
		ShopPriorityCountry struct {
			Entity string `json:"entity"`
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Lingua struct {
				Name struct {
					Genitive      string `json:"genitive"`
					Preposition   string `json:"preposition"`
					Prepositional string `json:"prepositional"`
					Accusative    string `json:"accusative"`
				} `json:"name"`
			} `json:"lingua"`
		} `json:"shopPriorityCountry"`
		IsPriorityRegion bool `json:"isPriorityRegion"`
		IsCountrywide    bool `json:"isCountrywide"`
		IsAvailable      bool `json:"isAvailable"`
		HasPickup        bool `json:"hasPickup"`
		HasLocalStore    bool `json:"hasLocalStore"`
		HasPost          bool `json:"hasPost"`
		Region           struct {
			Entity string `json:"entity"`
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Lingua struct {
				Name struct {
					Genitive      string `json:"genitive"`
					Preposition   string `json:"preposition"`
					Prepositional string `json:"prepositional"`
					Accusative    string `json:"accusative"`
				} `json:"name"`
			} `json:"lingua"`
		} `json:"region"`
		Price struct {
			Currency           string `json:"currency"`
			Value              string `json:"value"`
			IsDeliveryIncluded bool   `json:"isDeliveryIncluded"`
		} `json:"price"`
		IsFree         bool `json:"isFree"`
		IsDownloadable bool `json:"isDownloadable"`
		InStock        bool `json:"inStock"`
		PostAvailable  bool `json:"postAvailable"`
		Options        []struct {
			Price struct {
				Currency           string `json:"currency"`
				Value              string `json:"value"`
				IsDeliveryIncluded bool   `json:"isDeliveryIncluded"`
			} `json:"price"`
			DayFrom     int    `json:"dayFrom"`
			DayTo       int    `json:"dayTo"`
			OrderBefore string `json:"orderBefore"`
			IsDefault   bool   `json:"isDefault"`
			ServiceID   string `json:"serviceId"`
			PartnerType string `json:"partnerType"`
			Region      struct {
				Entity string `json:"entity"`
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Lingua struct {
					Name struct {
						Genitive      string `json:"genitive"`
						Preposition   string `json:"preposition"`
						Prepositional string `json:"prepositional"`
						Accusative    string `json:"accusative"`
					} `json:"name"`
				} `json:"lingua"`
			} `json:"region"`
		} `json:"options"`
	} `json:"delivery"`
	Shop              int    `json:"shop"`
	ReturnPolicy      string `json:"returnPolicy"`
	WareID            string `json:"wareId"`
	ClassifierMagicID string `json:"classifierMagicId"`
	Prices            struct {
		Currency           string `json:"currency"`
		Value              string `json:"value"`
		IsDeliveryIncluded bool   `json:"isDeliveryIncluded"`
		RawValue           string `json:"rawValue"`
	} `json:"prices"`
	Manufacturer struct {
		Entity   string `json:"entity"`
		Warranty bool   `json:"warranty"`
		Code     string `json:"code"`
	} `json:"manufacturer"`
	Seller struct {
		Price                    string `json:"price"`
		Currency                 string `json:"currency"`
		SellerToUserExchangeRate int    `json:"sellerToUserExchangeRate"`
	} `json:"seller"`
	IsRecommendedByVendor bool `json:"isRecommendedByVendor"`
	Outlet                struct {
		Entity          string   `json:"entity"`
		ID              string   `json:"id"`
		Name            string   `json:"name"`
		Type            string   `json:"type"`
		Purpose         []string `json:"purpose"`
		ServiceID       int      `json:"serviceId"`
		Email           string   `json:"email"`
		IsMarketBranded bool     `json:"isMarketBranded"`
		Shop            struct {
			ID int `json:"id"`
		} `json:"shop"`
		Address struct {
			FullAddress  string `json:"fullAddress"`
			Country      string `json:"country"`
			Region       string `json:"region"`
			Locality     string `json:"locality"`
			Street       string `json:"street"`
			Km           string `json:"km"`
			Building     string `json:"building"`
			Block        string `json:"block"`
			Wing         string `json:"wing"`
			Estate       string `json:"estate"`
			Entrance     string `json:"entrance"`
			Floor        string `json:"floor"`
			Room         string `json:"room"`
			OfficeNumber string `json:"office_number"`
			Note         string `json:"note"`
		} `json:"address"`
		Telephones []struct {
			Entity          string `json:"entity"`
			CountryCode     string `json:"countryCode"`
			CityCode        string `json:"cityCode"`
			TelephoneNumber string `json:"telephoneNumber"`
			ExtensionNumber string `json:"extensionNumber"`
		} `json:"telephones"`
		WorkingTime []struct {
			DaysFrom  string `json:"daysFrom"`
			DaysTo    string `json:"daysTo"`
			HoursFrom string `json:"hoursFrom"`
			HoursTo   string `json:"hoursTo"`
		} `json:"workingTime"`
		SelfDeliveryRule struct {
			WorkInHoliday          bool   `json:"workInHoliday"`
			Currency               string `json:"currency"`
			Cost                   string `json:"cost"`
			PriceTo                string `json:"priceTo"`
			ShipperHumanReadableID string `json:"shipperHumanReadableId"`
		} `json:"selfDeliveryRule"`
		Region struct {
			Entity string `json:"entity"`
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Lingua struct {
				Name struct {
					Genitive      string `json:"genitive"`
					Preposition   string `json:"preposition"`
					Prepositional string `json:"prepositional"`
					Accusative    string `json:"accusative"`
				} `json:"name"`
			} `json:"lingua"`
		} `json:"region"`
		GpsCoord struct {
			Longitude string `json:"longitude"`
			Latitude  string `json:"latitude"`
		} `json:"gpsCoord"`
	} `json:"outlet"`
	BundleSettings struct {
		QuantityLimit struct {
			Minimum int `json:"minimum"`
			Step    int `json:"step"`
		} `json:"quantityLimit"`
	} `json:"bundleSettings"`
	PrepayEnabled    bool   `json:"prepayEnabled"`
	PromoCodeEnabled bool   `json:"promoCodeEnabled"`
	FeedGroupID      string `json:"feedGroupId"`
	IsFulfillment    bool   `json:"isFulfillment"`
}

type YMOfferStripped struct {
	Entity string `json:"entity"`
	Titles struct {
		Raw string `json:"raw"`
	} `json:"titles"`
	Description                    string `json:"description"`
	EligibleForBookingInUserRegion bool   `json:"eligibleForBookingInUserRegion"`

	Prices struct {
		Currency           string `json:"currency"`
		Value              string `json:"value"`
		IsDeliveryIncluded bool   `json:"isDeliveryIncluded"`
		RawValue           string `json:"rawValue"`
	} `json:"prices"`
	Seller struct {
		Price                    string `json:"price"`
		Currency                 string `json:"currency"`
		SellerToUserExchangeRate int    `json:"sellerToUserExchangeRate"`
	} `json:"seller"`
}

type YMProduct YMProductStripped
type YMOffer YMOfferStripped

type YMResultsResp struct {
	Data struct {
		Collections struct {
			Products map[string]YMProduct `json:"products"`
			Offers   map[string]YMOffer   `json:"offers"`
		} `json:"collections"`
		Entities []struct {
			// ID     int    `json:"id"`
			Schema string `json:"schema"`
		} `json:"entities"`
	} `json:"data"`
	Status bool `json:"status"`
}
