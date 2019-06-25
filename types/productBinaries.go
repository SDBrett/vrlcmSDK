package types

type RefreshPatchResponse struct {
	RequestId string `json:"requestId"`
}

type ProductSourceMapping struct {
	LocationType string   `json:"locationType"`
	BaseLocation string   `json:"baseLocation"`
	FilesToMap   []string `json:"filesToMap"`
}

type ProductBinarySourceDocument struct {
	DocumentLinks []string `json:"documentLinks"`
	Documents     map[string]DocumentSourceLocation	`json:"documents"`
	DocumentCount int `json:"documentCount"`
}

type DocumentSourceLocation struct {
	SourceLocation               string `json:"sourceLocation"`
	SourceType                   string `json:"sourceType"`
	NfsMountLocation             string `json:"nfsMountLocation"`
	DocumentVersion              int    `json:"documentVersion"`
	DocumentEpoch                int    `json:"documentEpoch"`
	DocumentKind                 string `json:"documentKind"`
	DocumentSelfLink             string `json:"documentSelfLink"`
}


type ProductBinarySourceLocation struct {
	SourceLocation   string   `json:"sourceLocation"`
	SourceType       string   `json:"sourceType"`
}

/*// https://192.168.17.128/lcm/gui/api/sourceLocation
//{
//  "sourceType": "Local",
//  "sourceLocation": "/data/ova"
//}
//{"sourceLocation":"/data/ova","sourceType":"Local","productOVAs":["VMware-vRO-Appliance-7.5.0.458-10110089_OVF10.ova"],
"nfsMountLocation":"/data/nfsfiles","documentVersion":0,"documentEpoch":0,
"documentKind":"com:vmware:vrealize:lcm:nxui:document:OVFDeploymentSourceConfig","documentSelfLink":"/lcm/gui/api/sourceLocation/2c551b23da979e7558aae62bffc48",
"documentUpdateTimeMicros":1559856469245001,"documentUpdateAction":"POST","documentExpirationTimeMicros":0,"documentOwner":"cd4e4611-c19f-403c-8de3-466af8227407","documentAuthPrincipalLink":"/core/authz/users/vLCMAdmin"}

*/

/*

{{Server}}/lcm/gui/api/sourceLocation?expand=true With only local source set
{
    "documentLinks": [
        "/lcm/gui/api/sourceLocation/2c551b23da979e7558ab19f912f90"
    ],
    "documents": {
        "/lcm/gui/api/sourceLocation/2c551b23da979e7558ab19f912f90": {
            "sourceLocation": "/data/ova1",
            "sourceType": "Local",
            "nfsMountLocation": "/data/nfsfiles",
            "documentVersion": 3,
            "documentEpoch": 0,
            "documentKind": "com:vmware:vrealize:lcm:nxui:document:OVFDeploymentSourceConfig",
            "documentSelfLink": "/lcm/gui/api/sourceLocation/2c551b23da979e7558ab19f912f90",
            "documentUpdateTimeMicros": 1559871757323003,
            "documentUpdateAction": "PUT",
            "documentExpirationTimeMicros": 0,
            "documentOwner": "cd4e4611-c19f-403c-8de3-466af8227407",
            "documentAuthPrincipalLink": "/core/authz/users/vLCMAdmin"
        }
    },
    "documentCount": 1,
    "queryTimeMicros": 1,
    "documentVersion": 0,
    "documentUpdateTimeMicros": 0,
    "documentExpirationTimeMicros": 0,
    "documentOwner": "cd4e4611-c19f-403c-8de3-466af8227407"
}
*/
