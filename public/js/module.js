mini.parse();
var treeModule = mini.get("treeModule");
var winModule = mini.get("winModule");
var formModule = new mini.Form("formModule");
var selectModule = undefined;

var gridFunction = mini.get("gridFunction");
var formFunction = new mini.Form("formFunction");
var winFunc = mini.get("winFunc");
var selectFunction = undefined;

function onModuleSelect(e) {
	selectModule = e.node;
	selectFunction = undefined;
	if (selectModule.Code=="system") {
		return;
	}
	SelectFunc();
}

function OpenModuleAdd() {
	if (!selectModule) {
		mini.alert("请选择父模块");
		return;
	}
	OpenAddForm(formModule, winModule, "添加模块", "icon-add", function() {
		var parId = selectModule.ID;
		var parName = selectModule.Name;
		mini.get("hidPID").setValue(parId);
		mini.get("btnModuleAdd").show();
		mini.get("btnModuleEdit").hide();
		$('#trDsType').css("display", "none");
		$('#spParentName').html(parName);
		SetAsInput("txtCode");
	});
}

function OpenModuleEdit() {
	if (!selectModule) {
		mini.alert("请选择模块");
		return;
	}
	var url = "/System/ModuleInfo?id=" + selectModule.ID;
	OpenEditForm(url, formModule, winModule, "编辑模块", "icon-edit", function() {
		var Name = treeModule.getParentNode(selectModule).Name;
		$('#spParentName').html(Name);
		mini.get("btnModuleAdd").hide();
		mini.get("btnModuleEdit").show();
		SetAsLabel("txtCode");
	}, function(data) {
		if (data.DisplayType == 1) {
			$('#trDsType').css("display", "");
		} else {
			$('#trDsType').css("display", "none");
		}
	});
}

function OnTypeChanged(e) {
	if (e.value == "1") {
		$('#trDsType').css("display", "");
	} else {
		$('#trDsType').css("display", "none");
	}
}
function onPathValidation(e)
{
    if (e.isValid) {
        var re = new RegExp("^\/[a-zA-Z]+");
        if (e.value.replace(' ','')=='') {
            e.isValid = true;
            return;
        }
        if (!re.test(e.value)) {
            e.errorText = "必须为：/ABC/DCF格式";
            e.isValid = false;
        }
    }
}
function AddModule() {
	if (!CheckForm(formModule))
		return;
	var data = formModule.getData();
	Ajax({
		url: "/System/ModuleInsert",
		type: "post",
		data: data,
		success: function(msg) {
			if (msg.Result == 0) {
				ShowTips("保存成功");
				HideWin(winModule);
				var newNode = {
					ID: msg.Message,
					Name: data.Name,
					ParentID: data.ParentID
				};

				treeModule.addNode(newNode, 0, selectModule);
			}
		}
	});
}

function EditModule() {
	if (!CheckForm(formModule))
		return;
	var data = formModule.getData();
	Ajax({
		url: "/System/ModuleUpdate",
		type: "post",
		data: data,
		success: function(msg) {
			if (msg.Result == 0) {
				ShowTips("保存成功");
				HideWin(winModule);
				var newNode = {
					RoleName: data.RoleName,
				};
				treeModule.updateNode(selectModule, newNode);
			}
		}
	});
}

function DeleteModule() {
	if (!selectModule) {
		mini.alert("请选择模块");
		return;
	}
	if (!treeModule.isLeaf(selectModule)) {
		mini.alert("有子模块不能删除");
		return;
	}
	mini.confirm("是否删除？", "提示", function(action) {
		if (action == "ok") {
			var url = "/System/ModuleDelete?id=" + selectModule.ID;
			Ajax({
				url: url,
				type: "post",
				success: function(msg) {
					if (msg.Result == 0) {
						ShowTips("删除成功");
						treeModule.removeNode(selectModule);
					}
				},
			});
		}
	});

}
//=======================================================================
function OnFunctionSelect(e) {
	selectFunction = e.record;
}

function SelectFunc() {
	var moduleid = selectModule.ID;
	if (selectModule.Code == "system")
		return;
	gridFunction.load({
		moduleid: moduleid
	});
}

function OpenFunctionAdd() {
	if (!selectModule) {
		mini.alert("请选择所属模块");
		return;
	}
	OpenAddForm(formFunction, winFunc, "添加模块", "icon-add", function() {
		var moduleID = selectModule.ID;
		var moduleCode = selectModule.Code;
		var moduleName = selectModule.Name;
		mini.get("hidFunModuleID").setValue(moduleID);
		mini.get("hidFunModuleCode").setValue(moduleCode);
		mini.get("btnFunctionAdd").show();
		mini.get("btnFunctionEdit").hide();
		$('#spModuleName').html(moduleName);
		SetAsInput("txtFunCode");
	});
}

function OpenFunctionEdit() {
	if (!selectFunction) {
		mini.alert("请选择功能点");
		return;
	}
	var url = "/System/FuncInfo?id=" + selectFunction.ID;
	OpenEditForm(url, formFunction, winFunc, "编辑模块", "icon-edit", function() {
		var Name = treeModule.getParentNode(selectModule).Name;
		$('#spModuleName').html(Name);
		mini.get("btnFunctionAdd").hide();
		mini.get("btnFunctionEdit").show();
		SetAsLabel("txtFunCode");
	});
}

function AddFunction() {
	if (!CheckForm(formFunction))
		return;
	var data = formFunction.getData();
	Ajax({
		url: "/System/FuncInsert",
		type: "post",
		data: data,
		success: function(msg) {
			if (msg.Result == 0) {
				ShowTips("保存成功");
				HideWin(winFunc);
				data.ID = msg.Message;
				gridFunction.addRow(data);
			}
		}
	});
}

function EditFunction() {
	if (!CheckForm(formFunction))
		return;
	var data = formFunction.getData();
	Ajax({
		url: "/System/FuncUpdate",
		type: "post",
		data: data,
		success: function(msg) {
			if (msg.Result == 0) {
				ShowTips("保存成功");
				HideWin(winFunc);
				gridFunction.updateRow(selectFunction, data);
			}
		}
	});
}

function DeleteFunction() {
	if (!selectFunction) {
		mini.alert("请选择功能点");
		return;
	}
	mini.confirm("是否删除？", "提示", function(action) {
		if (action == "ok") {
			var url = "/System/FuncDelete?id=" + selectFunction.ID;
			Ajax({
				url: url,
				type: "post",
				success: function(msg) {
					if (msg.Result == 0) {
						ShowTips("删除成功");
						gridFunction.removeRow(selectFunction);
					}
				},
			});
		}
	});

}