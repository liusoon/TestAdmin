package InterfaceTestPartCtl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"salotto/model/vo"
	"salotto/service/InterfaceTestPartSrv"
	"salotto/utils"
	"salotto/utils/qjson"
)

func AddInterface(c *gin.Context) {
	var interfaceInfo vo.InterfaceInfoVO
	//var interfaceInfo InterfaceTestPartEntity.TInterfaceInfo
	//c.ShouldBind(&interfaceInfo)
	//InterfaceTestPartSrv.ItfTestSrv.AddInterface(&interfaceInfo)
	//res, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Println(string(res))
	c.ShouldBind(&interfaceInfo)
	fmt.Println(interfaceInfo)
	InterfaceTestPartSrv.ItfTestSrv.AddInterface(&interfaceInfo)
	utils.ResponseOkWithMsg(c, "新增成功", nil)
}

func EditInterface(c *gin.Context) {
	var interfaceInfo vo.InterfaceInfoVO
	c.ShouldBind(&interfaceInfo)
	if err := InterfaceTestPartSrv.ItfTestSrv.EditInterface(&interfaceInfo); err != nil {
		utils.ResponseOkWithMsg(c, "修改失败", nil)
	} else {
		utils.ResponseOkWithMsg(c, "修改成功", nil)
	}
}

func DelInterface(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if err := InterfaceTestPartSrv.ItfTestSrv.DelInterface(&reqInfo); err != nil {
		utils.ResponseOkWithMsg(c, "删除失败", err)
	}
	utils.ResponseOkWithMsg(c, "删除成功", nil)
}

func GetInterfaceList(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if res, err := InterfaceTestPartSrv.ItfTestSrv.GetInterfaceList(&reqInfo); err != nil {
		utils.ResponseOkWithMsg(c, "查询失败", nil)
	} else {
		utils.ResponseOk(c, res)
	}
}

func ImportSwagger(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	InterfaceTestPartSrv.ItfTestSrv.ImportSwagger(&reqInfo)
	utils.ResponseOkWithMsg(c, "导入成功", nil)
}

func GetInterfaceSelectOptions(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if res, err := InterfaceTestPartSrv.ItfTestSrv.GetInterfaceSelectOptions(&reqInfo); err != nil {
		utils.ResponseOkWithMsg(c, "查询失败", nil)
	} else {
		utils.ResponseOk(c, res)
	}
}

func GetSingleInterfaceInfo(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if res, err := InterfaceTestPartSrv.ItfTestSrv.GetSingleInterfaceInfo(&reqInfo); err != nil {
		utils.ResponseOkWithMsg(c, "查询失败", nil)
	} else {
		utils.ResponseOk(c, res)
	}
}
