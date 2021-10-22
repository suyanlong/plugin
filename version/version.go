package version

//var version control
var (
	Version   = "6.1.1"
	GitCommit string
	BuildTime string
)

//GetVersion 获取版本信息
func GetVersion() string {
	if GitCommit != "" {
		return Version + "-" + GitCommit
	}
	return Version
}
