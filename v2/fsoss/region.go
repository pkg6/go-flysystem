package fsoss

var (
	RegionCnHangzhou      = "oss-cn-hangzhou"
	RegionCnShanghai      = "oss-cn-shanghai"
	RegionCnNanjing       = "oss-cn-nanjing"
	RegionCnFuzhou        = "oss-cn-fuzhou"
	RegionCnQingdao       = "oss-cn-qingdao"
	RegionCnBeijing       = "oss-cn-beijing"
	RegionCnZhangjiakou   = "oss-cn-zhangjiakou"
	RegionCnHuhehaote     = "oss-cn-huhehaote"
	RegionCnWulanchabu    = "oss-cn-wulanchabu"
	RegionCnShenzhen      = "oss-cn-shenzhen"
	RegionCnHeyuan        = "oss-cn-heyuan"
	RegionCnGuangzhou     = "oss-cn-guangzhou"
	RegionCnChengdu       = "oss-cn-chengdu"
	RegionCnHongkong      = "oss-cn-hongkong"
	RegionUsWest1         = "oss-us-west-1"
	RegionUsEast1         = "oss-us-east-1"
	RegionApNortheast1    = "oss-ap-northeast-1"
	RegionApNortheast2    = "oss-ap-northeast-2"
	RegionApSoutheast1    = "oss-ap-southeast-1"
	RegionApSoutheast2    = "oss-ap-southeast-2"
	RegionApSoutheast3    = "oss-ap-southeast-3"
	RegionApSoutheast5    = "oss-ap-southeast-5"
	RegionApSoutheast6    = "oss-ap-southeast-6"
	RegionApSoutheast7    = "oss-ap-southeast-7"
	RegionApSouth1        = "oss-ap-south-1"
	RegionEuCentral1      = "oss-eu-central-1"
	RegionEuWest1         = "oss-eu-west-1"
	RegionMeEast1         = "oss-me-east-1"
	RegionRgChinaMainland = "oss-rg-china-mainland"
)

func Endpoint(regionID string) string {
	return regionID + "aliyuncs.com"
}
func InternalEndpoint(regionID string) string {
	return regionID + "-internal.aliyuncs.com"
}
