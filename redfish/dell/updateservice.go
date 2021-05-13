package dell

import "encoding/json"

type DellUpdateService struct {
	Actions DellUpdateServiceActions
}

type DellUpdateServiceActions struct {
	InstallUpon []string
	Target      string
}

func (d *DellUpdateServiceActions) UnmarshalJSON(b []byte) error {
	var t struct {
		DellUpdateService struct {
			InstallUpon []string `json:"InstallUpon@Redfish.AllowableValues"`
			Target      string
		} `json:"DellUpdateService.v1_0_0#DellUpdateService.Install"`
	}
	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	d.InstallUpon = t.DellUpdateService.InstallUpon
	d.Target = t.DellUpdateService.Target

	return nil
}
