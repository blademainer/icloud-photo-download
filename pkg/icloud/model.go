package icloud

type LoginRequest struct {
	AccountName string   `json:"accountName"`
	RememberMe  bool     `json:"rememberMe"`
	Password    string   `json:"password"`
	TrustTokens []string `json:"trustTokens"`
}

type LoginResponse struct {
	AccountName string   `json:"accountName"`
	RememberMe  bool     `json:"rememberMe"`
	Password    string   `json:"password"`
	TrustTokens []string `json:"trustTokens"`
}

type UserInfo struct {
	DsInfo struct {
		LastName                           string   `json:"lastName"`
		ICDPEnabled                        bool     `json:"iCDPEnabled"`
		TantorMigrated                     bool     `json:"tantorMigrated"`
		Dsid                               string   `json:"dsid"`
		HsaEnabled                         bool     `json:"hsaEnabled"`
		IsHideMyEmailSubscriptionActive    bool     `json:"isHideMyEmailSubscriptionActive"`
		IroncadeMigrated                   bool     `json:"ironcadeMigrated"`
		Locale                             string   `json:"locale"`
		BrZoneConsolidated                 bool     `json:"brZoneConsolidated"`
		IsManagedAppleID                   bool     `json:"isManagedAppleID"`
		IsCustomDomainsFeatureAvailable    bool     `json:"isCustomDomainsFeatureAvailable"`
		IsHideMyEmailFeatureAvailable      bool     `json:"isHideMyEmailFeatureAvailable"`
		ContinueOnDeviceEligibleDeviceInfo []string `json:"ContinueOnDeviceEligibleDeviceInfo"`
		GilliganInvited                    bool     `json:"gilligan-invited"`
		AppleIdAliases                     []string `json:"appleIdAliases"`
		HsaVersion                         int      `json:"hsaVersion"`
		UbiquityEOLEnabled                 bool     `json:"ubiquityEOLEnabled"`
		IsPaidDeveloper                    bool     `json:"isPaidDeveloper"`
		CountryCode                        string   `json:"countryCode"`
		NotificationId                     string   `json:"notificationId"`
		PrimaryEmailVerified               bool     `json:"primaryEmailVerified"`
		ADsID                              string   `json:"aDsID"`
		Locked                             bool     `json:"locked"`
		HasICloudQualifyingDevice          bool     `json:"hasICloudQualifyingDevice"`
		PrimaryEmail                       string   `json:"primaryEmail"`
		AppleIdEntries                     []struct {
			IsPrimary bool   `json:"isPrimary"`
			Type      string `json:"type"`
			Value     string `json:"value"`
		} `json:"appleIdEntries"`
		GilliganEnabled bool   `json:"gilligan-enabled"`
		FullName        string `json:"fullName"`
		MailFlags       struct {
			IsThreadingAvailable           bool `json:"isThreadingAvailable"`
			IsSearchV2Provisioned          bool `json:"isSearchV2Provisioned"`
			IsCKMail                       bool `json:"isCKMail"`
			IsMppSupportedInCurrentCountry bool `json:"isMppSupportedInCurrentCountry"`
		} `json:"mailFlags"`
		LanguageCode         string `json:"languageCode"`
		AppleId              string `json:"appleId"`
		HasUnreleasedOS      bool   `json:"hasUnreleasedOS"`
		AnalyticsOptInStatus bool   `json:"analyticsOptInStatus"`
		FirstName            string `json:"firstName"`
		ICloudAppleIdAlias   string `json:"iCloudAppleIdAlias"`
		NotesMigrated        bool   `json:"notesMigrated"`
		BeneficiaryInfo      struct {
			IsBeneficiary bool `json:"isBeneficiary"`
		} `json:"beneficiaryInfo"`
		HasPaymentInfo bool   `json:"hasPaymentInfo"`
		PcsDeleted     bool   `json:"pcsDeleted"`
		AppleIdAlias   string `json:"appleIdAlias"`
		BrMigrated     bool   `json:"brMigrated"`
		StatusCode     int    `json:"statusCode"`
		FamilyEligible bool   `json:"familyEligible"`
	} `json:"dsInfo"`
	HasMinimumDeviceForPhotosWeb bool `json:"hasMinimumDeviceForPhotosWeb"`
	ICDPEnabled                  bool `json:"iCDPEnabled"`
	Webservices                  struct {
		Reminders struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"reminders"`
		Notes struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"notes"`
		Mail struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"mail"`
		Ckdatabasews struct {
			PcsRequired bool   `json:"pcsRequired"`
			Url         string `json:"url"`
			Status      string `json:"status"`
		} `json:"ckdatabasews"`
		Photosupload struct {
			PcsRequired bool   `json:"pcsRequired"`
			Url         string `json:"url"`
			Status      string `json:"status"`
		} `json:"photosupload"`
		Mcc struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"mcc"`
		Photos struct {
			PcsRequired bool   `json:"pcsRequired"`
			UploadUrl   string `json:"uploadUrl"`
			Url         string `json:"url"`
			Status      string `json:"status"`
		} `json:"photos"`
		Drivews struct {
			PcsRequired bool   `json:"pcsRequired"`
			Url         string `json:"url"`
			Status      string `json:"status"`
		} `json:"drivews"`
		Uploadimagews struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"uploadimagews"`
		Schoolwork struct {
		} `json:"schoolwork"`
		Cksharews struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"cksharews"`
		Findme struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"findme"`
		Ckdeviceservice struct {
			Url string `json:"url"`
		} `json:"ckdeviceservice"`
		Iworkthumbnailws struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"iworkthumbnailws"`
		Mccgateway struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"mccgateway"`
		Calendar struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"calendar"`
		Docws struct {
			PcsRequired bool   `json:"pcsRequired"`
			Url         string `json:"url"`
			Status      string `json:"status"`
		} `json:"docws"`
		Settings struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"settings"`
		Premiummailsettings struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"premiummailsettings"`
		Ubiquity struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"ubiquity"`
		Streams struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"streams"`
		Keyvalue struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"keyvalue"`
		Mpp struct {
		} `json:"mpp"`
		Archivews struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"archivews"`
		Push struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"push"`
		Iwmb struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"iwmb"`
		Iworkexportws struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"iworkexportws"`
		Sharedlibrary struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"sharedlibrary"`
		Geows struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"geows"`
		Account struct {
			ICloudEnv struct {
				ShortId   string `json:"shortId"`
				VipSuffix string `json:"vipSuffix"`
			} `json:"iCloudEnv"`
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"account"`
		Contacts struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"contacts"`
		Developerapi struct {
			Url    string `json:"url"`
			Status string `json:"status"`
		} `json:"developerapi"`
	} `json:"webservices"`
	PcsEnabled        bool `json:"pcsEnabled"`
	TermsUpdateNeeded bool `json:"termsUpdateNeeded"`
	ConfigBag         struct {
		Urls struct {
			AccountCreateUI     string `json:"accountCreateUI"`
			AccountLoginUI      string `json:"accountLoginUI"`
			AccountLogin        string `json:"accountLogin"`
			AccountRepairUI     string `json:"accountRepairUI"`
			DownloadICloudTerms string `json:"downloadICloudTerms"`
			RepairDone          string `json:"repairDone"`
			AccountAuthorizeUI  string `json:"accountAuthorizeUI"`
			VettingUrlForEmail  string `json:"vettingUrlForEmail"`
			AccountCreate       string `json:"accountCreate"`
			GetICloudTerms      string `json:"getICloudTerms"`
			VettingUrlForPhone  string `json:"vettingUrlForPhone"`
		} `json:"urls"`
		AccountCreateEnabled bool `json:"accountCreateEnabled"`
	} `json:"configBag"`
	HsaTrustedBrowser            bool     `json:"hsaTrustedBrowser"`
	AppsOrder                    []string `json:"appsOrder"`
	Version                      int      `json:"version"`
	IsExtendedLogin              bool     `json:"isExtendedLogin"`
	PcsServiceIdentitiesIncluded bool     `json:"pcsServiceIdentitiesIncluded"`
	IsRepairNeeded               bool     `json:"isRepairNeeded"`
	HsaChallengeRequired         bool     `json:"hsaChallengeRequired"`
	RequestInfo                  struct {
		Country  string `json:"country"`
		TimeZone string `json:"timeZone"`
		Region   string `json:"region"`
	} `json:"requestInfo"`
	PcsDeleted bool `json:"pcsDeleted"`
	ICloudInfo struct {
		SafariBookmarksHasMigratedToCloudKit bool `json:"SafariBookmarksHasMigratedToCloudKit"`
	} `json:"iCloudInfo"`
	Apps struct {
		Calendar struct {
		} `json:"calendar"`
		Reminders struct {
		} `json:"reminders"`
		Keynote struct {
			IsQualifiedForBeta bool `json:"isQualifiedForBeta"`
		} `json:"keynote"`
		Settings struct {
			CanLaunchWithOneFactor bool `json:"canLaunchWithOneFactor"`
		} `json:"settings"`
		Pages struct {
			IsQualifiedForBeta bool `json:"isQualifiedForBeta"`
		} `json:"pages"`
		Mail struct {
		} `json:"mail"`
		Notes3 struct {
		} `json:"notes3"`
		Find struct {
			CanLaunchWithOneFactor bool `json:"canLaunchWithOneFactor"`
		} `json:"find"`
		Iclouddrive struct {
		} `json:"iclouddrive"`
		Numbers struct {
			IsQualifiedForBeta bool `json:"isQualifiedForBeta"`
		} `json:"numbers"`
		Photos struct {
		} `json:"photos"`
		Contacts struct {
		} `json:"contacts"`
	} `json:"apps"`
}
